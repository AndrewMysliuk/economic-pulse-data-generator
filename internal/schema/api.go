package schema

import (
	"encoding/json"
	"fmt"
)

type APIErrorResponse struct {
	Error *APIError `json:"error"`
}

type APIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

type OpenAIResponse struct {
	ID     string          `json:"id"`
	Object string          `json:"object"`
	Output []OpenAIMessage `json:"output"`
}

type OpenAIMessage struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Status  string          `json:"status"`
	Role    string          `json:"role,omitempty"`
	Action  json.RawMessage `json:"action,omitempty"`
	Content []OpenAIContent `json:"content,omitempty"`
}

type OpenAIContent struct {
	Type        string        `json:"type"`
	Text        string        `json:"text,omitempty"`
	Annotations []interface{} `json:"annotations,omitempty"`
	Logprobs    []interface{} `json:"logprobs,omitempty"`
}

func (r *OpenAIResponse) FirstText() (string, error) {
	for _, out := range r.Output {
		if len(out.Content) > 0 {
			for _, c := range out.Content {
				if c.Type == "output_text" && c.Text != "" {
					return c.Text, nil
				}
			}
		}
	}
	return "", fmt.Errorf("no output_text content found in OpenAI response")
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: %s (type=%s, param=%s, code=%s)", e.Message, e.Type, e.Param, e.Code)
}
