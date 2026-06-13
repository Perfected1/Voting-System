package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Candidate represents a person that can receive votes
type Candidate struct {
	ID    int
	Name  string
	Votes int
}

// getCandidates returns the initial list of candidates
func getCandidates() []Candidate {
	return []Candidate{
		{ID: 1, Name: "Alice", Votes: 0},
		{ID: 2, Name: "Bob", Votes: 0},
		{ID: 3, Name: "Charlie", Votes: 0},
	}
}

// displayCandidates prints all candidates and their vote counts
func displayCandidates(candidates []Candidate) {
	fmt.Println("\nAvailable Candidates:")
	for _, c := range candidates {
		fmt.Printf("ID: %d | Name: %s | Votes: %d\n", c.ID, c.Name, c.Votes)
	}
}

// vote finds a candidate by ID and increments their vote count
func vote(candidates []Candidate, id int) ([]Candidate, bool) {
	for i, c := range candidates {
		if c.ID == id {
			candidates[i].Votes++
			return candidates, true
		}
	}

	// Return false if no matching candidate is found
	return candidates, false
}

func main() {
	// Load candidates into memory
	candidates := getCandidates()

	// Create reader for user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Voting System")

	for {
		// Show available candidates
		displayCandidates(candidates)

		// Ask user for candidate ID
		fmt.Print("\nEnter Candidate ID to vote (or 0 to exit): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Convert input string to integer
		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Exit application
		if id == 0 {
			fmt.Println("Exiting voting system...")
			break
		}

		// Record vote
		updated, ok := vote(candidates, id)
		if !ok {
			fmt.Println("Candidate not found. Try again.")
			continue
		}

		// Update candidate list with new vote count
		candidates = updated

		fmt.Println("Vote recorded successfully!")
	}
}