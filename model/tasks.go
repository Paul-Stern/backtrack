package model

import (
	"context"
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Tasks []Task

func GetTasks(limit int) Tasks {
	if DB == nil {
		log.Info().Msg("DB is nil")
		return nil
	}
	ctx := context.Background()
	tasks, err := gorm.G[Task](DB).Order("ID desc").Limit(limit).Find(ctx)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to list tasks")
	}
	if tasks == nil {
		log.Info().Msg("No tasks found")
		return nil
	}
	return tasks
}

func Today() Tasks {
	// get Today's date
	td := time.Now().Round(time.Hour)
	if DB == nil {
		log.Info().Msg("DB is nil")
		return nil
	}
	ctx := context.Background()
	tasks, err := gorm.G[Task](DB).Where("time_finished >= ?", td).Order("ID desc").Limit(10).Find(ctx)
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

func (tasks Tasks) List() (l *tview.List) {
	l = tview.NewList()
	for _, t := range tasks {
		d, err := t.Duration()
		if err != nil {
			log.Error().Err(err).Caller().Msg("Failed to get duration")
			d = 0
		}
		l.AddItem(t.Message, d.String(), rune(0), nil)
	}
	return l
}
