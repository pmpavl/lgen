package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/lgen-core/model"
	log "github.com/pmpavl/lgen-log"
	"github.com/pmpavl/lgen/tex"
	"github.com/rs/zerolog"
)

type Tex interface {
	TemplateEnrich(template *model.Template, enrich *tex.Enrich) *model.Template
	GenerateDocument(template *model.Template, tasks []*model.Task) *tex.Document
}

type Storage interface {
	ReadTaskByID(ctx context.Context, id string) (*model.Task, error)
	ReadTasksByIDs(ctx context.Context, ids []string) ([]*model.Task, error)

	ReadTemplateByID(ctx context.Context, id string) (*model.Template, error)
	ReadTemplatesByIDs(ctx context.Context, ids []string) ([]*model.Template, error)
	ReadTemplateByName(ctx context.Context, name string) (*model.Template, error)
}

type Handler struct {
	logger *zerolog.Logger

	storate Storage
	tex     Tex
}

func Get(storage Storage, tex Tex) *Handler {
	return &Handler{
		logger:  log.For("handler"),
		storate: storage,
		tex:     tex,
	}
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
		Str("uri", c.Request.URL.Path).
		Int("code", httpResponseCode).
		Err(err).
		Msg("error response")

	return
}
