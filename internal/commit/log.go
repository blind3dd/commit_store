package commit

import (
	"fmt"
	"sync"
)

// Commit struct contains mutex and records
// to be commited as Logs in the storage
type Commit struct {
	mu      *sync.Mutex
	commits []Record
}

// Record struct defined Value to rw
// and offset for that value to be read
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"` // time conversion
}

// NewCommit instantiates a new Commit data struct
// that returns a new instance of the Commit
func (c *Commit) NewCommit() *Commit {

	return &Commit{
		//mu:      &sync.Mutex{},
		commits: make([]Record, 1),
	}
}

// NewCommit without a pointer receiver
func NewCommit() *Commit {
	logRecords := make([]Record, 1) // slice of Records (1 cap, len)

	return &Commit{commits: logRecords}
}

// Write adds new commit to the "parameter store"
// writing adjust offset - append new commit to slice of Records
func (c *Commit) Write(commit Record) (uint64, error) {
	defer c.mu.Unlock()
	c.mu.Lock()
	commit.Offset = uint64(len(c.commits))
	c.commits = append(c.commits, commit)

	return commit.Offset, nil
}

// Read reads value based on the offset (time based)
// and returns Record if offset exists
func (c *Commit) Read(offset uint64) (Record, error) {
	// use mutex fo reading to make sure program is thread safe when reading
	defer c.mu.Unlock()
	c.mu.Lock()
	if offset >= uint64(len(c.commits)) {
		return Record{}, fmt.Errorf("invalid offset defined")
	}

	return c.commits[offset], nil
}

//func timeToUint(time time.Duration) (uint64, error) {
//
//	ut := uint64(time)
//
//	return ut, nil
//}

//func main() {
//	mu := &sync.Mutex{}
//	//if _, err := timeToUint(time.Unix(0, 0)); err != nil {
//	//	panic(err)
//	//}
//	//log := Record{
//	//	Value:  make([]byte, 0),
//	//	Offset: tm,
//	//}
//	//TODO context and mutex for proper read write goroutines
//	ncl := NewCommit(mu)
//	fmt.Println(ncl)
//
//}
