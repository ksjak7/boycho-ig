package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Tile struct {
	position      Position
	texture       *sdl.Texture
	width         int32
	height        int32
	game          *Game
	adjacentTiles []*Tile
}

func createTile(x int32, y int32, texture *sdl.Texture, game *Game) Tile {
	return Tile{Position{x, y}, texture, game.TILE_WIDTH, game.TILE_HEIGHT, game, []*Tile{}}
}
func (tile *Tile) getPosition() Position {
	return tile.position
}
func (tile *Tile) move(x int32, y int32) {
	tile.position.x += x
	tile.position.y += y
}
func (tile *Tile) render(renderer *sdl.Renderer, displayRect *sdl.Rect) {
	if tile.texture == nil || renderer == nil {
		os.Exit(5)
	}
	displayRect.X = tile.position.x*tile.width/2 + tile.position.y*tile.width/2 + (tile.game.WINDOW_WIDTH/2 - tile.width/2)
	displayRect.Y = tile.position.y*tile.game.TILE_HEIGHT/2 - tile.position.x*tile.game.TILE_HEIGHT/2 + (tile.game.WINDOW_HEIGHT/2 - tile.height/2)
	displayRect.W = tile.width
	displayRect.H = tile.height
	err := renderer.Copy(tile.texture, nil, displayRect)
	if err != nil {
		return
	}
}

/*
	func shortestPath(start Tile, end Tile) []Tile {
		openList := [...]Tile{start}
		closedList := [...]Tile{}
		g, h := 0, 0

		for len(openList) > 0 {

		}
	}
*/
func (tile *Tile) withinAdjacentTiles(foreignTile *Tile) bool {
	for _, entity := range tile.adjacentTiles {
		if foreignTile == entity {
			return true
		}
	}
	return false
}
