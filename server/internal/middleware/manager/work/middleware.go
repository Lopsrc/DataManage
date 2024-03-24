package workmiddleware

import manager1 "server/protos/gen/go/manager"

const (
	emptyVar 		= 0
	invalidUserID 	= "user id is invalid"
	invalidID 		= "id is invalid"
	invalidName 	= "name is invalid"
	invalidDate 	= "date is invalid"
	invalidPrice 	= "price is invalid"
	invalidTime 	= "time is invalid"
	invalidPenalty 	= "penalty is invalid"
)

func HandleCreate(req *manager1.CreateWorkRequest)(string, error) {

	if req.GetUserId() == emptyVar {
		return invalidUserID, nil
	}
	if req.GetPrice() == emptyVar {
		return invalidPrice, nil
	}
	if req.GetTime() == emptyVar {
		return invalidTime, nil
	}
	if req.GetPenalty() == emptyVar {
		return invalidPenalty, nil
	}
	if req.GetDate() == "" {
		return invalidDate, nil
	}
	if req.GetName() == "" {
		return invalidName, nil
	}
	return "", nil
}

func HandleUpdate(req *manager1.UpdateWorkRequest)(string, error) {

	if req.GetId() == emptyVar {
		return invalidID, nil
	}
	if req.GetPrice() == emptyVar {
		return invalidPrice, nil
	}
	if req.GetTime() == emptyVar {
		return invalidTime, nil
	}
	if req.GetPenalty() == emptyVar {
		return invalidPenalty, nil
	}
	if req.GetDate() == "" {
		return invalidDate, nil
	}
	if req.GetName() == "" {
		return invalidName, nil
	}
	return "", nil
}

func HandleGet(req *manager1.GetWorkRequest)(string, error) {

	if req.GetUserId() == emptyVar {
		return invalidUserID, nil
	}
	if req.GetName() == "" {
		return invalidName, nil
	}
	return "", nil
}

func HandleGetByDate(req *manager1.GetByDateWorkRequest)(string, error) {

	if req.GetUserId() == emptyVar {
		return invalidUserID, nil
	}
	if req.GetName() == "" {
		return invalidName, nil
	}
	if req.GetDate() == "" {
		return invalidDate, nil
	}
	return "", nil
}

func HandleDelete(req *manager1.DeleteWorkRequest)(string, error) {

	if req.GetId() == emptyVar {
		return invalidID, nil
	}
	return "", nil
}