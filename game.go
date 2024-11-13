package main //created package name 
 
import rl "github.com/gen2brain/raylib-go/raylib"

func main() { //function called when game runs
	rl.InitWindow(800, 450, "raylib [core] example - basic window") //initialize window
	defer rl.CloseWindow() //close window when game is closed

	rl.SetTargetFPS(60) //set target frames per second

	for !rl.WindowShouldClose() { //loop until window is closed
		rl.BeginDrawing() //begin drawing

		rl.ClearBackground(rl.RayWhite) //drawing over new frame
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray) //draw text, position, size, color

		rl.EndDrawing() //end drawing
	}
}