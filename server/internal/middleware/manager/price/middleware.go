package pricemiddleware

import (
	"fmt"
	manager1 "server/protos/gen/go/manager"
)

const (
	emptyVar 		= 0
	invalidUserID 	= "user id is invalid"
	invalidPrice	= "price is invalid"
)
// HandleCreate is a function that handles the creation of a new price.
func HandleCreate(req *manager1.CreatePriceRequest) error {
	if req.GetUserId() == emptyVar || req.GetUserId() < 0{
		return fmt.Errorf(invalidUserID)
	}
	if req.GetPrice() == emptyVar || req.GetPrice() < 0{
        return fmt.Errorf(invalidPrice)
    }
	return nil
}
// HandleUpdate is a function that handles the updating of an existing price.
func HandleUpdate(req *manager1.UpdatePriceRequest) error {
	if req.GetUserId() == emptyVar || req.GetUserId() < 0{
		return fmt.Errorf(invalidUserID)
	}
	if req.GetPrice() == emptyVar || req.GetPrice() < 0{
        return fmt.Errorf(invalidPrice)
    }
    return nil
}
// HandleDelete is a function that handles the deleting of an existing price.
func HandleGet(req *manager1.GetPriceRequest) error {
	if req.GetUserId() == emptyVar || req.GetUserId() < 0{
		return fmt.Errorf(invalidUserID)
	}
    return nil
}