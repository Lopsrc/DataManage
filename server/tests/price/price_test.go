package price_test

import (
	"testing"

	manager1 "server/protos/gen/go/manager"
	suite "server/server/tests/suit"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

const (
	emptyVar 		= 0
	invalidUserID 	= "user id is invalid"
	invalidPrice	= "price is invalid"
	errAlreadyExist = "entity already exists"
)

func TestCreate_HappyPath(t *testing.T){
	ctx, st := suite.NewSuit(t)

	userID := int64(gofakeit.Number(100,1000))
	create := manager1.CreatePriceRequest{
        UserId: userID,
        Price: 3000,
    }
	resp, err := st.Price.Create(ctx, &create)
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)
}

func TestUpdate_HappyPath(t *testing.T){
	ctx, st := suite.NewSuit(t)

	// Create.
	userID := int64(gofakeit.Number(100,1000))
	create := manager1.CreatePriceRequest{
        UserId: userID,
        Price: 3000,
    }
	resp, err := st.Price.Create(ctx, &create)
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)
	// Update.
    respUp, err := st.Price.Update(ctx, &manager1.UpdatePriceRequest{
        UserId: userID,
        Price: 1000,
    })
    assert.NoError(t, err)
    assert.True(t, respUp.IsUpdate)
}

func TestGet_HappyPath(t *testing.T){
	ctx, st := suite.NewSuit(t)

	userID := int64(gofakeit.Number(100,1000))
	create := manager1.CreatePriceRequest{
        UserId: userID,
        Price: 3000,
    }
	resp, err := st.Price.Create(ctx, &create)
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)
	
    respGet, err := st.Price.Get(ctx, &manager1.GetPriceRequest{
        UserId: 1,
    })
    assert.NoError(t, err)
    assert.NotEmpty(t, respGet.Price)
}

func TestCreateDublicateUserID(t *testing.T) {
	ctx, st := suite.NewSuit(t)

	userID := int64(gofakeit.Number(100,1000))
	resp, err := st.Price.Create(ctx, &manager1.CreatePriceRequest{
        UserId: userID,
        Price: 3000,
    })
	assert.NoError(t, err)
	assert.True(t, resp.IsCreate)

	_, err = st.Price.Create(ctx, &manager1.CreatePriceRequest{
        UserId: userID,
        Price: 3000,
    })
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errAlreadyExist)
}
func TestCreate_FailCases(t *testing.T){
	ctx, st := suite.NewSuit(t)

    cases := []struct {
        description string
        userID int64
        price int64
        msgError string
    }{
        {
            description: "invalid user id",
            userID: -1,
            price: 3000,
            msgError: invalidUserID,
        },
        {
            description: "invalid user id(empty)",
            userID: 0,
            price: 3000,
            msgError: invalidUserID,
        },
        {
            description: "invalid price",
            userID: 1,
            price: 0,
            msgError: invalidPrice,
        },
		{
			description: "invalid price",
            userID: 1,
            price: -1,
            msgError: invalidPrice,
		},
    }
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
            _, err := st.Price.Create(ctx, &manager1.CreatePriceRequest{
                UserId: c.userID,
                Price: c.price,
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
        userID int64
        price int64
        msgError string
    }{
        {
            description: "invalid user id",
            userID: -1,
            price: 3000,
            msgError: invalidUserID,
        },
        {
            description: "invalid user id(empty)",
            userID: 0,
            price: 3000,
            msgError: invalidUserID,
        },
        {
            description: "invalid price",
            userID: 1,
            price: 0,
            msgError: invalidPrice,
        },
		{
			description: "invalid price",
            userID: 1,
            price: -1,
            msgError: invalidPrice,
		},
    }
	for _, c := range cases {
		t.Run(c.description, func(t *testing.T){
            _, err := st.Price.Update(ctx, &manager1.UpdatePriceRequest{
                UserId: c.userID,
                Price: c.price,
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
        msgError string
    }{
        {
            description: "invalid user id",
            userID: -1,
            msgError: invalidUserID,
        },
        {
            description: "invalid user id(empty)",
            userID: 0,
            msgError: invalidUserID,
        },
    }
    for _, c := range cases {
        t.Run(c.description, func(t *testing.T){
            _, err := st.Price.Get(ctx, &manager1.GetPriceRequest{
                UserId: c.userID,
            })
            assert.Error(t, err)
            assert.Contains(t, err.Error(), c.msgError)
        })
    }
}