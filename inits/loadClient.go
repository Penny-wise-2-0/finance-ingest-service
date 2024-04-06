package inits

import (
	"os"

	"github.com/plaid/plaid-go/v23/plaid"
)



var PlaidClient *plaid.APIClient

func LoadPlaidClient()  {
	clientID := os.Getenv("PLAID_CLIENT_ID")
	secret := os.Getenv("PLAID_SECRET")
	environment := plaid.Sandbox 
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", clientID)
	configuration.AddDefaultHeader("PLAID-SECRET", secret)
	configuration.UseEnvironment(environment)
	client := plaid.NewAPIClient(configuration)

	PlaidClient = client
}
