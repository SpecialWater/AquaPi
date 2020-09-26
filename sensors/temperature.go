package main //required of any executable program

import ( //our Go packages for this project
	"fmt"
	"github.com/yryz/ds18b20"
)

func Temp(sensors []string) {

	for {
		sensor := sensors[0]
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			t = t*9/5 + 32
			fmt.Printf("temperature: %.2f°F\n", t)
		}
	}
}

