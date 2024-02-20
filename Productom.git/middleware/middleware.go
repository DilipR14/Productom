package middleware

import(
	"net/http"
	
	token "github.com/DilipR14/Productom.git/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc{
	return func(c *gin.Context){

		ClientToken:= c.Reruest.Handler.Get("token")

		if ClientToken ==""{
			c.json(http.StatusInternalServerError,gin.H{"error":"No authorization header provided"})
			c.About()
			return
		}
		claims,err := token.ValidateToken(ClientToken)
		if err != "" {
			c.json(http.StatusInternalServerError, gin.H{"error":err})
		    c.Abort()
		    return
		}

		c.set("email", claims.Email)
		c.set("uid", claims.Uid)
		c.Next()
	}
}