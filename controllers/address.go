package controllers

import(
	"time"
	"context"
	"log"
	"errors"
	"net/http"
	"github.com/DilipR14/Productom.git/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-dirver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)
	

func AddAddress() gin.HandlerFunc{

	return func(c. *gin.Context){
	user_Id :=c.Query("id")

		if user_id ==""{
			c.Header("Content-Type", "application/json")
			c.json(http.StatusNotFound, gin.H{"error":"Invalid code"})
			c.Abort()
			return
		}
		address, err := ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJson(500, "Internal Server Error")
		}	
		var addresses models.Address

		addresses.Address_id = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil{
			c.IndentedJson(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancal = context.WithTimeOut(context.Background(), 100*time.Second)

		match_filter := bson.D{{key:"$match", Value: bson.D{primitive.E{key:_"id", Value:address}}}}
		unwind := bson.D{{key:"$unwind", Value: bson.D{primitive.E{key:"path", value:"$address"}}}}
		group := bson.D{{key:"$group", Value:bson.D{primitive.E{key:"_id", Value:"$address_id"},{key:"count", Value:bson.D{primitive.E{key:"$sum", Value: 1}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.pipeline{match_filter, unwind, group})
		if err != nil{
			c.IndentedJson(500, "Internal Server error")
		}

		var addressinfo []bson.M 
		if err = pointcursor.All(ctx, &addressinfo); err !=nil{
			panic(err)
		}
		var size int32
		for _, address_no := range addressinfo{
			count:= address_no["count"]
			size = count.(int32)
		}
		if size < 2{
			filter := bson.D{primitive.E{key:"_id", Value: address}}
			update := bson.D{{key:"$push", Value: bson.D{primitive.E{Key:"address", Value:addresses}}}}
			_,err:= UserCollection.UpdateOne(ctx, filter, update)

			if err != nil{
				fmt.Println(err)
			}

		}else{
			c.IndentedJson(400, "Not Allowed")
		}

		defer cancal()
		ctx.Done()

	}
}

func EditHomeAddress() gin.HandlerFunc{

	return func(c. *gin.Context){
		user_Id :=c.Query("id")

		if user_id ==""{
			c.Header("Content-Type", "application/json")
			c.json(http.StatusNotFound, gin.H{"error":"Invalid"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJson(500, "Internal Server Error")
		}
		
		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err!= nil{
			c.IndentedJson(http.statusBadRequest, err.Error())
		}
		var ctx,cancal = context.WithTimeOut(context.Background(),100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", value:"usert_id"}}
		Update := bson.D{{key:"$set", Value: bson.D{primitive.E{key:"address.0.house_name", Value: editaddress.House},{Key:"address.0.street_name", Value:editaddress.street},{key:"address.0.city_name", Value: editaddress.city},{key:"address.0.pin_code", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpateOne(ctx, filter, Upate)
		if err != nil{
			c.IndentedJson(500, "Something went wrong")
			return
		}
		defer cancal()
		ctx.Done()
		c.IndentedJson(200, "Successfully updated the home address")
	}
		
}

func EditWorkAddress() gin.HandlerFunc{
	return func(c. *gin.Context){
		user_Id :=c.Query("id")

		if user_id ==""{
			c.Header("Content-Type", "application/json")
			c.json(http.StatusNotFound, gin.H{"error":"Invalid"})
			c.Abort()
			return
		}

		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJson(500, "Internal Server Error")
		}
		
		var editaddress models.Address
		if err := c.BindJSON(&editaddress); err!= nil{
			c.IndentedJson(http.statusBadRequest, err.Error())
		}
		var ctx,cancal = context.WithTimeOut(context.Background(),100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", value:"usert_id"}}
		Update := bson.D{{key:"$set", Value: bson.D{primitive.E{key:"address.1.house_name", Value: editaddress.House},{Key:"address.1.street_name", Value:editaddress.street},{key:"address.1.city_name", Value: editaddress.city},{key:"address.1.pin_code", Value: editaddress.Pincode}}}}
		_, err = UserCollection.UpateOne(ctx, filter, Upate)
		if err != nil{
			c.IndentedJson(500, "Something went wrong")
			return
		}
		defer cancal()
		ctx.Done()
		c.IndentedJson(200, "Successfully updated the work address")
	}
		
}


func DeleteAddress() gin.HandlerFunc{

	return func(c * gin.context){
		user_Id :=c.Query("id")

		if user_id ==""{
			c.Header("Content-Type", "application/json")
			c.json(http.StatusNotFound, gin.H{"error":"Invalid Search Index"})
			c.Abort()
			return
		}
		addresses := make([]models.Address,0)
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil{
			c.IndentedJson(500, "Internal Server Error")
		}

		var ctx,cancal = context.WithTimeOut(context.Background(),100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{key:"_id", value:"usert_id"}}
		Update := bson.D{{key:"$set", Value: bson.D{primitive.E{key:"address", value: addresses}}}}
		_, err = UserCollection.UpateOne(ctx, filter, Upate)
		
		if err !=nil{
			c.IndentedJson(404, "Wrong command")
			return
		}
		defer cancal() 
		ctx.Done()
		c.IndentedJson(200, "Successfully Deleted")
	}
}
