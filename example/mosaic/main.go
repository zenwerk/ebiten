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

package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	_ "image/jpeg"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

const mosaicRatio = 16

type Game struct {
	gophersTexture      *ebiten.Texture
	gophersRenderTarget *ebiten.RenderTarget
}

func (g *Game) Update(r *ebiten.RenderTarget) error {
	geo := ebiten.ScaleGeometry(1.0/mosaicRatio, 1.0/mosaicRatio)
	ebiten.DrawWholeTexture(g.gophersRenderTarget, g.gophersTexture, geo, ebiten.ColorMatrixI())

	geo = ebiten.ScaleGeometry(mosaicRatio/2.0, mosaicRatio/2.0)
	ebiten.DrawWholeTexture(r, g.gophersRenderTarget.Texture(), geo, ebiten.ColorMatrixI())
	return nil
}

func main() {
	g := new(Game)
	var err error
	g.gophersTexture, _, err = ebitenutil.NewTextureFromFile("images/gophers.jpg", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	w, h := g.gophersTexture.Size()
	g.gophersRenderTarget, err = ebiten.NewRenderTarget(w/mosaicRatio, h/mosaicRatio, ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(g.Update, screenWidth, screenHeight, 2, "Mosaic (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}