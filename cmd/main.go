package main

import (
	"git.dmp.true.th/DMP-Public/zerolog-wrapper/logger"
	"github.com/pkg/errors"
)

func main() {
	logger.New("debug")
	logger.Get().Debug().Str("name", "gam").Msg("Test")
	logger.Get().Err(errors.New("error")).Str("name", "gam").Msg("Test")
	logger.Get().Info().Str("name", "gam").Msg("Test")
	logger.Get().Warn().Str("name", "gam").Msg("Test")
}
