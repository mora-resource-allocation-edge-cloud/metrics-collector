package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"metrics-collector/models"
)

type VideoReproductionHandler interface {
	StoreMetric(metric *models.VideoReproduction) (*models.VideoReproduction, error)
	FindAllBetweenDateInterval(startTime uint64, endTime uint64) (*[]models.VideoReproduction, error)
	FindVideoBetweenDateInterval(startTime uint64, endTime uint64, videoUrl string)
}

type videoReproductionMongoHandler struct {
	BaseMongoHandler
}

func (h *videoReproductionMongoHandler) StoreMetric(d *models.VideoReproduction) (*models.VideoReproduction, error) {
	// force id to be its zero-value
	d.ID = primitive.ObjectID{}
	result, err := h.collection.InsertOne(context.TODO(), d)
	if err != nil {
		panic(err)
	}
	d.ID = result.InsertedID.(primitive.ObjectID)
	return d, err
}

func (h *videoReproductionMongoHandler) FindAllBetweenDateInterval(startTime uint64, endTime uint64) (*[]models.VideoReproduction, error) {
	panic("implement me")
}

func (h *videoReproductionMongoHandler) FindVideoBetweenDateInterval(startTime uint64, endTime uint64, videoUrl string) {
	panic("implement me")
}
