package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shape interface {
	Draw(*ebiten.Image)
	GetCurPos() utils.Position
	GetProgress() int
	IncProgress()
	ResetProgress()
	GetSpeed() float32
	SeekTarget(float32) bool
	SetTarget(utils.Position)
	RotateToTarget(utils.Position)
	GetRotation() float64
	GetCurHp() int32
	TakeDamage(uint32) int32
	HealHP(uint32) int32
	GetRadius() float32
}

// returns pointer to the furthest shape in the given slice
// returns nil if there was no max (given slice empty)
func GetFurthestShape(shapes []Shape) (int, *Shape) {
	idx, furthestPtr := utils.Max(shapes, func(shape Shape) int {
		return shape.GetProgress()
	})
	return idx, furthestPtr
}
