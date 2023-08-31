package main

import (
	"context"
	"fmt"
	"log"
	cus "Netxd_Customer/Netxd_Customer" 
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure()) // checking append running
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := cus.NewCustomerServiceClient(conn)
	

	//response, err := client.CreateCustomer(context.Background(), &cus.Customer{FirstName: "GURU",LastName: "Akash",BankId: "ABCBANK",Balance: 55000.00,})
	response,err := client.TransferAmount(context.TODO(),&cus.Transaction{FromCustomer: "64ef0a92e7c1a8ad65217f7c",ToCoustomer: "64f043edc472394a24883e5f",Amount: 500.00})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response)
}
