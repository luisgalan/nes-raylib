package ui

import "github.com/gen2brain/raylib-go/raylib"

// 9-patch ui panel struct
type Panel struct {
    Texture    rl.Texture2D
    BorderSize float32
}

// Draw 9-patch ui panel
func DrawPanel(panel Panel, destination rl.Rectangle) {
    texWidth := float32(panel.Texture.Width)
    texHeight := float32(panel.Texture.Height)
    border := panel.BorderSize

    /*
         0 1      2 3  x
        0+-+------+-+
         | |      | |
        1+-+------+-+
         | |      | |
         | |      | |
         | |      | |
        2+-+------+-+ \
         | |      | |  > BorderSize
        3+-+------+-+ /

        y
    */

    xSourceCoords := [4]float32{
        0,
        border,
        texWidth - border,
        texWidth,
    }
    ySourceCoords := [4]float32{
        0,
        border,
        texHeight - border,
        texHeight,
    }

    xDestCoords := [4]float32{
        destination.X,
        destination.X + border * PixelSize,
        destination.X + destination.Width - (border * PixelSize),
        destination.X + destination.Width,
    }
    yDestCoords := [4]float32{
        destination.Y,
        destination.Y + border * PixelSize,
        destination.Y + destination.Height - (border * PixelSize),
        destination.Y + destination.Height,
    }

    for xIndex := 0; xIndex < 3; xIndex++ {
        for yIndex := 0; yIndex < 3; yIndex++ {
            rl.DrawTexturePro(
                panel.Texture,
                rl.Rectangle{
                    xSourceCoords[xIndex],
                    ySourceCoords[yIndex],
                    xSourceCoords[xIndex + 1] - xSourceCoords[xIndex],
                    ySourceCoords[yIndex + 1] - ySourceCoords[yIndex],
                },
                rl.Rectangle{
                    xDestCoords[xIndex],
                    yDestCoords[yIndex],
                    xDestCoords[xIndex + 1] - xDestCoords[xIndex],
                    yDestCoords[yIndex + 1] - yDestCoords[yIndex],
                },
                rl.Vector2{0, 0},
                0,
                rl.White,
            )
        }
    }
}

