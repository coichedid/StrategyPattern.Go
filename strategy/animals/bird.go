package animals

import (
    "github.com/coichedid/StrategyPattern.Go/strategy"
    
    "fmt"
)

type (
    //Bird is the bird strategy
    Bird struct {}
)

//Speech strategy implementation
func (bird *Bird) Speech()  {
    fmt.Println("Piu Piu Piu")
}

func init()  {
    strategy.RegisterAnimal("bird", (*Bird)(nil))
    fmt.Printf("Bird registered.")
}