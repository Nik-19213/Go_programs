// youtube-stats/youtube/youtube.go
package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Response models the JSON structure
// that we get back from the YouTube API
type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

// Items stores the ID + Statistics for
// a given channel
type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

// Stats stores the information we care about
// so how many views the channel has, how many subscribers
// how many videos, etc.
type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

// GetSubscribers fetches the YouTube channel statistics (views, subscribers, video count)
func GetSubscribers() (Items, error) {
	var response Response
	// We want to craft a new GET request that will include the query parameters we want
	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return Items{}, err
	}

	// Here we define the query parameters and their respective values
	q := req.URL.Query()
	// Fetch API key and channel ID from environment variables
	q.Add("key", "YOUR_API_HERE")
	q.Add("id", "YOUTUBE_CHANNEL_ID")
	q.Add("part", "statistics")
	req.URL.RawQuery = q.Encode()

	// Making the HTTP request to the YouTube API
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request to YouTube API:", err)
		return Items{}, err
	}
	defer resp.Body.Close()

	// Log the response status for debugging
	fmt.Println("Response Status:", resp.Status)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Items{}, err
	}

	// Unmarshal the JSON response into the Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return Items{}, err
	}

	// Check if Items array is empty (i.e., no data available)
	if len(response.Items) == 0 {
		return Items{}, fmt.Errorf("no channel data found")
	}

	// Return the first item in the response, which contains the statistics
	return response.Items[0], nil
}
