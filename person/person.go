package person

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PersonJSON struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Ages int64  `json:"age"`
}

var ServiceClient PersonServiceClient
var ctx context.Context
var cancel context.CancelFunc
var DefaultDialOption grpc.DialOption = grpc.WithTransportCredentials(insecure.NewCredentials())

func connect(host, port string, option ...grpc.DialOption) error {
	if len(option) == 0 {
		option = append(option, DefaultDialOption)
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), DefaultDialOption)

	if err == nil {
		ServiceClient = NewPersonServiceClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	}

	return err
}

func GetAll() ([]PersonJSON, error) {
	people := []PersonJSON{}
	start := time.Now().UnixMilli()
	log.Println(start)
	allPerson, err := ServiceClient.GetAll(ctx, &Empty{})
	end := time.Now().UnixMilli()
	log.Println(end)
	log.Println(end - start)

	if err != nil {
		log.Fatalf("Person service error: %v", err)
	} else {

		convert(allPerson.People, &people)

		if err == nil {
			return people, err
		}
	}

	return nil, err
}

func convert(i interface{}, o interface{}) error {
	b, err := json.Marshal(i)

	if err == nil {
		return json.Unmarshal(b, &o)
	}

	return err
}
