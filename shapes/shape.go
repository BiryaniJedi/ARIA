package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	EPSILON float32 = 1e-4
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
