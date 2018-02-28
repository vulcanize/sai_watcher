package pep

import (
	"github.com/vulcanize/vulcanizedb/libraries/shared"
)

func HandlerInitializers() []shared.HandlerInitializer {
	return []shared.HandlerInitializer{
		NewPepHandler,
	}
}