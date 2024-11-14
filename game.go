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
	playerMoving bool //player moving
	playerDir int //player direction
	playerUp, playerDown, playerLeft, playerRight bool //player up, down, left, right
	playerFrame int //player frame

	frameCount int //frame count

	playerSpeed float32 = 3 //player speed

	musicPaused bool //music paused
	music rl.Music //music

	cam rl.Camera2D //camera
)

func drawScene() { //draw scene
	rl.DrawTexture(grassSprite, 100, 50, rl.White) //draw grass sprite
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0 , rl.White) //draw player sprite
}

func input() { //get input
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp){ //check if W or up key is pressed
		playerMoving = true //player moving
		playerDir = 1 //player direction
		playerUp = true //player up
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown){ //check if S or down key is pressed
		playerMoving = true //player moving
		playerDir = 0 //player direction
		playerDown = true //player down
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight){ //check if D or right key is pressed
		playerMoving = true //player moving
		playerDir = 3 //player direction
		playerRight = true //player right
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft){ //check if A or left key is pressed
		playerMoving = true //player moving
		playerDir = 2 //player direction
		playerLeft = true //player left
	}
	if rl.IsKeyPressed(rl.KeyQ){ //check if Q key is pressed
		musicPaused = !musicPaused //toggle music paused
	}
}

func update(){ //update game
	running = !rl.WindowShouldClose() //check if window should close

	playerSrc.X = 0 //reset player source rectangle x

	if playerMoving { //check if player moving
		if playerUp { //check if player up
			playerDest.Y -= playerSpeed //move player up
		}
		if playerDown { //check if player down
			playerDest.Y += playerSpeed //move player down
		}
		if playerLeft { //check if player left
			playerDest.X -= playerSpeed //move player left
		}
		if playerRight { //check if player right
			playerDest.X += playerSpeed //move player right
		}
		if frameCount % 8 == 0{ //check if frame count is divisible by 8
			playerFrame++ //increment player frame
		}
		playerSrc.X = playerSrc.Width * float32(playerFrame) //update player source rectangle x

	}

	frameCount++ //increment frame count
	if frameCount > 3 { 
		frameCount = 0 //reset frame count
	}

	playerSrc.Y = playerSrc.Height * float32(playerDir) //update player source rectangle y

	rl.UpdateMusicStream(music) //update music stream
	if(musicPaused){ //check if music is paused
		rl.PauseMusicStream(music) //pause music
	}else{
		rl.ResumeMusicStream(music) //resume music
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))) //update camera target

	playerMoving = false //player not moving
	playerUp, playerDown, playerLeft, playerRight = false, false, false, false //player up, down, left, right

}

func init(){ //initialize game
	rl.InitWindow(screenWidth, screenHeight, "Game") //initialize window with title
	rl.SetExitKey(0) //set exit key to 0 
	rl.SetTargetFPS(60) //set target frames per second

	grassSprite = rl.LoadTexture("sprites/tilesets/Grass_tiles_v2.png") //load grass sprite
	playerSprite = rl.LoadTexture("sprites/char/basic-char-spritesheet.png") //load player sprite

	playerSrc = rl.NewRectangle(0, 0, 48, 48) //player source rectangle
	playerDest = rl.NewRectangle(200, 200, 100, 100) //player destination rectangle

	rl.InitAudioDevice() //initialize audio device
	music = rl.LoadMusicStream("audio/music/Avery's Farm.mp3") //load music
	musicPaused = false //music paused
	rl.PlayMusicStream(music) //play music

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0, 1) //initialize camera


}

func quit(){ //quit game
	rl.UnloadTexture(grassSprite) //unload grass sprite
	rl.UnloadTexture(playerSprite) //unload player sprite
	rl.UnloadMusicStream(music) //unload music
	rl.CloseAudioDevice() //close audio device
	rl.CloseWindow() //close window
}

func render(){ //render game
	rl.BeginDrawing() //begin drawing
	rl.ClearBackground(bgColor) //drawing over new frame
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing() //end drawing
}


func main() { //function called when game runs

	for running { //loop until window is closed
		input() //get input
		update() //update game
		render() //render game
	}

	quit() //quit game
}