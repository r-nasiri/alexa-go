package model

type PlainTextHint struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type HintDirective struct {
	Type string        `json:"type"`
	Hint PlainTextHint `json:"hint"`
}

type ConfirmIntentDirective struct {
	Type          string  `json:"type"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
}

type ConfirmSlotDirective struct {
	Type          string  `json:"type"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
	SlotToConfirm string  `json:"slotToConfirm"`
}

type DelegateDirective struct {
	Type          string  `json:"type"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
}

type ElicitSlotDirective struct {
	Type          string  `json:"type"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
	SlotToElicit  string  `json:"slotToElicit"`
}
