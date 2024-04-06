package middleware

import (
	"fmt"
	"log"
	"net/http"


	"github.com/Penny-wise-2-0/ingest-service/inits"
	"github.com/Penny-wise-2-0/ingest-service/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Logger (c *gin.Context) {
	
	log.Printf("Received %s request to %s from %s", c.Request.Method, c.Request.URL.Path, c.Request.RemoteAddr)
	c.Next()
}

func CorsMiddleware() gin.HandlerFunc {
    config := cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowWildcard:    true,
        AllowOriginFunc: func(origin string) bool {
            return true
        },
        MaxAge: 86400,
    }

    return cors.New(config)
}


func GetAccessToken(c *gin.Context) {

	userID := c.Param("userID")

	var plaidAccessToken  models.PlaidAccessToken
	result :=  inits.DB.Where("user_id = ?", userID).First(&plaidAccessToken)
	if result.Error != nil {
		fmt.Println("access token not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "access token not found"})
	}
	
	fmt.Println("plaidtoken:", plaidAccessToken.Token)
	fmt.Println("done")
	c.Set("plaidAccessToken", plaidAccessToken.Token)
	c.Next()
}