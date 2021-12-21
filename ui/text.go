package ui
import "github.com/gen2brain/raylib-go/raylib"

// Draw pixel perfect text
func DrawText(font rl.Font, text string, position rl.Vector2, spacing float32) {
    rl.DrawTextEx(
        font,
        text,
        position,
        float32(font.BaseSize * PixelSize),
        spacing * PixelSize,
        rl.Black,
    )
}

func DrawTextBox(panel Panel, destination rl.Rectangle, font rl.Font, text string, spacing float32, highlighted bool) {

    DrawPanel(panel, destination)

    // Calculate position of centered text
    textSize := rl.MeasureTextEx(
        font,
        text,
        float32(font.BaseSize * PixelSize),
        spacing * PixelSize,
    )
    xOffset := (destination.Width - textSize.X) / 2;
    yOffset := (destination.Height - textSize.Y) / 2;

    // Snap offsets for crisp pixel-perfectness
    xOffset = float32(int32(xOffset / PixelSize) * PixelSize)
    yOffset = float32(int32(yOffset / PixelSize) * PixelSize)

    position := rl.Vector2{
        destination.X + xOffset,
        destination.Y + yOffset,
    }

    if highlighted {
        rl.DrawRectangleLinesEx(destination, PixelSize, rl.Color{255, 128, 208, 255})
    }

    DrawText(font, text, position, spacing)

}

