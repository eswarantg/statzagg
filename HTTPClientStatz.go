//Package statzagg - aggregate statitics posted by other clients
package statzagg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

//HTTPClientStatz - client stats
type HTTPClientStatz struct {
	httpstat.Result           //HTTP statistic
	BegClock        time.Time //Tx Begin
	EndClock        time.Time //Tx Completed
	ID              string    //Client ID
	URL             string    //URL
	Status          int       //HTTP status code
	Bytes           int64     //HTTP response bytes
	Err             error     //Client error
	CdnHeaders      string    //cdn record for tracing
}

//String - output of user readable
func (h *HTTPClientStatz) String() string {
	var ret string
	err := "-"
	if h.Err != nil {
		err = h.Err.Error()
	}
	//EndClock: ID, URL, Status, Bytes, Duration, Err
	//  DNSLookup, Connect, TLSHandShake, ServerProcessing, StartTransfer
	ret = fmt.Sprintf("%v: %v %v %v %v %v %v %v %v %v %v %v %v",
		h.EndClock.UTC().Format("2006-01-02T15:04:05.000Z07:00"), h.ID, h.URL, h.Status, h.Bytes, h.EndClock.Sub(h.BegClock).Milliseconds(), err,
		h.DNSLookup.Milliseconds(), h.Connect.Milliseconds(), h.TLSHandshake.Milliseconds(), h.ServerProcessing.Milliseconds(), h.StartTransfer.Milliseconds(), h.CdnHeaders,
	)
	return ret
}

//ReadHTTPHeader - reads required details from header
func (h *HTTPClientStatz) ReadHTTPHeader(hdr *http.Header) error {
	//TBD - Fill CdnHeaders
	return nil
}
