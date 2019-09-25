package dnsapi

import (
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

type HeaderAuthentication struct {
	ID     string
	Secret string
}

func (h HeaderAuthentication) AuthenticateRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetHeaderParam("PIR5-ID", h.ID); err != nil {
		return err
	}
	if err := r.SetHeaderParam("PIR5-SECRET", h.Secret); err != nil {
		return err
	}
	return nil

}
