package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nashkispace/ns-todolist/serverSide/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var taskCollection *mongo.Collection = OpenCollection(Client, "dolist")

func AddTask(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(task)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	task.ID = primitive.NewObjectID()
	result, insertErr := taskCollection.InsertOne(ctx, task)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "order item was not created"})
		fmt.Println(insertErr)
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
}

func GetTasks(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var tasks []bson.M
	cursor, err := taskCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	defer cancel()
	fmt.Println(tasks)
	c.JSON(http.StatusOK, tasks)

}

// func GetTasksByIngredient(c *gin.Context) {
// 	ingredient := c.Params.ByName("id")
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var tasks []bson.M
// 	cursor, err := taskCollection.Find(ctx, bson.M{"ingredients": ingredient})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}
// 	if err = cursor.All(ctx, &tasks); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cancel()
// 	fmt.Println(tasks)

// 	c.JSON(http.StatusOK, tasks)
// }

// func GetTaskById(c *gin.Context) {
// 	TaskID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(TaskID)

// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var task bson.M
// 	if err := taskCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&task); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cancel()
// 	fmt.Println(task)
// 	c.JSON(http.StatusOK, task)

// }

// func UpdateIngredient(c *gin.Context) {
// 	entryID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(entryID)
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 	type Ingredient struct {
// 		Ingredients *string `json:"ingredients"`
// 	}
// 	var ingredient Ingredient

// 	if err := c.BindJSON(&ingredient); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
// 		bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}},
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cancel()
// 	c.JSON(http.StatusOK, result.ModifiedCount)
// }

// func UpdateEntry(c *gin.Context) {
// 	entryID := c.Params.ByName("id")
// 	docID, _ := primitive.ObjectIDFromHex(entryID)
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var entry models.Entry

// 	if err := c.BindJSON(&entry); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}

// 	validationErr := validate.Struct(entry)
// 	if validationErr != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
// 		fmt.Println(validationErr)
// 		return
// 	}

// 	result, err := entryCollection.ReplaceOne(
// 		ctx,
// 		bson.M{"_id": docID},
// 		bson.M{
// 			"dish":        entry.Dish,
// 			"fat":         entry.Fat,
// 			"ingredients": entry.Ingredients,
// 			"calories":    entry.Calories,
// 		},
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		fmt.Println(err)
// 		return
// 	}
// 	defer cancel()
// 	c.JSON(http.StatusOK, result.ModifiedCount)

// }

func DeleteTask(c *gin.Context) {
	taskID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(taskID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := taskCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, result.DeletedCount)
}
