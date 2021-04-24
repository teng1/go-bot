package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	log "github.com/sirupsen/logrus"
)

func InMemClone() {
	// InMemoryExample the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	log.Info("git clone https://github.com/go-git/go-billy")

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/go-git/go-billy",
	})

	log.Error(err)

	// Gets the HEAD history from HEAD, just like this command:
	log.Info("git log")

	// ... retrieves the branch pointed by HEAD
	log.Info(r.Head())
	// ref, err := r.Head()
	// ... retrieves the commit history
	// cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

	// ... just iterates over the commits, printing it
	// err = cIter.ForEach(func(c *object.Commit) error {
	// 	fmt.Println(c)
	// 	return nil
	// })

}
