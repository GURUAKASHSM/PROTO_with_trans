package controller

import (
	customerinterface "Netxd_Customer/Customer_DAL/Customer_Interface"
	customermodel "Netxd_Customer/Customer_DAL/Customer_Model"
	Cus "Netxd_Customer/Netxd_Customer"
	"context"
	"log"
)

type RPCServer struct {
	Cus.UnimplementedCustomerServiceServer
}

var (
	CustomerService    customerinterface.ICustomer
	TransactionService customerinterface.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *Cus.Customer) (*Cus.CustomerResponse, error) {
	dbProfile := &customermodel.Customer{FirstName: req.FirstName, LastName: req.LastName, Bank_id: req.BankId, Balance: float64(req.Balance)}
	result, err := CustomerService.CreateCustomer(dbProfile)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &Cus.CustomerResponse{
			FirstName: result.FirstName,
		}
		return responseCustomer, nil
	}
}
func (s *RPCServer) TransferAmount(ctx context.Context, req *Cus.Transaction) (*Cus.TransactionResponse, error) {
	dbTransaction := &customermodel.Tranaction{From_coustomer: req.FromCustomer, To_coustomer: req.ToCoustomer, Amount: float64(req.Amount)}
	response, err := CustomerService.TransferAmount(dbTransaction)
	if err != nil {
		log.Fatal(err)
	}
	result := &Cus.TransactionResponse{
		FromCustomer: response.From_coustomer,
		ToCoustomer:  response.To_coustomer,
		Amount:       float32(response.Amount),
	}
	return result, err

}
