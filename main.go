package main

import (
	"fmt",
	"go_load_balancer/config"
)

func loadConfig(file string) (Config, error) {
	var config Config

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {
	fmt.Println("Hello, World!")
}
