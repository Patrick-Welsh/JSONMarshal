package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Define map of strings; key value pairs
	// String map to string key
	data := map[string]string{
		"Black sea nettle":    "Chrysaora achlyos",
		"Atlantic sea nettle": "Chrysaora fuscescens",
	}

	b, err := json.MarshalIndent(data, "", "")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}
