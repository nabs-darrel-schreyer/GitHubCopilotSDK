package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	copilot "github.com/github/copilot-sdk/go"
)

const (
	green = "\x1b[32m"
	reset = "\x1b[0m"
)

// resolveCopilotCLI finds a Copilot CLI binary. The Go SDK does not auto-bundle
// a runtime like Node/.NET; without an embedded bundle it needs COPILOT_CLI_PATH
// or an explicit StdioConnection.Path (PATH alone is not enough for managed start).
func resolveCopilotCLI() string {
	if p := os.Getenv("COPILOT_CLI_PATH"); p != "" {
		return p
	}
	if p, err := exec.LookPath("copilot"); err == nil {
		return p
	}
	if runtime.GOOS == "windows" {
		home, err := os.UserHomeDir()
		if err == nil {
			candidates := []string{
				filepath.Join(home, "AppData", "Local", "GitHubCopilotCLI", "copilot.exe"),
				filepath.Join(home, "AppData", "Local", "GitHub CLI", "copilot", "copilot.exe"),
				filepath.Join(home, "AppData", "Roaming", "npm", "copilot.cmd"),
			}
			for _, c := range candidates {
				if st, err := os.Stat(c); err == nil && !st.IsDir() {
					return c
				}
			}
		}
	}
	return ""
}

func newClient() *copilot.Client {
	if path := resolveCopilotCLI(); path != "" {
		return copilot.NewClient(&copilot.ClientOptions{
			Connection: copilot.StdioConnection{Path: path},
		})
	}
	return copilot.NewClient(nil)
}

func assistantContent(response *copilot.SessionEvent) (string, error) {
	if response == nil {
		return "", fmt.Errorf("no assistant response received")
	}

	data, ok := response.Data.(*copilot.AssistantMessageData)
	if !ok {
		return "", fmt.Errorf("expected assistant message data, got %T", response.Data)
	}

	return data.Content, nil
}

func main() {
	fmt.Println("Starting a GitHub Copilot session...")

	ctx := context.Background()
	client := newClient()
	if err := client.Start(ctx); err != nil {
		log.Fatalf("%v\nHint: install GitHub Copilot CLI and set COPILOT_CLI_PATH to copilot.exe, or ensure it is discoverable under %%LOCALAPPDATA%%\\GitHubCopilotCLI\\", err)
	}
	defer client.Stop()

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model:               "auto",
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Disconnect()

	fmt.Println("Sending message to GitHub Copilot...")
	response, err := session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: "Calculate the product of 2 and 4.",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received response from GitHub Copilot...")

	content, err := assistantContent(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Response content: %s%s%s\n", green, content, reset)
	fmt.Println("GitHub Copilot session ended.")
	os.Exit(0)
}
