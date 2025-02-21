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
            var flag bool = true;
            for i := range len(*activeBullets) {
                if (*activeBullets)[i].direction && (*activeBullets)[i].height >= 2 {  
                    for j := range len(*activeRocks) {
                        if ((*activeBullets)[i].height-1 != (*activeRocks)[j].height) || ((*activeBullets)[i].height+1 != (*activeRocks)[j].height){
                            (*activeBullets)[i].height--;
                            flag = true;
                        } else {
                            flag = false;
                            break;
                        }
                    }
                    if flag {
                        newBullets = append(newBullets, (*activeBullets)[i]);
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
