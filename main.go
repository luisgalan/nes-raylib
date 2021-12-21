package main

import (
    "github.com/gen2brain/raylib-go/raylib"
    "nes-raylib/ui"
    "math"
)

func drawMovingPanel(panel ui.Panel, position rl.Vector2) {
    time := rl.GetTime()
    width := 100 + 30 * float32(math.Sin(float64(time)))
    height := 100 + 30 * float32(math.Cos(float64(time)))
    ui.DrawPanel(panel, rl.Rectangle{position.X, position.Y, width, height})
}

func drawItems(selectedItem int, items []string, panel ui.Panel, font rl.Font, position rl.Vector2) {
    width := float32(280)
    height := float32(80)
    padding := float32(20)
    for i := 0; i < len(items); i++ {
        x := position.X
        y := position.Y + float32(i) * (height + padding)
        ui.DrawTextBox(
            panel,
            rl.Rectangle{x, y, width, height},
            font,
            items[i],
            1,
            selectedItem == i,
        )
    }
}

func main() {
    rl.InitWindow(ui.ScreenWidth, ui.ScreenHeight, "window title goes here")
    rl.SetTargetFPS(60)

    panel := ui.Panel{
        rl.LoadTexture("assets/textbox.png"),
        float32(2),
    }
    font := rl.LoadFont("assets/dogicapixel.tff")

    items := []string{
        "item 1",
        "item 2 :)",
        "one last item",
    }
    selectedItem := 0

    // Loop as long as the user hasn't pressed quit
    for !rl.WindowShouldClose() {

        // Process input
        if rl.IsKeyPressed(rl.KeyS) && selectedItem + 1 < len(items) {
            selectedItem = selectedItem + 1
        }
        if rl.IsKeyPressed(rl.KeyW) && selectedItem - 1 >= 0 {
            selectedItem = selectedItem - 1
        }

        // Draw current frame
        rl.BeginDrawing() // Tell raylib that the frame starts here

        // Make the background light blue
        rl.ClearBackground(rl.Color{180, 180, 255, 255})

        // Draw a moveable panel. We can make panels have arbitrary
        // sizes without worrying about weird scaling :)
        drawMovingPanel(panel, rl.Vector2{500, 50})

        ui.DrawTextBox(
            panel,
            rl.Rectangle{20, 20, 300, 100},
            font,
            "centered text",
            1,
            false,
        )

        drawItems(selectedItem, items, panel, font, rl.Vector2{50, 180})

        // Tell raylib that we're done drawing the frame. Raylib now sends all of our
        // draw calls to the gpu and the frame gets rendered to our screen.
        rl.EndDrawing()
    }
}

