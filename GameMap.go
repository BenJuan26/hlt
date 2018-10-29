package hlt

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/BenJuan26/hlt/input"
)

// GameMap - Top level structure for the game map
type GameMap struct {
	width  int
	height int
	Cells  [][]*MapCell
}

// GetWidth returns the map's width
func (gm *GameMap) GetWidth() int {
	return gm.width
}

// GetHeight returns the map's height
func (gm *GameMap) GetHeight() int {
	return gm.height
}

func (gm *GameMap) String() string {
	return fmt.Sprintf("GameMap{height=%d,width=%d,cells=%d}", gm.height, gm.width, len(gm.Cells))
}

// AbsInt returns the absolute value of the given integer
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// CellsByHalite returns a list of MapCells within the given radius
// sorted by Halite in descending order
func (gm *GameMap) CellsByHalite(center *Position, radius int) []*MapCell {
	var list []*MapCell
	for _, row := range gm.Cells {
		for _, cell := range row {
			dist := gm.CalculateDistance(center, cell.Pos)
			if dist <= radius {
				list = append(list, cell)
			}
		}
	}

	sort.Sort(sort.Reverse(ByHalite(list)))

	return list
}

// NewGameMap - Creates an empty map
func NewGameMap(width int, height int) *GameMap {
	cells := make([][]*MapCell, height)
	for i := range cells {
		cells[i] = make([]*MapCell, width)
	}
	return &GameMap{width, height, cells}
}

// AtPosition - Returns the mapcell at the given position
func (gm *GameMap) AtPosition(position *Position) *MapCell {
	return gm.Cells[position.y][position.x]
}

// AtEntity - Returns the mapcell that the entity occupies
func (gm *GameMap) AtEntity(entity *Entity) *MapCell {
	return gm.AtPosition(entity.Pos)
}

// GenerateGameMap - Creates new game map from input data
func GenerateGameMap() *GameMap {
	var input = input.GetInstance()
	var width, _ = input.GetInt()
	var height, _ = input.GetInt()
	var gameMap = NewGameMap(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var halite, _ = input.GetInt()
			gameMap.Cells[y][x] = &MapCell{&Position{x, y}, halite, nil, nil}
		}
	}
	return gameMap
}

// Normalize -
func (gm *GameMap) Normalize(position *Position) *Position {
	return &Position{
		((position.x % gm.width) + gm.width) % gm.width,
		((position.y % gm.height) + gm.height) % gm.height}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// CalculateDistance - Normalizes the data points and then returns the calculated distance between them
func (gm *GameMap) CalculateDistance(source *Position, target *Position) int {
	return gm.calculateDistance(gm.Normalize(source), gm.Normalize(target))
}

func (gm *GameMap) calculateDistance(source *Position, target *Position) int {
	var dx = abs(source.x - target.x)
	var dy = abs(source.y - target.y)
	var toroidalDx = min(dx, gm.width-dx)
	var toroidalDy = min(dy, gm.height-dy)
	return toroidalDx + toroidalDy
}

func (gm *GameMap) at(position *Position) *MapCell {
	return gm.Cells[position.y][position.x]
}

// NaiveNavigate -
func (gm *GameMap) NaiveNavigate(ship *Ship, destination *Position) *Direction {
	var unsafeMoves = gm.GetUnsafeMoves(ship.E.Pos, destination)
	for _, direction := range unsafeMoves {
		var targetPos, _ = ship.E.Pos.DirectionalOffset(direction)
		targetPos = gm.Normalize(targetPos)
		if !gm.at(targetPos).IsOccupied() {
			return direction
		}
	}
	return Still()
}

// GetUnsafeMoves - Returns the list of moves that might result in collisions
func (gm *GameMap) GetUnsafeMoves(source *Position, destination *Position) []*Direction {
	return gm.unsafeMoves(gm.Normalize(source), gm.Normalize(destination))
}

func (gm *GameMap) unsafeMoves(source *Position, destination *Position) []*Direction {
	var dx = abs(source.x - destination.x)
	var dy = abs(source.y - destination.y)
	var wrappedDx = gm.width - dx
	var wrappedDy = gm.height - dy
	var xDirection = Still()
	if source.x < destination.x {
		if dx > wrappedDx {
			xDirection = West()
		} else {
			xDirection = East()
		}
	} else if source.x > destination.x {
		if dx < wrappedDx {
			xDirection = West()
		} else {
			xDirection = East()
		}
	}
	var yDirection = Still()
	if source.y < destination.y {
		if dy > wrappedDy {
			yDirection = North()
		} else {
			yDirection = South()
		}
	} else if source.y > destination.y {
		if dy < wrappedDy {
			yDirection = North()
		} else {
			yDirection = South()
		}
	}

	if rand.Intn(2) == 0 {
		return append(append([]*Direction{}, xDirection), yDirection)
	}

	return append(append([]*Direction{}, yDirection), xDirection)
}

// Update -
func (gm *GameMap) Update() {
	for y := 0; y < gm.height; y++ {
		for x := 0; x < gm.width; x++ {
			gm.Cells[y][x].ship = nil
		}
	}
	var input = input.GetInstance()
	var updateCount, _ = input.GetInt()
	for i := 0; i < updateCount; i++ {
		var x, _ = input.GetInt()
		var y, _ = input.GetInt()
		gm.Cells[y][x].Halite, _ = input.GetInt()
	}
}
