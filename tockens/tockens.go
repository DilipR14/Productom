package tockens

import(
	"context"
	"log"
	"os"
	"time"

	"github.com/DilipR14/Productom.git/models"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-dirver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/options"
	
	"github.com/DilipR14/Productom.git/tokens"
	
)

type singnedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.userData(database.Client, "Users")
var SECRET_KEY= os.Getenv("SECRET_KEY")

func TokenGenerator(Email string, firstname string, lastname string, uid string)(singnedToken string, signedrefreshtoken string, err error){
	claims := &SingnedDetails{
		Email : email,
		First_Name : firstname,
		Last_Name : lastname,
		Uid : uid,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshclaims := &SingnedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour *time.Duration(168)).Unix(),
		},
	}

	tocken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SingnedString([]byte(SECRET_KEY))

	if err != nil{
		return"","",err
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SingnedString([]byte(SECRET_KEY))
	
	if err != nil{
		log.Panic(err)
	}
	return token, refreshtoken, err


}
func ValidateToken(singnedToken string)(claims *singnedDetails, msg string){
	token, err := jwt.ParseWithClaims(singnedToken, &SingnedDetails{},func(token *jwt.Token)(interface{}, error)){
		return[]byte(SECRET_KEY),nil
	}

	if err != nil{
		msg = err.Error()
		return
	}

	claims, ok := tocken.claims.(*SingnedDetails)
	if !ok{
		msg = "the token in invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().unix(){
		msg = "token is alredy expired"
		return
	}
	return claims, msg

}

func UpdateAllToken(singnedtoken string,signedrefreshtoken string, userid string){

	var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
	var updateobj primitive.D

	updateobj = append (updateobj, bson.E{key:"token", value : signedrefreshtoken})
	updateobj = append (updateobj, bson.E{key:"refresh_token", value : signedrefreshtoken})
	updated_at, _ := time .parse(time.RFC3339, time.Now().Format.RFC3339)
	updateobj = append (updateobj, bson.E{key:"updatedat", value : updated_at})
	
	upsert := true

	filter := bson.M{"user_id": userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	UserData.UpdateOne(ctx, filter, bson.D{
		{key:"$set", value:updateobj},
	},
      &opt)

    defer cancel()

	if err != nil{
		log.Panic(err)
		return
	}

}
