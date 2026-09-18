package main

import (
	"testing"

	copilot "github.com/github/copilot-sdk/go"
)

func TestAssistantContentReturnsAssistantMessageContent(t *testing.T) {
	content, err := assistantContent(&copilot.SessionEvent{
		Data: &copilot.AssistantMessageData{
			Content:   "done",
			MessageID: "assistant-1",
		},
	})
	if err != nil {
		t.Fatalf("assistantContent returned error: %v", err)
	}
	if content != "done" {
		t.Fatalf("assistantContent returned %q, want %q", content, "done")
	}
}

func TestAssistantContentRejectsNilResponse(t *testing.T) {
	if _, err := assistantContent(nil); err == nil {
		t.Fatal("assistantContent returned nil error for nil response")
	}
}

func TestAssistantContentRejectsUnexpectedResponseData(t *testing.T) {
	if _, err := assistantContent(&copilot.SessionEvent{Data: &copilot.SessionIdleData{}}); err == nil {
		t.Fatal("assistantContent returned nil error for unexpected response data")
	}
}
