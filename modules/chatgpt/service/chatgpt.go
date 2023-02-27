package service

type ChatGPTService struct {
	OPENAI_API_KEY      string
	OPENAI_Organization string
}

// NewChatGPTService creates and returns a new ChatGPTService object.
func NewChatGPTService(apikey string, organization string) *ChatGPTService {
	return &ChatGPTService{
		OPENAI_API_KEY:      apikey,
		OPENAI_Organization: organization,
	}
}
