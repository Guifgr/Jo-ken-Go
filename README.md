# Rock Paper Scissors (Jo-Ken-Go)

A simple command-line Rock Paper Scissors game written in Go.

## Features
- Two-player mode via terminal input
- Clean and simple logic using Go maps

## How to Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Guifgr/Jo-ken-Go.git
   cd Jo-ken-Go
   ```
2. **Run the game:**
   ```bash
   go run main.go
   ```

## How to Play
- When prompted, each player types their move: `Rock`, `Paper`, or `Scissors` (case-insensitive).
- The game will announce the winner or if it's a draw.

## Example
```
Welcome to Rock Paper Scissors!
Player 1, choose (Rock, Paper, Scissors): rock
Player 2, choose (Rock, Paper, Scissors): paper
Paper wins!
```

## Project Structure
```
Jo-ken-Go/
├── enum/         # Move definitions (Rock, Paper, Scissors)
├── game/         # Game logic
├── main.go       # Entry point
├── go.mod        # Go module file
```

## License
MIT
