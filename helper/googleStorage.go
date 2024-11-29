package helper

import (
	"app/config"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
)

func UploadAndResizeImage(ctx context.Context, fileHeader *multipart.FileHeader, path string) (string, error) {
	// Open the file from the multipart.FileHeader
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Resize the image using the helper function
	buf, newFileName, err := ResizeImage(file, fileHeader.Filename, fileHeader.Size)
	if err != nil {
		return "", fmt.Errorf("failed to resize image: %v", err)
	}

	// Upload the resized image to GCS using the original UploadFileGCS function
	publicURL, err := UploadFileGCSFromImageSetPath(ctx, buf, newFileName, path)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to GCS: %v", err)
	}

	// Check if publicURL is nil before dereferencing
	if publicURL == nil {
		return "", fmt.Errorf("upload to GCS failed: received nil URL")
	}

	return *publicURL, nil
}

func ResizeImage(file multipart.File, fileName string, originalSize int64) (*bytes.Buffer, string, error) {
	ext := filepath.Ext(fileName)

	// Read the file into memory
	var img image.Image
	var err error

	// Read the file into a buffer so it can be reused
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file: %v", err)
	}

	// Decode the image based on file extension
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(bytes.NewReader(fileBytes))
	case ".png":
		img, err = png.Decode(bytes.NewReader(fileBytes))
	case ".gif":
		img, err = gif.Decode(bytes.NewReader(fileBytes))
	case ".webp":
		img, err = webp.Decode(bytes.NewReader(fileBytes))
	default:
		// Attempt to decode using the generic image.Decode
		img, _, err = image.Decode(bytes.NewReader(fileBytes))
	}

	// If decoding fails, return an error stating that the file is not a valid image
	if err != nil {
		return nil, "", fmt.Errorf("file is not a valid image: %v", err)
	}

	// Determine the resizing dimensions based on the file size
	var newWidth uint
	if originalSize > 1*1024*1024 { // If the original size is greater than 1MB
		newWidth = 800 // Resize to 800px width; adjust as needed
	} else {
		newWidth = uint(img.Bounds().Dx()) // Keep the original width
	}

	// Resize the image
	newImg := resize.Resize(newWidth, 0, img, resize.Lanczos3)

	// Create a buffer to hold the encoded image
	buf := new(bytes.Buffer)

	// Always encode as JPEG to ensure consistent output format
	err = jpeg.Encode(buf, newImg, &jpeg.Options{Quality: 80})
	if err != nil {
		return nil, "", fmt.Errorf("failed to encode image to jpeg: %v", err)
	}

	// Generate a unique filename
	newFileName := uuid.New().String() + ".jpg"

	return buf, newFileName, nil
}

func UploadFileGCSFromImageSetPath(ctx context.Context, buf *bytes.Buffer, fileName string, path string) (*string, error) {
	bucket := os.Getenv("BUCKET_NAME")
	storageClient, err := storage.NewClient(ctx, config.StorageConfig()...)
	if err != nil {
		return nil, err
	}

	// Define the object name in GCS
	filePath := path + "/" + fileName

	obj := storageClient.Bucket(bucket).Object(filePath)

	// Write the buffer to the GCS object
	writer := obj.NewWriter(ctx)
	if _, err := buf.WriteTo(writer); err != nil {
		return nil, fmt.Errorf("failed to write buffer to GCS: %v", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close GCS writer: %v", err)
	}

	// Make the object publicly accessible
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, fmt.Errorf("failed to set ACL on GCS object: %v", err)
	}

	// Return the public URL
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucket, filePath)
	return &publicURL, nil
}
