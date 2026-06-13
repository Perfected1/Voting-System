package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"voting-system/voting"
)

// displayCandidates prints all candidates and their vote counts
func displayCandidates(candidates []voting.Candidate) {
	fmt.Println("\nAvailable Candidates:")
	for _, c := range candidates {
		fmt.Printf("ID: %d | Name: %s | Votes: %d\n", c.ID, c.Name, c.Votes)
	}
}

func main() {
	candidates := voting.GetCandidates()

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🗳️ Voting System")

	for {
		displayCandidates(candidates)

		fmt.Print("\nEnter Candidate ID to vote (or 0 to exit): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if id == 0 {
			fmt.Println("Exiting voting system...")
			break
		}

		updated, ok := voting.Vote(candidates, id)
		if !ok {
			fmt.Println("Candidate not found. Try again.")
			continue
		}

		candidates = updated

		fmt.Println("Vote recorded successfully!")
	}
}