package main

import (
    "time"
)

func BulletLocation (activeBullets *[]bullet, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return
        default:
            var newBullets []bullet;
            for i := range *activeBullets {
                if ((*activeBullets)[i].direction && (*activeBullets)[i].height >= 2) || (!(*activeBullets)[i].direction && (*activeBullets)[i].height <= terminalHeight - 1) {  
                    (*activeBullets)[i].height--;
                    newBullets = append(newBullets, (*activeBullets)[i]);
                }
            }
            *activeBullets = newBullets;
        }
        time.Sleep(100 * time.Millisecond);
    }
}

func BulletCreate (activeBullets *[]bullet, currentHeight *int, currentWidth *int, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            (*activeBullets) = append((*activeBullets), bullet{direction: true, height: *currentHeight - 1, width: *currentWidth});
			time.Sleep(100 * time.Millisecond);
		}
	}
}
