package pricemiddleware

import manager1 "server/protos/gen/go/manager"

const (
	emptyVar 		= 0
	invalidUserID 	= "user id is invalid"
	invalidPrice	= "price is invalid"
)

func HandleCreate(req *manager1.CreatePriceRequest) (string, error) {
	if req.GetUserId() == emptyVar{
		return invalidUserID, nil
	}
	if req.GetPrice() == emptyVar{
        return invalidPrice, nil
    }
	return "", nil
}

func HandleUpdate(req *manager1.UpdatePriceRequest) (string, error) {
	if req.GetUserId() == emptyVar{
		return invalidUserID, nil
	}
	if req.GetPrice() == emptyVar{
        return invalidPrice, nil
    }
    return "", nil
}

func HandleGet(req *manager1.GetPriceRequest) (string, error) {
	if req.GetUserId() == emptyVar{
		return invalidUserID, nil
	}
    return "", nil
}