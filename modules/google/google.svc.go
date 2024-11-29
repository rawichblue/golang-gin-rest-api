package google

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func newService() *GoogleOAuthService {
	service := &GoogleOAuthService{
		oauth: nil,
	}

	kuy := os.Getenv("CLIENT_ID")

	log.Printf("kuy : %s", kuy)
	if err := service.Register(context.Background(), os.Getenv("REDIRECT_URL"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")); err != nil {
		panic(err)
	}

	return &GoogleOAuthService{
		oauth: service.oauth,
	}
}

type GoogleOAuthService struct {
	oauth *oauth2.Config
}

func (gos *GoogleOAuthService) Oauth() *oauth2.Config {

	return gos.oauth
}

func (gos *GoogleOAuthService) Register(ctx context.Context, redirectURL string, clientID string, clientSecret string) error {

	gos.oauth = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return nil
}
