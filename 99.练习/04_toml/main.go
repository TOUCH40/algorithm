package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type songInfo struct {
	Name     string
	Duration int
}

type config struct {
	Bc   string
	Song songInfo
}

func test_toml() {
	c := recover
	var cg config
	var cpath string = "./example.toml"
	if _, err := toml.DecodeFile(cpath, &cg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v %v\n", cg.Bc, cg.Song.Name)
}

func main() {
	test_toml()
}
