package main

import "github.com/ryomak/musicbox/cmd"
import "os"

func main() {
	app := cmd.New("musicbox", "BGM", "1.0.0")
	app.Run(os.Args)
}
