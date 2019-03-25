package response

import "github.com/r-nasiri/alexa-go/model"

// ResponseBuilder represents the struct that builds response
type ResponseBuilder struct {
	Response model.Response
}

// Speak sets the speech that needs to be rendered to the user.
func (rb *ResponseBuilder) Speak(speechOutput string) *ResponseBuilder {
	rb.Response.OutputSpeech = &model.OutputSpeech{
		Type: "SSML",
		SSML: "<speak>" + speechOutput + "</speak>",
	}
	return rb
}

// Reprompt sets the output speech to use if a re-prompt is necessary.
func (rb *ResponseBuilder) Reprompt(repromptSpeechOutput string) *ResponseBuilder {
	rb.Response.Reprompt =
		&model.Reprompt{
			OutputSpeech: &model.OutputSpeech{
				Type: "SSML",
				SSML: "<speak>" + repromptSpeechOutput + "</speak>",
			},
		}
	return rb
}

// WithSimpleCard sets the card to be rendered to the app.
func (rb *ResponseBuilder) WithSimpleCard(cardTitle, cardContent string) *ResponseBuilder {
	rb.Response.Card = &model.Card{
		Type:    "Simple",
		Title:   cardTitle,
		Content: cardContent,
	}
	return rb
}

// AddDelegateDirective
func (rb *ResponseBuilder) AddDelegateDirective(intent *model.Intent) *ResponseBuilder {
	delegateDirective := model.DelegateDirective{
		Type:          "Dialog.Delegate",
		UpdatedIntent: intent,
	}
	rb.addDirective(delegateDirective)

	return rb
}

// AddElicitSlotDirective
func (rb *ResponseBuilder) AddElicitSlotDirective(slotToElicit string, intent *model.Intent) *ResponseBuilder {
	elicitSlotDirective := model.ElicitSlotDirective{
		Type:          "Dialog.ElicitSlot",
		UpdatedIntent: intent,
		SlotToElicit:  slotToElicit,
	}
	rb.addDirective(elicitSlotDirective)

	return rb
}

// AddConfirmSlotDirective
func (rb *ResponseBuilder) AddConfirmSlotDirective(slotToConfirm string, intent *model.Intent) *ResponseBuilder {
	confirmSlotDirective := model.ConfirmSlotDirective{
		Type:          "Dialog.ConfirmSlot",
		UpdatedIntent: intent,
		SlotToConfirm: slotToConfirm,
	}
	rb.addDirective(confirmSlotDirective)

	return rb
}

// AddConfirmIntentDirective
func (rb *ResponseBuilder) AddConfirmIntentDirective(intent *model.Intent) *ResponseBuilder {
	confirmIntentDirective := model.ConfirmIntentDirective{
		Type:          "Dialog.ConfirmIntent",
		UpdatedIntent: intent,
	}
	rb.addDirective(confirmIntentDirective)

	return rb
}

// AddHintDirective
func (rb *ResponseBuilder) AddHintDirective(text string) *ResponseBuilder {

	hint := model.PlainTextHint{Type: "PlainText", Text: text}
	hintDirective := model.HintDirective{Type: "Hint", Hint: hint}
	rb.addDirective(hintDirective)

	return rb
}

// addDirective
func (rb *ResponseBuilder) addDirective(directive interface{}) *ResponseBuilder {
	rb.Response.Directives = append(rb.Response.Directives, directive)
	return rb
}

// WithShouldEndSession sets the property that indicates what should happen after Alexa speaks the response.
func (rb *ResponseBuilder) WithShouldEndSession(flag bool) *ResponseBuilder {
	rb.Response.ShouldSessionEnd = flag
	return rb
}

// GetResponse returns the final response
func (rb ResponseBuilder) GetResponse() model.Response {
	return rb.Response
}
