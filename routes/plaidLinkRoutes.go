package routes

import (
	"github.com/Penny-wise-2-0/ingest-service/controllers"
	
	"github.com/gin-gonic/gin"
)


func PlaidLinkRoutes (r *gin.Engine) {
	plaidLinkRoutes := r.Group("/plaid-link")
	
	{
		plaidLinkRoutes.GET("/link-token/:userID", controllers.CreateLinkToken)
		plaidLinkRoutes.POST("/exchange-public-token", controllers.ExchangePublicToken)	
		
	}

}
