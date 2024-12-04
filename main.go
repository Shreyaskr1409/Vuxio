package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Player struct {
	*Sprite
	Health uint
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}

type Game struct {
	player  *Player
	enemies []*Enemy
}

func (g *Game) Update() error {
	// reaction to key presses
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.X += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.player.Y += 2
	}

	for _, sprite := range g.enemies {

		if !sprite.FollowsPlayer {
			continue
		}

		if sprite.X < g.player.X {
			sprite.X += 1
		} else if sprite.X > g.player.X {
			sprite.X -= 1
		}
		if sprite.Y > g.player.Y {
			sprite.Y -= 1
		} else if sprite.Y < g.player.Y {
			sprite.Y += 1
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	//ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(
		g.player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	opts.GeoM.Reset()

	for _, enemy := range g.enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)

		screen.DrawImage(
			enemy.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")
	if err != nil {
		log.Fatal(err)
	}

	skeletonImg, _, err := ebitenutil.NewImageFromFile("assets/images/skeleton.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(&Game{
		player: &Player{
			&Sprite{
				Img: playerImg,
				X:   50,
				Y:   50,
			},
			50,
		},
		enemies: []*Enemy{
			{
				&Sprite{
					Img: skeletonImg,
					X:   100,
					Y:   100,
				},
				true,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   150,
					Y:   100,
				},
				false,
			},
			{
				&Sprite{
					Img: skeletonImg,
					X:   100,
					Y:   150,
				},
				false,
			},
		},
	}); err != nil {
		log.Fatal(err)
	}
}
