package model

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Tasks []Task

func GetTasks() Tasks {
	if DB == nil {
		log.Info().Msg("DB is nil")
		return nil
	}
	ctx := context.Background()
	tasks, err := gorm.G[Task](DB).Find(ctx)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to list tasks")
	}
	if tasks == nil {
		log.Info().Msg("No tasks found")
		return nil
	}
	return tasks
}

func (t Tasks) String() string {
	b := strings.Builder{}
	for i, task := range t {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(task.String())
	}
	return b.String()
}
