package main //required of any executable program

import ( //our Go packages for this project
	"fmt"
	"time"
	//"os"

	"github.com/stianeikeland/go-rpio/v4"
)

func BlinkBlue(pin rpio.Pin) {
	
	for i := 0; i < 4; i++ {
		pin.Toggle()
		fmt.Println("Blue Blink")
		time.Sleep(time.Second / 5)
	}
}

func BlinkRed(pin rpio.Pin) {
	
	for i := 0; i < 4; i++  {
		pin.Toggle()
		fmt.Println("Red Blink")
		time.Sleep(time.Second / 5)
	}
}

