package main

import (
	_ "Application/Game"
	_ "Drive/Disk/Model"
	_ "Drive/Http/Model"
	_ "Drive/MySQL/Model"
	"fmt"
	"path/filepath"
	"os"
	"strings"
	"Drive"
)

var lof = fmt.Println

func main() {

	Drive.Init("D:/workspace/games/FishGame/")
	//Drive.Init(getCurrentDirectory())
	select {}
}
func getCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)+"/"
}