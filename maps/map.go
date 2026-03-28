package maps

import (
	// "github.com/BiryaniJedi/ARIA/game"
	"image/color"

	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Map struct {
	Shapes   []shapes.Shape
	Path     []shapes.Position
	StartPos shapes.Position
}

func NewMap(initX, initY float32) *Map {
	initPos := shapes.Position{X: initX, Y: initY}
	return &Map{
		StartPos: initPos,
		Path:     []shapes.Position{initPos},
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	for i := range len(m.Path) - 1 {
		vector.StrokeLine(screen, m.Path[i].X, m.Path[i].Y, m.Path[i+1].X, m.Path[i+1].Y, 3, color.White, false)
	}
	for _, shape := range m.Shapes {
		shape.Draw(screen)
	}
}

func (m *Map) AppendPositions(positions ...shapes.Position) {
	m.Path = append(m.Path, positions...)
}

func (m *Map) AppendShapes(shapes ...shapes.Shape) {
	m.Shapes = append(m.Shapes, shapes...)
}
