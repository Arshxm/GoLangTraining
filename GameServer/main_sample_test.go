package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestSampleGameCreation(t *testing.T) {
// 	g, err := NewGame([]int{})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, g)
// }

// func TestSampleAddPlayer(t *testing.T) {
// 	g, err := NewGame([]int{1, 2, 3})
// 	assert.Nil(t, err)

// 	err = g.ConnectPlayer("Cyn")
// 	assert.Nil(t, err)
// }

// func TestSampleGetPlayer(t *testing.T) {
// 	g, err := NewGame([]int{1, 2, 3})
// 	assert.Nil(t, err)

// 	err = g.ConnectPlayer("Cyn")
// 	assert.Nil(t, err)

// 	p, err := g.GetPlayer("CyN")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, p)
// }

// func TestSampleAddMap(t *testing.T) {
// 	g, err := NewGame([]int{1, 2, 3})
// 	assert.Nil(t, err)

// 	err = g.AddMap(4)
// 	assert.Nil(t, err)

// 	// Test adding invalid map
// 	err = g.AddMap(0)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "invalid map id", err.Error())

// 	// Test adding negative map
// 	err = g.AddMap(-1)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "invalid map id", err.Error())

// 	// Test adding existing map
// 	err = g.AddMap(1)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "map already exists", err.Error())
// }

// func TestSamplePlayerSwitch(t *testing.T) {
// 	g, err := NewGame([]int{1, 2, 3})
// 	assert.Nil(t, err)

// 	err = g.ConnectPlayer("alice")
// 	assert.Nil(t, err)

// 	err = g.SwitchPlayerMap("alice", 1)
// 	assert.Nil(t, err)

// 	player, err := g.GetPlayer("alice")
// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, player.mapId)

// 	// Test switching to map 0 (remove from map)
// 	err = g.SwitchPlayerMap("alice", 0)
// 	assert.Nil(t, err)
// 	assert.Equal(t, 0, player.mapId)
// 	assert.Nil(t, player.currentMap)
// }

// func TestSampleMessageHandling(t *testing.T) {
// 	g, err := NewGame([]int{1})
// 	assert.Nil(t, err)

// 	err = g.ConnectPlayer("alice")
// 	assert.Nil(t, err)

// 	err = g.ConnectPlayer("bob")
// 	assert.Nil(t, err)

// 	alice, _ := g.GetPlayer("alice")
// 	bob, _ := g.GetPlayer("bob")

// 	// Put both in same map
// 	g.SwitchPlayerMap("alice", 1)
// 	g.SwitchPlayerMap("bob", 1)

// 	// Alice sends message
// 	err = alice.SendMessage("hello")
// 	assert.Nil(t, err)

// 	// Test that bob has the channel (message handling is async)
// 	assert.NotNil(t, bob.GetChannel())

// 	// Test ReadMessage functionality
// 	msg, err := bob.ReadMessage()
// 	if err == nil {
// 		assert.Equal(t, "Alice says: hello", msg)
// 	}
// }
