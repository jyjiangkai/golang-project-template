package client

type TestClientInterface interface {
	Get() string
	Update(name string) string
}