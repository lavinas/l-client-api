// Package domain provides data structure to client api system
package domain

//  Link represents a relationship beetween two clients
//
// A link allows a series of use cases in applications that use this system.
// It's also used as one of main axes for reporting.
//
// A link can be classified in determined type.
//
// swagger:model
type Link struct {
	// the id for this link
	// required: true
	Id uint64 `json:"id"`
	// the id of the client root or master of this relationship
	// required: true
	ClientRootId uint64 `json:"client_root_id" validate:"required"`
	// the id of the client related to the master of this relationship
	// required: true
	ClientNodeId uint64 `json:"client_node_id" validate:"required"`
	// the type that describe of this relationship
	// required: false
	LinkType LinkType `json:"link_type"`
}

//  LinkType describe the type of relation beewtween clients
//
// swagger:model
type LinkType struct {
	// the id for this link type
	// required: true
	Id uint64 `json:"id"`
	// the name or description of this type of link
	// required: true
	Name string `json:"name" validate:"required"`
}
