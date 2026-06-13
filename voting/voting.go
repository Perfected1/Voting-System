package voting

// Candidate represents a person that can receive votes
type Candidate struct {
	ID    int
	Name  string
	Votes int
}

// GetCandidates returns the initial list of candidates
func GetCandidates() []Candidate {
	return []Candidate{
		{ID: 1, Name: "Alice", Votes: 0},
		{ID: 2, Name: "Bob", Votes: 0},
		{ID: 3, Name: "Charlie", Votes: 0},
	}
}

// Vote finds a candidate by ID and increments their vote count
func Vote(candidates []Candidate, id int) ([]Candidate, bool) {
	for i, c := range candidates {
		if c.ID == id {
			candidates[i].Votes++
			return candidates, true
		}
	}

	return candidates, false
}