package cli

import (
    "github.com/coichedid/StrategyPattern.Go/strategy"
    "github.com/coichedid/StrategyPattern.Go/strategy/animals"
    
    "fmt"
)

type (
    //Cli command line interface
    Cli struct {}
    
    iCli interface {
        ChooseAnimal(string)
    }
)

//ChooseAnimal choose an animal from Strategy Controller
func (cli *Cli) ChooseAnimal(animalType string)  {
    fmt.Println(strategy.Animals)
    petShop := strategy.NewPetShop()
    fmt.Printf("Finding animal '%s'\n", animalType)
    petShop.ChooseAnimal(animalType)
}

//NewCli basic constructor
func NewCli() *Cli {
    animals.ImportAnimals()
    fmt.Println("Animals Registered.")
    return &Cli{}
}