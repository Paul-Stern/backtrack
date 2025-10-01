package model

import (
	"context"
	"errors"
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
