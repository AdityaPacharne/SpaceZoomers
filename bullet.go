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
            bulletMutex.Lock();
            rockMutex.Lock();
            var newBullets []bullet;
            for i := range *activeBullets {
                if (*activeBullets)[i].direction && (*activeBullets)[i].height >= 2 {  
                    var flag bool = true;
                    for j := range *activeRocks {
                        if (*activeBullets)[i].height-1 == (*activeRocks)[j].height && (*activeBullets)[i].width == (*activeRocks)[j].width {
                            flag = false
                            break
                        }
                    }
                    if flag {
                        (*activeBullets)[i].height--;
                        newBullets = append(newBullets, (*activeBullets)[i])
                    }
                }
            }
            *activeBullets = newBullets;
            rockMutex.Unlock();
            bulletMutex.Unlock();
        }
        time.Sleep(50 * time.Millisecond);
    }
}

func BulletCreate (activeBullets *[]bullet, spaceship *spaceshipstruct, spaceshipBullet bool, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
            spaceshipMutex.Lock();
            var newBullet bullet = bullet{
                direction: true,
                height: (*spaceship).height - 1,
                width: (*spaceship).width,
            }
            spaceshipMutex.Unlock();

            bulletMutex.Lock();
            (*activeBullets) = append((*activeBullets), newBullet);
            bulletMutex.Unlock();
			time.Sleep(200 * time.Millisecond);
		}
	}
}
