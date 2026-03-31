package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/vector"
)

type Rect struct {
	pos      utils.Position
	target   utils.Position
	sprite   *ebiten.Image
	rotation float64
	progress int
	maxHp    int32
	curHp    int32
	radius   float32
	speed    float32 // per second
}

func NewRect(
	radius, speed float32,
	pos, target utils.Position,
	sprite *ebiten.Image,
	maxHp int32,
	curHp int32,
) *Rect {
	ret := &Rect{
		pos,
		target,
		sprite,
		0,
		0,
		maxHp,
		curHp,
		radius,
		speed,
	}
	ret.RotateToTarget(ret.target)
	return ret
}

func (r *Rect) Draw(screen *ebiten.Image) {
	utils.DrawSprite(screen, r.sprite, r.pos, r.rotation)
}

func (r *Rect) GetCurPos() utils.Position { return r.pos }

func (r *Rect) GetRadius() float32 {
	return r.radius
}

func (r *Rect) GetProgress() int { return r.progress }
func (r *Rect) IncProgress()     { r.progress++ }
func (r *Rect) ResetProgress()   { r.progress = 0 }

func (r *Rect) GetSpeed() float32 { return r.speed }

func (r *Rect) GetRotation() float64 { return r.rotation }

func (r *Rect) SetTarget(newPos utils.Position)         { r.target = newPos }
func (r *Rect) RotateToTarget(targetPos utils.Position) { r.rotation = r.pos.AngleToPos(targetPos) }

// Returns true if the rectangle has reached the target position
func (r *Rect) SeekTarget(toMove float32) bool {
	toTarget := utils.Distance(r.pos, r.target)
	if toTarget < utils.EPSILON || toMove >= toTarget {
		r.pos = r.target
		return true
	}

	ratio := toMove / toTarget
	distX := r.target.X - r.pos.X
	distY := r.target.Y - r.pos.Y
	r.pos.X += distX * ratio
	r.pos.Y += distY * ratio
	return false
}

func (r *Rect) GetCurHp() int32 { return r.curHp }
func (r *Rect) TakeDamage(damage uint32) int32 {
	r.curHp -= int32(damage)
	return r.curHp
}
func (r *Rect) HealHP(healing uint32) int32 {
	r.curHp += int32(healing)
	return r.curHp
}
