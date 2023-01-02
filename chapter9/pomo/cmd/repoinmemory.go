package cmd

import (
	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro"
	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
