#A golang AppAnnie Library

##usage
```go
package main

import (
	"github.com/athurg/appannie_api"
	"log"
)
func main(){
	apiKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	apiVer := "" //default is v1.2
	client := appannie.New(apiKey, apiVer)
	accounts, err := client.Accounts()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(accounts)
}
```

About the parameter, please refer [AppAnnie's API Doc](https://support.appannie.com/hc/en-us/categories/202773667-API) for detail.
Other problems please [create an issue here](https://github.com/athurg/appannie_api/issues/new)

###TODO
Because of accounts limit, many apis need to be more testing with special apikey.
