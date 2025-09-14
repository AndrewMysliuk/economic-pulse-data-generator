package schema

type SearchAndSummarizeResponse struct {
	ID     string                     `json:"id"`
	Object string                     `json:"object"`
	Output []SearchAndSummarizeOutput `json:"output"`
}

type SearchAndSummarizeOutput struct {
	Content []SearchAndSummarizeContent `json:"content"`
}

type SearchAndSummarizeContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}
