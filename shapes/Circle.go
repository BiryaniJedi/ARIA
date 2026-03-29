package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Circle struct {
	CenterPos utils.Position
	Radius    float32
	Color     color.Color
	Progress  int
	Speed     float32 // per second
	Target    utils.Position
	Sprite    *ebiten.Image
	Rotation  float64
}

func NewCircle(
	radius float32,
	color color.Color,
	speed float32,
	centerPos, target utils.Position,
	sprite *ebiten.Image,
) *Circle {
	circ := &Circle{
		centerPos,
		radius,
		color,
		0,
		speed,
		target,
		sprite,
		0,
	}
	circ.RotateToTarget(circ.Target)
	return circ
}

func (c *Circle) Draw(screen *ebiten.Image) {
	utils.DrawSprite(screen, c.Sprite, c.CenterPos, c.Rotation)
}

func (c *Circle) GetCurPos() utils.Position       { return c.CenterPos }
func (c *Circle) GetProgress() int                { return c.Progress }
func (c *Circle) IncProgress()                    { c.Progress++ }
func (c *Circle) ResetProgress()                  { c.Progress = 0 }
func (c *Circle) GetSpeed() float32               { return c.Speed }
func (c *Circle) GetRotation() float64            { return c.Rotation }
func (c *Circle) SetTarget(newPos utils.Position) { c.Target = newPos }
func (c *Circle) RotateToTarget(targetPos utils.Position) {
	c.Rotation = c.CenterPos.AngleToPos(targetPos)
}

// Returns true if the rectangle has reached the target position
func (c *Circle) SeekTarget(toMove float32) bool {
	toTarget := utils.Distance(c.CenterPos, c.Target)
	if toTarget < EPSILON || toMove >= toTarget {
		c.CenterPos = c.Target
		return true
	}

	ratio := toMove / toTarget
	distX := c.Target.X - c.CenterPos.X
	distY := c.Target.Y - c.CenterPos.Y
	c.CenterPos.X += distX * ratio
	c.CenterPos.Y += distY * ratio
	return false
}
