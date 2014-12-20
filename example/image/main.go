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
	"math"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	count           int
	horizontalCount int
	verticalCount   int
	gophersTexture  *ebiten.Texture
}

func (g *Game) Update(r *ebiten.RenderTarget) error {
	g.count++
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.horizontalCount--
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.horizontalCount++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.verticalCount--
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.verticalCount++
	}

	w, h := g.gophersTexture.Size()
	geo := ebiten.TranslateGeometry(-float64(w)/2, -float64(h)/2)
	scaleX := 0.5 * math.Pow(1.05, float64(g.horizontalCount))
	scaleY := 0.5 * math.Pow(1.05, float64(g.verticalCount))
	geo.Concat(ebiten.ScaleGeometry(scaleX, scaleY))
	geo.Concat(ebiten.RotateGeometry(float64(g.count%720) * 2 * math.Pi / 720))
	geo.Concat(ebiten.TranslateGeometry(screenWidth/2, screenHeight/2))
	//clr := ebiten.RotateHue(float64(g.count%180) * 2 * math.Pi / 180)
	clr := ebiten.ColorMatrixI()
	ebiten.DrawWholeTexture(r, g.gophersTexture, geo, clr)
	return nil
}

func main() {
	g := new(Game)
	var err error
	g.gophersTexture, _, err = ebitenutil.NewTextureFromFile("images/gophers.jpg", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(g.Update, screenWidth, screenHeight, 2, "Image (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}