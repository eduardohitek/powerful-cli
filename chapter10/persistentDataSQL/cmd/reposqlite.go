package cmd

import (
	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro"
	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}
	return repo, nil
}
