package controllers

import (
	"fmt"
	"Context"
	"log"
	"time"
	"net/http"
	
	"github.com/DilipR14/Productom.git/database"
	"github.com/DilipR14/Productom.git/models"
	generate"github.com/DilipR14/Productom.git/tokens"
	"github.com/go-playground/validator.v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-dirver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "User")
var ProdCollection  *mongo.Collection = database.ProductData(database.Client, "Products")
var validate = vaildator.New()

func HashPassword(Password string) string{

	bcrypt.GenerateFromPassword	([]byte(password), 14)
	if err != nil{
		log.panic(err)
	}	
	return string(bytes)
}

func VerifyPassword(UserPassword string, givenPassword string) (bool, string){

	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(UserPassword))
	valid := true
	msg := ""

	if err != nil{
		msg = "Log or Password is incorrect"
		Valid = false
	}
	return Valid, msg
}

func Signup() gin.HandlerFunc{

	return func(c *gin.context){
		var ctx, cancel = context.withTimeOut(context.Background(), 100*time.second)
	 defer cancel()

		var user models.UserData
		if err := c.BindJSON(&user); err != nil{
			c.JSON{http.statusBadRequest, gin.H{"error": validationErr}}
			return
		}

		count, err = UserCollection.countDocuments(ctx, bson.M{"email": user.Email})
		if err != nil{
			log.Panic(err)
		    c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		    return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"user already exists"})
		}

		//

		count, err = UserCollection.countDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this phone no. is already used"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.parse(time.RFC3339, time.Nmw().Format(time.RFC3339))
		user.Upate_At, _ = time.parse(time.RFC3339, time.Nmw().Format(time.RFC3339))
		User_ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		tocken, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, *user.User_Id )
		user.token = &token
		user.Refesh_Token = &refresh_token
		user.UserCart = make([]models.ProductUser,0)
		user.Address_Details = make([]model.Address,0)
		user.Order_Status = make([]model.Order,0)
		_, Inserterr := UserCollection.InsertOne(ctx, use)

		if Inserterr != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error" : "the user did not get created"})
			return
		}
		defer cancal()

		c.JSON(http.StatusCreated, "Successfully singned in!")

	}	
}

func Login() gin.HandlerFunc{
	ret func(c *gin.context){
		var ctx, cancal = context.withTimeOut(context.Background(), 100*time.second)
		defer cancal()

		var user models.User
		if err := c.BindJSON(&user); err != nil{
			c.json(http.statusBadRequest, gin.H{"error" : err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancal()

		if err != nil{
			c.json(http.StatusInternalServerError, gin.H{"error" : "login or password incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancal()

		if !PasswordIsValid{
			c.json(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshtoken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)
		defer cancal()

		c.json(http.StatusFound, founduser) 
	}
}

func ProductViewerAdmin() gin.HandlerFunc{
}

func SearchProduct() gin.HandlerFunc{

	return func(c *gin.context){

		var productlist []models.product
		var ctx, cancal = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancal()

		ProdCollection.find(ctx, bson.D{{}})
		if err != nil{
			c.IndentedJson(http.StatusInternalServerError, "something went wrong, please try after some time")
			return
		}
		err = cursor.All(ctx, &productlist)

		if err!= nil{
			log.println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cursor.close()

		if err := cursor.err(); err != nil{
			log.Println(err)
			c.IndentedJson(400, "invalid")
			return
		}
		defer cancal()
		c.IndentedJson(200, productlist)
	}
}

func SearchProductByQuery() gin.HandlerFunc{
	return func(c *gin.Context){
		var SearchProduct[]models.product
		queryparam := c.Queryparam("name")

		if queryparam ==""{
			log.Println("query is empty")
			c.Header("Content-type", "application/json")
			c.json(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.About()
			return
		}

		var ctx, cancal = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancal()

		searchquerydb, err := ProdCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryparam}})

		if err != nil{
			c.IndentedJson(404, "something went wrong fetching the data")
			return
		}

		err = searchquerydb.All(ctx. &SearchProducts)
		if err != nil{
			log.Println(err)
			c.IndentedJson(404, "invalid")
			return
		}

		defer searchquerydb.close(ctx)

		if err :=searchquerydb.Err(); err != nil{
			log.Println(err)
			c.IndentedJson(404, "Invalid")
			return
		}

		defer cancal()
		c.IndentedJson(200, SearchProducts)
	}
} 