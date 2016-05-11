package animals

import (
    "github.com/coichedid/StrategyPattern.Go/strategy"
    
    "fmt"
)

type (
    //Dog is the dog strategy
    Dog struct {}
)

//Speech strategy implementation
func (dog *Dog) Speech()  {
    fmt.Println("Au Au Au")
}

func init()  {
    strategy.RegisterAnimal("dog", (*Dog)(nil))
    fmt.Printf("Dog registered.")
}