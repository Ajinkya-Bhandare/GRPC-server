/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	pb "restrauntpb"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.

type server struct {
	pb.UnimplementedAdd_RecieveDataServer
}

type restraunt struct {
	User_name string
	Address   string
	Phone     string
}

// SayHello implements helloworld.GreeterServer
func (s *server) AddRestraunt(ctx context.Context, in *pb.Restraunt) (*pb.InsertResult, error) {
	log.Printf("Received: %v", in)

	fmt.Printf("Name = %s\n", in.Name)
	fmt.Printf("Address = %s\n", in.Address)
	fmt.Printf("Phone = %s\n", in.Phone)

	//Create and insert values in database
	// res := restraunt{User_name: in.Name, Address: in.Address, Phone: in.Phone}
	// db.Create(&res)

	return &pb.InsertResult{Result: true}, nil
}

func (s *server) GetRestraunt(ctx context.Context, in *pb.Restraunt_Request) (*pb.Restraunt_Request, error) {

	fmt.Println("Processing Data")
	dsn := "username:password@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("error")
	}

	//Retrieve all values from database
	var products []restraunt
	db.Find(&products)

	// Print all Restraunts
	// for value := range products {
	// 	p := products[value]
	// 	fmt.Println(p)}

	// Convert object to JSON to send message to client
	urlsJson, _ := json.Marshal(products)
	// fmt.Println(string(urlsJson))
	urls := &pb.Restraunt_Request{Message: string(urlsJson)}
	return urls, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdd_RecieveDataServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
