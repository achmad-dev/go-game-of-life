package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/achmad-dev/go-game-of-life/internal/utils"
)

func main() {
	fmt.Println("game of life graphic in go using conway algorithm")
	game := utils.NewGameOfLife()
	game.Init()
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rname := rd.Float64()
	name := fmt.Sprintf("%v.gif", rname)
	if err := game.GenerateGIF(name, 100, 50); err != nil {
		log.Fatal(err)
	}
}
