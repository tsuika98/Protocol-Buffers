package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	// create an instance of the person type
	p := &Person{
		Name: "tsuika",
		Age:  22,
		SocialMediaFollowers: &SocialMediaFollowers{
			Youtube: 2500,
			Twitter: 500,
		},
	}

	log.Print("Person before marshal")
	printPersonInfo(p)

	// marshall it into protobuff format
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Person after marshalling: ", data)
	fmt.Println("")

	// unmarshal protocol buffer object into a person type
	newP := &Person{}
	err = proto.Unmarshal(data, newP)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Person after unmarshalling")
	printPersonInfo(p)

}

func printPersonInfo(person *Person) {
	// use generated methods in person.pb file to retrieve field names of the person type
	log.Print("Name: ", person.GetName())
	log.Print("Age: ", person.GetAge())
	log.Print("SocialFollowing: ", person.GetSocialMediaFollowers())
	fmt.Println("")
}
