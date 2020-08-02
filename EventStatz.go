//Package statzagg - aggregate statitics posted by other clients
package statzagg

import (
	"fmt"
	"time"
)

//EventStats - Events stats
type EventStats struct {
	EventClock time.Time     //Tx Begin
	ID         string        //Client ID
	Name       string        //Name of event
	Err        error         //Error
	Values     []interface{} //Value
}

//String - output of user readable
func (h *EventStats) String() string {
	var ret string
	err := "-"
	if h.Err != nil {
		err = h.Err.Error()
	}
	//EndClock: ID, error, value
	ret = fmt.Sprintf("%v: %v %v %v ",
		h.EventClock.UTC().Format("2006-01-02T15:04:05.000Z07:00"), h.ID, h.Name, err,
	)
	for _, val := range h.Values {
		ret += fmt.Sprintf("%v ", val)
	}
	return ret
}
