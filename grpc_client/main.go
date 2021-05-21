package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	pb "restrauntpb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type restraunt struct {
	User_Name string
	Address   string
	Phone     string
}

// func main() {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewAddDataClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	// res := pb.Restraunt{}
// 	r, err := c.AddRestraunt(ctx, &pb.Restraunt{Name: "name", Address: "address", Phone: "phonmesadf"})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	fmt.Println(r)
// }

//send data to add new entry
func handler(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//Retrieve data from request and store in variable restraunt
	var res restraunt
	s, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error")
	}
	// fmt.Println(string(s))
	json.Unmarshal([]byte(s), &res)

	//Connect to the server

	defer conn.Close()
	c := pb.NewAdd_RecieveDataClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.AddRestraunt(ctx, &pb.Restraunt{Name: res.User_Name, Address: res.Phone, Phone: res.Phone})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//write result and send result to frontend
	fmt.Println(result)

	fmt.Printf("Name = %s\n", res.User_Name)
	fmt.Printf("Address = %s\n", res.Address)
	fmt.Printf("Phone = %s\n", res.Phone)

}

type response struct {
	array []restraunt
}

//recieve all the entries from the server
func get_data(w http.ResponseWriter, r *http.Request) {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewAdd_RecieveDataClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.GetRestraunt(ctx, &pb.Restraunt_Request{Message: "GET"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//write result and send result to frontend
	res := response{}
	// res := response{}
	err = json.Unmarshal([]byte(result.Message), &res.array)

	// for value := range res.array {
	// p := res.array[value]
	// fmt.Println(p)
	// }

	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// w.WriteHeader(200)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write([]byte(result.Message))
}

func main() {
	// setupRoutes()
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", get_data)
	http.ListenAndServe(":8080", nil)

}
