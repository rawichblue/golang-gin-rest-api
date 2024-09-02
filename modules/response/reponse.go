package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Paginate struct {
	From  int64 `json:"from"`
	Size  int64 `json:"size"`
	Total int64 `json:"total"`
}

type Response struct {
	Status StatusResponse `json:"status"`
	Data   any            `json:"data,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message any    `json:"message"`
}

type ResponseError struct {
	Status StatusResponse `json:"status"`
	Error  Error          `json:"error"`
}

type ResponsePaginate struct {
	Status   StatusResponse `json:"status"`
	Data     any            `json:"data,omitempty"`
	Paginate any            `json:"paginate"`
}

func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{StatusResponse{
		Code:    200,
		Message: "Success",
	}, data})
}

func SuccessWithPaginate(ctx *gin.Context, data any, paginate any) {
	ctx.JSON(http.StatusOK, ResponsePaginate{StatusResponse{
		Code:    200,
		Message: "success",
	}, data, paginate})
}

func BadRequest(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusBadRequest, StatusResponse{
		Code:    400,
		Message: message.(string),
	})
}

func InternalError(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    404,
		Message: message.(string),
	})
}

func Unauthorized(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    401,
		Message: message.(string),
	})
}

func Forbidden(ctx *gin.Context, message any, payloadCode ...string) {
	ctx.JSON(http.StatusInternalServerError, StatusResponse{
		Code:    403,
		Message: message.(string),
	})
}

// type Pagination struct {
// 	CurrentPage int `json:"current_page"`
// 	PerPage     int `json:"per_page"`
// 	TotalPages  int `json:"total_pages"`
// 	Total       int `json:"total"`
// }

// type ResponsePaginate struct {
// 	Status     StatusResponse `json:"status"`
// 	Data       any            `json:"data,omitempty"`
// 	Pagination Pagination     `json:"pagination"`
// }

// type ResponsePaginate0 struct {
// 	Status     StatusResponse `json:"status"`
// 	Data       any            `json:"data,omitempty"`
// 	Pagination any            `json:"pagination"`
// }

// func SuccessWithPaginate(ctx *gin.Context, data any, paginate response.Paginate) {
// 	ctx.JSON(http.StatusOK, response.ResponsePaginate{
// 		Status: StatusResponse{
// 			Code:    200,
// 			Message: "Success",
// 		},
// 		Data:     data,
// 		Paginate: paginate,
// 	})
// }

// func SuccessWithPaginate(ctx *gin.Context, data any, pagination Pagination) {
// 	if pagination.Total == 0 {
// 		ctx.JSON(http.StatusOK, ResponsePaginate0{
// 			Status: StatusResponse{
// 				Code:    200,
// 				Message: "Success",
// 			},
// 			Data:       []any{},
// 			Pagination: gin.H{},
// 		})
// 		return
// 	} else {
// 		ctx.JSON(http.StatusOK, ResponsePaginate{
// 			Status: StatusResponse{
// 				Code:    200,
// 				Message: "Success",
// 			},
// 			Data:       data,
// 			Pagination: pagination,
// 		})
// 	}
// }
