package workmiddleware

import (
	"fmt"
	manager1 "server/protos/gen/go/manager"
	"time"
)

const (
	emptyVar 		= 0
	invalidUserID 	= "user id is invalid"
	invalidID 		= "id is invalid"
	invalidName 	= "name is invalid"
	invalidDate 	= "date is invalid"
	invalidTime 	= "time is invalid"
	invalidPenalty 	= "penalty is invalid"
)

func HandleCreate(req *manager1.CreateWorkRequest) error {

	if req.GetUserId() == emptyVar {
		return fmt.Errorf(invalidUserID)
	}
	if req.GetTime() == emptyVar {
		return fmt.Errorf(invalidTime)
	}
	if req.GetPenalty() < 0 {
		return fmt.Errorf(invalidPenalty)
	}
	if req.GetDate() == "" {
		return fmt.Errorf(invalidDate)
	}
	if req.GetName() == "" {
		return fmt.Errorf(invalidName)
	}
	if _, err := time.Parse("2006-01-02", req.GetDate()); err != nil {
		return fmt.Errorf(invalidDate)
    }
	return nil
}

func HandleUpdate(req *manager1.UpdateWorkRequest) error {

	if req.GetId() == emptyVar {
		return fmt.Errorf(invalidID)
	}
	if req.GetTime() == emptyVar {
		return fmt.Errorf(invalidTime)
	}
	if req.GetPenalty() < 0 {
		return fmt.Errorf(invalidPenalty)
	}
	if req.GetDate() == "" {
		return fmt.Errorf(invalidDate)
	}
	if req.GetName() == "" {
		return fmt.Errorf(invalidName)
	}
	if _, err := time.Parse("2006-01-02", req.GetDate()); err != nil {
		return fmt.Errorf(invalidDate)
    }
	return nil
}

func HandleGet(req *manager1.GetWorkRequest) error {

	if req.GetUserId() == emptyVar {
		return fmt.Errorf(invalidUserID)
	}
	if req.GetName() == "" {
		return fmt.Errorf(invalidName)
	}
	return nil
}

func HandleGetByDate(req *manager1.GetByDateWorkRequest) error {

	if req.GetUserId() == emptyVar {
		return fmt.Errorf(invalidUserID)
	}
	if req.GetName() == "" {
		return fmt.Errorf(invalidName)
	}
	if req.GetDate() == "" {
		return fmt.Errorf(invalidDate)
	}
	return nil
}

func HandleDelete(req *manager1.DeleteWorkRequest) error {

	if req.GetId() == emptyVar {
		return fmt.Errorf(invalidID)
	}
	return nil
}