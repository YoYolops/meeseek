package main

import (
	"fmt"

	"github.com/yoyolops/meeseek/keeper"
)

func main() {
	keeper, err := keeper.SetupKeeper()
	if err != nil {
		fmt.Print(err)
		panic("PANICKING DUE TO ERROR WHILE SETUPING MEESEEK KEEPER")
	}

	keeper.Retrieve()
}
