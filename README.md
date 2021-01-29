# BandGo

네이버 밴드 API

## Install
```shell
go get github.com/Beta5051/BandGo
```

## Example
```go
package main

import (
	"github.com/Beta5051/BandGo"
	"log"
)

func main() {
	client := BandGo.New("YOUR TOKEN", false)

	bands, err := client.GetBands()
	if err != nil {
		log.Fatal(err)
	}

	for _, band := range bands {
		_, _, err = client.CreatePost(band.BandKey, "test", false)
		if err != nil {
			log.Println(err)
		}
	}
}
```