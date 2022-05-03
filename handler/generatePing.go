package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneratePingRequest struct{}

type GeneratePingResponse struct {
	UriPdf string `json:"uriPdf"`
	UriTex string `json:"uriTex"`
}

func (h *Handler) GeneratePing(c *gin.Context) {
	var request GeneratePingRequest

	if err := c.BindJSON(&request); err != nil {
		h.errorResponse(c, http.StatusBadRequest, err)

		return
	}

	response := GeneratePingResponse{
		UriPdf: "/static/gen/ping.pdf",
		UriTex: "/static/tex/ping.tex",
	}

	h.okResponse(c, response)
}
