package towers

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Tower interface {
	Draw(*ebiten.Image)
}

type TriTower struct {
	Pos         utils.Position
	Base        float32
	Height      float32
	Color       color.Color
	RangeRadius float32
	Selected    bool
	Kills       uint32
}

func (t *TriTower) Draw(screen *ebiten.Image) {
}
