package main

import (
	"context"
	"fmt"
	"os"

	"github.com/utekaravinash/gopaapi5/api"
	"github.com/utekaravinash/gopaapi5/v2"
)

const tmpl = `Id: %s
Name: %s
-----------------------------
`

func main() {

	// Get Access Key, Secret Key and Associate Tag from the environment variables
	accessKey := os.Getenv("PA_ACCESS_KEY")
	secretKey := os.Getenv("PA_SECRET_KEY")
	associateTag := os.Getenv("PA_ASSOCIATE_TAG")

	// Initiate gopaapi5 Client
	client, err := gopaapi5.NewClient(accessKey, secretKey, associateTag, api.UnitedStates)
	if err != nil {
		panic(err)
	}

	// Construct request parameters for GetBrowseNodes operation
	params := api.GetBrowseNodesParams{
		BrowseNodeIds: []string{
			"6960520011",
			"281407",
		},
		Resources: []api.Resource{
			api.BrowseNodesAncestor,
			api.BrowseNodesChildren,
		},
		LanguagesOfPreference: []api.Language{api.EnglishUnitedStates},
	}

	// Call GetBrowseNodes operation
	response, err := client.GetBrowseNodes(context.Background(), &params)
	if err != nil {
		panic(err)
	}

	// Loop over browse nodes in response
	for _, node := range response.BrowseNodesResult.BrowseNodes {
		fmt.Printf(tmpl, node.Id, node.DisplayName)
	}
}
