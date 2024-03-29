package work_test

import (
	"testing"

	manager1 "server/protos/gen/go/manager"
	suite "server/server/tests/suit"

	"github.com/stretchr/testify/assert"
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

func TestCreate_HappyPath(t *testing.T) {
	ctx, st := suite.NewSuit(t)

	create := manager1.CreateWorkRequest{
        Name: "test",
        Date: "2013-09-09",
        Time: 8,
		Penalty: 100,
		UserId:  1,
	}
	resp, err := st.Work.Create(ctx, &create)
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)
}

func TestGet_HappyPath(t *testing.T) {
	name := "testGet"
	ctx, st := suite.NewSuit(t)
	// Create a new entity.
    create := manager1.CreateWorkRequest{
		Name: name,
        Date: "2013-09-09",
        Time: 8,
        Penalty: 100,
        UserId:  1,
	}
	resp, err := st.Work.Create(ctx, &create)
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)
	// Get the entity.
	getWork := manager1.GetWorkRequest{
		Name: name,
        UserId: 1,
    }
	respGet, err := st.Work.GetAll(ctx, &getWork)
	assert.NoError(t, err)
	assert.NotEmpty(t, respGet.ListWorks)
}

func TestGetByDate_HappyPath(t *testing.T) {
    name := "testGetByDate"
    ctx, st := suite.NewSuit(t)
    // Create a new entity.
    create := manager1.CreateWorkRequest{
        Name: name,
        Date: "2013-04-04",
        Time: 8,
        Penalty: 100,
        UserId:  1,
    }
    resp, err := st.Work.Create(ctx, &create)
    assert.NoError(t, err)
    assert.True(t, resp.IsCreate)
    // Get the entity.
    getWork := manager1.GetByDateWorkRequest{
        Name: name,
        UserId: 1,
        Date: "April",
    }
	respGet, err := st.Work.GetAllByDate(ctx, &getWork)
	assert.NoError(t, err)
	assert.NotEmpty(t, respGet.ListWorks)
}

func TestUpdate_HappyPath(t *testing.T) {
    name := "testUpdate"
    ctx, st := suite.NewSuit(t)
    // Create a new entity.
    create := manager1.CreateWorkRequest{
        Name: name,
        Date: "2013-09-09",
        Time: 8,
        Penalty: 100,
        UserId:  1,
    }
    resp, err := st.Work.Create(ctx, &create)
    assert.NoError(t, err)
    assert.True(t, resp.IsCreate)
	// Get the entity.
	getWork := manager1.GetWorkRequest{
		Name: name,
        UserId: 1,
    }
	respGet, err := st.Work.GetAll(ctx, &getWork)
	assert.NoError(t, err)
	assert.NotEmpty(t, respGet.ListWorks)
    // Update the entity.
	update := manager1.UpdateWorkRequest{
		Id: respGet.ListWorks[0].Id,
        Name: name,
        Date: "2013-09-09",
        Time: 8,
        Penalty: 100,
    }
    respUpdate, err := st.Work.Update(ctx, &update)
    assert.NoError(t, err)
    assert.True(t, respUpdate.IsUpdate)
}

func TestDelete_HappyPath(t *testing.T) {
	var userID int64  
    name := "testDelete"
    userID = 1
    ctx, st := suite.NewSuit(t)
    // Create a new entity.
    create := manager1.CreateWorkRequest{
        Name: name,
        Date: "2013-09-09",
        Time: 8,
        Penalty: 100,
        UserId:  1,
    }
    resp, err := st.Work.Create(ctx, &create)
    assert.NoError(t, err)
    assert.True(t, resp.IsCreate)
	// Get the entity.
	getWork := manager1.GetWorkRequest{
		Name: name,
        UserId: userID,
    }
	respGet, err := st.Work.GetAll(ctx, &getWork)
	assert.NoError(t, err)
	assert.NotEmpty(t, respGet.ListWorks)
	// Delete the entity.
	deleteWork := manager1.DeleteWorkRequest{
        Id: respGet.ListWorks[0].Id,
    }
    respDelete, err := st.Work.Delete(ctx, &deleteWork)
    assert.NoError(t, err)
    assert.True(t, respDelete.IsDel)
}

func TestCreate_FailCases(t *testing.T){
	ctx, st := suite.NewSuit(t)
	
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
            description: "invalid user id(empty)",
            userId: 0,
			name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidUserID,
        },
		{
            description: "invalid user id",
            userId: -1,
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
			description: "invalid date(empty)",
            userId: 1,
            name: "workspace",
            date: "",
            time: 8,
            penalty: 0,
            msgError: invalidDate,
		},
		{
			description: "invalid date",
            userId: 1,
            name: "workspace",
            date: "<INVALID-DATE>",
            time: 8,
            penalty: 0,
            msgError: invalidDate,
		},
		{
			description: "invalid time(empty)",
            userId: 1,
            name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: 0,
            msgError: invalidTime,
		},
		{
            description: "invalid time",
            userId: 1,
			name: "workspace",
            date: "2013-01-01",
            time: -1,
            penalty: 0,
            msgError: invalidTime,
        },
		{
			description: "invalid penalty",
            userId: 1,
            name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: -1,
            msgError: invalidPenalty,
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
			
            _, err := st.Work.Create(ctx, &manager1.CreateWorkRequest{
                Name: c.name,
                Date: c.date,
                Time: c.time,
                Penalty: c.penalty,
                UserId: c.userId,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
		})
	}
}

func TestUpdate_FailCases(t *testing.T){
	ctx, st := suite.NewSuit(t)

    cases := []struct {
		description string
		id int64
		name string
		date string
		time int32
		penalty int64
		msgError string
	}{
		{
			description: "invalid id",
            id: -1,
            name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidID,
		},
		{
			description: "invalid id(empty)",
            id: 0,
            name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidID,
		},
		{
			description: "invalid name",
            id: 1,
            name: "",
            date: "2013-01-01",
            time: 8,
            penalty: 0,
            msgError: invalidName,
		},
		{
			description: "invalid date(empty)",
            id: 1,
            name: "workspace",
            date: "",
            time: 8,
            penalty: 0,
            msgError: invalidDate,
		},
		{
			description: "invalid date",
            id: 1,
            name: "workspace",
            date: "<INVALID-DATE>",
            time: 8,
            penalty: 0,
            msgError: invalidDate,
		},
		{
			description: "invalid time(empty)",
            id: 1,
            name: "workspace",
            date: "2013-01-01",
            time: 0,
            penalty: 0,
            msgError: invalidTime,
		},
		{
			description: "invalid time",
            id: 1,
            name: "workspace",
            date: "2013-01-01",
            time: -1,
            penalty: 0,
            msgError: invalidTime,
		},
		{
			description: "invalid penalty",
            id: 1,
            name: "workspace",
            date: "2013-01-01",
            time: 8,
            penalty: -1,
            msgError: invalidPenalty,
		},
	}
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
            _, err := st.Work.Update(ctx, &manager1.UpdateWorkRequest{
                Id: c.id,
                Name: c.name,
                Date: c.date,
                Time: c.time,
                Penalty: c.penalty,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
        })
	}
}

func TestGet_FailCases(t *testing.T){
	ctx, st := suite.NewSuit(t)

    cases := []struct {
        description string
        userID int64
		name string
        msgError string
    }{
        {
            description: "invalid id",
            userID: -1,
			name: "workspace",
            msgError: invalidUserID,
        },
        {
            description: "invalid id(empty)",
            userID: 0,
			name: "workspace",
            msgError: invalidUserID,
        },
		{
			description: "invalid name",
            userID: 1,
            name: "",
            msgError: invalidName,
		},
    }
    for _, c := range cases {
        t.Run(c.description, func(t *testing.T){
            _, err := st.Work.GetAll(ctx, &manager1.GetWorkRequest{
                UserId: c.userID,
				Name: c.name,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
        })
    }
}

func TestGetByDate_failCases(t *testing.T){
	ctx, st := suite.NewSuit(t)

    cases := []struct {
        description string
        userID int64
		name string
        date string
        msgError string
    }{
        {
            description: "invalid id",
            userID: -1,
			name: "workspace",
            date: "2013-01-01",
            msgError: invalidUserID,
        },
        {
            description: "invalid id(empty)",
            userID: 0,
			name: "workspace",
            date: "2013-01-01",
            msgError: invalidUserID,
        },
        {
            description: "invalid date(empty)",
            userID: 1,
			name: "workspace",
            date: "",
            msgError: invalidDate,
        },
		{
			description: "invalid date",
            userID: 1,
            name: "workspace",
            date: "<INVALID-DATE>",
            msgError: invalidDate,
		},
		{
			description: "invalid name",
            userID: 1,
            name: "",
            date: "2013-01-01",
            msgError: invalidName,
		},
    }
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
            _, err := st.Work.GetAllByDate(ctx, &manager1.GetByDateWorkRequest{
                UserId: c.userID,
				Name: c.name,
                Date: c.date,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
        })
	}
}

func TestDelete_FailCases(t *testing.T){
	ctx, st := suite.NewSuit(t)

    cases := []struct {
        description string
        id int64
        msgError string
    }{
        {
            description: "invalid id",
            id: -1,           
            msgError: invalidID,
        },
        {
            description: "invalid id(empty)",
            id: 0,
            msgError: invalidID,
        },
    }
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
            _, err := st.Work.Delete(ctx, &manager1.DeleteWorkRequest{
                Id: c.id,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
        })
	}
}