package main //required of any executable program

import ( //our Go packages for this project
	"time"
	"github.com/stianeikeland/go-rpio/v4"
	"github.com/yryz/ds18b20"
	"fmt"
	"os"

)

var (
	pin = rpio.Pin(21)
)

func main() {

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
	pin.Output()
	
	go Temp(sensors)
	go Blink(pin)

	for {
		time.Sleep(time.Second / 5)
	}

}
