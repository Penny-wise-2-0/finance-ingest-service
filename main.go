package main

import (
	"github.com/Penny-wise-2-0/finance-ingest-service/inits"
	"github.com/Penny-wise-2-0/finance-ingest-service/middleware"

	"github.com/Penny-wise-2-0/finance-ingest-service/routes"

	"github.com/gin-gonic/gin"
)


func init() {
	inits.LoadEnvVars()
	inits.ConnectToDB()
	inits.LoadPlaidClient()
	
}

func main() {
	r := gin.Default()
	r.Use(middleware.Logger)
	r.Use(middleware.CorsMiddleware())
	
	
	routes.InitializeRoutes(r)
	
	

	r.Run() 
}