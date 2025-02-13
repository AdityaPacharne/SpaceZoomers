package main

func playerInput(screen *[][]string, *currentHeight int, *currentWidth int, terminalHeight int, terminalWidth int, quit chan bool) {
	fd := int(os.Stdin.Fd())
	oldState, _ := term.MakeRaw(fd) // Enable raw mode
	defer term.Restore(fd, oldState) // Restore on exit

	reader := bufio.NewReader(os.Stdin);
	for {
		input, _ := reader.ReadByte();
		if input == 27 { // Escape key
			next1, _ := reader.ReadByte()
			next2, _ := reader.ReadByte()
			if next1 == 91 { // '['
				if next2 == 67 {
					fmt.Println("Right Arrow")
				} else if next2 == 68 {
					fmt.Println("Left Arrow")
				}
			}
		} else if input == 'q' { // Quit on 'q'
			break
		} else {
			fmt.Printf("You typed: %c\n", input)
		}

		var inputString string = string(input);
		if inputString == "q" {
			quit <- true;
			break;
		}
		spaceship(screen, currentHeight, currentWidth, inputString, terminalHeight, terminalWidth);
	}
}
