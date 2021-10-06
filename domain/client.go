package domain
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
	Name string `json:"name"`
	// the document number for this client
	// required: false
	Document uint64 `json:"document"`
	// the email address for this client
	// required: false
	// example: user@provider.net
	Email string `json:"email"`
	// the cell number for this client
	// required: false
	Cell uint64 `json:"phone"`
	// the unified password for this client
	// required: false
	Password string `json:"password"`
}


// validate Client fields
func (c *Client) Validate() error {
	return nil
}
