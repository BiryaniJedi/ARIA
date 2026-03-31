package utils

import (
	"bytes"
	"cmp"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

const (
	SCREEN_WIDTH          = 900
	SCREEN_HEIGHT         = 600
	EPSILON       float32 = 1e-4
)

type Position struct {
	X float32
	Y float32
}

func (p *Position) OutOfBounds() bool {
	if p.X < 0 || p.X > SCREEN_WIDTH {
		return true
	}
	if p.Y < 0 || p.Y > SCREEN_HEIGHT {
		return true
	}
	return false
}

func (p *Position) AngleToPos(pos Position) float64 {
	dY := pos.Y - p.Y
	dX := pos.X - p.X
	return math.Atan2(float64(dY), float64(dX))
}

func Dot(a, b Position) float32 { return a.X*b.X + a.Y*b.Y }

func Distance(p1 Position, p2 Position) float32 {
	dist := math.Sqrt(
		math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2),
	)
	return float32(dist)
}

func Clamp(val, minimum, maximum float32) float32 {
	if val < minimum {
		return minimum
	}
	if val > maximum {
		return maximum
	}
	return val
}

func LoadSprite(data []byte) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func DrawSprite(screen, sprite *ebiten.Image, pos Position, angle float64) {
	w, h := sprite.Bounds().Dx(), sprite.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(angle)
	op.GeoM.Translate(float64(pos.X), float64(pos.Y))
	screen.DrawImage(sprite, op)
}

func CirclesIntersect(pos1 Position, rad1 float32, pos2 Position, rad2 float32) bool {
	return Distance(pos1, pos2) <= rad1+rad2
}

// Generic functional style Filter function. return length of filtered slice
func Filter[T any](xs []T, fn func(T) bool) []T {
	ret := []T{}
	for _, x := range xs {
		if fn(x) {
			ret = append(ret, x)
		}
	}
	return ret
}

// Generic function for getting the max value for some t given a transformation to a comparable
// Used for getting the closest shape to a tower, or a shape that's furthest along, etc.
// Returns index of the max item in the slice, and a pointer to that item
// Returns -1, nil if the length of the input slice is 0
func Max[T any, E cmp.Ordered](xs []T, fn func(T) E) (int, *T) {
	if len(xs) == 0 {
		return -1, nil
	}
	curMaxIdx := 0
	for i := 1; i < len(xs); i++ {
		if fn(xs[i]) > fn(xs[curMaxIdx]) {
			curMaxIdx = i
		}
	}
	return curMaxIdx, &(xs[curMaxIdx])
}
