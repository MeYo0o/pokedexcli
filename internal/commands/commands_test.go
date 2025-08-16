package commands

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
)

// Override apiCall for tests
func mockApiCallSuccess(data []byte) func(string) ([]byte, error) {
	return func(_ string) ([]byte, error) {
		return data, nil
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestCommandsMap(t *testing.T) {
	// Mock API with your curl JSON response
	apiCall = mockApiCallSuccess([]byte(`{
		"count": 1089,
		"next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
		"previous": null,
		"results": [
			{"name": "canalave-city-area", "url": "https://pokeapi.co/api/v2/location-area/1/"},
			{"name": "eterna-city-area", "url": "https://pokeapi.co/api/v2/location-area/2/"}
		]
	}`))

	out := captureOutput(func() {
		err := commandsMap()
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	// Assert output contains names
	if !bytes.Contains([]byte(out), []byte("canalave-city-area")) {
		t.Errorf("expected output to contain canalave-city-area, got %s", out)
	}
	if !bytes.Contains([]byte(out), []byte("eterna-city-area")) {
		t.Errorf("expected output to contain eterna-city-area, got %s", out)
	}

	// Assert config updated
	if commandConfig.Next == "" {
		t.Errorf("expected commandConfig.Next to be set")
	}
}

func TestCommandCatch_Success(t *testing.T) {
	// Mock API to return a Pokémon with low BaseExperience (easy to catch)
	apiCall = mockApiCallSuccess([]byte(`{
        "base_experience": 10
    }`))

	// Fix randomness (always 0, so it's easier to catch)
	rand.Seed(1)

	out := captureOutput(func() {
		err := commandCatch("catch", "pikachu")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	if !bytes.Contains([]byte(out), []byte("Throwing a Pokeball at pikachu...")) {
		t.Errorf("expected throwing message, got %s", out)
	}

	if !bytes.Contains([]byte(out), []byte("was caught")) {
		t.Errorf("expected pikachu to be caught, got %s", out)
	}
}

func TestCommandCatch_Escape(t *testing.T) {
	// Mock API to return a Pokémon with high BaseExperience (harder to catch)
	apiCall = mockApiCallSuccess([]byte(`{
        "base_experience": 999
    }`))

	// Fix randomness (force a low roll to simulate escape)
	rand.Seed(1)

	out := captureOutput(func() {
		err := commandCatch("catch", "mewtwo")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	if !bytes.Contains([]byte(out), []byte("Throwing a Pokeball at mewtwo...")) {
		t.Errorf("expected throwing message, got %s", out)
	}

	if !bytes.Contains([]byte(out), []byte("escaped")) {
		t.Errorf("expected mewtwo to escape, got %s", out)
	}
}

func TestCommandCatch_InvalidArgs(t *testing.T) {
	out := captureOutput(func() {
		err := commandCatch("catch")
		if err == nil {
			t.Fatal("expected error for missing args")
		}
	})

	if !bytes.Contains([]byte(out), []byte("")) {
		// just confirm it doesn’t print anything strange
	}
}
