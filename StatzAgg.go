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
	PostEventStats(context.Context, *EventStats)           //PostEventStats - Post Event Tx
}

//NullStatzAgg - Ignores Statz
type NullStatzAgg struct {
}

//PostHTTPClientStats - Post Http Tx
func (n *NullStatzAgg) PostHTTPClientStats(ctx context.Context, statz *HTTPClientStatz) {
}

//PostEventStats - Post Event Tx
func (n *NullStatzAgg) PostEventStats(ctx context.Context, statz *EventStats) {
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

func (l *LogStatzAgg) toLogLineHTTPStatz(statz *HTTPClientStatz) string {
	return statz.String() + "\n"
}

func (l *LogStatzAgg) toLogLineEventStatz(statz *EventStats) string {
	return statz.String() + "\n"
}

//PostHTTPClientStats - post stats
func (l *LogStatzAgg) PostHTTPClientStats(ctx context.Context, statz *HTTPClientStatz) {
	logline := l.toLogLineHTTPStatz(statz)
	l.wtrMutex.Lock()
	defer l.wtrMutex.Unlock()
	//if it takes too long ... handle ctx.cancel
	l.wtr.Write([]byte(logline))
}

//PostEventStats - Post Event Tx
func (l *LogStatzAgg) PostEventStats(ctx context.Context, statz *EventStats) {
	logline := l.toLogLineEventStatz(statz)
	l.wtrMutex.Lock()
	defer l.wtrMutex.Unlock()
	//if it takes too long ... handle ctx.cancel
	l.wtr.Write([]byte(logline))
}
