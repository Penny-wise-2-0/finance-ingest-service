package routes

import (
	"github.com/Penny-wise-2-0/ingest-service/controllers"
	"github.com/gin-gonic/gin"
)


func BudgetRoutes (r *gin.Engine) {
	budgetRoutes := r.Group("/budgets")
	
	budgetRoutes.POST("", controllers.CreateBudget)
	budgetRoutes.POST("/get", controllers.GetBudgets)
	budgetRoutes.DELETE("", controllers.DeleteBudget)
	budgetRoutes.PUT("", controllers.UpdateBudget)
	
}


