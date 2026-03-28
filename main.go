package main

import (
	"github.com/BiryaniJedi/ARIA/game"
	"github.com/BiryaniJedi/ARIA/maps"
	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"time"
)

func main() {
	curMap := maps.NewMap(300, 400)
	curMap.AppendPositions(
		shapes.Position{X: 600, Y: 400},
		shapes.Position{X: 300, Y: 0},
		shapes.Position{X: 300, Y: 400},
		shapes.Position{X: 500, Y: 599},
		shapes.Position{X: 100, Y: 100},
	)

	rect := &shapes.Rect{
		Pos:    curMap.StartPos,
		Width:  15,
		Height: 15,
		Color:  color.RGBA{0xFF, 0, 0, 0xFF},
		Speed:  100,
		Target: curMap.Path[1],
	}
	circ := &shapes.Circle{
		CenterPos: curMap.StartPos,
		Radius:    15,
		Color:     color.RGBA{0, 0xFF, 0, 0xFF},
		Speed:     200,
		Target:    curMap.Path[1],
	}

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
