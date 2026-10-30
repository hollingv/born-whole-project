package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const transcribeAPI = "https://transcribeyoutube.com/api/transcript"

// finalTranscriptErrors are error codes from the TranscribeYouTube API where
// retrying will never help — the video genuinely has no captions or does not exist.
var finalTranscriptErrors = map[string]bool{
	"NO_CAPTIONS":         true,
	"INVALID_VIDEO_ID":    true,
	"INVALID_URL":         true,
	"URL_TOO_LONG":        true,
	"INVALID_LANG":        true,
	"INVALID_LANG_FORMAT": true,
	"PAYLOAD_TOO_LARGE":   true,
}

type transcribeRequest struct {
	URL  string `json:"url"`
	Lang string `json:"lang"`
}

type transcribeResponse struct {
	OK         bool   `json:"ok"`
	VideoID    string `json:"videoId"`
	Title      string `json:"title"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	RetryAfter int    `json:"retryAfter"`
	Transcript []struct {
		Text  string  `json:"text"`
		Start float64 `json:"start"`
		Dur   float64 `json:"dur"`
	} `json:"transcript"`
}

var transcribeHTTPClient = &http.Client{Timeout: 30 * time.Second}

// fetchTranscript calls the TranscribeYouTube API for the given video ID and
// returns the transcript as plain text. Rate limiting (429) is handled
// automatically by sleeping the retryAfter duration and retrying. Retriable
// error codes (EXTRACTION_FAILED, AUTOMATED_ACCESS, etc.) are retried once
// after 60 seconds. Final error codes are returned as errors without retry.
func fetchTranscript(videoID string) (string, error) {
	url := "https://www.youtube.com/shorts/" + videoID
	return doTranscribeRequest(url, false)
}

// doTranscribeRequest is the internal implementation that supports one retry.
func doTranscribeRequest(videoURL string, isRetry bool) (string, error) {
	body, _ := json.Marshal(transcribeRequest{URL: videoURL, Lang: "en"})

	resp, err := transcribeHTTPClient.Post(transcribeAPI, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("POST: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	// 429 — honour the retryAfter value and retry indefinitely.
	if resp.StatusCode == 429 {
		var r transcribeResponse
		_ = json.Unmarshal(raw, &r)
		wait := r.RetryAfter
		if wait <= 0 {
			wait = 60
		}
		fmt.Printf("  [ RATE LIMITED ] sleeping %ds before retry...\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
		return doTranscribeRequest(videoURL, isRetry)
	}

	var r transcribeResponse
	if err := json.Unmarshal(raw, &r); err != nil {
		return "", fmt.Errorf("parse response (HTTP %d): %w", resp.StatusCode, err)
	}

	if r.OK {
		var parts []string
		for _, seg := range r.Transcript {
			if t := strings.TrimSpace(seg.Text); t != "" {
				parts = append(parts, t)
			}
		}
		return strings.Join(parts, " "), nil
	}

	// Final error — retrying will not help.
	if finalTranscriptErrors[r.Error] {
		return "", fmt.Errorf("%s: %s", r.Error, r.Message)
	}

	// Retriable error — one attempt after 60 seconds.
	if !isRetry {
		fmt.Printf("  [ RETRY ] %s — waiting 60s...\n", r.Error)
		time.Sleep(60 * time.Second)
		return doTranscribeRequest(videoURL, true)
	}

	return "", fmt.Errorf("%s: %s", r.Error, r.Message)
}
