package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartEventMarshaling(t *testing.T) {

	// read json from file
	inputJSON, err := ioutil.ReadFile("./testdata/startEvent.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	// de-serialize into CognitoEvent
	var inputEvent RequestEnvelope
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	// serialize to json
	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	assert.JSONEq(t, string(inputJSON), string(outputJSON))
}

func TestIntentEventMarshaling(t *testing.T) {

	// read json from file
	inputJSON, err := ioutil.ReadFile("./testdata/intent-mycoloris.json")
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	// de-serialize into CognitoEvent
	var inputEvent RequestEnvelope
	if err := json.Unmarshal(inputJSON, &inputEvent); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}
	assert.Equal(t, inputEvent.Request.Intent.Slots["Color"].Value, "blue")

	// serialize to json
	outputJSON, err := json.Marshal(inputEvent)
	if err != nil {
		t.Errorf("could not marshal event. details: %v", err)
	}

	assert.JSONEq(t, string(inputJSON), string(outputJSON))
}
