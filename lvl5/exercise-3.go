package main

import "fmt"

func exec3() {

	type vehicle struct {
		doors int
		color string
	}

	type pickupTruck struct {
		fourWeels bool
		vehicle
	}

	type sedan struct {
		luxuryModel bool
		vehicle
	}

	monsterTruck := pickupTruck{
		fourWeels: true,
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
	}

	bmw := sedan{
		luxuryModel: true,
		vehicle: vehicle{
			doors: 4,
			color: "Light-Blue",
		},
	}

	fmt.Println(monsterTruck)
	fmt.Println(monsterTruck.vehicle.doors)
	fmt.Println(bmw)
	fmt.Println(bmw.luxuryModel)

	fmt.Println()
}
