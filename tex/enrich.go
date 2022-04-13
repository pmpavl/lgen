package tex

import (
	"fmt"
	"strings"

	"github.com/pmpavl/lgen-core/model"
)

const (
	placeReplacementLeafletTheme string = `LEAFLET_THEME`
	placeReplacementLeafletClass string = `LEAFLET_CLASS`
)

type Enrich struct {
	Theme string
	Class int
}

func (t *Tex) TemplateEnrich(template *model.Template, enrich *Enrich) *model.Template {
	begin := template.Layout.Begin

	if count := strings.Count(begin, placeReplacementLeafletTheme); count != 1 {
		t.logger.Debug().Str("templateName", template.Name).Msgf("%d place for replacement leaflet theme", count)
	} else {
		begin = strings.ReplaceAll(
			begin,
			placeReplacementLeafletTheme,
			fmt.Sprintf(`\newcommand{\theme}{%s}`, enrich.Theme),
		)
	}

	if count := strings.Count(begin, placeReplacementLeafletClass); count != 1 {
		t.logger.Debug().Str("templateName", template.Name).Msgf("%d place for replacement leaflet class", count)
	} else {
		begin = strings.ReplaceAll(
			begin,
			placeReplacementLeafletClass,
			fmt.Sprintf(`\newcommand{\class}{%d}`, enrich.Class),
		)
	}

	template.Layout.Begin = begin

	return template
}
