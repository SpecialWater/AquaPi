package main //required of any executable program

import ( //our Go packages for this project
	"fmt"
	"time"
	//"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func Blink(pin rpio.Pin) {
	
	// Toggle pin 20 times
	for {
		pin.Toggle()
		fmt.Println("sleep")
		time.Sleep(time.Second / 5)
	}
}

