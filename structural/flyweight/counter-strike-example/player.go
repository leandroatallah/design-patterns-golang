package main

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dresstype string) *Player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dresstype)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
