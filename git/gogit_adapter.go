package git

import (
	"strings"

	"github.com/meinto/glow"
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// GoGitAdapter implemented with go-git
type goGitAdapter struct{}

// CurrentBranch returns the current branch name
func (a goGitAdapter) CurrentBranch() (glow.Branch, error) {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error opening repository")
	}

	headRef, err := r.Head()
	if err != nil {
		return nil, errors.Wrap(err, "error getting current branch")
	}

	refName := string(headRef.Name())
	return glow.NewBranch(refName)
}

// BranchList returns a list of avalilable branches
func (s goGitAdapter) BranchList() ([]glow.Branch, error) {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error opening repository")
	}

	refList, err := r.References()
	if err != nil {
		return nil, errors.Wrap(err, "error getting ref list")
	}

	branches := make([]glow.Branch, 0)
	refPrefix := "refs/heads/"
	refList.ForEach(func(ref *plumbing.Reference) error {
		refName := ref.Name().String()
		if strings.HasPrefix(refName, refPrefix) {
			b, _ := glow.NewBranch(refName)
			branches = append(branches, b)
		}
		return nil
	})

	return nil, errors.New("")
}

// Fetch changes
func (a goGitAdapter) Fetch() error {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return errors.Wrap(err, "error while fetching")
	}
	return r.Fetch(&git.FetchOptions{})
}

// Create a new branch
func (a goGitAdapter) Create(b glow.Branch) error {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return errors.Wrap(err, "error opening repository")
	}

	headRef, err := r.Head()
	if err != nil {
		return errors.Wrap(err, "error getting current branch")
	}

	refName := string(headRef.Name())
	if !b.CreationIsAllowedFrom(refName) {
		return errors.New("You are not on the develop branch.\nPlease switch branch...\n")
	}

	ref := plumbing.NewHashReference(plumbing.ReferenceName(b.BranchName()), headRef.Hash())

	err = r.Storer.SetReference(ref)
	return errors.Wrap(err, "error while creating branch")
}

// Checkout a branch
func (a goGitAdapter) Checkout(b glow.Branch) error {
	r, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		return errors.Wrap(err, "error opening repository")
	}

	w, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, "error getting worktree")
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(b.BranchName()),
	})
	return errors.Wrapf(err, "error while checkout branch %s", b.BranchName())
}
