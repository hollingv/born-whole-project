package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const aiModel = "@cf/meta/llama-3.1-8b-instruct-fast"

const systemPrompt = `You are a helpful assistant answering questions about circumcision,
bodily autonomy, and children's rights. Only use the text provided below as context.
Do not use any outside knowledge. Be concise, factual, and compassionate.
If the provided context does not contain enough information to answer, say so explicitly.`

type cfAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type cfAIRequest struct {
	Messages  []cfAIMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

type cfAIResponse struct {
	Result struct {
		Response string `json:"response"`
	} `json:"result"`
	Success bool `json:"success"`
}

// callCloudflareAI sends a question and context chunks to the Cloudflare AI API.
func callCloudflareAI(question string, chunks []string) (string, error) {
	accountID := os.Getenv(projectPrefix + "_CF_ACCOUNT_ID")
	apiToken := os.Getenv(projectPrefix + "_CF_API_TOKEN")
	if accountID == "" || apiToken == "" {
		return "", fmt.Errorf("%s_CF_ACCOUNT_ID and %s_CF_API_TOKEN env vars not set", projectPrefix, projectPrefix)
	}

	context := strings.Join(chunks, "\n\n")
	req := cfAIRequest{
		MaxTokens: 1024,
		Messages: []cfAIMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf("Context:\n%s\n\nQuestion: %s", context, question)},
		},
	}
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", accountID, aiModel)
	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var aiResp cfAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return "", err
	}
	if !aiResp.Success {
		return "", fmt.Errorf("Cloudflare AI returned success=false")
	}
	return aiResp.Result.Response, nil
}
