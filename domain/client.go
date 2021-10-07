// Package domain provides data structure to client api system
package domain

import (
	"github.com/go-playground/validator/v10"
)

//  Client represents basic informations of a client
//
// A client is the security principal for this application.
// It's also used as one of main axes for reporting.
//
// A client can have links with whom they can be connected in some form.
//
// swagger:model
type Client struct {
	// the id for this client
	// required: true
	Id uint64 `json:"id"`
	// the name for this client
	// required: true
	Name string `json:"name" validate:"required"`
	// the document number for this client
	// required: false
	Document uint64 `json:"document" validate:"required"`
	// the email address for this client
	// required: true
	// example: user@provider.net
	Email string `json:"email" validate:"required, email"`
	// the cell number for this client in the i164 pattern
	// required: false
	CellPhoneI164 uint64 `json:"cell_phone_i164"`
	// the unified password for this client
	// required: false
	Password string `json:"password" validate:"base64"`
}

// Validate is a Client method that validate if the fields are in the right expected format
func (c *Client) Validate() error {
	v := validator.New()
	return v.Struct(c)
}
