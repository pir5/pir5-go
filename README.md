# pir5-go
this is pir5 api client by go.

## generate
```
$ cd github.com/pir5/pdns-api
$ swagger generate client -f docs/swagger.yaml -m model
```

## example
```go
package main

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/k0kubun/pp"
	"github.com/pir5/pir5-go/dnsapi"
	"github.com/pir5/pir5-go/dnsapi/operations"
)

func main() {
	client := dnsapi.NewHTTPClient(strfmt.Default)

	id := int64(2)
	p := operations.NewGetDomainsParamsWithTimeout(10 * time.Second)
	p.ID = &id
	res, err := client.Operations.GetDomains(p, dnsapi.HeaderAuthentication{
                // Header authentication(PIR5-ID,PIR5-SECRET)
		ID:     "pir-5-id",
		Secret: "pir-5-secret",
	})
	pp.Println(err)
	pp.Println(res)
}
```
