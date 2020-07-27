package statzagg

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestLogStats(t *testing.T) {
	urls := []string{
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/init.mp4",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/1.m4s",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/2.m4s",
	}
	logs := NewLogStatzAgg(os.Stderr)
	for i, url := range urls {
		go func(url string, i int) {
			ctx := context.TODO()
			stats := helperGetURL(ctx, url, "client"+strconv.Itoa(i), 10*time.Second)
			logs.PostHTTPClientStats(ctx, stats)
		}(url, i)
	}
}
