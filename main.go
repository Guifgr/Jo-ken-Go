package main

import (
	"bufio"
	"fmt"
	"jo-ken-go/enum"
	"jo-ken-go/game"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Rock Paper Scissors!")
	fmt.Print("Player 1, choose (Rock, Paper, Scissors): ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	// Clear the terminal so Player 2 can't see Player 1's input
	clearTerminal()

	fmt.Print("Player 2, choose (Rock, Paper, Scissors): ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	var move1, move2 enum.Move
	switch strings.ToLower(input1) {
	case "rock":
		move1 = enum.Rock
	case "paper":
		move1 = enum.Paper
	case "scissors":
		move1 = enum.Scissors
	default:
		fmt.Println("Player 1 chose an invalid move.")
		return
	}

	switch strings.ToLower(input2) {
	case "rock":
		move2 = enum.Rock
	case "paper":
		move2 = enum.Paper
	case "scissors":
		move2 = enum.Scissors
	default:
		fmt.Println("Player 2 chose an invalid move.")
		return
	}

	game.JoKenGo(move1, move2)
}

// clearTerminal clears the console in a cross-platform way.
// On Windows it runs 'cmd /c cls', on other platforms it runs 'clear'.
func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// ignore error; clearing is best-effort
	_ = cmd.Run()
}
