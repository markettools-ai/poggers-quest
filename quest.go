package main

import (
	"encoding/json"
	"fmt"

	"github.com/markettools-ai/poggers"
)

type Quest struct {
	Name  string                   `json:"name"`
	NPC   map[string]interface{}   `json:"npc"`
	Steps []map[string]interface{} `json:"steps"`
	Loot  []map[string]interface{} `json:"loot"`
}

var quest Quest
var promptBuilder poggers.PromptBuilder

// Generates a quest based on the provided request body
func GenerateQuest(name string) (Quest, error) {
	// Create a new prompt builder
	promptBuilder = poggers.NewPromptBuilder(
		poggers.PromptBuilderOptions{
			Annotations: map[string]string{
				"QuestInfo": "the quest information",
				"input":     "The quest name is \"" + name + "\".", // Could also be a JSON string
			},
			// Callbacks for processing the prompts
			OnBeforeProcess: nil,
			OnAfterProcess:  PromptCallback,
		},
	)
	// Process prompts
	err := promptBuilder.ProcessBatchFromDir("./quest")
	quest.Name = name
	return quest, err
}

// Callback function that processes the result of a prompt
func PromptCallback(name string, constants map[string]string, messages []poggers.Message) error {
	// Define model
	model := "gpt-4o"
	if m, ok := constants["MODEL"]; ok {
		model = m
	}
	// Send the messages to the OpenAI API
	result, err := SendMessages(messages, model)
	if err != nil {
		return err
	}
	// Switch between prompts
	switch name {
	case "loot":
		return handleLoot(result)
	case "npc":
		return handleNPC(result)
	case "steps":
		return handleSteps(result)
	}
	return fmt.Errorf("unknown prompt name: %s", name)
}

// handleLoot processes the loot prompt result
func handleLoot(result string) error {
	var loot []map[string]interface{}
	if err := json.Unmarshal([]byte(result), &loot); err != nil {
		return fmt.Errorf("failed to unmarshal loot: %w", err)
	}
	quest.Loot = loot
	promptBuilder.SetAnnotation("loot", result)
	return nil
}

// handleNPC processes the NPC prompt result
func handleNPC(result string) error {
	var npc map[string]interface{}
	if err := json.Unmarshal([]byte(result), &npc); err != nil {
		return fmt.Errorf("failed to unmarshal npc: %w", err)
	}
	quest.NPC = npc
	promptBuilder.SetAnnotation("npc", result)
	return nil
}

// handleSteps processes the steps prompt result
func handleSteps(result string) error {
	var steps []map[string]interface{}
	if err := json.Unmarshal([]byte(result), &steps); err != nil {
		return fmt.Errorf("failed to unmarshal steps: %w", err)
	}
	quest.Steps = steps
	return nil
}
