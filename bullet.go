package main

type bullet struct {
    direction bool
	xaxis int
	yaxis int
}

func bulletExistence (activeBullets *[]bullet, terminalHeight int) bool {
    var newBullets []bullet;
    for i := range *activeBullets {
        if (activeBullets[i].direction && activeBullets[i].yaxis >= 2) || (!activeBullets[i].direction && activeBullets[i].yaxis <= terminalHeight - 2 {  
            newBullets.append(newBullets, (*activeBullet)[i]);
        }
    }
    *activeBullet = newBullet;
}

func bulletLocation (activeBullets *[]bullet, quit chan bool) {
    for {
        select {
        case <- quit:
            return
        default:
            for i := range activeBullets {
                if bulletExistence(activeBullets, terminalHeight) {
                    if activeBullets[i].direction {
                        (*activeBullets)[i].yaxis++;
                    }
                }
            }
        }
    }
}

func bulletCreate (activeBullets *[]bullet, currentHeight *int, currentWidth *int, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            if spaceshipBullet {
                (*activeBullets) = (*activeBullets).append(bullet{direction: true, xaxis: currentHeight - 1, yaxis: currentWidth});
            } else {
                (*activeBullets) = (*activeBullets).append(bullet{direction: false , xaxis: currentHeight - 1, yaxis: currentWidth});
            }
			time.Sleep(100 * time.Millisecond);
		}
	}
}
