package commit

import (
	"fmt"
	"github.com/blind3dd/commit_store/internal/entities"
	"sync"
)

type Commit struct {
	mu      sync.Mutex
	Commits []entities.Commit
}

// NewCommit without a pointer receiver
func NewCommit() *Commit {
	return &Commit{}
}

// Write adds new commit to the "parameter store"
// writing adjust offset - append new commit to slice of Records
func (c *Commit) Write(commit entities.Commit) (uint64, error) {
	defer c.mu.Unlock()
	c.mu.Lock()
	commit.Offset = uint64(len(c.Commits))
	c.Commits = append(c.Commits, commit)

	return commit.Offset, nil
}

// Read reads value based on the offset (time based)
// and returns Record if offset exists
func (c *Commit) Read(offset uint64) (entities.Commit, error) {
	defer c.mu.Unlock()
	c.mu.Lock()
	if offset >= uint64(len(c.Commits)) {
		return entities.Commit{}, OffsetNotFoundError
	}

	return c.Commits[offset], nil
}

var OffsetNotFoundError = fmt.Errorf("invalid offset defined")
