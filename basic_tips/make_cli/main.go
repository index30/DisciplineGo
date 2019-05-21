package main
import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprintf(os.Stderr, "[usage] %s menu1 menu2...", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf(os.Args[rand.Intn(c) + 1])
}