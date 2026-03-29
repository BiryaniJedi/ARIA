package shapes

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	Pos      utils.Position
	Target   utils.Position
	Sprite   *ebiten.Image
	Rotation float64
	Progress int
	MaxHp    int32
	CurHp    int32
	Radius   float32
	Speed    float32 // per second
}

func NewCircle(
	radius float32,
	speed float32,
	centerPos, target utils.Position,
	sprite *ebiten.Image,
	maxHp int32,
	curHp int32,
) *Circle {
	circ := &Circle{
		centerPos,
		target,
		sprite,
		0,
		0,
		maxHp,
		curHp,
		radius,
		speed,
	}
	circ.RotateToTarget(circ.Target)
	return circ
}

func (c *Circle) Draw(screen *ebiten.Image) {
	utils.DrawSprite(screen, c.Sprite, c.Pos, c.Rotation)
}

func (c *Circle) GetCurPos() utils.Position { return c.Pos }

func (c *Circle) GetRadius() float32 {
	return c.Radius
}

func (c *Circle) GetProgress() int { return c.Progress }
func (c *Circle) IncProgress()     { c.Progress++ }
func (c *Circle) ResetProgress()   { c.Progress = 0 }

func (c *Circle) GetSpeed() float32 { return c.Speed }

func (c *Circle) GetRotation() float64 { return c.Rotation }

func (c *Circle) SetTarget(newPos utils.Position) { c.Target = newPos }
func (c *Circle) RotateToTarget(targetPos utils.Position) {
	c.Rotation = c.Pos.AngleToPos(targetPos)
}

// Returns true if the rectangle has reached the target position
func (c *Circle) SeekTarget(toMove float32) bool {
	toTarget := utils.Distance(c.Pos, c.Target)
	if toTarget < EPSILON || toMove >= toTarget {
		c.Pos = c.Target
		return true
	}

	ratio := toMove / toTarget
	distX := c.Target.X - c.Pos.X
	distY := c.Target.Y - c.Pos.Y
	c.Pos.X += distX * ratio
	c.Pos.Y += distY * ratio
	return false
}

func (c *Circle) GetCurHp() int32 { return c.CurHp }
func (c *Circle) TakeDamage(damage uint32) int32 {
	c.CurHp -= int32(damage)
	return c.CurHp
}
func (c *Circle) HealHP(healing uint32) int32 {
	c.CurHp += int32(healing)
	return c.CurHp
}
