package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/pmpavl/lgen-core/constant"
	"github.com/pmpavl/lgen/tex"
)

const DefaultDirectory string = "./gen"

type GenerateLeafletRequest struct {
	TemplateName        string   `json:"templateName,omitempty"`
	TemplateEnrichTheme string   `json:"templateEnrichTheme,omitempty"`
	TemplateEnrichClass int      `json:"templateEnrichClass,omitempty"`
	TaskIDs             []string `json:"taskIds"`
}

func (h *Handler) GenerateLeaflet(c *gin.Context) { // nolint
	var request GenerateLeafletRequest

	if err := c.BindJSON(&request); err != nil {
		h.errorResponse(c, http.StatusBadRequest, err)

		return
	}

	// If TemplateName is empty set DefaultTemplateName
	if request.TemplateName == "" {
		request.TemplateName = constant.DefaultTemplateName
	}

	// Read template by name from storage
	template, err := h.storate.ReadTemplateByName(c.Request.Context(), request.TemplateName)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err)

		return
	}

	// Template enrich
	template = h.tex.TemplateEnrich(template, &tex.Enrich{
		Theme: request.TemplateEnrichTheme,
		Class: request.TemplateEnrichClass,
	})

	// Read tasks by id from storage
	tasks, err := h.storate.ReadTasksByIDs(c.Request.Context(), request.TaskIDs)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err)

		return
	}

	// If tasks is empty return error
	if len(tasks) == 0 {
		err = errors.New("tasks not found")
		h.errorResponse(c, http.StatusNotFound, err)

		return
	}

	// Filling out a document
	document := h.tex.GenerateDocument(template, tasks)

	// Create save directory if not exist
	if _, err := os.Stat(DefaultDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(DefaultDirectory, os.ModeDir|0744); err != nil { // nolint
			h.errorResponse(c, http.StatusInternalServerError, err)

			return
		}
	}

	// File path
	filePath := fmt.Sprintf("%s/%s.tex", DefaultDirectory, document.Name)

	// Tex file entry
	file, _ := os.Create(filePath)
	file.WriteString(document.Tex)

	// Generate pdf from tex
	if err := exec.Command(
		document.CommandGenerator,
		fmt.Sprintf("-output-directory=%s", DefaultDirectory),
		filePath,
	).Run(); err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err)

		return
	}

	// Remove .log and .aux
	os.Remove(fmt.Sprintf("%s/%s.log", DefaultDirectory, document.Name))
	os.Remove(fmt.Sprintf("%s/%s.aux", DefaultDirectory, document.Name))
}
