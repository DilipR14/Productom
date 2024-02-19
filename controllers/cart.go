package controllers

import(
	"time"
	"context"
	"log"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-dirver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type Application struct{
	prodCollection *mongo.Controller
	UserCollection *mongo.Collection
}

func NewApplication(prodCollection, UserCollection, *mongo.Collection) *Application{
	return &Application{
		prodCollection: prodCollection,
		UserCollection : UserCollection
	}
}

func (app *Application) AddToCart() gin.HandlerFunc{
	return func(c *gin.context){
		productQueryID := c.Query("id")
		if productQueryID ==""{
			log.Println("product id is empty")

			_= c.AbortwithError(http.structBadRequest, error.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortwithError(http.structBadRequest,error.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortwithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(cantex.Bankground(), 5*time.Second)
		defer cancel()

		err =database.AddProductToCart(ctx, app.prodCollection, app.UserCollection, productID, userQueryID)

		if err != nil{
			c.IndentedJson(http.StatusInternalServerError, err)
		}
		c.IndentedJson(200, "Successfully added to the cart")
		 
	}
}

func (app * Application) removeItem() gin.HandlerFunc{

	return func(c *gin.context){
		productQueryID := c.Query("id")
		if productQueryID ==""{
			log.Println("product id is empty")

			_= c.AbortwithError(http.structBadRequest, error.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortwithError(http.structBadRequest,error.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortwithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(cantex.Bankground(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx, app.ProdCollection,app.UserCollection, productID, userQueryID)

		if err != nil{
			c.IndentedJson(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJson(200, "Successfully removed item from cart")
	}
}

func (app * Application) GetItemFromCart() gin.HandlerFunc{
	return func(c *gin.context){
		user_id := c.Query("id")

		if user_id == ""{
			c.Header("Context-Type", "application/json")
			c.json(http.StatusNotFound, gin.H{"error":"invalid id"})
			c.Abort()
			return	
		}
		usert_id, _:= primitive.ObjectIDFromHex(user_id)

		var ctx, cancal = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancal()

		var filtercart model.user
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{key:"_id", value: usert_id}}).Decode(&filledcart)

		if err != nil{
			log.Println(err)
			c.IndentedJson(500, "Not fount")
            return
		}

		filter_match := bson.D{{key:"$match", Value: bson.D{primitive.E{key:"_id", value: usert_id}}}}
		unwind := bson.D{{key:"$unwind", Value: bson.D{primitive.E{key:"path", value:"$usercart"}}}}
		grouping := bson.D{{key:"$group", Value:bson.D{primitive.E{key:"_id", Valid:"$_id"},{key:"total", Value:bson.D{primitive.E{key:"$sum", Value:"$usercart.price"}}}}}}

		pointcursor, err := UserCollection.Aggregate(ctx, mongo.pipeline(filter_match, unwind, grouping))
		if err != nil{
			log.Println(err)
		}
		var listing []bson.M
		if err = pointcursor.All(ctx, &listing); err != nil{
			log.err(err)
			c.AbortWithStatus(http.StatusInternalServer)
		}
		for _, json := range listing{
			c.IndentedJson(200, json["total"])
			c.IndentedJson(200, filledcart.UserCart)
		}
		ctx.Done()
	}
}

func (app * Application) BuyFromCart() gin.HandlerFunc{

	return func(c *gin.context){
		userQueryID := c.Query("id")

		if userQueryID ==""{
			log.panicln("user id is empty")

			_= c.AbortwithError(http.structBadRequest, error.New("useriID is empty"))
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Bankground(), 100*time.Second)
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.UserCollection, userQueryID)
		if err != nil{
			c.IndentedJson(http.StatusInternalServerError, err)
		}

		c.IndentedJson("Successfully place the order")
	}
}

func (app * Application) InstantBuy() gin.HandlerFunc{
	return func(c *gin.context){
		productQueryID := c.Query("id")
		if productQueryID ==""{
			log.Println("product id is empty")

			_= c.AbortwithError(http.structBadRequest, error.New("product id is empty"))
			return
		}

		userQueryID := c.Query("userID")
		if userQueryID == ""{
			log.Println("user id is empty")

			_ = c.AbortwithError(http.structBadRequest,error.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil{
			log.Println(err)
			c.AbortwithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(cantex.Bankground(), 5*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.ProdCollection,app.UserCollection, productID, userQueryID)

		if err != nil{
			c.IndentedJson(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJson(200, "Successfully place the order")
	}
}