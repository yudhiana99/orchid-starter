package clients

import (
	internalClient "orchid-starter/internal/clients/internal"
)

type Client struct {
	InternalClient internalClient.InternalClientInterface
}

func NewClient() *Client {
	return &Client{
		InternalClient: internalClient.NewInternalClient(),
	}
}
