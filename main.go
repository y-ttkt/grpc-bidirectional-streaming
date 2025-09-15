package main

import (
	"fmt"
	"github.com/y-ttkt/grpc-bidirectional-streaming/game"
)

func main() {
	b := game.Board{Cells: make([][]game.Color, 10)}
	for i := 0; i < 10; i++ {
		b.Cells[i] = make([]game.Color, 10)
	}
	for i := 0; i < 10; i++ {
		b.Cells[0][i] = game.Wall
	}
	for i := 1; i < 9; i++ {
		b.Cells[i][0] = game.Wall
		b.Cells[i][9] = game.Wall
	}
	for i := 0; i < 9; i++ {
		b.Cells[9][i] = game.Wall
	}
	fmt.Println(b)
}
