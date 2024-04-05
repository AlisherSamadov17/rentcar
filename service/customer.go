package service

import (
	"context"
	"fmt"
	"log"
	"rent-car/api/models"
	"rent-car/storage"
)


type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{
		storage: storage,
	}
}

func (cs customerService) Create(ctx context.Context, customer models.Customer) (string,error) {
	pkey,err := cs.storage.Customer().Create(ctx,customer)
	if err != nil {
		log.Println("ERROR in service layer while creating customer", err.Error())
		fmt.Println("ERROR in service layer while creating customer", err.Error())
		return "", err
	}
	return pkey,nil
}

func (cs customerService) Update(ctx context.Context, customer models.Customer) (string,error) {
	pkey, err := cs.storage.Customer().UpdateCustomer(ctx,customer)
	if err != nil {
		log.Println("ERROR in service layer while updating customer",err.Error())
		return "",err
	}
	return pkey,nil
}

func (cs customerService) GetByIDCustomer(ctx context.Context, id string) (models.Customer,error) {
	customer, err := cs.storage.Customer().GetByID(ctx,id)
	if err != nil {
		log.Println("ERROR in service layer while getting by id customer",err.Error())
		return models.Customer{},err
	}
	return customer,nil
}

func (cs customerService) Delete(ctx context.Context, id string) (error) {
	err := cs.storage.Customer().Delete(ctx,id)
	if err != nil {
		log.Println("ERROR in service layer while deleting customer",err.Error())
		return err
	}
	return nil
}

func (cs customerService) GetCustomerAll(ctx context.Context,customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	customers, err := cs.storage.Customer().GetAllCustomer(ctx,customer)
	if err != nil {
		log.Println("ERROR in service layer while getting all customers",err.Error())
		return customers,err
	}
	return customers,nil
}

func (u customerService) UpdatePassword(ctx context.Context, customer models.PasswordOfCustomer) (string, error) {

	pKey, err := u.storage.Customer().UpdateCustomerPassword(ctx, customer)
	if err != nil {
		fmt.Println("ERROR in service layer while updating Customer", err.Error())
		return "", err
	}

	return pKey, nil
}