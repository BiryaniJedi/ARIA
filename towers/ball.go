package towers

import (
	// "github.com/BiryaniJedi/ARIA/shapes"
	// "github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type TargetMethod uint8

const (
	FixedPath TargetMethod = iota
	SeekTarget
)

type Projectile interface {
	Draw(*ebiten.Image)
	GetDefaultVelo() int32
	GetCurVelo() int32
	SetVelo(int32)

	// Moves the projectile to it's next position based on
	// it's targeting method and the input float32 (amount to move).
	// Returns false if the projectile needs to be despawned for
	// whatever reason (projectile will handle)
	NextPos(float32) bool
}

type Ball interface {
	GetMaxHits() *uint32
}
