package main

import "fmt"

// Candidate represents the person being voted for

type Candidate struct {
	ID 		int
	Name 	string
	Votes 	int
}

//getCandidates returns initial list of available candidates

func getCandidates() []Candidate {
	return []Candidate{
		{ID: 1, Name: "Alice", Votes: 0},
		{ID: 2, Name: "Bob", Votes: 0},
		{ID: 3, Name: "Charlie", Votes: 0},
	}
}

func main() {
	fmt.Println("🗳️ Voting System Initialized\n")

	candidates := getCandidates()

	fmt.Println("Available Candidates:")
	for _, c := range candidates {
		fmt.Printf("ID: %d | Name: %s | Votes: %d\n", c.ID, c.Name, c.Votes)
	}
}