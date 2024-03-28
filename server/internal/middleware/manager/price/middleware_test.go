package pricemiddleware

import (
	manager1 "server/protos/gen/go/manager"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate_HappyPath(t *testing.T) {
	err := HandleCreate(&manager1.CreatePriceRequest{
		UserId: 1,
        Price: 10,
	})
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T) {
	err := HandleUpdate(&manager1.UpdatePriceRequest{
		UserId: 1,
        Price: 10,
	})
	assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T) {
	err := HandleGet(&manager1.GetPriceRequest{
		UserId: 1,
	})
	assert.NoError(t, err)
}

func TestCreate_FailCases(t *testing.T) {

	cases := []struct {
		name string
        id int64
		price int64
		msgError string
    }{
		{
			name: "invalid id",
            id: 0,
            price: 10,
            msgError: invalidUserID,
		},
		{
			name: "invalid price",
            id: 1,
            price: 0,
            msgError: invalidPrice,
		},
		{
			name: "invalid userId and price",
            id: 0,
            price: 0,
            msgError: invalidUserID,
		},
		{
			name: "invalid price",
            id: 1,
            price: -1,
            msgError: invalidPrice,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleCreate(&manager1.CreatePriceRequest{
                UserId: c.id,
                Price: c.price,
            })
            assert.EqualError(t, err, c.msgError)
        })
    }
}

func TestUpdate_FailCases(t *testing.T) {
	cases := []struct {
		name string
        id int64
		price int64
		msgError string
    }{
		{
			name: "invalid id",
            id: 0,
            price: 10,
            msgError: invalidUserID,
		},
		{
			name: "invalid price",
            id: 1,
            price: 0,
            msgError: invalidPrice,
		},
		{
			name: "invalid userId and price",
            id: 0,
            price: 0,
            msgError: invalidUserID,
		},
		{
			name: "invalid price",
            id: 1,
            price: -1,
            msgError: invalidPrice,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleUpdate(&manager1.UpdatePriceRequest{
                UserId: c.id,
                Price: c.price,
            })
            assert.EqualError(t, err, c.msgError)
        })
    }
}

func TestGet_FailCases(t *testing.T) {
	cases := []struct{
		name string
        userId int64
        msgError string
    }{
        {
            name: "invalid id",
            userId: 0,
            msgError: invalidUserID,
        },
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
            err := HandleGet(&manager1.GetPriceRequest{
                UserId: c.userId,
            })
            assert.EqualError(t, err, c.msgError)
        })
	}
}