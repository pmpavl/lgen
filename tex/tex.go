package tex

import (
	"fmt"
	"time"

	"github.com/pmpavl/lgen-core/model"
	log "github.com/pmpavl/lgen-log"
	"github.com/rs/zerolog"
)

type Tex struct {
	logger *zerolog.Logger
}

func Get() *Tex {
	return &Tex{
		logger: log.For("tex"),
	}
}

const (
	defaultCommandGenerator string = "pdflatex"
	figureDirectory         string = "./figure"
)

type Document struct {
	Name string
	Tex  string

	CommandGenerator string
}

func (t *Tex) GenerateDocument(template *model.Template, tasks []*model.Task) *Document {
	// TODO: Think about how best to fill out the document
	tex := template.Layout.Begin

	for _, task := range tasks {
		if task.Figure != "" {
			figurePath := fmt.Sprintf("%s/%s", figureDirectory, task.Figure)
			task.Task += fmt.Sprintf(`$$\includegraphics[height=50px]{%s}$$`, figurePath)
		}

		tex += fmt.Sprintf(template.Layout.Task,
			task.Task,
			task.Answer,
			task.Help,
			task.Solution,
		)
	}

	tex += template.Layout.End

	return &Document{
		Name:             fmt.Sprintf("document_%d", time.Now().UnixMilli()),
		Tex:              tex,
		CommandGenerator: defaultCommandGenerator,
	}
}
