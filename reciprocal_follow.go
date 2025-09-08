package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

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


func extractUsernamesFromHTML(r io.Reader) ([]string, error) {
	var usernames []string

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "https://www.instagram.com/") {
					parts := strings.Split(attr.Val, "/")
					if len(parts) > 3 {
						username := parts[3]
						if username != "" {
							usernames = append(usernames, username)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return usernames, nil
}


func main() {
	printGradientBanner()

	followingFile, err := os.Open("html/following.html")
	if err != nil {
		fmt.Println("Error opening following.html:", err)
		return
	}
	defer followingFile.Close()

	followersFile, err := os.Open("html/followers_1.html")
	if err != nil {
		fmt.Println("Error opening followers.html:", err)
		return
	}
	defer followersFile.Close()

	following, err := extractUsernamesFromHTML(followingFile)
	if err != nil {
		fmt.Println("Error parsing following.html:", err)
		return
	}
	followers, err := extractUsernamesFromHTML(followersFile)
	if err != nil {
		fmt.Println("Error parsing followers.html:", err)
		return
	}

	// Convert to sets
	followingSet := make(map[string]bool)
	for _, user := range following {
		followingSet[user] = true
	}

	followersSet := make(map[string]bool)
	for _, user := range followers {
		followersSet[user] = true
	}

	// Compare
	fmt.Println("✅ Reciprocal Followers:")
	for _, user := range following {
		if followersSet[user] {
			fmt.Println("  -", user)
		}
	}

	fmt.Println("\n❌ Not Following You Back:")
	for _, user := range following {
		if !followersSet[user] {
			fmt.Println("  -", user)
		}
	}
}

