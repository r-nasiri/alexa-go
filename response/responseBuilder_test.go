package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElicitDirective(t *testing.T) {
	rb := ResponseBuilder{}
	resp := rb.AddElicitSlotDirective("test", nil).GetResponse()
	outputJSON, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("could not marshal response. details: %v", err)
	}
	assert.Contains(t, string(outputJSON), "slotToElicit")
	assert.Contains(t, string(outputJSON), "Dialog.ElicitSlot")
}

func TestMultiDirectives(t *testing.T) {
	rb := ResponseBuilder{}
	resp := rb.AddElicitSlotDirective("coffeeRoast", nil).
		AddConfirmSlotDirective("flavor", nil).
		GetResponse()
	outputJSON, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("could not marshal response. details: %v", err)
	}
	assert.Contains(t, string(outputJSON), "coffeeRoast")
	assert.Contains(t, string(outputJSON), "slotToElicit")
	assert.Contains(t, string(outputJSON), "Dialog.ElicitSlot")

	assert.Contains(t, string(outputJSON), "flavor")
	assert.Contains(t, string(outputJSON), "slotToConfirm")
	assert.Contains(t, string(outputJSON), "Dialog.ConfirmSlot")
}
