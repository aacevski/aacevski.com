package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ContributionDay struct {
	Date              string
	ContributionCount int
	Level             int // 0-4 for intensity
}

type ContributionWeek struct {
	Days []ContributionDay
}

type ContributionData struct {
	Weeks              []ContributionWeek
	TotalContributions int
}

// GraphQL query to fetch contribution data
const contributionsQuery = `
query($userName:String!) {
  user(login: $userName) {
    contributionsCollection {
      contributionCalendar {
        totalContributions
        weeks {
          contributionDays {
            date
            contributionCount
            contributionLevel
          }
        }
      }
    }
  }
}
`

// GetContributions fetches GitHub contribution data for a user
func GetContributions(username, token string) (*ContributionData, error) {
	query := map[string]interface{}{
		"query": contributionsQuery,
		"variables": map[string]string{
			"userName": username,
		},
	}

	jsonData, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response
	var result struct {
		Data struct {
			User struct {
				ContributionsCollection struct {
					ContributionCalendar struct {
						TotalContributions int `json:"totalContributions"`
						Weeks              []struct {
							ContributionDays []struct {
								Date              string `json:"date"`
								ContributionCount int    `json:"contributionCount"`
								ContributionLevel string `json:"contributionLevel"`
							} `json:"contributionDays"`
						} `json:"weeks"`
					} `json:"contributionCalendar"`
				} `json:"contributionsCollection"`
			} `json:"user"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// Convert to our data structure
	data := &ContributionData{
		TotalContributions: result.Data.User.ContributionsCollection.ContributionCalendar.TotalContributions,
		Weeks:              make([]ContributionWeek, 0),
	}

	for _, week := range result.Data.User.ContributionsCollection.ContributionCalendar.Weeks {
		cWeek := ContributionWeek{
			Days: make([]ContributionDay, 0),
		}
		for _, day := range week.ContributionDays {
			level := contributionLevelToInt(day.ContributionLevel)
			cWeek.Days = append(cWeek.Days, ContributionDay{
				Date:              day.Date,
				ContributionCount: day.ContributionCount,
				Level:             level,
			})
		}
		data.Weeks = append(data.Weeks, cWeek)
	}

	return data, nil
}

func contributionLevelToInt(level string) int {
	switch level {
	case "NONE":
		return 0
	case "FIRST_QUARTILE":
		return 1
	case "SECOND_QUARTILE":
		return 2
	case "THIRD_QUARTILE":
		return 3
	case "FOURTH_QUARTILE":
		return 4
	default:
		return 0
	}
}
