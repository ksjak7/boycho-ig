package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
)

const (
	TILE_HEIGHT int32 = TILE_WIDTH * 9 / 28
	TILE_WIDTH  int32 = 56
)

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

	window, err = sdl.CreateWindow("Input", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1600, 900, sdl.WINDOW_FULLSCREEN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
			tile := createTile(int32(i), int32(j), TILE_WIDTH, TILE_HEIGHT, tex)
			visibleEntities = append(visibleEntities, &tile)
		}
	}

	centerTile := createTile(4, 20, TILE_WIDTH, TILE_HEIGHT, tex2)

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
						for _, entity := range visibleEntities {
							entity.movePos(TILE_WIDTH/2, -TILE_HEIGHT)
						}
					}
					if keyCode == sdl.K_d {
						for _, entity := range visibleEntities {
							entity.movePos(-TILE_WIDTH/2, TILE_HEIGHT)
						}
					}
					if keyCode == sdl.K_w {
						for _, entity := range visibleEntities {
							entity.movePos(TILE_WIDTH/2, TILE_HEIGHT)
						}
					}
					if keyCode == sdl.K_s {
						for _, entity := range visibleEntities {
							entity.movePos(-TILE_WIDTH/2, -TILE_HEIGHT)
						}
					}
					fmt.Println("(", -visibleEntities[0].getX()/TILE_WIDTH, ", ", visibleEntities[0].getY()/TILE_HEIGHT, ")")
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
	x, y, w, h int32
	tex        *sdl.Texture
}

func createTile(x int32, y int32, w int32, h int32, tex *sdl.Texture) Tile {
	return Tile{x*TILE_WIDTH + TILE_WIDTH/2 + (y % 2 * TILE_WIDTH / 2), y*TILE_HEIGHT + TILE_HEIGHT, w, h, tex}
}
func (tile *Tile) getX() (x int32) {
	return tile.x
}
func (tile *Tile) getY() (y int32) {
	return tile.y
}
func (tile *Tile) setPos(x int32, y int32) {
	tile.x = x
	tile.y = y
}
func (tile *Tile) movePos(dx int32, dy int32) {
	tile.x += dx
	tile.y += dy
}
func (tile *Tile) render(renderer *sdl.Renderer, displayRect *sdl.Rect) {
	if tile.tex == nil || renderer == nil {
		os.Exit(5)
	}
	displayRect.X = tile.x - tile.w/2
	displayRect.Y = tile.y - tile.h/2
	displayRect.W = tile.w
	displayRect.H = tile.h
	err := renderer.Copy(tile.tex, nil, displayRect)
	if err != nil {
		return
	}
}

type Visible interface {
	getX() (x int32)
	getY() (y int32)
	setPos(x int32, y int32)
	movePos(dx int32, dy int32)
	render(renderer *sdl.Renderer, displayRect *sdl.Rect)
}
