package main

import (
	"fmt"
	"log"
	"protobuf-lesson/pb"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"ProjectX": &pb.Company_Project{}},
		Profile:     &pb.Employee_Text{Text: "My name is Suzuki"},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	//binData, err := proto.Marshal(employee)
	//if err != nil {
	//	log.Fatalln("Can't serialize", err)
	//}
	//
	//if err := os.WriteFile("test.bin", binData, 0666); err != nil {
	//	log.Fatalln("Can't write", err)
	//}
	//
	//in, err := os.ReadFile("test.bin")
	//if err != nil {
	//	log.Fatalln("Can't read file", err)
	//}
	//
	//readEmployee := &pb.Employee{}
	//
	//err = proto.Unmarshal(in, readEmployee)
	//if err != nil {
	//	log.Fatalln("Can't deserialize", err)
	//}
	//
	//fmt.Println(readEmployee)

	m := jsonpb.Marshaler{}
	out, err := m.MarshalToString(employee)
	if err != nil {
		log.Fatalln("Can't marshal to json", err)
	}

	//fmt.Println(out)

	readEmployee := &pb.Employee{}
	if err := jsonpb.UnmarshalString(out, readEmployee); err != nil {
		log.Fatalln("Can't unmarshal from json", err)
	}

	fmt.Println(readEmployee)
}
