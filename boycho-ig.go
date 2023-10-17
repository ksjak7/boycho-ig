package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math"
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
			tile := createTile(int32(i), int32(j), game.TILE_TEXTURE, game)
			allTiles = append(allTiles, &tile)
		}
	}
	for _, entity := range allTiles {
		for _, entityJ := range allTiles {
			xDistance := int(math.Abs(float64(entity.position.x) - float64(entityJ.position.x)))
			yDistance := int(math.Abs(float64(entity.position.y) - float64(entityJ.position.y)))
			//fmt.Println(entity.position, entityJ.position, xDistance, yDistance)
			if entity != entityJ && (!(xDistance <= 1.0 && yDistance <= 1.0) && (xDistance <= 1.0 || yDistance <= 1.0)) && !entity.withinAdjacentTiles(entityJ) {
				entity.adjacentTiles = append(entity.adjacentTiles, entityJ)
			}
		}
		visibleEntities = append(visibleEntities, entity)
	}

	for _, entity := range allTiles {
		fmt.Println(len(entity.adjacentTiles))
	}

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

					}
					if keyCode == sdl.K_d {

					}
					if keyCode == sdl.K_w {

					}
					if keyCode == sdl.K_s {

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
