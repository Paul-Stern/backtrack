package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TimeFinished time.Time
	Message      string
}

func NewTask(msg string) *Task {
	return &Task{
		TimeFinished: time.Now(),
		Message:      msg,
	}
}

func (t *Task) Add() (err error) {
	ctx := context.Background()
	if DB == nil {
		err = errors.New("db is nil")
		log.Error().Err(err).Caller().Msg("Failed to add task")
		return err
	}
	if t == nil {
		err = errors.New("task is nil")
		log.Error().Err(err).Caller().Msg("Failed to add task")
		return err
	}
	log.Debug().Interface("t", t).Msg("Adding task")
	err = gorm.G[Task](DB).Create(ctx, t)
	if err != nil {
		log.Error().Err(err).Caller().Msg("Failed to add task")
	}
	log.Info().Interface("t", t).Msg("Task added")
	return err
}

func (t *Task) Duration() (dur time.Duration, err error) {
	ctx := context.Background()
	if t.ID <= 1 {
		return 0, nil
	}
	pt, err := gorm.G[Task](DB).Where("ID = ?", t.ID-1).First(ctx)
	if err != nil {
		// log.Error().Err(err).Caller().Msg("Failed to get previous task")
		return 0, err
	}
	return t.TimeFinished.Sub(pt.TimeFinished).Round(time.Second), err
}

func (t *Task) String() string {
	et, err := t.Duration()
	if err != nil {
		return t.Message
	}
	return fmt.Sprintf("%d\t%s\t%s", t.ID, et.String(), t.Message)
}
