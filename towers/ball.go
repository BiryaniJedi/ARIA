package towers

import (
	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/BiryaniJedi/ARIA/utils"
)

type Ball interface {
	StartFrom(utils.Position, utils.Position)
	GetMaxHits() uint32
	HitShape(*shapes.Shape) int32
}
