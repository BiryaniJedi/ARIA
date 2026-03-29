package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/BiryaniJedi/ARIA/maps"
	"github.com/hajimehoshi/ebiten/v2"
	"slices"
)

const (
	SCREEN_WIDTH  = 900
	SCREEN_HEIGHT = 600
)

type Game struct {
	StartTime time.Time
	LastTime  time.Time
	CurMap    *maps.Map
}

func (g *Game) Update() error {
	g.updateShapes(time.Since(g.LastTime))
	g.LastTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{230, 185, 161, 255})
	g.CurMap.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func (g *Game) updateShapes(deltaTime time.Duration) {
	shapes := &(g.CurMap.Shapes)
	for i := range len(*shapes) {
		curShape := (*shapes)[i]
		distanceToMove := float32(deltaTime.Seconds()) * curShape.GetSpeed()
		if reached := curShape.SeekTarget(distanceToMove); reached {
			curShape.IncProgress()
			//if shape has reached the final target, delete it
			if curShape.GetProgress() == len(g.CurMap.Path) {
				*shapes = slices.Delete(*shapes, i, i+1)
				fmt.Println("Deleting!!")
				continue
			}
			nextPos := g.CurMap.Path[curShape.GetProgress()]
			fmt.Printf("Next pos: %v\n", nextPos)

			curShape.RotateToTarget(nextPos)
			fmt.Printf("Rotation set to %f\n", curShape.GetRotation())
			curShape.SetTarget(nextPos)
		}
	}
}
