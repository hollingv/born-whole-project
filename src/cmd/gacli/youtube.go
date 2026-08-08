package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"
)

// ytChannelResp is the response from the YouTube channels API.
type ytChannelResp struct {
	Items []struct {
		ID string `json:"id"`
	} `json:"items"`
}

// ytSearchResp is the response from the YouTube search API.
type ytSearchResp struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"snippet"`
	} `json:"items"`
	NextPageToken string `json:"nextPageToken"`
}

// buildYouTubeResources fetches Shorts from YouTube and saves them to resources.json.
func buildYouTubeResources() {
	// Always write an empty resources.json so the file exists after kb-build.
	if err := os.WriteFile(resourcesPath, []byte("[]\n"), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error initialising %s: %v\n", resourcesPath, err)
	}

	envKey := projectPrefix + "_YT_API_KEY"
	apiKey := os.Getenv(envKey)
	if apiKey == "" {
		fmt.Printf("[ WARN ] %s not set — skipping YouTube Shorts fetch\n", envKey)
		return
	}
	fmt.Println("Fetching YouTube Shorts...")
	resources, err := fetchYouTubeShorts(apiKey, youtubeChannel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching YouTube Shorts: %v\n", err)
		return
	}
	data, _ := json.MarshalIndent(resources, "", "  ")
	if err := os.WriteFile(resourcesPath, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", resourcesPath, err)
		return
	}
	fmt.Printf("Written %s (%d shorts)\n", resourcesPath, len(resources))
}

// fetchYouTubeShorts retrieves Shorts from a YouTube channel published in the past 3 months.
func fetchYouTubeShorts(apiKey, handle string) ([]Resource, error) {
	channelID, err := resolveChannelID(apiKey, handle)
	if err != nil {
		return nil, err
	}
	return searchShorts(apiKey, channelID)
}

// resolveChannelID returns the YouTube channel ID for a given handle.
func resolveChannelID(apiKey, handle string) (string, error) {
	chURL := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/channels?forHandle=%s&part=id&key=%s",
		url.QueryEscape(handle), apiKey,
	)
	resp, err := httpClient.Get(chURL)
	if err != nil {
		return "", fmt.Errorf("channel lookup: %w", err)
	}
	defer resp.Body.Close()

	var chResp ytChannelResp
	if err := json.NewDecoder(resp.Body).Decode(&chResp); err != nil {
		return "", fmt.Errorf("channel decode: %w", err)
	}
	if len(chResp.Items) == 0 {
		return "", fmt.Errorf("channel not found for handle: %s", handle)
	}
	channelID := chResp.Items[0].ID
	fmt.Printf("  Channel ID: %s\n", channelID)
	return channelID, nil
}

// searchShorts pages through the YouTube search API and returns all matching Shorts.
func searchShorts(apiKey, channelID string) ([]Resource, error) {
	publishedAfter := time.Now().AddDate(0, -3, 0).UTC().Format(time.RFC3339)
	var resources []Resource
	pageToken := ""
	for {
		searchURL := fmt.Sprintf(
			"https://www.googleapis.com/youtube/v3/search?channelId=%s&type=video&videoDuration=short&part=snippet&maxResults=50&publishedAfter=%s&order=date&key=%s",
			channelID, url.QueryEscape(publishedAfter), apiKey,
		)
		if pageToken != "" {
			searchURL += "&pageToken=" + pageToken
		}

		resp, err := httpClient.Get(searchURL)
		if err != nil {
			return nil, fmt.Errorf("search: %w", err)
		}
		defer resp.Body.Close()

		var searchResp ytSearchResp
		if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
			return nil, fmt.Errorf("search decode: %w", err)
		}

		for _, item := range searchResp.Items {
			resources = append(resources, Resource{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				VideoID:     item.ID.VideoID,
			})
		}

		if searchResp.NextPageToken == "" {
			break
		}
		pageToken = searchResp.NextPageToken
	}
	return resources, nil
}
