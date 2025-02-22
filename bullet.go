package main

import (
    "time"
)

func BulletLocation (activeBullets *[]bullet, activeRocks *[]rocks, terminalHeight int, quit chan bool) {
    for {
        select {
        case <- quit:
            return
        default:
            var newBullets []bullet;
            for i := range len(*activeBullets) {
                if (*activeBullets)[i].direction && (*activeBullets)[i].height >= 2 {  
                    var flag bool = true;
                    for j := range len(*activeRocks) {
                        if (*activeBullets)[i].height-1 == (*activeRocks)[j].height && (*activeBullets)[i].width == (*activeRocks)[j].width {
                            flag = false
                            break
                        }
                    }
                    if flag {
                        (*activeBullets)[i].height-- // Move bullet if no collision
                        newBullets = append(newBullets, (*activeBullets)[i])
                    }
                }
            }
            *activeBullets = newBullets;
        }
        time.Sleep(100 * time.Millisecond);
    }
}

func BulletCreate (activeBullets *[]bullet, spaceship *spaceshipstruct, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            (*activeBullets) = append((*activeBullets), bullet{direction: true, height: (*spaceship).height- 1, width: (*spaceship).width});
			time.Sleep(100 * time.Millisecond);
		}
	}
}
