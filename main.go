package main

import (
	"fmt"

	"github.com/gleisonbs/poker-golang/game/table"
)

func main() {
	fmt.Println((5 + 1) % 6)

	newTable, _ := table.New(6)
	newTable.SitPlayers()
	newTable.DealStartingHands()
	newTable.CollectBlinds()

}
