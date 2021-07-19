package client

type ClientInterface interface {
	Get() string
	Update(name string) string
}

type Client struct {
	Name   string
}

// NewClient
func NewClient() *Client {
	return &Client{
		"jiangkai",
	}
}

// Get
func (c *Client) Get() string {
	return c.Name
}

// Update
func (c *Client) Update(name string) string {
	c.Name = name
	return c.Name
}
