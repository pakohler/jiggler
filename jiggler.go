package main

import (
	"flag"
	"math/rand"
	"math"
	"time"
	"log"
	"os"
	"github.com/go-vgo/robotgo"
)

func main() {
	var jigglemult = flag.Int("pixels", 5, "maximum amount of x or y jiggle")
	var delay = flag.Duration("delay", 15*time.Second, "time between jiggles")
	var click = flag.Bool("click", false, "whether to click or not")
	flag.Parse()
	log := log.New(os.Stdout, "jiggler: ", log.Ldate | log.Ltime | log.Lshortfile)
	for {
		dx := int(math.Round(rand.NormFloat64() * float64(*jigglemult)))
		dy := int(math.Round(rand.NormFloat64() * float64(*jigglemult)))
		log.Printf("moving mouse %dpx horizontally and %dpx vertically", dx, dy)
		robotgo.MoveSmoothRelative(dx, dy)
		if *click {
			robotgo.Click()
		}
		time.Sleep(1 * *delay)
	}
}
