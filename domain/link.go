package domain

//  Link represents the relationship beetween clients
type Link struct {
	// x
	RootId   uint64   `json:"root_id"`
	NodeId   uint64   `json:"node_id"`
	LinkType LinkType `json:"link_type"`
}

//  LinkType describe the type of relation beewtween clients
type LinkType struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
