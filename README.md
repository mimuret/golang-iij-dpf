# IIJ DNS Platform Service API for Go
[![codecov](https://codecov.io/gh/mimuret/golang-iij-dpf/branch/main/graph/badge.svg?token=CTIBL2M5YD)](https://codecov.io/gh/mimuret/golang-iij-dpf)

- API Library for [IIJ DNS Platform Service](https://www.iij.ad.jp/en/biz/dns-pfm/).
- This is not an official IIJ software.

## Usage
```
package main

import (
	"fmt"
	"os"

	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/apis/core"
)

func main() {
	token := os.Getenv("DPF_TOKEN")
	cl := api.NewClient(token, "", nil)

	zoneList := &core.ZoneList{}
	searchParam := &core.ZoneListSearchKeywords{Name: api.KeywordsString{"example.jp"}}
	req, err := cl.ListAll(zoneList, searchParam)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RequestID: %s\n", req)
	for _, item := range zoneList.Items {
		fmt.Println(item)
	}
}
```