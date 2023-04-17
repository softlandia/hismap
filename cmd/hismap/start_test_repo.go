package main

import (
	"github.com/rs/zerolog"
	"github.com/softlandia/hismap/service"
)

func testRepo(logger zerolog.Logger, repo service.Repositories) error {
	if err := repo.Items.Test(); err != nil {
		logger.Error().Err(err).Msg("<<< test table 'items'")
		return err
	}
	return nil
}
