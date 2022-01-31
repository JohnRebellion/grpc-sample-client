package grpc

import (
	"context"
	"fmt"
	"grpc-sample-client/person"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Context context.Context
var Cancel context.CancelFunc
var DefaultDialOption grpc.DialOption = grpc.WithTransportCredentials(insecure.NewCredentials())

type Service struct {
	Service interface{}
	Client  ServiceClient
	Connect func(cc grpc.ClientConnInterface) interface{}
	Name    string
}

type ServiceClient struct{}

func (service *Service) Register(connection *grpc.ClientConn) {
	service.Service = service.Connect(connection)
}

type Client struct {
	Host        string
	Port        string
	DialOptions []grpc.DialOption
	Connection  *grpc.ClientConn
	Error       error
	Services    []Service
}

func (c *Client) New() error {
	if len(c.DialOptions) == 0 {
		c.DialOptions = append(c.DialOptions, DefaultDialOption)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.Host, c.Port), DefaultDialOption)

	if err == nil {
		c.Connection = conn

		for _, service := range c.Services {
			service.Register(c.Connection)
		}
	}

	c.Error = err
	return err
}

func a() {
	personServiceClient := new(person.PersonServiceClient)
	personClient := Client{
		Host: "localhost",
		Port: "8111",
		Services: []Service{
			Service{
				Client: ServiceClient{},
				Name:   "person",
				Connect: func(cc grpc.ClientConnInterface) interface{} {
					return personServiceClient
				},
			},
		},
	}

	err := personClient.New()

	if err == nil {
		personClient
	}
}
