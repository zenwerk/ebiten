/*
Copyright 2014 Hajime Hoshi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blocks

import (
	"github.com/hajimehoshi/ebiten"
	"sync"
)

type Size struct {
	Width  int
	Height int
}

// TODO: Should they be global??
var texturePaths = map[string]string{}
var renderTargetSizes = map[string]Size{}

const ScreenWidth = 256
const ScreenHeight = 240

type GameState struct {
	SceneManager *SceneManager
	Input        *Input
}

type Game struct {
	once         sync.Once
	sceneManager *SceneManager
	input        *Input
	textures     *Textures
}

func NewGame() *Game {
	game := &Game{
		sceneManager: NewSceneManager(NewTitleScene()),
		input:        NewInput(),
	}
	return game
}

func (game *Game) isInitialized() bool {
	if game.textures == nil {
		return false
	}
	for name := range texturePaths {
		if !game.textures.Has(name) {
			return false
		}
	}
	for name := range renderTargetSizes {
		if !game.textures.Has(name) {
			return false
		}
	}
	return true
}

func (game *Game) Update(r *ebiten.RenderTarget) error {
	game.once.Do(func() {
		game.textures = NewTextures()
		for name, path := range texturePaths {
			game.textures.RequestTexture(name, path)
		}
		for name, size := range renderTargetSizes {
			game.textures.RequestRenderTarget(name, size)
		}
	})
	if !game.isInitialized() {
		return nil
	}
	game.input.Update()
	game.sceneManager.Update(&GameState{
		SceneManager: game.sceneManager,
		Input:        game.input,
	})
	game.sceneManager.Draw(r, game.textures)
	return nil
}