package utils

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
)

type Position struct {
	X float32
	Y float32
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
