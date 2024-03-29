package workservice

import (
	"context"

	models "server/server/internal/models/work"
	"server/server/internal/services/work/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate_HappyPath(t *testing.T){
	rep := mocks.NewWorkRepository(t)
	rep.On("Create", context.Background(),  &models.CreateWork{}).Return(nil).Once()
	err := rep.Create(context.Background(), &models.CreateWork{})
	assert.NoError(t, err)
}

func TestUpdate_HappyPath(t *testing.T){
	rep := mocks.NewWorkRepository(t)
	rep.On("Update", context.Background(),  &models.UpdateWork{}).Return(nil).Once()
	err := rep.Update(context.Background(), &models.UpdateWork{})
	assert.NoError(t, err)
}

func TestGet_HappyPath(t *testing.T){
	rep := mocks.NewWorkRepository(t)
    rep.On("GetAll", context.Background(),  &models.GetAllWork{}).Return([]models.Work{}, nil).Once()
    _, err := rep.GetAll(context.Background(), &models.GetAllWork{})
    assert.NoError(t, err)
}

func TestGetByDate_HappyPath(t *testing.T){
	rep := mocks.NewWorkRepository(t)
    rep.On("GetAll", context.Background(),  &models.GetAllWork{}).Return([]models.Work{}, nil).Once()
    w, err := rep.GetAll(context.Background(), &models.GetAllWork{})
    assert.NoError(t, err)
	w = SortByMonth(&w, "april")
}

func TestDelete_HappyPath(t *testing.T){
	rep := mocks.NewWorkRepository(t)
    rep.On("Delete", context.Background(),  &models.DeleteWork{}).Return(nil).Once()
    err := rep.Delete(context.Background(), &models.DeleteWork{})
    assert.NoError(t, err)
}