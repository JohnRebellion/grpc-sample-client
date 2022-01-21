package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-sample-client/envRouting"
	"grpc-sample-client/person"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	envRouting.LoadEnv()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", envRouting.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		personServiceClient := person.NewPersonServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		allPerson, err := personServiceClient.GetAll(ctx, &person.Empty{})

		if err != nil {
			log.Fatalf("Person service error: %v", err)
		} else {
			people := []person.PersonJSON{}
			convert(allPerson.People, &people)
			peopleJSON, err := json.Marshal(people)

			if err == nil {
				log.Printf("People: %s", string(peopleJSON))
			}
		}
	}
}

func convert(i interface{}, o interface{}) error {
	b, err := json.Marshal(i)

	if err == nil {
		return json.Unmarshal(b, &o)
	}

	return err
}
