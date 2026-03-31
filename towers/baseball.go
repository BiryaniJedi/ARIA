package towers

import (
	_ "embed"
	// "fmt"
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

//go:embed assets/bigBaseball.png
var ballSriteBytes []byte
var ballSpriteImg *ebiten.Image

func init() {
	ballSpriteImg = utils.LoadSprite(ballSriteBytes)
}

// implements types.Ball and types.Projectile interfaces
type Baseball struct {
	parent *Tower

	sprite *ebiten.Image

	curPos utils.Position

	lifespanSecs  *uint32
	spawnTime     *time.Time
	maxHits       *uint32
	hitsRemaining *uint32

	angle float64

	targetingMethod TargetMethod
	initRatio       float32 //ratio of initial distanceToMove (speed*)
	dx              float32
	dy              float32

	damage  uint32
	defVelo int32
	curVelo int32
}

func NewBaseball(parent Tower, target utils.Position) *Baseball {
	parentPos := parent.GetPos()
	return &Baseball{
		parent:          &parent,
		curPos:          parentPos,
		lifespanSecs:    new(uint32(4)),
		spawnTime:       new(time.Now()),
		maxHits:         new(uint32(1)),
		hitsRemaining:   new(uint32(1)),
		angle:           parentPos.AngleToPos(target),
		targetingMethod: FixedPath,
		dx:              target.X - parentPos.X,
		dy:              target.Y - parentPos.Y,
		damage:          100,
		defVelo:         500,
		curVelo:         500,
		sprite:          ballSpriteImg,
	}
}

func (b *Baseball) Draw(screen *ebiten.Image) {
	utils.DrawSprite(screen, b.sprite, b.curPos, b.angle)
}

// Projectile interface methods
func (b *Baseball) GetDefaultVelo() int32 { return b.defVelo }
func (b *Baseball) GetCurVelo() int32     { return b.curVelo }
func (b *Baseball) SetVelo(newVelo int32) { b.curVelo = newVelo }

// Moves the Baseball to it's next position, returns true if good or false if needs to be despawned
func (b *Baseball) NextPos(toMove float32) bool {
	hypotenuse := math.Sqrt(math.Pow(float64(b.dx), 2) + math.Pow(float64(b.dy), 2))
	ratio := toMove / float32(hypotenuse)
	b.curPos.X += b.dx * ratio
	b.curPos.Y += b.dy * ratio
	return !b.curPos.OutOfBounds()
}

// Ball interface methods
func (b *Baseball) GetMaxHits() *uint32 { return b.maxHits }
