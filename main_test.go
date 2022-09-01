package main

import (
	"reflect"
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
		{name: "monday", input: "monday", want: dayOnThisWeek(time.Monday)},
		{name: "friday", input: "friday", want: dayOnThisWeek(time.Friday)},
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

func Test_convertToInterval(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    daysToLog
		wantErr bool
	}{
		{name: "empty", input: "", want: daysToLog{time.Now().Truncate(24 * time.Hour).UTC(): struct{}{}}},
		{name: "just monday", input: "monday", want: dayOnThisWeek(time.Monday)},
		// {name: "monday to friday", input: "monday-friday", want: dayOnThisWeek(time.Monday)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToDays(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertToIntervals incorrect result; want: %+v, got: %+v", tt.want, got)
			}
		})
	}
}

func dayOnThisWeek(day time.Weekday) time.Time {
	return time.Now().Add(daysDiff(day, time.Now().Weekday()))
}

func daysDiff(a, b time.Weekday) time.Duration {
	diff := a - b
	return time.Duration(diff*24) * time.Hour
}
