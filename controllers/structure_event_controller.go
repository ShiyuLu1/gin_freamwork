package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_mvc/conf/database"
	"go_mvc/models"
	"net/http"
	"time"
)

var (
	client  = database.InitMongo()
	mongoDB *mongo.Database
	coll    *mongo.Collection
	err     error
	cursor  *mongo.Cursor
)

func GetEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 选择数据库 monitor_platform
		mongoDB = client.Database("monitor_platform")
		// 选择表 107886707_structured_events
		coll = mongoDB.Collection("107886707_structured_events")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// 过滤条件
		filter := bson.M{}
		cursor, err = coll.Find(ctx, filter, options.Find().SetLimit(10))
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				models.BaseResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
					Data:    map[string]interface{}{"list": []string{}},
				})
			return
		}

		defer cursor.Close(ctx)
		var structureEvents []models.StructureEvent
		for cursor.Next(ctx) {
			var singleEvent models.StructureEvent
			if err = cursor.Decode(&singleEvent); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					models.BaseResponse{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
						Data:    map[string]interface{}{"list": []string{}}})
			}

			structureEvents = append(structureEvents, singleEvent)
		}

		c.JSON(
			http.StatusOK,
			models.BaseResponse{
				Code:    0,
				Message: "success",
				Data:    map[string]interface{}{"list": structureEvents}})
	}
}
