package invasion

import (
	"math/rand"
	"time"
)

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	West  Direction = "west"
	East  Direction = "east"
)

var directions = [4]Direction{North, South, West, East}

func (d Direction) isValid() bool {
	switch d {
	case North, South, West, East:
		return true
	}
	return false
}

type city struct {
	name                     string
	alien                    *alien
	north, south, east, west *city
}

func (c city) String() string {
	return c.name
}

// getRandomNeighbourCity randomly selects a neighbour city.
func (c *city) getRandomNeighbourCity() *city {
	rand.Seed(time.Now().UnixNano())

	// Shuffle the direction for random selection.
	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	// Get a non null neighbour city.
	for _, direction := range directions {
		switch direction {
		case North:
			if c.north == nil {
				continue
			}
			return c.north
		case South:
			if c.south == nil {
				continue
			}
			return c.south
		case West:
			if c.west == nil {
				continue
			}
			return c.west
		case East:
			if c.east == nil {
				continue
			}
			return c.east
		}
	}

	return nil
}
