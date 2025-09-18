package game

import (
	"fmt"
	"jo-ken-go/enum"
)

func JoKenGo(move1, move2 enum.Move) {
	beats := map[enum.Move]enum.Move{
		enum.Rock:     enum.Scissors,
		enum.Paper:    enum.Rock,
		enum.Scissors: enum.Paper,
	}

	if move1 == move2 {
		fmt.Println("Draw!")
		return
	}

	if beats[move1] == move2 {
		PrintWinner(move1)
	} else {
		PrintWinner(move2)
	}
}

func PrintWinner(move enum.Move) {
	switch move {
	case enum.Paper:
		fmt.Println("Paper wins!")
	case enum.Rock:
		fmt.Println("Rock wins!")
	case enum.Scissors:
		fmt.Println("Scissors wins!")
	}
}
