package tex

import (
	"github.com/pmpavl/lgen/pkg/log"
	"github.com/rs/zerolog"
)

const PackageName string = "tex"

type Tex struct {
	logger *zerolog.Logger
}

func Get() *Tex {
	return &Tex{
		logger: log.For(PackageName),
	}
}
