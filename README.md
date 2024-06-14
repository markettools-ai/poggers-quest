<p align="center">
  <img src="https://github.com/markettools-ai/poggers-quest/assets/20731019/e9099fe5-62a9-4a71-a08e-a2b6f5a97cd2" width="100">
  <h1 align="center">Welcome to poggers-quest!</h1>
</p>

This repository showcases the usage of [poggers](https://github.com/markettools-ai/poggers), a library and its IDL for creating and managing AI prompts.

> Make sure to download poggers' [VSCode Extension](https://marketplace.visualstudio.com/items?itemName=markettools-ai.poggers-prompt)!

## Pre-requisites
- [Golang 1.16+](https://golang.org/dl/)
- An [OpenAI](https://platform.openai.com/signup) API key

## Usage
Simply run:
```bash
go run . -name="In Search of the Lost Ice Cream Temple" -key="YOUR_OPENAI_API_KEY"
```
to get a response like this:
```json
{
  "name": "In Search of the Lost Ice Cream Temple",
  "npc": {
    "gender": "female",
    "name": "Elaria Snow",
    "profession": "explorer",
    "race": "elf"
  },
  "steps": [
    {
      "description": "Speak with Elaria Snow to begin the quest.",
      "line": "Hello adventurer, I need your help to find the legendary Lost Ice Cream Temple.",
      "step": "talk to npc"
    },
    {
      "description": "Travel to the Frozen Tundra to find clues about the temple's location.",
      "step": "go to location"
    },
    {
      "description": "Defeat the Ice Goblins guarding the entrance to the temple.",
      "step": "kill enemy"
    },
    {
      "description": "Confront and defeat the Frost Giant inside the temple.",
      "step": "kill greater enemy"
    },
    {
      "description": "Return to Elaria Snow to report your success and claim your reward.",
      "line": "Thank you, brave hero! Here's your reward: the Frostbite Blade, some gold coins, a Temple Amulet, and a Glacial Shield.",
      "step": "talk to npc"
    }
  ],
  "loot": [
    {
      "description": "A sword that deals extra frost damage.",
      "name": "Frostbite Blade",
      "type": "sword"
    },
    {
      "amount": 12,
      "description": "Shiny gold coins.",
      "type": "gold"
    },
    {
      "description": "An amulet that guides you to the nearest ice cream.",
      "name": "Temple Amulet",
      "type": "misc"
    },
    {
      "description": "A shield that protects against frost-based attacks.",
      "name": "Glacial Shield",
      "type": "armor"
    }
  ]
}
```

## Explanation
The idea of this library is to optimize the process of creating AI prompts while keeping them human-readable and consistent. In this case, there are three different generations to achieve the full quest object:
- The Loot, rewarded at the end of the quest,
- The NPC, who gives the quest to the player,
- and The Steps, a list of actions that tie the beginning and the end of the quest together.

### IDL
The `./quest/*.prompt` files are written in an IDL (Interface Definition Language) that describes the structure of the prompt. This IDL supports comments, schema definitions, dynamic annotations, and more.

As we're interacting with OpenAI, the prompt uses the system-user-assistant pattern.

> Note: comments in the prompt files are also passed to the AI. This can be useful for debugging or providing additional context or instructions.

### Prompts Order and Dependency
Some prompts depend on others to be generated. For example, `1_steps.prompt` depends on both `0_npc.prompt` and `0_loot.prompt`, which are independent from each other and, therefore, generated concurrently.

By adding a prefix to the prompt files, we can define the order and concurrency which they should be generated by **poggers**.

### Templating
The IDL supports templating through annotations. For example, the result of both `0_npc.prompt` and `0_loot.prompt` are used in `1_steps.prompt` as `@npc` and `@loot`, respectively.

These values are passed in to the prompt builder in one of the callback functions.

All prompts also use the `@input` annotation (the quest name), which is passed right at the beginning of the execution.

> Note: not all annotations are explicitly defined on this code. Some are default values, for example, `@OutputSchema`, which is replaced by `"This is the output schema:"`. You can see those in the `annotations.go` file in the `poggers` library.