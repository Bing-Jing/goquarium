package termination

import (
	"math/rand"
)

func UpMovement(t *Termination, e *Entity, position Position) Position {
	position.Y -= 1
	return position
}

func DownMovement(t *Termination, e *Entity, position Position) Position {
	position.Y += 1
	return position
}

func LeftMovement(t *Termination, e *Entity, position Position) Position {
	position.X -= 1
	return position
}

func RightMovement(t *Termination, e *Entity, position Position) Position {
	position.X += 1
	return position
}
func UpRightMovement(t *Termination, e *Entity, position Position) Position{
	position.X += 3
	position.Y += 1
	return position
}
func DownRightMovement(t *Termination, e *Entity, position Position) Position{
	position.X += 3
	position.Y -= 1
	return position
}
func UpLeftMovement(t *Termination, e *Entity, position Position) Position{
	position.X -= 3
	position.Y += 1
	return position
}
func DownLeftMovement(t *Termination, e *Entity, position Position) Position{
	position.X -= 3
	position.Y -= 1
	return position
}
func RandomRightMovement(t *Termination, e *Entity, position Position) Position{
	position.X += 3
	position.Y += rand.Intn(4)-2
	return position
}
func RandomLeftMovement(t *Termination, e *Entity, position Position) Position{
	position.X -= 3
	position.Y += rand.Intn(4)-2
	return position
}
