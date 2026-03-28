package shapes

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	EPSILON float32 = 1e-4
)

type Shape interface {
	Draw(*ebiten.Image)
	GetCurPos() Position
	GetProgress() int
	IncProgress()
	ResetProgress()
	GetSpeed() float32
	SeekTarget(float32) bool
	SetTarget(Position)
}

type Position struct {
	X float32
	Y float32
}

func distance(p1 Position, p2 Position) float32 {
	dist := math.Sqrt(
		math.Pow(math.Abs(float64(p2.X-p1.X)), 2) + math.Pow(math.Abs(float64(p2.Y-p1.Y)), 2),
	)
	return float32(dist)
}

type Rect struct {
	Pos      Position
	Width    float32
	Height   float32
	Color    color.Color
	Progress int
	Speed    float32 // per second
	Target   Position
}

func (r *Rect) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, r.Pos.X, r.Pos.Y, r.Width, r.Height, r.Color, false)
}

func (r *Rect) GetCurPos() Position       { return r.Pos }
func (r *Rect) GetProgress() int          { return r.Progress }
func (r *Rect) IncProgress()              { r.Progress++ }
func (r *Rect) ResetProgress()            { r.Progress = 0 }
func (r *Rect) GetSpeed() float32         { return r.Speed }
func (r *Rect) SetTarget(newPos Position) { r.Target = newPos }

// Returns true if the rectangle has reached the target position
func (r *Rect) SeekTarget(toMove float32) bool {
	toTarget := distance(r.Pos, r.Target)
	if toTarget < EPSILON || toMove >= toTarget {
		r.Pos = r.Target
		return true
	}

	ratio := toMove / toTarget
	distX := r.Target.X - r.Pos.X
	distY := r.Target.Y - r.Pos.Y
	r.Pos.X += distX * ratio
	r.Pos.Y += distY * ratio
	return false
}

type Circle struct {
	CenterPos Position
	Radius    float32
	Color     color.Color
	Progress  int
	Speed     float32 // per second
	Target    Position
}

func (c *Circle) Draw(screen *ebiten.Image) {
	vector.FillCircle(screen, c.CenterPos.X, c.CenterPos.Y, c.Radius, c.Color, false)
}

func (c *Circle) GetCurPos() Position       { return c.CenterPos }
func (c *Circle) GetProgress() int          { return c.Progress }
func (c *Circle) IncProgress()              { c.Progress++ }
func (c *Circle) ResetProgress()            { c.Progress = 0 }
func (c *Circle) GetSpeed() float32         { return c.Speed }
func (c *Circle) SetTarget(newPos Position) { c.Target = newPos }

// Returns true if the rectangle has reached the target position
func (c *Circle) SeekTarget(toMove float32) bool {
	toTarget := distance(c.CenterPos, c.Target)
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
