package main //created package name 
 
import rl "github.com/gen2brain/raylib-go/raylib" //import raylib

const (
	screenWidth = 1000 //screen width
	screenHeight = 480 //screen height
)

var (
	running = true //keeps track of main game loop
	bgColor = rl.NewColor(147,211,196,255) //background color

	grassSprite rl.Texture2D //grass sprite
	playerSprite rl.Texture2D //player sprite
	playerSrc    rl.Rectangle  //player source rectangle
	playerDest   rl.Rectangle  //player destination rectangle

	playerSpeed float32 = 3 //player speed
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0 , rl.White)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp){
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown){
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight){
		playerDest.X += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft){
		playerDest.X -= playerSpeed
	}
}

func update(){
	running = !rl.WindowShouldClose()
}

func init(){
	rl.InitWindow(screenWidth, screenHeight, "Game") //initialize window with title
	rl.SetExitKey(0) //set exit key to 0 
	rl.SetTargetFPS(60) //set target frames per second

	grassSprite = rl.LoadTexture("sprites/tilesets/Grass_tiles_v2.png")
	playerSprite = rl.LoadTexture("sprites/char/basic-char-spritesheet.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)


}

func quit(){
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func render(){
	rl.BeginDrawing() //begin drawing
	rl.ClearBackground(bgColor) //drawing over new frame

	drawScene()
	rl.EndDrawing() //end drawing
}


func main() { //function called when game runs

	for running { //loop until window is closed
		input()
		update()
		render()
	}

	quit()
}