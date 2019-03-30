package model

// RequestEnvelope represents the top level data that is sent to the skill.
type RequestEnvelope struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Request *Request `json:"request"`
	Context *Context `json:"context"`
}

// Session represents the session data from the request.
// It provides additional context associated with the request.
type Session struct {
	New        bool                   `json:"new"`
	SessionID  string                 `json:"sessionId"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
}

// Context represents provides your skill with information about the current state
// of the Alexa service and device at the time the request is sent to your service.
type Context struct {
	System struct {
		Device struct {
			DeviceID            string `json:"deviceId,omitempty"`
			SupportedInterfaces struct {
				AudioPlayer struct {
				} `json:"AudioPlayer,omitempty"`
			} `json:"supportedInterfaces"`
		} `json:"device"`
		Application struct {
			ApplicationID string `json:"applicationId"`
		} `json:"application"`
		User struct {
			UserID      string          `json:"userId"`
			AccessToken string          `json:"accessToken,omitempty"`
			Permissions UserPermissions `json:"permissions,omitempty"`
		} `json:"user"`
		APIEndpoint    string `json:"apiEndpoint,omitempty"`
		APIAccessToken string `json:"apiAccessToken,omitempty"`
	} `json:"System"`
	AudioPlayer struct {
		PlayerActivity       string `json:"playerActivity"`
		Token                string `json:"token,omitempty"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
	} `json:"AudioPlayer"`
}

type UserPermissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

// Request represents the details of the user's request.
// It should contain all different request types.
type Request struct {
	LaunchRequest
	IntentRequest
	SkillDisabledRequest
	Type      string `json:"type"`
	RequestID string `json:"requestId"`
	Locale    string `json:"locale"`
	Timestamp string `json:"timestamp"`
}

// LaunchRequest represents that a user made a request to an Alexa skill, but did not provide a specific intent.
type LaunchRequest struct {
}

// IntentRequest represents a request made to a skill based on what the user wants to do.
type IntentRequest struct {
	Intent *Intent `json:"intent,omitempty"`
}

// SkillDisabledRequest
type SkillDisabledRequest struct {
	EventCreationTime   string `json:"eventCreationTime,omitempty"`
	EventPublishingTime string `json:"eventPublishingTime,omitempty"`
}

// Intent represents what the user wants.
type Intent struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus,omitempty"`
	Slots              map[string]IntentSlot `json:"slots"`
}

const (
	ConfirmationStatusNone      string = "NONE"
	ConfirmationStatusDenied    string = "DENIED"
	ConfirmationStatusConfirmed string = "CONFIRMED"
)

// IntentSlot contains the data for one Slot
type IntentSlot struct {
	Name               string       `json:"name"`
	ConfirmationStatus string       `json:"confirmationStatus,omitempty"`
	Value              string       `json:"value"`
	Resolutions        *Resolutions `json:"resolutions,omitempty"`
}

// Resolutions contains the results of resolving the words captured from the user's utterance.
type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority"`
		Status    struct {
			Code string `json:"code"`
		} `json:"status"`
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
