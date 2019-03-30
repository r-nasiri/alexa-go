package alexa

import (
	"context"
	"fmt"

	"github.com/r-nasiri/alexa-go/model"
	"github.com/r-nasiri/alexa-go/response"
)

// lambdaFunctionType is the Lambda handler signature
type lambdaFunctionType = func(ctx context.Context, re model.RequestEnvelope) (model.ResponseEnvelope, error)

// RequestHandler interface  should be fulfilled by skill's request handlers
type RequestHandler interface {
	CanHandle(handlerInput HandlerInput) bool
	Handle(HandlerInput HandlerInput) (model.Response, error)
}

// CustomSkill represents a collection of request handlers for the skill
type CustomSkill struct {
	handlers []RequestHandler
}

// AddRequestHandlers adds request handlers to the skill
func (cs *CustomSkill) AddRequestHandlers(handlers ...RequestHandler) *CustomSkill {
	cs.handlers = append(cs.handlers, handlers...)
	return cs
}

// Lambda returns a function that implements lambda function
func (cs CustomSkill) Lambda() lambdaFunctionType {
	return func(ctx context.Context, re model.RequestEnvelope) (model.ResponseEnvelope, error) {
		responseEnvelope := model.ResponseEnvelope{
			Version: "1.0",
		}
		for _, handler := range cs.handlers {
			handlerInput := HandlerInput{
				RequestEnvelope:   re,
				AttributesManager: BasicAttributeManger{Attributes: re.Session.Attributes},
			}
			if handler.CanHandle(handlerInput) {
				fmt.Println("handeling request", re.Session.SessionID)
				response, err := handler.Handle(handlerInput)
				responseEnvelope.Response = &response
				responseEnvelope.SessionAttributes = handlerInput.AttributesManager.GetSessionAttributes()
				return responseEnvelope, err
			}
		}

		return responseEnvelope, fmt.Errorf("Cannot find any handlers")

	}

}

type AttributesManager interface {
	GetSessionAttributes() map[string]interface{}
	SetSessionAttributes(map[string]interface{})
}

// HandlerInput represents the struct that is passed to each request handler
type HandlerInput struct {
	ResponseBuilder   response.ResponseBuilder
	RequestEnvelope   model.RequestEnvelope
	AttributesManager AttributesManager
}
