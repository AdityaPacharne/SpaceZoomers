package main

func spaceship (screen *[][]string, currentHeight *int, currentWidth *int, direction string, terminalHeight int, terminalWidth int) {
	if direction == "k" {
		if *currentWidth < terminalWidth - 1 {
			(*screen)[*currentHeight][*currentWidth] = "_";
			(*screen)[*currentHeight][*currentWidth+1] = "H";
			*currentWidth += 1;
		}
	} else if direction == "j" {
		if *currentWidth > 0 {
			(*screen)[*currentHeight][*currentWidth] = "_";
			(*screen)[*currentHeight][*currentWidth-1] = "H";
			*currentWidth -= 1;
		}
	} else {
		fmt.Println("Invalid Input");
	}
}

