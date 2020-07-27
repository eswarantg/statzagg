//Package statzagg - aggregate statitics posted by other clients
package statzagg

import (
	"context"
	"io"
	"sync"
)

//StatzAgg - Aggregates Statz
type StatzAgg interface {
	PostHTTPClientStats(context.Context, *HTTPClientStatz) //PostHTTPClientStats - Post Http Tx
}

//NullStatzAgg - Ignores Statz
type NullStatzAgg struct {
}

//PostHTTPClientStats - Post Http Tx
func (n *NullStatzAgg) PostHTTPClientStats(ctx context.Context, statz *HTTPClientStatz) {
}

//LogStatzAgg - Write logs of metrics
type LogStatzAgg struct {
	wtr      io.Writer
	wtrMutex sync.Mutex
}

//NewLogStatzAgg - Create LogStatzAgg
func NewLogStatzAgg(wtr io.Writer) *LogStatzAgg {
	ret := &LogStatzAgg{}
	ret.wtr = wtr
	return ret
}

func (l *LogStatzAgg) toLogLine(statz *HTTPClientStatz) string {
	return statz.String()
}

//PostHTTPClientStats - post stats
func (l *LogStatzAgg) PostHTTPClientStats(ctx context.Context, statz *HTTPClientStatz) {
	logline := l.toLogLine(statz)
	l.wtrMutex.Lock()
	defer l.wtrMutex.Unlock()
	//if it takes too long ... handle ctx.cancel
	l.wtr.Write([]byte(logline))
}
