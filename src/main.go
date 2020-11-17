package main

import (
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	// create an instance of the person type
	p := &Person{
		Name: "tsuika",
		Age:  22,
	}

	// marshall it into protobuff format
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Protobuf data: ", data)
}
