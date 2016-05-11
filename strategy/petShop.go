package strategy

import (
    "reflect"
    "fmt"
)

var (
    //Animals list of strategy implementations
    Animals = make(map[string]reflect.Type)
)

type (
    //PetShop is a strategy controller
    PetShop struct{}
    
    iPetShop interface {
        ChooseAnimal(string)
    }
    
    //IAnimal Strategy Interface
    IAnimal interface {
        Speech()
    }
)

//ChooseAnimal select strategy based on animal type
func (petShop *PetShop) ChooseAnimal(animalType string)  {
    animal := reflect.New(Animals[animalType]).Elem().Interface().(IAnimal)
    fmt.Printf("Found... ")
    animal.Speech()
}

//RegisterAnimal register a new strategy based on animal type
func RegisterAnimal(animalType string, typedNil IAnimal) {
    fmt.Printf("Registering %s\n",animalType)
    t := reflect.TypeOf(typedNil)
    Animals[animalType] = t
    fmt.Printf("Registering %s... %d registered\n",animalType,len(Animals))
}

//NewPetShop basic constructor
func NewPetShop() *PetShop {
    return &PetShop{}
}