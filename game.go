package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"sync"
)

type spaceship struct {
    health int
    xaxis int
    yaxis int
}

type Rocks struct {
	xaxis int
	yaxis int
}

type bullet struct {
    direction bool
	xaxis int
	yaxis int
}

// Code to fetch terminal size
func getTerminalSize() (terminalWidth, terminalHeight int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return 80,24;
	}
	return int(ws.Col), int(ws.Row)
}

func main() {
    // Player Enters
	fmt.Println("Welcome Player");
	fmt.Println("Enter q to exit");

    // Fetching the terminal size
	var terminalWidth, terminalHeight int = getTerminalSize();

    // Creating a nested slice of strings to perfectly mimic user's terminal as screen
	screen := make([][]string, terminalHeight);
	for i := range screen {
    	screen[i] = make([]string, terminalWidth)
		for j := range screen[i] {
			screen[i][j] = "_";
		}
	}
    
    // Placing our spaceship as a single character for now at the centre of the screen
	var currentHeight int = terminalHeight-1;
	var currentWidth int = terminalWidth/2;

    // Creating a slice of type bullet struct to save the activeBullets present on Users screen
	var activeBullets []bullet;

    // Bool Channel to check if user has to quit
	quit := make(chan bool);
    spaceshipDirection := make(chan string);
	
    // A waitgroup to synchronize multiple goroutines, it waits for 3 goroutines to finish 
	var wg sync.WaitGroup
	wg.Add(4);

    // Goroutine to Print the screen
	go func() {
		defer wg.Done();
		Render(screen, &activeBullets, terminalWidth, currentHeight, currentWidth, spaceshipDirection, quit);
	}();

    // Goroutine to take player input and modify the position of spaceship accordingly
	go func() {
		defer wg.Done();
		PlayerInput(spaceshipDirection, &currentHeight, &currentWidth, terminalHeight, terminalWidth, quit);
	}();

    // Go routine to create bullet every specific duration
    // It is made so the user dont have to shoot, the spaceship will keep firing and the user has to aim
	go func() {
		defer wg.Done();
		BulletCreate(&activeBullets, &currentHeight, &currentWidth, true, quit);
	}();

	go func() {
		defer wg.Done();
        BulletLocation(&activeBullets, terminalHeight, quit);
	}();

    // WaitGroup waits....
	wg.Wait();

    // A little thanks for playing
    // Thanks for playing my game
    // And thanks to you too, to take time from your day to read my dumb code
	fmt.Println("Thanks for playing the game");
}
