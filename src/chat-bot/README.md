# Chat-Bot Application

Welcome to the Chat-Bot application! This project is designed to provide an interactive chat experience using various APIs.

## Features

* Chat with an LLM via the terminal
* Keeps conversation in context window, provides option to clear conversation
* Swap between Anthropic, OpenAI, and a Mock for the chat client

## Setting API Keys

To use the Chat-Bot application, you need to set up environment variables of API keys for the services you intend to use.

```bash
export OPENAI_API_KEY=<your_key>
export ANTHROPIC_API_KEY=<your_key>
```

## Getting Started

To the run application from this directory run:

```sh
go run main.go
```
