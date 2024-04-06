package routes

import (
	
	"github.com/gin-gonic/gin"
)


func InitializeRoutes (r *gin.Engine) {

	BudgetRoutes(r)
	PlaidLinkRoutes(r)
	PlaidProductRoutes(r)
}




