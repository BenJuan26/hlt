package hlt

import "fmt"

// MapCell - Position on a map
type MapCell struct {
	Pos       *Position
	Halite    int
	ship      *Ship
	structure *Entity
}

// ByHalite implements sort.Interface to sort MapCell objects by their Halite
type ByHalite []*MapCell

func (s ByHalite) Len() int           { return len(s) }
func (s ByHalite) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByHalite) Less(i, j int) bool { return s[i].Halite < s[j].Halite }

func (m *MapCell) String() string {
	return fmt.Sprintf("Map{halite=%d}", m.Halite)
}

// IsEmpty - Returns true if no structure nor ship is present
func (m *MapCell) IsEmpty() bool {
	return m.ship == nil && m.structure == nil
}

// IsOccupied - Returns true if ship is present
func (m *MapCell) IsOccupied() bool {
	return m.ship != nil
}

// HasStructure - Returns true if structure is present
func (m *MapCell) HasStructure() bool {
	return m.structure != nil
}

// MarkUnsafe - Marks a mapcell as unsafe, occupying it with a ship
func (m *MapCell) MarkUnsafe(s *Ship) {
	m.ship = s
}
