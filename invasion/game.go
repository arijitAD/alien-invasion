package invasion

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

const totalMoves = 10000

type Game struct {
	cityMap  map[string]*city
	alienMap map[string]*alien
	fileChan chan string
}

func NewGame() *Game {
	return &Game{
		cityMap:  make(map[string]*city),
		alienMap: make(map[string]*alien),
		fileChan: make(chan string, 1000),
	}
}

func (g *Game) Init(path string, alienCount int) error {
	var wg errgroup.Group

	// Read data from file.
	wg.Go(func() error {
		defer func() {
			close(g.fileChan)
		}()
		return asyncLineReader(path, g.fileChan)
	})

	// Process city values and store it in game.
	wg.Go(func() error {
		return g.processCity()
	})

	if err := wg.Wait(); err != nil {
		return fmt.Errorf("failed to init game: %w", err)
	}

	if len(g.cityMap) < alienCount {
		return fmt.Errorf("invalid input: number of aliens(%d) cannot be greater than city(%d)", alienCount, len(g.cityMap))
	}

	g.assignAliens(alienCount)
	return nil
}

// Start will start moving all the aliens in the game
func (g *Game) Start() {
	// aliens can only move max upto 10000 moves
	for i := 0; i < totalMoves; i++ {

		// if all aliens are killed then game will stop
		if len(g.alienMap) == 0 {
			break
		}

		for _, alien := range g.alienMap {
			randCity := alien.city.getRandomNeighbourCity()

			// This is only possible when the alien is trapped.
			if randCity == nil {
				continue
			}

			alien.city.alien = nil

			if randCity.alien == nil {
				randCity.alien = alien
				alien.city = randCity
				continue
			}

			g.fightAliens(randCity, alien, randCity.alien)
		}
	}
}

func (g *Game) PrintState() {
	for _, ct := range g.cityMap {
		var msg []string
		msg = append(msg, ct.name)
		if ct.north != nil {
			msg = append(msg, fmt.Sprintf("north=%s", ct.north))
		}
		if ct.south != nil {
			msg = append(msg, fmt.Sprintf("south=%s", ct.south))
		}
		if ct.west != nil {
			msg = append(msg, fmt.Sprintf("west=%s", ct.west))
		}
		if ct.east != nil {
			msg = append(msg, fmt.Sprintf("east=%s", ct.east))
		}

		log.Println(strings.Join(msg, " "))
	}
}

func (g *Game) assignAliens(alienCount int) {
	g.alienMap = generateAliens(alienCount)

	var cities []*city
	for _, city := range g.cityMap {
		if len(cities) == len(g.alienMap) {
			break
		}
		cities = append(cities, city)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cities), func(i, j int) { cities[i], cities[j] = cities[j], cities[i] })

	var idx int
	for _, alien := range g.alienMap {
		alien.city = cities[idx]
		cities[idx].alien = alien
		log.Printf("alien %s started in city %s", alien,  cities[idx])
		idx++
	}
}

func (g *Game) processCity() error {
	for v := range g.fileChan {
		if v == "" {
			continue
		}

		tokens := strings.Split(v, " ")
		city := g.getOrCreateCity(tokens[0])

		if len(tokens) == 1 {
			continue
		}

		for _, token := range tokens[1:] {
			neighbourCity := strings.Split(token, "=")
			if len(neighbourCity) != 2 {
				return fmt.Errorf("incorrect data format for city direction %s", v)
			}

			direction := Direction(neighbourCity[0])
			if !direction.isValid() {
				return fmt.Errorf("invalid direction %s", neighbourCity[0])
			}

			g.setNeighbourCity(city, direction, neighbourCity[1])
		}
	}
	return nil
}

func (g *Game) getOrCreateCity(cityName string) *city {
	if city, ok := g.cityMap[cityName]; ok {
		return city
	}

	newCity := &city{name: cityName}
	g.cityMap[cityName] = newCity
	return newCity
}

func (g *Game) setNeighbourCity(c *city, d Direction, cityName string) {
	switch d {
	case North:
		ct := g.getOrCreateCity(cityName)
		c.north = ct
		ct.south = c
	case South:
		ct := g.getOrCreateCity(cityName)
		c.south = ct
		ct.north = c
	case West:
		ct := g.getOrCreateCity(cityName)
		c.west = ct
		ct.east = c
	case East:
		ct := g.getOrCreateCity(cityName)
		c.east = ct
		ct.west = c
	}
}

func (g *Game) fightAliens(city *city, alien1, alien2 *alien) {
	log.Printf("%s has been destroyed by alien %s and %s \n", city, alien1, alien2)

	if city.east != nil {
		city.east.west = nil
	}

	if city.north != nil {
		city.north.south = nil
	}

	if city.west != nil {
		city.west.east = nil
	}

	if city.south != nil {
		city.south.north = nil
	}

	delete(g.cityMap, city.name)
	delete(g.alienMap, alien1.name)
	delete(g.alienMap, alien2.name)
}