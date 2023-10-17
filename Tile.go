package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"os"
)

type Tile struct {
	position      Position
	texture       *sdl.Texture
	width         int32
	height        int32
	game          *Game
	adjacentTiles []*Tile
	cost          float64
	priority      float64
}

func createTile(x int32, y int32, texture *sdl.Texture, game *Game) Tile {
	return Tile{Position{x, y}, texture, game.TILE_WIDTH, game.TILE_HEIGHT, game, []*Tile{}, 1, 0}
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

func shortestPath(start *Tile, end *Tile) (map[*Tile]*Tile, map[*Tile]float64) {
	openList := []*Tile{start}
	current := openList[0]
	parent := map[*Tile]*Tile{}
	currentCost := map[*Tile]float64{}
	parent[start] = nil
	currentCost[start] = 0
	currentIndex := 0

	for len(openList) > 0 {
		for index, entity := range openList {
			if entity.priority < current.priority {
				current = entity
				openList[currentIndex] = openList[len(openList)-1]
				openList = openList[:len(openList)-1]
				currentIndex = index
			}
		}
		current = openList[0]
		openList[0] = openList[len(openList)-1]
		openList = openList[:len(openList)-1]

		if current == end {
			return parent, currentCost
		}

		for _, child := range current.adjacentTiles {
			new_cost := currentCost[current] + child.cost
			if !childInKeys(child, currentCost) || new_cost < currentCost[child] {
				currentCost[child] = new_cost
				child.priority = new_cost + child.heuristic(end)
				openList = append(openList, child)
				parent[child] = current
			}
		}
		fmt.Println(len(openList))
	}
	return parent, currentCost
}

func (tile *Tile) withinAdjacentTiles(foreignTile *Tile) bool {
	for _, entity := range tile.adjacentTiles {
		if foreignTile == entity {
			return true
		}
	}
	return false
}

func childInKeys(tile *Tile, costMap map[*Tile]float64) bool {
	for key, _ := range costMap {
		if tile == key {
			return true
		}
	}
	return false
}

func (tile *Tile) heuristic(end *Tile) float64 {
	return -math.Abs(float64(tile.position.x-end.position.x)) -
		math.Abs(float64(tile.position.y-end.position.y))
}
