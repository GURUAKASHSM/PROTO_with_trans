package main

import (
	customerconfig "Netxd_Customer/Customer_Connection/Customer_DAL_config"
	customerconstants "Netxd_Customer/Customer_Connection/Customer_DAL_constants"
	customerservice "Netxd_Customer/Customer_DAL/Customer_Service"
	controller "Netxd_Customer/Netxd_Controllers"
	cus "Netxd_Customer/Netxd_Customer"
	"context"
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := customerconfig.Collection
	controller.CustomerService = customerservice.InitCustomerService(CustomerCollection, context.Background())
	TransactionCollection := customerconfig.Collection
	controller.TransactionService = customerservice.InitCustomerService(TransactionCollection,context.Background())
}
func main() {
	mongoclient, err := customerconfig.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", customerconstants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	cus.RegisterCustomerServiceServer(s, &controller.RPCServer{})

	fmt.Println("Server listening on", customerconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
