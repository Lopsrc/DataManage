package price

import (
	"context"
	"testing"

	"server/server/internal/grpc/manager/price/mocks"
	models "server/server/internal/models/price"

	"github.com/stretchr/testify/assert"
)


func TestCreate_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
	createPrice := models.CreatePrice{}
	service.On("Create", context.Background(), createPrice).Return(nil)
	err := service.Create(context.Background(), createPrice)
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
	updatePrice := models.UpdatePrice{}
	service.On("Update", context.Background(), updatePrice).Return(nil)
	err := service.Update(context.Background(), updatePrice)
	assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T){
	service := mocks.NewPriceService(t)
	getPrice := models.GetPrice{}
	service.On("Get", context.Background(), getPrice).Return(models.Prices{}, nil)
	_, err := service.Get(context.Background(), getPrice)
	assert.NoError(t, err)
}
