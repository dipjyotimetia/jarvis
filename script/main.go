package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-billy/v5/memfs"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user information:", err)
		return
	}

	sshKeyPath := filepath.Join(usr.HomeDir, ".ssh", "id_ed25519")

	sshAuth, err := ssh.NewPublicKeysFromFile("git", sshKeyPath, "")
	if err != nil {
		fmt.Println("Error loading SSH key:", err)
		return
	}

	repo, err := gogit.Clone(memory.NewStorage(), memfs.New(), &gogit.CloneOptions{
		Auth:          sshAuth,
		URL:           "git@github.com:dipjyotimetia/jarvis.git",
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName("refs/heads/main"),
		SingleBranch:  true,
	})
	if err != nil {
		log.Fatal(err)
	}

	tagrefs, err := repo.Tags()
	if err != nil {
		fmt.Println("Error getting tags:", err)
		return
	}

	versions := make([]*semver.Version, 0)
	tagrefs.ForEach(func(t *plumbing.Reference) error {
		tagName := t.Name().Short()
		v := semver.MustParse(tagName)
		versions = append(versions, v)
		return nil
	})

	if len(versions) == 0 {
		fmt.Println("No valid SemVer tags found.")
		return
	}

	sort.Sort(semver.Collection(versions))
	latestTag := len(versions) - 1
	fmt.Println("Latest Tag:", versions[latestTag].String())
}
