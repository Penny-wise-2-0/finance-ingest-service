package main

import (
	
	"github.com/Penny-wise-2-0/ingest-service/inits"
	"github.com/Penny-wise-2-0/ingest-service/models"
)

func init (){
inits.LoadEnvVars()
inits.ConnectToDB()

}

func main() {
	inits.DB.AutoMigrate(&models.Budget{})
	inits.DB.AutoMigrate(&models.PlaidAccessToken{})
	inits.DB.AutoMigrate(&models.Transaction{})
	inits.DB.AutoMigrate(&models.Cursor{})

}

