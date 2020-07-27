package statzagg

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func helperGetURL(ctx context.Context, url string, id string, timeout time.Duration) (statz *HTTPClientStatz) {
	var req *http.Request
	var err error
	statz = &HTTPClientStatz{}
	statz.ID = id
	statz.URL = url

	client := &http.Client{
		Timeout: timeout,
	}
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		statz.Err = fmt.Errorf("NewRequest Fail: %w", err)
		return
	}
	//Pass cancel context
	req = req.WithContext(ctx)
	// Create go-httpstat powered context and pass it to http.Request
	sCtx := httpstat.WithHTTPStat(req.Context(), &statz.Result)
	req = req.WithContext(sCtx)

	//Start a clock for overall duration compute
	statz.BegClock = time.Now()
	defer func() {
		statz.EndClock = time.Now()
	}()

	res, err := client.Do(req)
	if err != nil {
		statz.Err = fmt.Errorf("client.Do Fail: %w", err)
		return
	}
	defer res.Body.Close()
	err = statz.ReadHTTPHeader(&res.Header)
	if err != nil {
		statz.Err = fmt.Errorf("statz.ReadHeader Fail: %w", err)
		return
	}
	statz.Status = res.StatusCode
	//Check for error
	if res.StatusCode != http.StatusOK {
		//Error
		statz.Err = fmt.Errorf("res.StatusCode!=OK")
		return
	}
	written, err := io.Copy(ioutil.Discard, res.Body)
	if err != nil {
		statz.Err = fmt.Errorf("io.Copy Fail: %w", err)
		return
	}
	statz.Bytes += written
	return
}

func TestRequestStats(t *testing.T) {
	urls := []string{
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/init.mp4",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/1.m4s",
		"https://livesim.dashif.org/dash/vod/testpic_2s/V300/2.m4s",
	}
	for i, url := range urls {
		go func(url string, i int) {
			stats := helperGetURL(context.TODO(), url, "client"+strconv.Itoa(i), 10*time.Second)
			t.Logf("%v\n", stats.String())
		}(url, i)
	}
}
