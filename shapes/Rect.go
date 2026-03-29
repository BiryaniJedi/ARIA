package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Rect struct {
	pos      utils.Position
	width    float32
	height   float32
	color    color.Color
	progress int
	speed    float32 // per second
	target   utils.Position
	sprite   *ebiten.Image
	rotation float64
}

func (r *Rect) Draw(screen *ebiten.Image) {
	utils.DrawSprite(screen, r.sprite, r.pos, r.rotation)
}

func (r *Rect) GetCurPos() utils.Position               { return r.pos }
func (r *Rect) GetProgress() int                        { return r.progress }
func (r *Rect) IncProgress()                            { r.progress++ }
func (r *Rect) ResetProgress()                          { r.progress = 0 }
func (r *Rect) GetSpeed() float32                       { return r.speed }
func (r *Rect) GetRotation() float64                    { return r.rotation }
func (r *Rect) SetTarget(newPos utils.Position)         { r.target = newPos }
func (r *Rect) RotateToTarget(targetPos utils.Position) { r.rotation = r.pos.AngleToPos(targetPos) }

// Returns true if the rectangle has reached the target position
func (r *Rect) SeekTarget(toMove float32) bool {
	toTarget := utils.Distance(r.pos, r.target)
	if toTarget < EPSILON || toMove >= toTarget {
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

func NewRect(
	width, height float32,
	color color.Color,
	speed float32,
	pos, target utils.Position,
	sprite *ebiten.Image,
) *Rect {
	ret := &Rect{
		pos,
		width,
		height,
		color,
		0,
		speed,
		target,
		sprite,
		0,
	}
	ret.RotateToTarget(ret.target)
	return ret
}
