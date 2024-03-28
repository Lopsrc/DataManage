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

func HandleCreate(req *manager1.CreatePriceRequest) error {
	if req.GetUserId() == emptyVar{
		return fmt.Errorf(invalidUserID)
	}
	if req.GetPrice() == emptyVar{
        return fmt.Errorf(invalidPrice)
    }
	if req.GetPrice() < 0{
        return fmt.Errorf(invalidPrice)
    }
	return nil
}

func HandleUpdate(req *manager1.UpdatePriceRequest) error {
	if req.GetUserId() == emptyVar{
		return fmt.Errorf(invalidUserID)
	}
	if req.GetPrice() == emptyVar{
        return fmt.Errorf(invalidPrice)
    }
	if req.GetPrice() < 0{
        return fmt.Errorf(invalidPrice)
    }
    return nil
}

func HandleGet(req *manager1.GetPriceRequest) error {
	if req.GetUserId() == emptyVar{
		return fmt.Errorf(invalidUserID)
	}
    return nil
}