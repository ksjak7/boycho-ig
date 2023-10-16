package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
)

var WINDOW_WIDTH,
	WINDOW_HEIGHT,
	TILE_HEIGHT,
	TILE_WIDTH int32

func main() {
	runtime.LockOSThread()
	var window *sdl.Window
	var renderer *sdl.Renderer
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

	window, err = sdl.CreateWindow("Input", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1600, 900, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	WINDOW_WIDTH, WINDOW_HEIGHT = window.GetSize()
	TILE_WIDTH, TILE_HEIGHT = 112, 36
	tex, _ := img.LoadTexture(renderer, "./res/img/tile.png")
	tex2, _ := img.LoadTexture(renderer, "./res/img/tile.png")
	tex.SetColorMod(148, 169, 255)
	tex2.SetColorMod(0, 0, 0)
	var displayRect sdl.Rect = sdl.Rect{
		X: 0,
		Y: 0,
		W: 0,
		H: 0,
	}

	var startTime time.Time
	var frameTime int64

	var visibleEntities []Visible
	for i := 0; i < 9; i++ {
		for j := 0; j < 36; j++ {
			tile := createTile(int32(i), int32(j), tex)
			visibleEntities = append(visibleEntities, &tile)
		}
	}

	centerTile := createTile(4, 18, tex2)
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
					fmt.Println(centerTile.position.x, " ", centerTile.position.y)
				}
				if t.State == sdl.RELEASED {

				}
			}
		}

		err = renderer.SetDrawColor(156, 156, 156, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		renderer.Clear()
		for _, entity := range visibleEntities {
			entity.render(renderer, &displayRect)
		}
		centerTile.render(renderer, &displayRect)
		renderer.Present()

		frameTime = time.Now().Sub(startTime).Milliseconds()
		if frameTime <= 16 {
			sdl.Delay(uint32(16 - frameTime))
		}
	}
	if err := window.Destroy(); err != nil {
		os.Exit(1)
	}
}

type Tile struct {
	position Position
	tex      *sdl.Texture
}

func createTile(x int32, y int32, tex *sdl.Texture) Tile {
	return Tile{Position{x, y}, tex}
}
func (tile *Tile) getPosition() Position {
	return tile.position
}
func (tile *Tile) move(x int32, y int32) {
	tile.position.x += x
	tile.position.y += y
}
func (tile *Tile) render(renderer *sdl.Renderer, displayRect *sdl.Rect) {
	if tile.tex == nil || renderer == nil {
		os.Exit(5)
	}
	displayRect.X = tile.position.x*TILE_WIDTH/2 + tile.position.y*TILE_WIDTH/2
	displayRect.Y = tile.position.y*TILE_HEIGHT/2 - tile.position.x*TILE_HEIGHT/2
	displayRect.W = TILE_WIDTH
	displayRect.H = TILE_HEIGHT
	err := renderer.Copy(tile.tex, nil, displayRect)
	if err != nil {
		return
	}
}

type Visible interface {
	render(renderer *sdl.Renderer, displayRect *sdl.Rect)
}

type Position struct {
	x int32
	y int32
}
