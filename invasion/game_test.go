package invasion

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame(t *testing.T) {
	invasion := NewGame()

	log.Println("initializing game")
	err := invasion.Init("../data/world-map.txt", 2)
	require.NoError(t, err)

	log.Println("Starting game")
	invasion.Start()

	log.Println("Game ended")
	invasion.PrintState()
}

func TestDirection(t *testing.T) {
	tc := []struct {
		value string
		valid bool
	}{
		{"north", true},
		{"south", true},
		{"west", true},
		{"east", true},
		{"random", false},
		{"", false},
	}

	for _, test := range tc {
		t.Run(test.value, func(t *testing.T) {
			direction := Direction(test.value)
			require.Equal(t, direction.isValid(), test.valid)
		})
	}
}
