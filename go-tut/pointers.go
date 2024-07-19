package main

import (
	"fmt"
)

type Player struct {
	health	int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}


func _() {
	player := Player{health: 100}
	fmt.Printf("before explosion %+v\n", player)
	// takeDamageFromExplosion(player, 10);
	player.takeDamageFromExplosion(10);
	fmt.Printf("before explosion %+v\n", player)
}
