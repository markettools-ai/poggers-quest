package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func main() {
	// Parse the command line arguments
	questName := flag.String("name", "Hunting down the Wendigo", "a descriptive name for your quest")
	openAIAPIKey = flag.String("key", "YOUR_OPENAI_API_KEY", "your OpenAI API key")
	flag.Parse()

	// Generate the quest
	quest, err := GenerateQuest(*questName)
	if err != nil {
		panic(err)
	}

	// Pretty print the generated quest
	questJSON, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(questJSON))
}
