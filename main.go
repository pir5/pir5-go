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
		ID:     "pir-5-id",
		Secret: "pir-5-secret",
	})
	pp.Println(err)
	pp.Println(res)
}
