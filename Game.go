package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

type Game struct {
	WINDOW_WIDTH,
	WINDOW_HEIGHT,
	TILE_HEIGHT,
	TILE_WIDTH int32
	TILE_TEXTURE          *sdl.Texture
	SELECTED_TILE_TEXTURE *sdl.Texture

	window   *sdl.Window
	renderer *sdl.Renderer
	err      error
}

func InitializeGame() *Game {
	game := Game{TILE_WIDTH: 168, TILE_HEIGHT: 54}
	_, game.err = sdl.ShowCursor(0)
	if game.err != nil {
		return nil
	}
	game.window, game.err = sdl.CreateWindow("Input", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 1600, 900, sdl.WINDOW_FULLSCREEN)
	if game.err != nil {
		fmt.Println(game.err)
		os.Exit(1)
	}

	game.renderer, game.err = sdl.CreateRenderer(game.window, -1, sdl.RENDERER_ACCELERATED)
	if game.err != nil {
		fmt.Println(game.err)
		os.Exit(1)
	}

	game.WINDOW_WIDTH, game.WINDOW_HEIGHT = game.window.GetSize()
	game.TILE_TEXTURE, _ = img.LoadTexture(game.renderer, "./res/img/tile.png")
	game.TILE_TEXTURE.SetColorMod(148, 169, 255)
	game.SELECTED_TILE_TEXTURE, _ = img.LoadTexture(game.renderer, "./res/img/tile.png")
	game.SELECTED_TILE_TEXTURE.SetColorMod(0, 0, 0)

	return &game
}