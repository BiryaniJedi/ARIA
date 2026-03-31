package towers

import (
	_ "embed"
	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

//go:embed assets/bigCat.png
var spriteBytes []byte
var spriteImg *ebiten.Image

func init() {
	spriteImg = utils.LoadSprite(spriteBytes)
}

type Tower interface {
	Draw(*ebiten.Image)
	GetPos() utils.Position
	GetAngle() float64
	SetAngle(float64)
	IsDeleted() bool
	GetLastShotTime() time.Time
	UpdateLastShotTime()
	ReadyToFire() bool
	GetShapesInRange(*[]shapes.Shape) []shapes.Shape
	GetNewProjPtr(*shapes.Shape) Projectile //
}

type D1Commit struct {
	pos            utils.Position
	sprite         *ebiten.Image
	lastShotTime   time.Time
	angle          float64
	targetingRange float32
	aps            float32 // attacks per second
}

func NewD1Commit(pos utils.Position, radius float32, attacksPerSecond float32) *D1Commit {
	return &D1Commit{
		pos,
		spriteImg,
		time.Now(),
		0,
		radius,
		0.5,
	}
}

func (dc *D1Commit) Draw(screen *ebiten.Image) { utils.DrawSprite(screen, dc.sprite, dc.pos, dc.angle) }

func (dc *D1Commit) GetPos() utils.Position    { return dc.pos }
func (dc *D1Commit) GetAngle() float64         { return dc.angle }
func (dc *D1Commit) SetAngle(newAngle float64) { dc.angle = newAngle }

func (dc *D1Commit) IsDeleted() bool { return false }

func (dc *D1Commit) GetLastShotTime() time.Time { return dc.lastShotTime }
func (dc *D1Commit) UpdateLastShotTime()        { dc.lastShotTime = time.Now() }

func (dc *D1Commit) GetNewProjPtr(targetShapePtr *shapes.Shape) Projectile {
	return NewBaseball(dc, (*targetShapePtr).GetCurPos())
}

func (dc *D1Commit) ReadyToFire() bool {
	if dc.aps < utils.EPSILON {
		return false
	}
	return time.Since(dc.lastShotTime).Seconds() >= float64(dc.aps)
}

// Filters the given list of shapes to only those whose circles intersect
// with the tower (within range for targeting)
func (dc *D1Commit) GetShapesInRange(shapesPtr *[]shapes.Shape) []shapes.Shape {
	return utils.Filter(*shapesPtr, func(shape shapes.Shape) bool {
		return utils.CirclesIntersect(dc.pos, dc.targetingRange, shape.GetCurPos(), shape.GetRadius())
	})
}
