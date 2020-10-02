package main //required of any executable program

import ( //our Go packages for this project
	"time"
	"github.com/yryz/ds18b20"
)



func main() {

	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}
	
	go Temp(sensors)

	for {
		time.Sleep(time.Second / 5)
	}

}
