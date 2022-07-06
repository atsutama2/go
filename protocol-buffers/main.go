package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"protocol-buffers/pb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "Suzuki@gmail.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "Suzuki daiki",
		},
		Birthday: &pb.Date{
			Year:  2022,
			Month: 7,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't write", err)
	}

	if err := ioutil.WriteFile("test.bin", binData, 0666); err != nil {
		log.Fatalln("Can't write", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't read file", err)
	}

	readEmployee := &pb.Employee{}

	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't deserialize", err)
	}

	fmt.Println(readEmployee)
}
