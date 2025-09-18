package main

import (
	"fmt"
	_ "image/png"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type firstGame struct {
	player    *ebiten.Image
	xloc      int
	yloc      int
	speed     int
	score     int
	treasures []*coinPile
}

type coinPile struct {
	pict *ebiten.Image
	xLoc float64
	yLoc float64
}

func NewCoin(maxX, maxY int, image *ebiten.Image) *coinPile {
	return &coinPile{
		pict: image,
		xLoc: float64(rand.Intn(maxX)),
		yLoc: float64(rand.Intn(maxY)),
	}
}

func (f *firstGame) Update() error {
	f.xloc += f.speed
	if f.xloc > (1000-f.player.Bounds().Dx()) || f.xloc < 0 {
		f.speed = -f.speed
	}
	return nil
}

func (f *firstGame) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Aqua)
	drawOps := &ebiten.DrawImageOptions{}
	for _, pile := range f.treasures {
		drawOps.GeoM.Reset()
		drawOps.GeoM.Translate(pile.xLoc, pile.yLoc)
		screen.DrawImage(pile.pict, drawOps)
	}
	drawOps.GeoM.Reset()
	drawOps.GeoM.Translate(float64(f.xloc), float64(f.yloc))
	screen.DrawImage(f.player, drawOps)
}

func (f *firstGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("First Class Example")
	playerImage, _, err := ebitenutil.NewImageFromFile("ship.png")
	moneyPiles := make([]*coinPile, 0)
	coinImage, _, err := ebitenutil.NewImageFromFile("coins.png")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		moneyPiles = append(moneyPiles, NewCoin(950, 950, coinImage))
	}
	ourGame := firstGame{
		player:    playerImage,
		xloc:      200,
		speed:     3,
		yloc:      400,
		treasures: moneyPiles,
	} //we will use the zero value for now
	err = ebiten.RunGame(&ourGame)
	if err != nil {
		fmt.Println("Failed to run game", err)
	}
}
