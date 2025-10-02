package model

import (
	"time"

	"github.com/rs/zerolog/log"
)

func TryParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		return t
	}
	t, err = time.Parse(time.DateOnly, s)
	if err == nil {
		return t
	}
	log.Fatal().Err(err).Msg("Failed to parse time")
	return time.Time{}
}
