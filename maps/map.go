package maps

import (
	// "github.com/BiryaniJedi/ARIA/game"
	"fmt"
	"image/color"
	"slices"

	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/BiryaniJedi/ARIA/towers"
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Map struct {
	Shapes   []shapes.Shape
	Path     []utils.Position
	StartPos utils.Position
	Towers   []*towers.Tower

	//Maps addresses of towers to slice of pointers of it's active projectiles
	Projectiles map[*towers.Tower][]*towers.Projectile
}

func NewMap(initX, initY float32) *Map {
	initPos := utils.Position{X: initX, Y: initY}
	return &Map{
		StartPos:    initPos,
		Path:        []utils.Position{initPos},
		Projectiles: make(map[*towers.Tower][]*towers.Projectile),
	}
}

func (m *Map) Draw(screen *ebiten.Image) {
	for i := range len(m.Path) - 1 {
		vector.StrokeLine(screen, m.Path[i].X, m.Path[i].Y, m.Path[i+1].X, m.Path[i+1].Y, 3, color.White, false)
	}
	for _, shape := range m.Shapes {
		shape.Draw(screen)
	}
	for _, tower := range m.Towers {
		(*tower).Draw(screen)
		projectiles, ok := m.Projectiles[tower]
		if !ok {
			continue
		}
		for _, projPtr := range projectiles {
			if projPtr != nil {
				(*projPtr).Draw(screen)
			}
		}
	}
}

func (m *Map) PrintProjsForTower(towerPtr *towers.Tower) {
	projectiles, ok := m.Projectiles[towerPtr]
	if !ok {
		fmt.Printf("No entry for tower %v\n", *towerPtr)
		return
	}
	fmt.Printf("Projectiles for tower at %p...\n", towerPtr)
	for _, projPtr := range projectiles {
		if projPtr != nil {
			fmt.Printf("\t- %v\n", *projPtr)
		} else {
			fmt.Println("\t- nil Value!")
		}
	}
}

func (m *Map) AppendPositions(positions ...utils.Position) {
	m.Path = append(m.Path, positions...)
}

func (m *Map) AppendShapes(shapes ...shapes.Shape) {
	m.Shapes = append(m.Shapes, shapes...)
}

func (m *Map) PushTower(tower towers.Tower) {
	m.Towers = append(m.Towers, &tower)
}

func (m *Map) PushProjectilePtr(towerPtr *towers.Tower, projPtr *towers.Projectile) {
	m.Projectiles[towerPtr] = append(m.Projectiles[towerPtr], projPtr)
	// m.PrintProjsForTower(towerPtr)
}

// returns true if shape existed before and was deleted, false if shape didn't exist
func (m *Map) DeleteShapeByPointer(targetShapePtr *shapes.Shape) bool {
	if len(m.Shapes) == 0 {
		return false
	}
	n := len(m.Shapes) - 1
	for i := n; i >= 0; i-- {
		curShapePtr := &(m.Shapes[i])
		if curShapePtr == targetShapePtr {
			m.Shapes = slices.Delete(m.Shapes, i, i+1)
			return true
		}
	}
	return false
}

func (m *Map) DeleteProjectileByPointer(towerPtr *towers.Tower, targetProjPtr *towers.Projectile) bool {
	projPtrs, ok := m.Projectiles[towerPtr]
	if !ok || len(projPtrs) == 0 {
		return false
	}
	n := len(projPtrs) - 1
	for i := n; i >= 0; i-- {
		curProjPtr := projPtrs[i]
		if curProjPtr == targetProjPtr {
			m.Projectiles[towerPtr] = slices.Delete(m.Projectiles[towerPtr], i, i+1)
			return true
		}
	}
	return false
}
