package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/lgen/pkg/log"
	"github.com/rs/zerolog"
)

const PackageName string = "handler"

type Storage interface{}

type Tex interface{}

type Handler struct {
	logger  *zerolog.Logger
	storate Storage
	tex     Tex
}

func Get(storage Storage, tex Tex) *Handler {
	return &Handler{
		logger:  log.For(PackageName),
		storate: storage,
		tex:     tex,
	}
}

func (h *Handler) okResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

type ErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func (h *Handler) errorResponse(c *gin.Context, httpResponseCode int, err error) {
	errorCodeMap := map[int]string{
		http.StatusBadRequest:          "bad_request",           // 400
		http.StatusNotFound:            "not_found",             // 404
		http.StatusInternalServerError: "internal_server_error", // 500
	}

	response := ErrorResponse{
		Code:  errorCodeMap[httpResponseCode],
		Error: err.Error(),
	}

	c.JSON(httpResponseCode, response)

	h.logger.Error().
		Err(err).
		Int("code", httpResponseCode).
		Str("uri", c.Request.URL.Path).
		Msg("error response")

	return
}
