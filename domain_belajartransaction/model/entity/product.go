package entity

import (
	"gogen-transaction/domain_belajartransaction/model/vo"
	"time"
)

type Product struct {
	ID      vo.ProductID `bson:"_id" json:"id"`
	Created time.Time    `bson:"created" json:"created"`
	Updated time.Time    `bson:"updated" json:"updated"`

	// edit or add new necessary field here ...
}

type ProductCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`

	// edit or add new necessary field for create request here ...

}

func (r ProductCreateRequest) Validate() error {

	// validate the create request here ...

	return nil
}

func NewProduct(req ProductCreateRequest) (*Product, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id, err := vo.NewProductID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Product
	obj.ID = id
	obj.Created = req.Now
	obj.Updated = req.Now

	// another field input here ...

	return &obj, nil
}

type ProductUpdateRequest struct {
	Now time.Time `json:"-"`

	// add new necessary field to update request here ...

}

func (r ProductUpdateRequest) Validate() error {

	// validate the update request here ...

	return nil
}

func (r *Product) Update(req ProductUpdateRequest) error {

	err := req.Validate()
	if err != nil {
		return err
	}

	r.Updated = req.Now

	// update field here ...

	return nil
}
