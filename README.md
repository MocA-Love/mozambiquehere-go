# Mozambiquehe.re API Client

This library is a Golang API wrapper for https://apexlegendsapi.com/ .

## Install

```golang
go get github.com/MocA-Love/mozambiquehere-go/v4
```

## Usage example

```golang
package main

import (
	"fmt"

	"github.com/MocA-Love/mozambiquehere-go/v4/client"
	"github.com/MocA-Love/mozambiquehere-go/v4/domain/models"
)

func main() {
	apexClient := client.DefaultClient("api_token")
	data, err := apexClient.GetUserByName(models.PlatformPC, "MocA-Love")

	if err != nil {
		panic(err)
	}

	fmt.Println(data.Global.Level) //38
}

```


