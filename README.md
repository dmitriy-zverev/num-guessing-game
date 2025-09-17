# Number Guessing Game

A simple command-line number guessing game written in Go where players try to guess a randomly generated number between 1 and 100.

## Features

- **Multiple Difficulty Levels**: Choose from Easy (10 chances), Medium (5 chances), or Hard (3 chances)
- **Score Tracking**: Automatically saves your best scores (fewest attempts and fastest time)
- **Hint System**: Get helpful hints on easier difficulty levels after multiple attempts
- **Input Validation**: Ensures valid number inputs within the correct range
- **Replay Option**: Continue playing multiple rounds without restarting the program

## Requirements

- Go 1.24.5 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/dmitriy-zverev/num-guessing-game.git
cd num-guessing-game
```

2. Run the game:
```bash
go run .
```

Alternatively, you can build the executable:
```bash
go build -o num-guessing-game
./num-guessing-game
```

## How to Play

1. **Start the Game**: Run the program and you'll see a welcome message along with your best scores (if any)

2. **Choose Difficulty**: Select your preferred difficulty level:
   - **Easy**: 10 chances to guess
   - **Medium**: 5 chances to guess  
   - **Hard**: 3 chances to guess

3. **Make Your Guesses**: Enter numbers between 1 and 100. The game will tell you if your guess is too high or too low

4. **Get Hints**: On Easy and Medium difficulties, you may receive hints about the number after several attempts

5. **Win or Lose**: 
   - **Win**: Guess correctly within the allowed attempts
   - **Lose**: Use all attempts without guessing correctly

6. **Play Again**: Choose whether to start a new game or exit

## Game Features

### Score System
- **Attempts**: Tracks the minimum number of attempts needed to win
- **Time**: Records the fastest completion time
- Scores are automatically saved to `scores.json`

### Hint System
- Available on Easy and Medium difficulties
- Hints are provided after 3+ attempts
- Two types of hints:
  - Last digit of the number
  - First digit of the number

### Input Validation
- Only accepts numbers between 1 and 100
- Invalid inputs don't count against your attempt limit
- Clear error messages for invalid inputs

## Project Structure

```
num-guessing-game/
├── main.go           # Entry point of the application
├── logic.go          # Core game logic and flow control
├── functions.go      # Utility functions for difficulty handling
├── validations.go    # Input validation and win/lose checking
├── storage.go        # Score persistence and file operations
├── consts.go         # Game constants and messages
├── scores.json       # High scores storage (created automatically)
├── go.mod           # Go module definition
└── README.md        # This file
```

## Code Overview

- **main.go**: Application entry point that initializes the game
- **logic.go**: Contains the main game loop, initialization, and flow control
- **functions.go**: Helper functions for converting difficulty levels
- **validations.go**: Input validation and game state checking
- **storage.go**: JSON-based score persistence system
- **consts.go**: Game configuration and user messages

## Example Gameplay

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have limited chances to guess the correct number.

Your best scores are:
    Attempts: 4
    Seconds: 12.34

Please select the difficulty level:
    1. Easy (10 chances)
    2. Medium (5 chances)
    3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

1. Enter your guess: 50
Incorrect! The number is greater than 50

2. Enter your guess: 75
Incorrect! The number is less than 75

3. Enter your guess: 62
Incorrect! The number is greater than 62

4. Enter your guess: 68
Congratulations! You guessed the correct number in 4 attempts

You have guessed in 8.45 seconds

Want to try again?
1. Yes
2. No

Start again: 2

Ok. It was nice to see you. Bye.
```

## Contributing

Feel free to fork this repository and submit pull requests for any improvements or bug fixes.

## License

This project is open source and available under no license.
