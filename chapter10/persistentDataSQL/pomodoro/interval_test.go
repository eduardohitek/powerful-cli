package pomodoro_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/eduardohitek/powerful-cli/chapter9/pomo/pomodoro"
)

func TestNewConfig(t *testing.T) {
	testCases := []struct {
		name     string
		input    [3]time.Duration
		expected pomodoro.IntervalConfig
	}{
		{
			name: "Default",
			expected: pomodoro.IntervalConfig{
				PomodoroDuration:   25 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
				LongBreakDuration:  15 * time.Minute,
			},
		},
		{
			name: "SingleInput",
			input: [3]time.Duration{
				20 * time.Minute,
			},
			expected: pomodoro.IntervalConfig{
				PomodoroDuration:   20 * time.Minute,
				ShortBreakDuration: 5 * time.Minute,
				LongBreakDuration:  15 * time.Minute,
			},
		},
		{
			name: "MultiInput",
			input: [3]time.Duration{
				20 * time.Minute,
				10 * time.Minute,
				12 * time.Minute,
			},
			expected: pomodoro.IntervalConfig{
				PomodoroDuration:   20 * time.Minute,
				ShortBreakDuration: 10 * time.Minute,
				LongBreakDuration:  12 * time.Minute,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo pomodoro.Repository
			config := pomodoro.NewConfig(
				repo, tc.input[0], tc.input[1], tc.input[2],
			)

			if config.PomodoroDuration != tc.expected.PomodoroDuration {
				t.Errorf("Expected Pomodoro Duration %q, got %q instead\n",
					tc.expected.PomodoroDuration, config.PomodoroDuration)
			}
			if config.ShortBreakDuration != tc.expected.ShortBreakDuration {
				t.Errorf("Expected ShortBreak Duration %q, got %q instead\n",
					tc.expected.ShortBreakDuration, config.ShortBreakDuration)
			}
			if config.LongBreakDuration != tc.expected.LongBreakDuration {
				t.Errorf("Expected LongBreak Duration %q, got %q instead\n",
					tc.expected.LongBreakDuration, config.LongBreakDuration)
			}
		})
	}
}

func TestGetInterval(t *testing.T) {
	repo, cleanup := getRepo(t)
	defer cleanup()

	const duration = 1 * time.Millisecond
	config := pomodoro.NewConfig(repo, 3*duration, duration, 2*duration)

	for i := 1; i <= 16; i++ {
		var (
			expCategory string
			expDuration time.Duration
		)

		switch {
		case i%2 != 0:
			expCategory = pomodoro.CategoryPomodoro
			expDuration = 3 * duration
		case i%8 == 0:
			expCategory = pomodoro.CategoryLongBreak
			expDuration = 2 * duration
		case i%2 == 0:
			expCategory = pomodoro.CategoryShortBreak
			expDuration = duration
		}

		testName := fmt.Sprintf("%s%d", expCategory, i)
		t.Run(testName, func(t *testing.T) {
			res, err := pomodoro.GetInterval(config)

			if err != nil {
				t.Errorf("Expected no error, got %q.\n", err)
			}

			noop := func(pomodoro.Interval) {}
			err = res.Start(context.Background(), config, noop, noop, noop)
			if err != nil {
				t.Fatal(err)
			}

			if res.Category != expCategory {
				t.Errorf("Expected category %q, got %q.\n", expCategory, res.Category)
			}

			if res.PlannedDurarion != expDuration {
				t.Errorf("Expected PlannedDuration %q, got %q.\n", expDuration, res.PlannedDurarion)
			}

			if res.State != pomodoro.StateNotStarted {
				t.Errorf("Expected State = %q, got %q.\n",
					pomodoro.StateNotStarted, res.State)
			}
			ui, err := repo.ByID(res.ID)
			if err != nil {
				t.Errorf("Expected no error. Got %q.\n", err)
			}
			if ui.State != pomodoro.StateDone {
				t.Errorf("Expected State = %q, got %q.\n",
					pomodoro.StateDone, res.State)
			}
		})
	}
}
