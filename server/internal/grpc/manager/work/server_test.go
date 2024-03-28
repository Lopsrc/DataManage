package work

import (
	"context"
	"testing"

	models "server/server/internal/models/work"

	"server/server/internal/grpc/manager/work/mocks"

	"github.com/stretchr/testify/assert"
)


func TestCreate_HappyPath(t *testing.T){
	service := mocks.NewWorkService(t)
	createWork := models.CreateWork{}
	service.On("Create", context.Background(), createWork).Return(nil)
	err := service.Create(context.Background(), createWork)
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T){
	service := mocks.NewWorkService(t)
	updateWork := models.UpdateWork{}
	service.On("Update", context.Background(), updateWork).Return(nil)
	err := service.Update(context.Background(), updateWork)
	assert.NoError(t, err)
}
func TestGet_HappyPath(t *testing.T){
	service := mocks.NewWorkService(t)
	getWork := models.GetAllWork{}
	service.On("Get", context.Background(), getWork).Return([]models.Work{}, nil)
	_, err := service.Get(context.Background(), getWork)
	assert.NoError(t, err)
}
func TestGetByDate_HappyPath(t *testing.T){
	service := mocks.NewWorkService(t)
	getWork := models.GetAllWorkByDate{}
	service.On("GetByDate", context.Background(), getWork).Return([]models.Work{}, nil)
	_, err := service.GetByDate(context.Background(), getWork)
	assert.NoError(t, err)
}
func TestDelete_HappyPath(t *testing.T){
	service := mocks.NewWorkService(t)
	deleteWork := models.DeleteWork{}
	service.On("Delete", context.Background(), deleteWork).Return(nil)
	err := service.Delete(context.Background(), deleteWork)
	assert.NoError(t, err)
}
