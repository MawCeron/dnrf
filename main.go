package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	posX   float32
	posY   float32
	offset float32
}

const roadLeftLimit = 103
const roadRightLimit = 197

type Dimensions struct {
	width  float32
	height float32
}

var auto = Dimensions{width: 10, height: 20}

var roadWidth = float32(roadRightLimit+auto.width) - roadLeftLimit
var roadCenter = roadLeftLimit + (roadWidth / 2)
var laneWidth = roadCenter - roadLeftLimit
var laneCenter = roadCenter + (laneWidth / 2)

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.posX > roadLeftLimit {
		g.posX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.posX < roadRightLimit {
		g.posX += 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Limites carretera
	vector.FillRect(screen, roadLeftLimit-2, 0, 2, 240, color.RGBA{R: 100, G: 100, B: 100, A: 255}, false)
	vector.FillRect(screen, roadRightLimit+auto.width, 0, 2, 240, color.RGBA{R: 100, G: 100, B: 100, A: 255}, false)

	var posInicial = float32(5)
	for i := 0; i < 11; i++ {
		vector.FillRect(screen, roadCenter-2, posInicial, 2, 10, color.RGBA{R: 100, G: 100, B: 100, A: 255}, false)
		posInicial += 30
	}

	// Auto
	vector.FillRect(screen, g.posX, g.posY, auto.width, auto.height, color.RGBA{R: 255, G: 255, B: 255, A: 255}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Definitely Not Road Fighter")

	// Posicion Inicial
	g := &Game{
		posX: laneCenter,
		posY: 197,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
