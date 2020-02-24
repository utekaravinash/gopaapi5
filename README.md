# gopaapi5 #

gopaapi5 is a Go client library for accessing the [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).

## Usage ##

Install package:

```bash
go get -u github.com/utekaravinash/gopaapi5
```

Here's an example for GetBrowseNodes operation:

```go
package main

import (
	"fmt"
	"os"

	"github.com/utekaravinash/gopaapi5"
	"github.com/utekaravinash/gopaapi5/api"
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
	response, err := client.GetBrowseNodes(&params)
	if err != nil {
		panic(err)
	}

	// Loop over browse nodes in response
	for _, node := range response.BrowseNodesResult.BrowseNodes {
		fmt.Printf(tmpl, node.Id, node.DisplayName)
	}
}
```

This client library exposes these operation for Amazon Product Advertising API 5.0:

- [GetBrowseNodes](https://github.com/utekaravinash/gopaapi5/blob/master/_examples/get_browse_nodes/main.go)
- [GetItems](https://github.com/utekaravinash/gopaapi5/blob/master/_examples/get_items/main.go)
- [GetVariations](https://github.com/utekaravinash/gopaapi5/blob/master/_examples/get_variations/main.go)
- [SearchItems](https://github.com/utekaravinash/gopaapi5/blob/master/_examples/search_items/main.go)


## Author ##

[Avinash Utekar](https://www.utekar.com/author/avinash/)

## License ##

[BSD 3-Clause](https://github.com/utekaravinash/gopaapi5/blob/master/LICENSE)
