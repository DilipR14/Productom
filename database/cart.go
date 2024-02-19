package database

import(
	"context"
	"errors"
	"log"
	"time"
	"github.com/DilipR14/Productom.git/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-dirver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"


)

var(
	ErrCanFindProduct = errors.New("can't find the product")
	ErrCanDecodeProducts = errors.New("can't find the product")
	ErrUserIdNotValid = errors.New("this user is not vaild")
	ErrCanUpdateUser = errors.New("cannot add this product to the cart")
	ErrCanRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCanGetItem = errors.New("was unable to get the item from the cart")
	ErrCanBuyCartItem = errors.New("Cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, UserCollection *mongo.collection, productID primitive.ObjectID, userID string) error{
	searchfromdb, err := prodCollection.find(ctx, bson.M{"_id": productID})
	if err != nil{
		log.Println(err)
		return ErrCanFindProduct
	}
	var productCart []models.ProductUser

	err = searchfromdb.All(ctx, &productCart)
	if err !=nil{
		log.Println(err)
		return ErrCanDecodeProducts
	}
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{key:"_id", value:id}}
	update := bson.D{{key:"$pust", Value: bson.D{primitive.E{key:"usercart", value:bson.D{{key:"$each", Value:productCart}}}}}}

	_, err = UserCollection.UpdateOne(ctx, filter, update)
	if err!= nil{
		return ErrUserIdIsNotValid
	}
	return nil
}

func RemoveItemToCart(ctx context.Context, prodCollection, UserCollection *mongo.collection, productID primitive.ObjectID, userID string) error{
	
	id, err := primitive.ObjectIDFromHex(usrID)
	if err!=nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D{primitive.E{key:"_id", value:id}}
	update := bson.D{"$pull"bson.M{"usercart":bson.M{"_id":productID}}}

	_,err := updateMany(ctx, filter, update)
	if err!=nil{
		return ErrCanRemoveItemCart
	}
	return nil

}

func BuyItemFromCart(ctx context.Context, UserCollection *mongo.collection, userID string) error{
	
	id, err := primitive.ObjectIDFromHex(usrID)
	if err!=nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	var getcartitems models.User
	var ordercart model.Order

	ordercart.Order_ID = primitive.NewObjectID()
	ordercart.Order_At = time.Now()
	ordercart.Order_Cart =make([]models.ProductUser, 0)
	ordercart.payment_Method.COD = true

	unwind := bson.D{{key :"$unwind", Value:bson.D{primitive.E{key:"path", Value:"$usercart"}}}}
	grouping := bson.D{{key:"$group", Value:bson.D{primitive.E{key:"_id", Value:"$_id"},{key:"total", Value:bson.D{primitive.E{key:"$sum", Value:"$usercart.price"}}}}}}
	
	currentresults, err:= UserCollection.Aggregate(ctx, mongo.pipeline{unwind, grouping})
	ctx.Done()
	if err!=nil{
		panic(err)
	}

	var getusercart []bson.make
	currentresults.All(ctx,&getusercart); err !=nil{
		panic(err)
	}

	var total_price int32

	for _, user_item:= range getusercart{
		price:= user_item["total"]
		total_price = price.(int32)
	}
	ordercart.price= int(total_price)

	filter := bson.D{primitive.E{key:"_id", value: id}}
	update := bson.D{{key:"$push", Value:bson.D{primitive.E{key:"orders", Value:ordercart}}}}

	_,err := UserCollection.UpdateMany(ctx, filter, update)
	
	if err!=nil{
		log.Println(err)
	}

	err =UserCollection.FindOne(ctx, bson.D{primitive.E{key :"_id", value:id}}).Decode(&getcartitems)
	if err != nil{
		log.Println(err)
	}
	filter2 := bson.D{primitive.E{key:"_id", Value:id}}
	update2 :=bson.M{"$push":bson.M{"order.$[].order_list":bson{"$each":getcartitems.UserCart}}}
	_,err := UserCollection.UpdateOne(ctx, filter2, update2)
	
	if err!=nil{
		log.Println(err)
	}

	usercart_empty:=make{[]models.ProductUser,0}
	filter3 := bson.D{primitive.E{key:"_id", Value:id}}
	update3 :=bson.D{{Key:"$set", Value:bson.D{primitive.E{key:"usercart", Value:usercart_empty}}}}
	_,err := UserCollection.UpdateOne(ctx, filter3, update3)
	if err != nil{
		return ErrCanBuyCartItem
	}

	return nil

}

func InstantBuyer(ctx context.Context, prodCollection, UserCollection *mongo.collection, productID primitive.ObjectID, userID string) error{

	id, err := primitive.ObjectIDFromHex(usrID)

	if err!=nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	var product_details models.ProductUser
	var order_details models.Order

	order_details.Order_ID= primitive.NewObjectID()
	order_details.Order_At = time.Now()
	order_details.Order_Cart =make([]models.ProductUser, 0)
	order_details.payment_Method.COD = true
	prodCollection.FindOne(ctx, bson.D{primitive.E{key:"_id", Value:productID}}).Decode(&product_details)

	if err!=nil{
		log.Println(err)
	}
	order_details.Price = product_details.price

	filter := bson.D{primitive.E{key:"_id", value: id}}
	update := bson.D{{key:"$push", Value:bson.D{primitive.E{key:"orders", Value:order_detail}}}}
	_,err := UserCollection.updateMany(ctx, filter, update)
	if err!=nil{
		log.Println(err)
	}

	filter2 := bson.D{primitive.E{key:"_id", value: id}}
	update2:= bson.M{"$push":bson.M{"orders.$[].order_list":product_details}}
	_,err := UserCollection.updateMany(ctx, filter2, update2)
	if err!=nil{
		log.Println(err)
	}
	return nil
}