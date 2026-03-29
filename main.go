package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/BiryaniJedi/ARIA/game"
	"github.com/BiryaniJedi/ARIA/maps"
	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/enemies/rectSprite.png
var rectSpriteBytes []byte
var rectSpriteImg *ebiten.Image

//go:embed assets/enemies/circSprite.png
var circleSpriteBytes []byte
var circleSpriteImg *ebiten.Image

func init() {
	rectSpriteImg = utils.LoadSprite(rectSpriteBytes)
	circleSpriteImg = utils.LoadSprite(circleSpriteBytes)
}

func main() {
	curMap := maps.NewMap(300, 400)
	curMap.AppendPositions(
		utils.Position{X: 600, Y: 400},
		utils.Position{X: 300, Y: 0},
		utils.Position{X: 300, Y: 400},
		utils.Position{X: 500, Y: 599},
		utils.Position{X: 100, Y: 100},
	)
	rect := shapes.NewRect(
		15,
		100,
		curMap.StartPos,
		curMap.Path[1],
		rectSpriteImg,
		100,
		100,
	)

	circ := shapes.NewCircle(
		15,
		200,
		curMap.StartPos,
		curMap.Path[1],
		circleSpriteImg,
		130,
		130,
	)

	curMap.AppendShapes(rect, circ)

	game := &game.Game{
		StartTime: time.Now(),
		LastTime:  time.Now(),
		CurMap:    curMap,
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("ARIA")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
