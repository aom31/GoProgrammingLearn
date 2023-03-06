package main

import (
	"fmt"
	"log"
)

type Citizen struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	CitizenID string `json:"citizen_id"`
}

//interface will have method , input request to interface will work with method they have

// 1 declare interface type and in interface can have method to do something
// ICitizenService is the interface for citizen service
type ICitizenService interface {
	//can add method for interface
	Validate(c *Citizen) bool
	CreateCitizenCard(c *Citizen) error
}

// 2 declare stuct for represent object
// the service to work with citizen data
type ThaiCitizenService struct {
}

// 4 create constructor for CitizenService use New+ Name-struct-object
// constructor is  function use for create object
func NewThaiCitizenService() *ThaiCitizenService {
	//So constructor will return pointer to own struct
	return &ThaiCitizenService{}
}

//3 declare function method in interface
//And function in interface will be method of struct
//use func (struct for object) Namefunction(input) output{}

// เป็นวิธีการ ผูกฟังก์ชั่น เพื่อเป็น method ของ struct
func (svc *ThaiCitizenService) Validate(c *Citizen) bool {
	return len(c.Firstname) > 0 && len(c.Lastname) > 0 && len(c.CitizenID) > 0

}

func (svc *ThaiCitizenService) CreateCitizenCard(c *Citizen) error {
	//TODO: request API to create citizen card
	fmt.Printf("Successfully create thai citizen card for ID=%s \n", c.CitizenID)
	return nil
}

//5

func main() {

	//step1 constructor function will return struct pointer
	citizen := NewCitizen("Thamakorn", "ketnoi", "1249949492009")
	citizenSvc := NewThaiCitizenService()

	//step2 function that accept interface, the argrument must implement function declare in interface
	err := CreateCitizenCard(citizenSvc, citizen)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCitizenCard(svc ICitizenService, c *Citizen) error {
	if !svc.Validate(c) {
		return fmt.Errorf("Citizen data is invalid")
	}
	err := svc.CreateCitizenCard(c)
	if err != nil {
		return err
	}
	return nil
}

// NewCitizen is constructor function for Citizen
func NewCitizen(firstname, lastname, citizenID string) *Citizen {
	return &Citizen{
		Firstname: firstname,
		Lastname:  lastname,
		CitizenID: citizenID,
	}
}
