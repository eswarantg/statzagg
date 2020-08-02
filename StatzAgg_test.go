package statzagg

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestHTTPLogStats(t *testing.T) {
	urls := []string{
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/init.mp4",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/1.m4s",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/2.m4s",
	}
	logs := NewLogStatzAgg(os.Stderr)
	for i, url := range urls {
		func(url string, i int) {
			ctx := context.TODO()
			stats := helperGetURL(ctx, url, "client"+strconv.Itoa(i), 10*time.Second)
			logs.PostHTTPClientStats(ctx, stats)
		}(url, i)
	}
}

func TestEventLogStats(t *testing.T) {
	events := []string{
		"Event1",
		"Event1",
		"Event1",
	}
	errorList := []error{
		nil,
		nil,
		nil,
	}
	values := [][]string{
		{"one"},
		{"two", "two"},
		{"three", "three", "three"},
	}
	logs := NewLogStatzAgg(os.Stderr)
	ctx := context.TODO()
	for i := range events {
		var e EventStats
		e.ID = strconv.FormatInt(int64(i), 10)
		e.Name = events[i]
		e.Err = errorList[i]
		e.Values = make([]interface{}, len(values[i]))
		for j, value := range values[i] {
			e.Values[j] = value
		}
		logs.PostEventStats(ctx, &e)
	}
}
