package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/r-nasiri/alexa-go"
	"github.com/r-nasiri/alexa-go/model"
)

type startHandler struct{}

func (sh startHandler) CanHandle(handlerInput alexa.HandlerInput) bool {
	return handlerInput.RequestEnvelope.Request.Type == "LaunchRequest"

}

func (sh startHandler) Handle(handlerInput alexa.HandlerInput) (model.Response, error) {
	responseBuilder := handlerInput.ResponseBuilder
	speechText := "Welcome to Alexa skill demo."
	responseBuilder.
		Speak(speechText).
		Reprompt(speechText)
	return responseBuilder.GetResponse(), nil
}

type helloHandler struct{}

func (sh helloHandler) CanHandle(handlerInput alexa.HandlerInput) bool {
	return handlerInput.RequestEnvelope.Request.Type == "IntentRequest" && handlerInput.RequestEnvelope.Request.Intent.Name == "HelloWorldIntent"

}

func (sh helloHandler) Handle(handlerInput alexa.HandlerInput) (model.Response, error) {
	responseBuilder := handlerInput.ResponseBuilder
	speechText := "Hello World!"
	responseBuilder.
		Speak(speechText).
		Reprompt(speechText)
	return responseBuilder.GetResponse(), nil
}

func main() {

	var alexaSkill alexa.CustomSkill

	alexaSkill.AddRequestHandlers(helloHandler{}, startHandler{})

	lambda.Start(alexaSkill.Lambda())
}
