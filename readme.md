# go-fireblocks-sdk

fireblocks rest api implementation of the go language version, **if you are using it in production, please rigorous testing**

[Fireblocks REST Api](https://docs.fireblocks.com/api/#introduction)


## example


	endpoint := "https://api.fireblocks.io"
	apiKey := ""
	secretPath := ""

	fbClient, err = NewFBClient(endpoint, apiKey, secretPath)
	if err != nil {
		panic(err)
	}


    




