package entity

import (
	"gogen-transaction/domain_belajartransaction/model/errorenum"
	"gogen-transaction/domain_belajartransaction/model/vo"
	"time"
)

type Order struct {
	ID      vo.OrderID `bson:"_id" json:"id"`
	Created time.Time  `bson:"created" json:"created"`
	Updated time.Time  `bson:"updated" json:"updated"`
	User    string     `bson:"user" json:"user"`
}

type OrderCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	User         string    `json:"user"`
}

func (r OrderCreateRequest) Validate() error {

	if r.User == "" {
		return errorenum.UserMustNotEmpty
	}

	return nil
}

func NewOrder(req OrderCreateRequest) (*Order, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewOrderID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Order
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now
	obj.User = req.User

	// another field input here ...

	return &obj, nil
}

type OrderUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r OrderUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Order) Update(req OrderUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
