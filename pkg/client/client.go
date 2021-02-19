package client

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"github.com/MaxPolarfox/toDoList/toDoList"
)

type Client interface {
	SayHello(ctx context.Context) error
}

type ToDoListClientImpl struct {
	client toDoList.ToDoListServiceClient
}

func NewToDoListClient() Client {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//defer conn.Close()

	c := toDoList.NewToDoListServiceClient(conn)

	return &ToDoListClientImpl{
		client: c,
	}
}

func (i *ToDoListClientImpl) SayHello(ctx context.Context) error {
	response, err := i.client.SayHello(context.Background(), &toDoList.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	return nil
}