package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneratePingRequest struct{}

type GeneratePingResponse struct {
	FileNamePdf string `json:"fileNamePdf"`
	FileNameTex string `json:"fileNameTex"`
}

func (h *Handler) GeneratePing(c *gin.Context) {
	var request GeneratePingRequest

	if err := c.BindJSON(&request); err != nil {
		h.errorResponse(c, http.StatusBadRequest, err)

		return
	}

	response := GeneratePingResponse{
		FileNamePdf: "ping.pdf",
		FileNameTex: "ping.tex",
	}

	h.okResponse(c, response)
}
