package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Penny-wise-2-0/ingest-service/inits"
	"github.com/Penny-wise-2-0/ingest-service/models"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v23/plaid"
)


func CreateLinkToken(c *gin.Context) {
	userID := c.Param("userID")

	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: userID,
	}

	request := plaid.NewLinkTokenCreateRequest(
		"PennyWise",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)
	request.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})
	
	
	res, _, err := inits.PlaidClient.PlaidApi.LinkTokenCreate(context.Background()).LinkTokenCreateRequest(*request).Execute()

	if err != nil {
		fmt.Printf("Failed to create link token: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link token"})
	}

	c.JSON (http.StatusOK, gin.H{"link_token": res.GetLinkToken()})
}



func ExchangePublicToken(c *gin.Context) {
	type Req struct {
		UserID string `json:"userID"`
		Token string `json:"token"`

	}

	var req Req

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ctx := context.Background()
	exchangePublicTokenReq := plaid.NewItemPublicTokenExchangeRequest(req.Token)
	exchangePublicTokenRes, _, err := inits.PlaidClient.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(*exchangePublicTokenReq).Execute()
	
	if err != nil {
		fmt.Println("Failed to exchange publix token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange public token"})
	}
	itemId := exchangePublicTokenRes.ItemId
	accessToken := exchangePublicTokenRes.AccessToken

	token := &models.PlaidAccessToken{
		UserId: req.UserID,
		Token: accessToken,
		ItemID: itemId,
	} 

	res := inits.DB.Create(&token)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": accessToken})



	
	
}