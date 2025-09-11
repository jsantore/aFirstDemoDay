package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type firstGame struct {
	player *ebiten.Image
	xloc   int
	yloc   int
	score  int
}

func (f firstGame) Update() error {
	return nil
}

func (f firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Aqua)
}

func (f firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	ourGame := firstGame{} //we will use the zero value for now
	err := ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
