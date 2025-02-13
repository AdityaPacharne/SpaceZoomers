// Function that renders the screen
func render(screen [][]string, activeBullets *[]bullet, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			fmt.Print("\033[H\033[2J")

            // Adds bullet onto the screen from activeBullets slice
            for _, tempBullet := range *activeBullets {
                screen[tempBullet.xaxis][tempBullet.yaxis] = "^";
            }

            // Printint the screen
			for _, row := range (screen) {
				fmt.Println(strings.Join(row, ""))
			}

			time.Sleep(10 * time.Millisecond);
		}
	}
}

