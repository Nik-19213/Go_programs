package main

//go get google.golang.org/protobuf@latest
//go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
//protoc --go_out=. --go_opt=paths=source_relative person.protos

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// let's go the other way and unmarshal
	// our byte array into an object we can modify
	// and use
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// print out our `newElliot` object
	// for good measure
	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.SocialFollowers.GetTwitter())
	fmt.Println(newElliot.SocialFollowers.GetYoutube())

}
