package vo

import (
	"fmt"
	"time"
)

type OrderID string

func NewOrderID(randomStringID string, now time.Time) (OrderID, error) {
	var obj = OrderID(fmt.Sprintf("ORD-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r OrderID) String() string {
	return string(r)
}
