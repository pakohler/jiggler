package main

import (
	"math/rand"
	"math"
	"time"
	"log"
	"os"
	"github.com/go-vgo/robotgo"
)

const jigglemult = 5
const delay = 15 * time.Second

func main() {
	log := log.New(os.Stdout, "jiggler: ", log.Ldate | log.Ltime | log.Lshortfile)
	for {
		dx := int(math.Round(rand.NormFloat64() * jigglemult))
		dy := int(math.Round(rand.NormFloat64() * jigglemult))
		log.Printf("moving mouse %dpx horizontally and %dpx vertically", dx, dy)
		robotgo.MoveSmoothRelative(dx, dy)
		time.Sleep(1 * delay)
	}
}
