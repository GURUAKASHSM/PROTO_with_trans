package customerservice

import (
	customerinterface "Netxd_Customer/Customer_DAL/Customer_Interface"
	customermodel "Netxd_Customer/Customer_DAL/Customer_Model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) customerinterface.ICustomer {
	return &CustomerService{collection, ctx}
}
func (p *CustomerService) CreateCustomer(customer *customermodel.Customer) (*customermodel.CustomerResponse, error) {
	customer.CreatedAt = time.Now().Format(time.Kitchen)
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = true
	res, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted", res.InsertedID)
	response := &customermodel.CustomerResponse{
		Customer_Id: res.InsertedID.(primitive.ObjectID),
		CreatedAt:   customer.CreatedAt,
	}
	fmt.Println(response)
	return response, nil
}
type BankAccount struct {
	balance1 float64
	balance2 float64
}
func (p *CustomerService)  TransferAmount(transaction *customermodel.Tranaction) ( *customermodel.TranactionResponse, error){
	var acc BankAccount // Initialize BankAccount outside the function

	id1, err := primitive.ObjectIDFromHex(transaction.From_coustomer)
	if err != nil {
		log.Fatal(err)
	}
	id2, err := primitive.ObjectIDFromHex(transaction.To_coustomer)
	if err != nil {
		log.Fatal(err)
	}
	response := &customermodel.TranactionResponse{
        From_coustomer: id1.String(),
		To_coustomer: id2.String(),
        Amount:        555.0,
    }
	filter1 := bson.M{"_id": id1}
	var account customermodel.Customer
	err = p.CustomerCollection.FindOne(context.Background(), filter1).Decode(&account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Senders Current balance: %.2f\n", account.Balance)
	fmt.Printf("Debiting Amount %.2f\n", transaction.Amount)
	if account.Balance <transaction.Amount {
		fmt.Println("Insufficient Funds")
		return response,nil
	}
	acc.balance1 = account.Balance - transaction.Amount // Initialize acc's balance
	fmt.Printf("Senders New balance: %.2f\n", acc.balance1)
	update1 := bson.M{"$set": bson.M{"balance": acc.balance1}}
	p.CustomerCollection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		log.Fatal(err)
	}
	filter2 := bson.M{"_id": id2}

	err = p.CustomerCollection.FindOne(context.Background(), filter2).Decode(&account)
	if err != nil {
		log.Fatal(err)
	}
	acc.balance2 = account.Balance + transaction.Amount

	fmt.Printf("Receiver Current balance: %.2f\n", account.Balance)
	fmt.Printf("Crediting %.2f\n", transaction.Amount)
	fmt.Printf("Receiver New balance: %.2f\n", acc.balance2)
    fmt.Println()
    fmt.Println()
	update2 := bson.M{"$set": bson.M{"balance": acc.balance2}}
	p.CustomerCollection.UpdateOne(context.Background(), filter2, update2)
	if err != nil {
		log.Fatal(err)
	}
	
	return response,nil
}
