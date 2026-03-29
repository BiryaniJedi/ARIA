package towers

import (
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tower interface {
	Draw(*ebiten.Image)
}

type D1Commit struct {
	Pos            utils.Position
	baseball       Ball
	TargetingRange float32
}

// func (t *TriTower) Draw(screen *ebiten.Image) {
// }
