package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type StringListData struct {
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

type Relationship struct {
	StringListData []StringListData `json:"string_list_data"`
}

type FollowingData struct {
	Following []Relationship `json:"relationships_following"`
}

type FollowersData struct {
	Followers []Relationship `json:"relationships_followers"`
}

const Reset = "\033[0m"

var gradient = []string{
	"\033[38;2;131;58;180m",  // Purple
	"\033[38;2;193;53;132m",  // Magenta
	"\033[38;2;225;48;108m",  // Pink
	"\033[38;2;253;29;29m",   // Red
	"\033[38;2;245;96;64m",   // Orange
	"\033[38;2;247;119;55m",  // Deep Orange
	"\033[38;2;252;175;69m",  // Yellow-Orange
	"\033[38;2;255;220;128m", // Light Yellow
}

func printGradientBanner() {
	banner := `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠀
⠀⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⠀
⠀⣿⣿⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⠀
⠀⣿⣿⣿⣿⡟⢀⣾⣿⣷⡖⠒⣀⣀⣤⣶⣾⣿⣷⣶⣤⣤⣴⣾⣆⠘⣿⣿⣿⠀
⠀⣿⣿⣿⡿⠁⣾⣿⣿⣿⣧⣀⠙⠛⠛⠛⠋⣈⠻⢿⣿⣿⣿⣿⣿⣧⠈⢿⣿⠀
⠀⣿⣿⣿⠁⣼⣿⠟⠻⠿⣿⣿⣿⣷⣶⣾⣿⠿⣷⣄⡉⠻⣿⣿⣿⣿⣧⠈⢿⠀
⠀⠛⠛⠃⠀⠻⠋⣠⣶⠄⠙⠛⢻⣿⣿⣿⣷⣦⣈⠛⢿⣦⣄⠙⠻⠛⠁⠀⠀⠀
⠀⠀⠀⠀⠀⢠⣾⠟⠁⣠⣿⠇⠈⠛⣿⣿⣈⠙⠿⣷⣦⡈⠛⠛⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠁⢠⣾⡿⠁⣠⣾⠆⠸⠉⠻⣷⣤⡈⠙⠿⠃⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⣾⠟⢁⣴⡶⠀⣤⡀⠙⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠰⣿⠟⠁⠀⠛⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
`

	lines := strings.Split(banner, "\n")

	for i, line := range lines {
		color := gradient[i%len(gradient)] // Cycle through gradient colors
		fmt.Println(color + line + Reset)
	}
}

func extractUsernames(relationships []Relationship) map[string]bool {
	users := make(map[string]bool)
	for _, rel := range relationships {
		if len(rel.StringListData) > 0 {
			users[rel.StringListData[0].Value] = true
		}
	}
	return users
}

func main() {
	printGradientBanner()

	// Load following.json
	followingFile, err := os.ReadFile("json/following.json")
	if err != nil {
		fmt.Println("Error reading following.json:", err)
		return
	}
	var followingData FollowingData
	json.Unmarshal(followingFile, &followingData)

	// Load followers.json
	followersFile, err := os.ReadFile("json/followers_1.json")
	if err != nil {
		fmt.Println("Error reading followers.json:", err)
		return
	}
	var followersData FollowersData
	json.Unmarshal(followersFile, &followersData)

	// Extract usernames
	followingUsers := extractUsernames(followingData.Following)
	followersUsers := extractUsernames(followersData.Followers)

	// Find reciprocal and non-followers
	fmt.Println("✅ Reciprocal Followers:")
	for user := range followingUsers {
		if followersUsers[user] {
			fmt.Println("  -", user)
		}
	}

	fmt.Println("\n❌ Not Following You Back:")
	for user := range followingUsers {
		if !followersUsers[user] {
			fmt.Println("  -", user)
		}
	}
}

