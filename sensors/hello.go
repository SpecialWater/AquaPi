package main //required of any executable program

import ( //our Go packages for this project
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
	"github.com/yryz/ds18b20"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(21)
)

func temp(sensors []string) {
	for {
		sensor := sensors[0]
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			t = t*9/5 + 32
			fmt.Printf("temperature: %.2f°F\n", t)
		}
	}
}

func blink() {
	// Toggle pin 20 times
	for {
		pin.Toggle()
		fmt.Println("sleep")
		time.Sleep(time.Second / 5)
	}
}

func main() {
	// Unmap gpio memory when done
	defer rpio.Close()

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set pin to output mode
	pin.Output()

	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	go temp(sensors)
	go blink()

	for {
		time.Sleep(time.Second / 5)
	}

}
