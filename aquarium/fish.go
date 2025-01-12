package goquarium

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ansoni/termination"
	"github.com/nsf/termbox-go"
)

type Goquarium struct {
	term     *termination.Termination
	surface  []*termination.Entity
	ripples  []*termination.Entity
	seaweeds []*termination.Entity
	fishes   []*termination.Entity
	castle   *termination.Entity
	bubbles  *termination.Entity
}

func (goq *Goquarium) generateBubbles() {
	term := goq.term
	for {
		rand.Seed(int64(time.Now().Nanosecond()))
		randX := random(0, term.Width)
		randY := random(17, term.Height)

		bubble := term.NewEntity(termination.Position{X: randX, Y: randY, Z: 10})
		bubble.Shape = bubbleShape
		bubble.MovementCallback = termination.UpMovement
		bubble.DefaultColor = 'c'
		bubble.MovesPerSecond = 5
		bubble.FramesPerSecond = 2
		bubble.DeathOnLastFrame = true
		time.Sleep(3 * time.Second)
	}
}

func (goq *Goquarium) deadFish(term *termination.Termination, entity *termination.Entity) {
	goq.addFish()
}

func (goq *Goquarium) deadWhale(term *termination.Termination, entity *termination.Entity) {
	goq.addWhale()
}

func (goq *Goquarium) deadShark(term *termination.Termination, entity *termination.Entity) {
	goq.addShark()
}
func (goq *Goquarium) deadSnake(term *termination.Termination, entity *termination.Entity) {
	goq.addSnake()
}

func (goq *Goquarium) addWhale() {
	direction := []string{"left", "right"}[random(0, 2)]
	width := goq.term.Width
	position := termination.Position{X: -10, Y: 0, Z: 50}
	if direction == "left" {
		position = termination.Position{X: width + 10, Y: 0, Z: 50}
	}

	whale := goq.term.NewEntity(position)
	whale.Shape = whaleShape
	whale.DeathOnOffScreen = true
	whale.ColorMask = whaleMask
	whale.ShapePath = direction
	whale.DefaultColor = 'b'
	whale.DeathCallback = goq.deadWhale
	if direction == "left" {
		whale.MovementCallback = termination.LeftMovement
	} else {
		whale.MovementCallback = termination.RightMovement
	}
	//fish.Death = goq.deadFish
	whale.MovesPerSecond = random(5,20)
	whale.FramesPerSecond = 10
}

func (goq *Goquarium) addShark() {
	direction := []string{"left", "right"}[random(0, 2)]
	width := goq.term.Width
	position := termination.Position{X: -53, Y: 0, Z: 50}
	if direction == "left" {
		position = termination.Position{X: width + 53, Y: 0, Z: 50}
	}

	shark := goq.term.NewEntity(position)
	shark.Shape = sharkShape
	shark.DeathOnOffScreen = true
	shark.ColorMask = sharkMask
	shark.ShapePath = direction
	shark.DefaultColor = 'c'
	shark.DeathCallback = goq.deadShark
	if direction == "left" {
		shark.MovementCallback = termination.LeftMovement
	} else {
		shark.MovementCallback = termination.RightMovement
	}
	//fish.Death = goq.deadFish
	shark.MovesPerSecond = random(2,10)
	shark.FramesPerSecond = 10
}
func (goq *Goquarium) addSnake(){
	direction := []string{"right","left"}[random(0,1)]
	width := goq.term.Width
	position := termination.Position{X: -53, Y: 0, Z: 50}
	if direction == "left" {
		position = termination.Position{X: width + 53, Y: 0, Z: 50}
	}

	snake := goq.term.NewEntity(position)
	snake.Shape = snakeShape
	snake.DeathOnOffScreen = true
	snake.ColorMask = snakeMask
	snake.ShapePath = direction
	snake.DefaultColor = 'W'
	snake.DeathCallback = goq.deadSnake
	if direction == "left" {
		snake.MovementCallback = termination.LeftMovement
	}else{
		snake.MovementCallback = termination.RightMovement
	}
	snake.MovesPerSecond = random(2,10)
	snake.FramesPerSecond = 10
}

func (goq *Goquarium) addFish() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fishSelection := random(0, len(fishShapes))
	fishShape := fishShapes[fishSelection]
	fishMask := fishMasks[fishSelection]

	direction := []string{"left", "right","left", "right","Randomleft", "Randomright", "upleft", "downleft", "downright", "upright"}[random(0, 10)]

	// we unfortunately have to iterate to get its height before insert...
	shapeData := []rune(fishShape["left"][0])
	shapeHeight := 1
	for _, char := range shapeData {
		if char == '\n' {
			shapeHeight += 1
		}
	}

	height := goq.term.Height - shapeHeight
	width := goq.term.Width
	randY := random(9, height)
	speed := random(7, 20)
	// TODO: lots of duplicate code here
	if direction == "left" || direction == "upleft" || direction == "downleft" || direction == "Randomleft" {
		fish := goq.term.NewEntity(termination.Position{X: width + 10, Y: randY, Z: 5})
		fish.Shape = fishShape
		fish.DeathOnOffScreen = true
		fish.ColorMask = fishMask
		fish.ShapePath = "left"
		if direction == "left" {
			fish.MovementCallback = termination.LeftMovement
		}else if direction == "upleft"{
			fish.MovementCallback = termination.UpLeftMovement
		}else if direction == "downleft"{
			fish.MovementCallback = termination.DownLeftMovement
		}else{
			fish.MovementCallback = termination.RandomLeftMovement
		}
		fish.DeathCallback = goq.deadFish
		fish.MovesPerSecond = speed
		fish.FramesPerSecond = 10
		goq.fishes = append(goq.fishes, fish)
	} else if direction == "right" || direction == "upright" || direction == "downright" {
		fish := goq.term.NewEntity(termination.Position{X: -10, Y: randY, Z: 5})
		fish.Shape = fishShape
		fish.DeathOnOffScreen = true
		fish.ColorMask = fishMask
		fish.ShapePath = "right"
		if direction == "right" {
			fish.MovementCallback = termination.RightMovement
		}else if direction == "upright"{
			fish.MovementCallback = termination.UpRightMovement
		}else if direction == "downright"{
			fish.MovementCallback = termination.DownRightMovement
		}else{
			fish.MovementCallback = termination.RandomRightMovement
		}
		fish.DeathCallback = goq.deadFish
		fish.MovesPerSecond = speed
		fish.FramesPerSecond = 10
		goq.fishes = append(goq.fishes, fish)
	}
}

func (goq *Goquarium) generateFishes() {
	// screenSize := (goq.term.Height - 9) * goq.term.Width
	// fishCount := int(screenSize / 200)

	// keep adding fish when we need
	for {
		goq.addFish()
		time.Sleep(1000 * time.Millisecond) // space out the fish a bit
	}
}

func (goq *Goquarium) setupEnvironment() {
	term := goq.term
	top_y := 5
	height := term.Height
	width := term.Width

	needed := width / 4
	for i := 0; i < needed; i++ {
		surface := term.NewEntity(termination.Position{X: i * 4, Y: top_y, Z: 10})
		surface.Shape = surfaceShape
		surface.ColorMask = surfaceMask
		surface.FramesPerSecond = 1
		goq.surface = append(goq.surface, surface)

		ripplePaths := []string{"a", "b", "c", "d"}
		rand := random(0, 4)
		ripples := term.NewEntity(termination.Position{X: i * 4, Y: top_y + 1, Z: 10})
		ripples.Shape = rippleShape
		ripples.ColorMask = waterMask
		ripples.ShapePath = ripplePaths[rand]
		ripples.FramesPerSecond = 1
		goq.ripples = append(goq.ripples, surface)
	}

	//castle! castle is 13 high... if we change it... gonna suck here
	goq.castle = term.NewEntity(termination.Position{X: width - 31, Y: height - 13, Z: -1})
	goq.castle.Shape = castleShape
	goq.castle.FramesPerSecond = 1
	goq.castle.ShapePath = "default"
	goq.castle.ColorMask = castleMask
	goq.castle.DefaultColor = 'W'

	//grass is a bit random
	seaweedCount := 15
	for i := 0; i < seaweedCount; i++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		seaweedHeight := random(10, 20)
		seaweedX := random(0, width)
		path := []string{"a", "b"}
		h := 0
		for j := seaweedHeight; j >= 0; j-- {
			x := seaweedX
			y := height - seaweedHeight
			if x > width {
				x = width
			}

			seaweed := term.NewEntity(termination.Position{X: x, Y: y + j, Z: 10})
			seaweed.Shape = seaweedShape
			seaweed.FramesPerSecond = 2
			seaweed.ShapePath = path[h]
			seaweed.ColorMask = seaweedMask
			goq.seaweeds = append(goq.seaweeds, seaweed)
			if h == 0 {
				h = 1
			} else {
				h = 0
			}
		}
	}

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func Fish() {
	rand.Seed(int64(time.Now().Nanosecond()))

	goquarium := Goquarium{}
	term := termination.New()
	//term.Debug="./debug.out"
	goquarium.term = term
	term.FramesPerSecond = 60
	goquarium.setupEnvironment()

	go goquarium.generateBubbles()
	go goquarium.generateFishes()
	goquarium.addWhale()
	goquarium.addShark()
	goquarium.addSnake()
	go term.Animate()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				term.Close()
				os.Exit(0)
			}
		case termbox.EventMouse:
			//update_mouse(mouse, &ev)
		}
	}
}

// utility function to check differencies in shape/mask
func check(s1, s2 []string) {
	if len(s1) != len(s2) {
		panic("foo")
	}
	for i := 0; i < len(s1); i++ {
		l1 := strings.Split(s1[i], "\n")
		l2 := strings.Split(s2[i], "\n")
		if len(l1) != len(l2) {
			panic(fmt.Sprintf("different in part %d", i))
		}
		for j := 0; j < len(l1); j++ {
			if len(l1) != len(l2) {
				panic(fmt.Sprintf("different in line %d of part %d", j, i))
			}
		}
	}
}
