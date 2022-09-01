package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_convertToDay(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{name: "empty", input: "", want: time.Now()},
		{name: "today", input: "today", want: time.Now()},
		{name: "yesterday", input: "yesterday", want: time.Now().Add(-24 * time.Hour)},
		{name: "monday", input: "monday", want: time.Now().Add(daysDiff(time.Now().Weekday(), time.Monday))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToDay(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want.Truncate(24*time.Hour).UTC(), got)
		})
	}
}

func daysDiff(a, b time.Weekday) time.Duration {
	diff := a - b
	return time.Duration(diff*24) * time.Hour
}
