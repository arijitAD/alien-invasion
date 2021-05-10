package invasion

import (
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type alien struct {
	name string
	city *city
}

func (a alien) String() string {
	return a.name
}

// generateAliens generates Aliens with random names.
func generateAliens(n int) map[string]*alien {
	randSrc := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomdata.CustomRand(randSrc)

	aliens := make(map[string]*alien, n)
	for i := 0; i < n; i++ {
		al := &alien{name: randomdata.FirstName(rand.Intn(2))}
		aliens[al.name] = al
	}

	return aliens
}
