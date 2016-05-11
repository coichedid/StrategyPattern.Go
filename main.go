package main

import (
    "github.com/coichedid/StrategyPattern.Go/cli"
    
    "os"
    "fmt"
)

func main()  {
    if len(os.Args) < 2 {
        fmt.Println("Choose an animal")
        fmt.Println("StrategyPattern.Go <animal>")
        os.Exit(-1)
    }
    cli := cli.NewCli()
    cli.ChooseAnimal(os.Args[1])
    os.Exit(1)
}