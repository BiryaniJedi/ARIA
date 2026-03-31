package game

import (
	// "fmt"
	"image/color"
	"time"

	"slices"

	"github.com/BiryaniJedi/ARIA/maps"
	"github.com/BiryaniJedi/ARIA/shapes"
	"github.com/BiryaniJedi/ARIA/towers"
	"github.com/BiryaniJedi/ARIA/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	StartTime time.Time
	LastTime  time.Time
	CurMap    *maps.Map
}

func (g *Game) Update() error {
	deltaTime := time.Since(g.LastTime)
	g.updateTowers(deltaTime)
	g.updateShapes(deltaTime)
	g.LastTime = time.Now()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{230, 185, 190, 255})
	g.CurMap.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return utils.SCREEN_WIDTH, utils.SCREEN_HEIGHT
}

func (g *Game) updateShapes(deltaTime time.Duration) {
	shapes := &(g.CurMap.Shapes)
	if len(*shapes) == 0 {
		return
	}
	n := len(*shapes) - 1
	for i := n; i >= 0; i-- {
		curShape := (*shapes)[i]

		if curShape.GetCurHp() <= 0 {
			*shapes = slices.Delete(*shapes, i, i+1)
			continue
		}

		distanceToMove := float32(deltaTime.Seconds()) * curShape.GetSpeed()
		if reached := curShape.SeekTarget(distanceToMove); reached {
			curShape.IncProgress()
			//if shape has reached the final target, delete it
			if curShape.GetProgress() == len(g.CurMap.Path) {
				*shapes = slices.Delete(*shapes, i, i+1)
				continue
			}
			nextPos := g.CurMap.Path[curShape.GetProgress()]

			curShape.RotateToTarget(nextPos)
			curShape.SetTarget(nextPos)
		}
	}
}

func (g *Game) fireTowerAt(towerPtr *towers.Tower, targetShapePtr *shapes.Shape) {
	//Appends a new projectile (of the appropriate type given the tower) to the map's projectile CurMap
	// and updates the lastshotTime of the tower
	newProj := (*towerPtr).GetNewProjPtr(targetShapePtr)
	g.CurMap.PushProjectilePtr(towerPtr, &newProj)
	(*towerPtr).UpdateLastShotTime()
	// g.CurMap.PrintProjsForTower(towerPtr)

}

func (g *Game) updateTowers(deltaTime time.Duration) {
	towers := &(g.CurMap.Towers)
	n := len(*towers) - 1
	for i := n; i >= 0; i-- {
		curTowerPtr := ((*towers)[i])
		curTowerPos := (*curTowerPtr).GetPos()
		if (*curTowerPtr).IsDeleted() {
			// TODO: Change from despawning projectiles on IsDeleted
			// to something else (maybe checking if parent ptr is nil and deleting
			// on reaching target)
			delete(g.CurMap.Projectiles, curTowerPtr)
			*towers = slices.Delete(*towers, i, i+1)
			continue
		}
		withinRange := (*curTowerPtr).GetShapesInRange(&(g.CurMap.Shapes))
		if len(withinRange) == 0 {
			// no enemies within range
			continue
		}

		_, furthestPtr := shapes.GetFurthestShape(withinRange)

		//rotate tower to look at target
		(*curTowerPtr).SetAngle(curTowerPos.AngleToPos((*furthestPtr).GetCurPos()))

		//if the tower is ready to shoot, fire at the furthestPtr
		// TODO: Various Targeting methods (will be handled in tower method anyways)
		if (*curTowerPtr).ReadyToFire() {
			g.fireTowerAt(curTowerPtr, furthestPtr)
		}

		//Update projectiles:
		projPtrLst, ok := g.CurMap.Projectiles[curTowerPtr]
		if !ok {
			continue
		}
		numProjectiles := len(projPtrLst) - 1
		for i := numProjectiles; i >= 0; i-- {
			curProjPtr := projPtrLst[i]
			if curProjPtr == nil {
				continue
			}
			toMove := deltaTime.Seconds() * float64((*curProjPtr).GetCurVelo())
			if validPos := (*curProjPtr).NextPos(float32(toMove)); !validPos {
				//projectile needs to be despawned
				g.CurMap.Projectiles[curTowerPtr] = slices.Delete(projPtrLst, i, i+1)
			}
			// TODO: Collision Detection stuff
		}
	}
}
