package main //required of any executable program

import ( //our Go packages for this project
	"fmt"
	"github.com/yryz/ds18b20"
	"github.com/stianeikeland/go-rpio/v4"
	"os"
)

var (
	pinBlue = rpio.Pin(21)
	pinRed = rpio.Pin(26)
)

func Temp(sensors []string) {
	defer rpio.Close()

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	// Set pin to output mode
	pinBlue.Output()
	pinRed.Output()

	for {
		sensor := sensors[0]
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			t = t*9/5 + 32
			fmt.Printf("temperature: %.2f°F\n", t)
			if t > 85 {
				BlinkRed(pinRed)
			} else if t < 70 {
				BlinkBlue(pinBlue)
			}
		}
	}
}

