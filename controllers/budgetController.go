package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/Penny-wise-2-0/ingest-service/inits"
	"github.com/Penny-wise-2-0/ingest-service/models"
	"github.com/gin-gonic/gin"
)

type BudgetResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Frequency string `json:"frequency"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateBudget (c *gin.Context) {
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result := inits.DB.Create(&budget)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, budget)

}
type RequestBody struct {
	ID string
}

func GetBudgets (c *gin.Context) {
	
	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var budgets []models.Budget
	result := inits.DB.Where("user_id = ?", req.ID).Find(&budgets)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println("these are the budgets", budgets)

	var responses []BudgetResponse
	for _, b := range budgets {
		responses = append(responses, BudgetResponse{
			ID:        fmt.Sprintf("%d", b.ID),
			UserID:    b.UserID,
			Frequency: b.Frequency,
			Category:  b.Category,
			Name:      b.Name,
			Amount:    b.Amount,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		})
	}
	c.JSON(http.StatusOK, responses)

}

func DeleteBudget (c *gin.Context) {
	var ID RequestBody
	if err:= c.ShouldBindJSON(&ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	var deleted models.Budget
	id, err := strconv.ParseInt(ID.ID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

	result := inits.DB.Unscoped().Delete(&deleted, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	fmt.Println(deleted)
	c.JSON(http.StatusOK,gin.H{"message": "succesfully deleted budget"} )
}

func UpdateBudget(c * gin.Context) {
	var budgetResponse BudgetResponse
    if err := c.ShouldBindJSON(&budgetResponse); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id, err := strconv.ParseInt(budgetResponse.ID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var budget models.Budget
    if err := inits.DB.First(&budget, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
        return
    }
    budget.UserID = budgetResponse.UserID
    budget.Frequency = budgetResponse.Frequency
    budget.Category = budgetResponse.Category
    budget.Name = budgetResponse.Name
    budget.Amount = budgetResponse.Amount

    if err := inits.DB.Save(&budget).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Budget updated successfully"})
}
