package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Penny-wise-2-0/finance-ingest-service/inits"
	"github.com/Penny-wise-2-0/finance-ingest-service/models"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v23/plaid"
	"gorm.io/gorm"
)

func SaveTransactions(c *gin.Context) {
	accessToken := c.MustGet("plaidAccessToken").(string)
	userID := c.Param("userID")
	fmt.Println("access token", accessToken)
	hasMore := true
	
	var existingCursor models.Cursor
	result := inits.DB.Where("user_id = ?", userID).First(&existingCursor)

	var cursor *string = nil
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("Error fetching cursor: ", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
		

	}else {
		cursor = &existingCursor.Cursor
	}


	ctx := context.Background()

	for hasMore {
		req := plaid.NewTransactionsSyncRequest(accessToken)
		if cursor != nil {
			req.SetCursor(*cursor)
		}

		res, _, err := inits.PlaidClient.PlaidApi.TransactionsSync(ctx).TransactionsSyncRequest(*req).Execute()
		if err != nil {
			fmt.Printf("Error getting transactions: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(res)
	
		for _, transaction := range res.GetAdded() {
			
			fmt.Println("transaction:", transaction)
			fmt.Printf("Type of category: %T\n", transaction.GetCategory())
			category, err := json.Marshal(transaction.GetCategory())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		
				
				newTransaction := models.Transaction{
				UserID: userID,
				TransactionID: transaction.GetTransactionId(),
				AccountID: transaction.GetAccountId(),
				Amount: transaction.GetAmount(),
				Category: string(category),
				Date: transaction.GetDate(),
				Name: transaction.GetName(),
				Address: transaction.Location.GetAddress(),
				City: transaction.Location.GetCity(),
				Pending: transaction.GetPending(),
				TransactionType: transaction.GetTransactionType(),
				ISOCurrencyCode: transaction.GetIsoCurrencyCode(),
				MerchantName: transaction.GetMerchantName(),
				PaymentChannel: transaction.GetPaymentChannel(),
			}

			result := inits.DB.Create(&newTransaction)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}

		for _, transaction := range res.GetModified() {
			fmt.Println(transaction)
			category, err := json.Marshal(transaction.GetCategory())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			updatedTransaction := models.Transaction{
				UserID: userID,
				TransactionID: transaction.GetTransactionId(),
				AccountID: transaction.GetAccountId(),
				Amount: transaction.GetAmount(),
				Category: string(category),
				Date: transaction.GetDate(),
				Name: transaction.GetName(),
				Address: transaction.Location.GetAddress(),
				City: transaction.Location.GetCity(),
				Pending: transaction.GetPending(),
				TransactionType: transaction.GetTransactionType(),
				ISOCurrencyCode: transaction.GetIsoCurrencyCode(),
				MerchantName: transaction.GetMerchantName(),
				PaymentChannel: transaction.GetPaymentChannel(),
			}

			result := inits.DB.Model(&models.Transaction{}).Where("transaction_id = ?", transaction.GetTransactionId()).Updates(updatedTransaction)

			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}

		for _, transactionID := range res.GetRemoved() {
			result := inits.DB.Where("transaction_id = ?", transactionID).Delete(&models.Transaction{})
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
		}

		hasMore = res.GetHasMore()
		if newCursor := res.GetNextCursor(); newCursor != "" {
			cursor = &newCursor

		} else {
			break
		}
	}
	saveCursor(userID, *cursor)
}

func saveCursor (userID string , cursor string) {
	var existingCursor models.Cursor
	result := inits.DB.Where("user_id = ?", userID).First(&existingCursor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			newCursor := models.Cursor{
				UserID: userID,
				Cursor: cursor,
			}
			result := inits.DB.Create(&newCursor)
			if result.Error != nil {
				fmt.Println("Error creeating cursor: ", result.Error)
			}
			
		}else {
			fmt.Println("error fetching cursor: ", result.Error)
		}
	}else {
		existingCursor.Cursor = cursor
		result := inits.DB.Save(&existingCursor)
		if result.Error != nil {
			fmt.Println("Error updating cursor: ", result.Error)
		}
	}
}




func GetTransactions (c *gin.Context) {
	userID := c.Param("userID")
	var transactions []models.Transaction
	result := inits.DB.Where("user_id = ?", userID).Find(&transactions)
	if result.Error != nil {
		fmt.Println("Error fetchign transactions: ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}