package priceservice

import (
	"context"
	"testing"

	"server/server/internal/grpc/manager/price/mocks"
	price "server/server/internal/models/price"

	"github.com/stretchr/testify/assert"
)

func TestCreate_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
    createPrice := price.CreatePrice{}
    service.On("Create", context.Background(), createPrice).Return(nil)
    err := service.Create(context.Background(), createPrice)
    assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
    updatePrice := price.UpdatePrice{}
    service.On("Update", context.Background(), updatePrice).Return(nil)
    err := service.Update(context.Background(), updatePrice)
    assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
    getPrice := price.GetPrice{}
    service.On("Get", context.Background(), getPrice).Return(price.Prices{}, nil)
    _, err := service.Get(context.Background(), getPrice)
    assert.NoError(t, err)
}
