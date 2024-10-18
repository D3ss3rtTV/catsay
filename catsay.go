package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ANSI color codes
var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[37m", // White
}

// Cat ASCII art variations
var cats = []string{
	"  /\\_/\\\n ( o.o )\n  > ^ <",
	"  /\\_/\\\n ( -.- )\n  > ^ <",
	"  /\\_/\\\n ( O.O )\n  > ^ <",
	"  /\\_/\\\n ( @.@ )\n  > ^ <",
	"  /\\_/\\\n ( ^_^ )\n  > ^ <",
	"  /\\_/\\\n ( *_* )\n  > ^ <",
	"  /\\_/\\\n ( ;_; )\n  > ^ <",
	"  /\\_/\\\n ( +_+ )\n  > ^ <",
	"  /\\_/\\\n ( >.< )\n  > ^ <",
	"  /\\_/\\\n ( ~_~ )\n  > ^ <",
	"  /\\_/\\\n ( 0_0 )\n  > ^ <",
	"  /\\_/\\\n ( ._. )\n  > ^ <",
	"  /\\_/\\\n ( 'o' )\n  > ^ <",
	"  /\\_/\\\n ( -o- )\n  > ^ <",
	"  /\\_/\\\n ( #_# )\n  > ^ <",
	"  /\\_/\\\n ( ^-^ )\n  > ^ <",
	"  /\\_/\\\n ( T_T )\n  > ^ <",
	"  /\\_/\\\n ( $_$ )\n  > ^ <",
	"  /\\_/\\\n ( @_@ )\n  > ^ <",
}

// Custom quotes array renamed to myQuotes
var myQuotes = []string{
	"All your base are belong to us.",
	"It’s a UNIX system! I know this!",
	"The cake is a lie.",
	"Bite my shiny metal ass!",
	"Hello, IT. Have you tried turning it off and on again?",
	"An idiot admires complexity, a genius admires simplicity.",
	"The internet! It’s become sentient!",
	"Hello, World!",
	"He who controls the spice controls the universe.",
	"Violence is the last refuge of the incompetent.",
}

// catsay prints a message inside a speech bubble with a cat below it in random color
func catsay(message string, rng *rand.Rand, typewriterSpeed int, disableTypewriter bool) {
	// Select random cat and color
	cat := cats[rng.Intn(len(cats))]
	color := colors[rng.Intn(len(colors))]

	// Create the speech bubble
	messageLine := fmt.Sprintf("| %s |", message)
	bubbleBorder := strings.Repeat("-", len(messageLine))

	// Print the speech bubble with typewriter effect if not disabled
	fmt.Println(color + bubbleBorder)
	if disableTypewriter {
		fmt.Println(messageLine)
	} else {
		typewriterEffect(messageLine, typewriterSpeed)
	}
	fmt.Println(bubbleBorder)
	fmt.Println(cat + "\033[0m") // Reset color after printing cat
}

// typewriterEffect prints text one character at a time with a delay to simulate a typewriter effect
func typewriterEffect(text string, speed int) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(time.Duration(speed) * time.Millisecond) // Speed adjustment for typewriter effect
	}
	fmt.Println()
}

func main() {
	// Create a new random number generator with a new source
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Define flags for custom message, typewriter speed, and disabling typewriter
	messageFlag := flag.String("m", "", "Custom message for the cat")
	speedFlag := flag.Int("speed", 50, "Set typewriter speed in milliseconds per character")
	disableTypewriterFlag := flag.Bool("disable-typewriter", false, "Disable typewriter effect")
	flag.Parse()

	// Set a random default message from myQuotes if the flag is not provided
	var message string
	if *messageFlag == "" {
		message = myQuotes[rng.Intn(len(myQuotes))]
	} else {
		message = *messageFlag
	}

	// Print one random cat with the message and custom options
	catsay(message, rng, *speedFlag, *disableTypewriterFlag)
}
