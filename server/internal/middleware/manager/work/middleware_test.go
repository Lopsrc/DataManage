package workmiddleware

import (
	manager1 "server/protos/gen/go/manager"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate_HappyPath(t *testing.T) {
	err := HandleCreate(&manager1.CreateWorkRequest{
		UserId: 1,
		Name: "workspace",
        Date: "2013-01-01",
		Time: 8,
		Penalty: 0,
    })
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T) {
	err := HandleUpdate(&manager1.UpdateWorkRequest{
		Name: "workspace",
        Date: "2013-01-01",
		Time: 8,
		Penalty: 0,
		Id: 1,
	})
	assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T) {
	err := HandleGet(&manager1.GetWorkRequest{
		UserId: 1,
		Name: "workspace",
	})
	assert.NoError(t, err)
}

func TestGetByDate_HappyPath(t *testing.T) {
	err := HandleGetByDate(&manager1.GetByDateWorkRequest{
		UserId: 1,
		Name: "workspace",
		Date: "April",
	})
	assert.NoError(t, err)
}

func TestDelete_HappyPath(t *testing.T) {
	err := HandleDelete(&manager1.DeleteWorkRequest{
        Id: 1,
    })
    assert.NoError(t, err)
}

func TestCreate_FailCases(t *testing.T) {

	cases := []struct {
		description string
		userId int64
		name string
		date string
		time int32
		penalty int64
		msgError string
    }{
		{
			description: "invalid user id",
            userId: 0,
			name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidUserID,
		},
		{
			description: "invalid name",
            userId: 1,
			name: "",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidName,
		},
		{
			description: "invalid date",
            userId: 1,
			name: "workspace",
            date: "",
            time: 8,
            penalty: 0,
            msgError: invalidDate,
		},
		{
			description: "invalid time",
            userId: 1,
			name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: 0,
            msgError: invalidTime,
		},
		{
			description: "invalid penalty",
            userId: 1,
			name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: -1,
            msgError: invalidPenalty,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleCreate(&manager1.CreateWorkRequest{
				UserId: c.userId,
                Name: c.name,
                Date: c.date,
                Time: c.time,
                Penalty: c.penalty,
            })
            assert.EqualError(t, err, c.msgError)
        })
    }
}

func TestUpdate_FailCases(t *testing.T) {
	cases := []struct {
		description string
		name string
		date string
		time int32
		penalty int64
		id int64
		msgError string
    }{
		{
			description: "invalid id",
			name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
			id: 0,
            msgError: invalidID,
		},
		{
			description: "invalid name",
			name: "",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
			id: 1,
            msgError: invalidName,
		},
		{
			description: "invalid date",
			name: "workspace",
            date: "",
            time: 8,
            penalty: 0,
			id: 1,
            msgError: invalidDate,
		},
		{
			description: "invalid time",
			name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: 0,
			id: 1,
            msgError: invalidTime,
		},
		{
			description: "invalid penalty",
			name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: -1,
			id: 1,
            msgError: invalidPenalty,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleUpdate(&manager1.UpdateWorkRequest{
				Id: c.id,
                Name: c.name,
                Date: c.date,
                Time: c.time,
                Penalty: c.penalty,
				
            })
            assert.EqualError(t, err, c.msgError)
        })
    }
}

func TestGet_FailCases(t *testing.T) {
	cases := []struct{
		description string
		name string
        userId int64
        msgError string
    }{
        {
			description: "invalid id",
            userId: 0,
			name: "workspace",
            msgError: invalidUserID,
        },
		{
			description: "invalid name",
            userId: 1,
			name: "",
            msgError: invalidName,
        },
		{
			description: "invalid both",
            userId: 0,
			name: "",
            msgError: invalidUserID,
        },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleGet(&manager1.GetWorkRequest{
                UserId: c.userId,
				Name: c.name,
            })
            assert.EqualError(t, err, c.msgError)
        })
	}
}

func TestGetByDate_FailCases(t *testing.T) {
	cases := []struct{
		description string
		name string
		date string
        userId int64
        msgError string
    }{
        {
			description: "invalid id",
            userId: 0,
			date: "2002-02-02",
			name: "workspace",
            msgError: invalidUserID,
        },
		{
			description: "invalid date",
            userId: 1,
			date: "",
			name: "workspace",
            msgError: invalidDate,
        },
		{
			description: "invalid name",
            userId: 1,
			date: "2002-02-02",
			name: "",
            msgError: invalidName,
        },
		{
			description: "invalid both",
            userId: 0,
			date: "2002-02-02",
			name: "",
            msgError: invalidUserID,
        },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleGet(&manager1.GetWorkRequest{
                UserId: c.userId,
				Name: c.name,
            })
            assert.EqualError(t, err, c.msgError)
        })
	}
}