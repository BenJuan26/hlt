package hlt

import (
	"fmt"

	"github.com/BenJuan26/hlt/gameconfig"
	"github.com/BenJuan26/hlt/input"
)

// Entity - Base entity structure
type Entity struct {
	id       int
	playerID int
	Pos      *Position
}

// GetID returns the entity ID
func (e *Entity) GetID() int {
	return e.id
}

// GetPlayerID returns the entity player ID
func (e *Entity) GetPlayerID() int {
	return e.playerID
}

/*********************************************************************************/

// Dropoff - Dropoff structure
type Dropoff struct {
	E *Entity
}

// NewDropoff - Creates and initializes a new Dropoff
func NewDropoff(playerID int) *Dropoff {
	var input = input.GetInstance()
	var dropoffID, _ = input.GetInt()
	var x, _ = input.GetInt()
	var y, _ = input.GetInt()
	return &Dropoff{&Entity{dropoffID, playerID, &Position{x, y}}}
}

/*********************************************************************************/

// Ship - Ship structure
type Ship struct {
	E      *Entity
	Halite int
}

// GetID returns the ID of the ship's entity
func (s *Ship) GetID() int {
	return s.E.GetID()
}

// GetPlayerID returns the Player ID of the ship's entity
func (s *Ship) GetPlayerID() int {
	return s.E.GetPlayerID()
}

func (s *Ship) String() string {
	return fmt.Sprintf("Ship{id=%d,%s,halite=%d}", s.GetID(), s.E.Pos, s.Halite)
}

// NewShip - Creates a new ship
func NewShip(playerID int) *Ship {
	var input = input.GetInstance()
	var shipID, _ = input.GetInt()
	var x, _ = input.GetInt()
	var y, _ = input.GetInt()
	var halite, _ = input.GetInt()
	return &Ship{&Entity{shipID, playerID, &Position{x, y}}, halite}
}

// IsFull - Returns true if the ship is full
func (s *Ship) IsFull() bool {
	var maxHalite, _ = gameconfig.GetInstance().GetInt(gameconfig.MaxHalite)
	return s.Halite > maxHalite
}

// MakeDropoff - Creates command to turn the ship into a dropoff
func (s *Ship) MakeDropoff() Command {
	return &TransformToDropoff{s.E.id}
}

// Move - Creates command to move ship
func (s *Ship) Move(d *Direction) Command {
	return &Move{s.E.id, d}
}

// StayStill - Creates command to keep ship still
func (s *Ship) StayStill() Command {
	return s.Move(Still())
}

/*********************************************************************************/

// Shipyard -
type Shipyard struct {
	E *Entity
}

// NewShipyard - Creates a new Shipyard
func NewShipyard(playerID int, position *Position) *Shipyard {
	return &Shipyard{&Entity{-1, playerID, position}}
}
