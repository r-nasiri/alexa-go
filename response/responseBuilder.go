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

// GetResponse returns the final response
func (rb ResponseBuilder) GetResponse() model.Response {
	return rb.Response
}
