package animals

import (
    "github.com/coichedid/StrategyPattern.Go/strategy"
    
    "fmt"
)

type (
    //Cat is the cat strategy
    Cat struct {}
)

//Speech strategy implementation
func (cat *Cat) Speech()  {
    fmt.Println("Miau Miau Miau")
}

func init()  {
    strategy.RegisterAnimal("cat", (*Cat)(nil))
    fmt.Printf("Cat registered.")
}