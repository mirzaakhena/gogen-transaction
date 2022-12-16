package vo

import (
	"fmt"
	"time"
)

type ProductID string

func NewProductID(randomStringID string, now time.Time) (ProductID, error) {
	var obj = ProductID(fmt.Sprintf("PRO-%s-%s", now.Format("060102"), randomStringID))

	// you may change it as necessary ...

	return obj, nil
}

func (r ProductID) String() string {
	return string(r)
}
