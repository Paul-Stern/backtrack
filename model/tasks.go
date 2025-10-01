package model

import (
	"context"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Tasks []Task

func GetTasks() Tasks {
	ctx := context.Background()
	tasks, err := gorm.G[Task](DB).Find(ctx)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to list tasks")
	}
	return tasks
}
