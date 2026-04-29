package request

import (
	"encoding/json"
	"testing"
)

func TestNewV3RequestIncludesMarkdownFilterAdditions(t *testing.T) {
	req := NewV3Request("run **grep** now", "voice-demo", "mp3", 24000, "seed-tts-2.0-standard", "req-id")

	raw, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal request error = %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal request error = %v", err)
	}

	reqParams, ok := payload["req_params"].(map[string]any)
	if !ok {
		t.Fatalf("req_params missing: %#v", payload)
	}
	if reqParams["text"] != "run **grep** now" {
		t.Fatalf("req_params.text = %#v", reqParams["text"])
	}

	additions, ok := payload["Additions"].(map[string]any)
	if !ok {
		t.Fatalf("Additions missing: %#v", payload)
	}
	if additions["disable_markdown_filter"] != true {
		t.Fatalf("disable_markdown_filter = %#v", additions["disable_markdown_filter"])
	}
}
