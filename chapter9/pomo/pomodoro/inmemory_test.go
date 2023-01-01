package pomodoro_test

import (
	"testing"

	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro"
	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()
	return repository.NewInMemoryRepo(), func() {}
}
