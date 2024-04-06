package routes

import (
	

	"github.com/Penny-wise-2-0/finance-ingest-service/controllers"
	"github.com/Penny-wise-2-0/finance-ingest-service/middleware"
	"github.com/gin-gonic/gin"
)


func PlaidProductRoutes(r *gin.Engine) {
	plaidProductRoutes := r.Group("/plaid-product")
	plaidProductRoutes.Use(middleware.GetAccessToken)
	
	{
		plaidProductRoutes.GET("/sync-transactions/:userID", controllers.SaveTransactions)
		plaidProductRoutes.GET("/get-transactions/:userID", controllers.GetTransactions)
	}
}