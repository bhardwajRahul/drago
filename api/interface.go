package api

import (
	"path"

	"github.com/seashell/drago/drago/structs"
)

const (
	interfacesPath = "/api/interfaces"
)

// Interfaces is a handle to the interfaces API
type Interfaces struct {
	client *Client
}

// Interfaces returns a handle on the interfaces endpoints.
func (c *Client) Interfaces() *Interfaces {
	return &Interfaces{client: c}
}

// Get :
func (n *Interfaces) Get(id string) (*structs.Interface, error) {

	var iface *structs.Interface
	err := n.client.getResource(interfacesPath, id, &iface)
	if err != nil {
		return nil, err
	}

	return iface, nil
}

// List :
func (n *Interfaces) List(filters map[string][]string) ([]*structs.InterfaceListStub, error) {

	var items []*structs.InterfaceListStub
	err := n.client.listResources(path.Join(interfacesPath, "/"), filters, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
