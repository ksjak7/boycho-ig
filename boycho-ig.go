package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.LockOSThread()
	var err error

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = img.Init(img.INIT_PNG)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	game := InitializeGame()

	var displayRect sdl.Rect = sdl.Rect{
		X: 0,
		Y: 0,
		W: 0,
		H: 0,
	}

	var startTime time.Time
	var frameTime int64

	var visibleEntities []Visible
	var allTiles []*Tile
	for i := -4; i < 5; i++ {
		for j := -4; j < 5; j++ {
			if rand.Int()%5 < 3 {
				tile := createTile(int32(i), int32(j), game.BASIC_TILE_TEXTURE, 1, game)
				allTiles = append(allTiles, &tile)
			} else {
				tile := createTile(int32(i), int32(j), game.GRASSY_TILE_TEXTURE, 2, game)
				allTiles = append(allTiles, &tile)
			}
		}
	}

	centerTile := createTile(0, 0, game.PLAYER_TILE_TEXTURE, -1, game)

	for _, entity := range allTiles {
		for _, entityJ := range allTiles {
			xDistance := int(math.Abs(float64(entity.position.x) - float64(entityJ.position.x)))
			yDistance := int(math.Abs(float64(entity.position.y) - float64(entityJ.position.y)))
			if entity != &centerTile && entity != entityJ && !entity.withinAdjacentTiles(entityJ) && ((xDistance == 1 && yDistance == 0) || (xDistance == 0 && yDistance == 1)) {
				entity.adjacentTiles = append(entity.adjacentTiles, entityJ)
			}
		}
		visibleEntities = append(visibleEntities, entity)
	}

	var start, end *Tile
	visibleEntities = append(visibleEntities, &centerTile)
	running := true
	for running {
		startTime = time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym

				if t.State == sdl.PRESSED {
					if keyCode == sdl.K_ESCAPE {
						os.Exit(1)
					}
					if keyCode == sdl.K_a {
						centerTile.move(-1, 0)
					}
					if keyCode == sdl.K_d {
						centerTile.move(1, 0)
					}
					if keyCode == sdl.K_w {
						centerTile.move(0, -1)
					}
					if keyCode == sdl.K_s {
						centerTile.move(0, 1)
					}
					if keyCode == sdl.K_SPACE {
						if start == nil {
							for _, entity := range allTiles {
								if entity.position == centerTile.position {
									start = entity
								}
							}
						} else {
							for _, entity := range allTiles {
								if entity.position == centerTile.position {
									end = entity
								}
							}

							parent, _ := shortestPath(start, end)
							for end != nil {
								end.texture = game.SELECTED_TILE_TEXTURE
								end = parent[end]
							}
							start = nil
						}
					}
				}
				if t.State == sdl.RELEASED {

				}
			}
		}

		err = game.renderer.SetDrawColor(156, 156, 156, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		game.renderer.Clear()
		for _, entity := range visibleEntities {
			entity.render(game.renderer, &displayRect)
		}
		game.renderer.Present()

		frameTime = time.Now().Sub(startTime).Milliseconds()
		if frameTime <= 16 {
			sdl.Delay(uint32(16 - frameTime))
		}
	}
	if err := game.window.Destroy(); err != nil {
		os.Exit(1)
	}
}

type Visible interface {
	render(renderer *sdl.Renderer, displayRect *sdl.Rect)
}
