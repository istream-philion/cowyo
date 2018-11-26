package main

import (
	"github.com/istream-philion/cowyo/config"
)

func main() {
	c, err := config.ParseFile("multisite_sample.toml")
	if err != nil {
		panic(err)
	}

	panic(c.ListenAndServe())
}
