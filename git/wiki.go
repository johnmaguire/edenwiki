package git

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

const (
	homePageName           = "Home"
	homePageInitialContent = `# Welcome to EdenWiki!

This is your home page. Feel free to edit it however you please.`
)

// ErrPageNotExists is returned when a page does not exist.
var ErrPageNotExists = errors.New("page does not exist")

// Wiki encapsulates a Git-backed wiki.
type Wiki struct {
	path string
	fs   billy.Filesystem
	r    *git.Repository
}

// CreateWiki creates a new wiki on-disk at the specified path and returns a new Wiki object for it.
func CreateWiki(path string) (*Wiki, error) {
	w := Wiki{path: path}

	// Create repository directory
	err := os.MkdirAll(path+"/.git", 0755)
	if err != nil {
		return nil, err
	}

	gitfs := osfs.New(path + "/.git")
	s := filesystem.NewStorage(gitfs, cache.NewObjectLRUDefault())

	w.fs = osfs.New(path)
	w.r, err = git.Init(s, w.fs)
	if err != nil {
		return nil, err
	}

	homeFileName := homePageName + ".md"
	f, err := w.fs.Create(homeFileName)
	if err != nil {
		return nil, err
	}

	_, err = f.Write([]byte(homePageInitialContent))
	if err != nil {
		return nil, err
	}

	wt, err := w.r.Worktree()
	if err != nil {
		return nil, err
	}

	_, err = wt.Add(homeFileName)
	if err != nil {
		return nil, err
	}

	author := object.Signature{
		Name: "EdenWiki",
	}
	_, err = wt.Commit("Initial commit", &git.CommitOptions{
		All:       true,
		Author:    &author,
		Committer: &author,
	})
	if err != nil {
		return nil, err
	}

	return &w, nil
}

// SetPage creates or updates a page with the given content, and commits the change.
func (w *Wiki) SetPage(name string, contents []byte) error {
	fileName := name + ".md"

	// Update or create the file (and set the appropriate commit message)
	commitMessage := fmt.Sprintf("Update %s", fileName)
	f, err := w.fs.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0666)
	if os.IsNotExist(err) {
		commitMessage = fmt.Sprintf("Create %s", fileName)
		f, err = w.fs.Create(fileName)
		if err != nil {
			return err
		}
	}

	_, err = f.Write(contents)
	if err != nil {
		return err
	}

	// Commit the change
	wt, err := w.r.Worktree()
	if err != nil {
		return err
	}

	_, err = wt.Add(fileName)
	if err != nil {
		return err
	}

	author := object.Signature{
		Name: "EdenWiki",
		When: time.Now(),
	}
	_, err = wt.Commit(commitMessage, &git.CommitOptions{
		All:       true,
		Author:    &author,
		Committer: &author,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetPage looks up a page in the Git repository and returns its contents.
func (w *Wiki) GetPage(name string) ([]byte, error) {
	fileName := name + ".md"

	f, err := w.fs.Open(fileName)
	switch {
	case os.IsNotExist(err):
		return nil, ErrPageNotExists
	case err != nil:
		return nil, err
	}

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (w *Wiki) ListPages() ([]string, error) {
	var pageNames []string

	err := filepath.WalkDir(w.path, func(path string, d fs.DirEntry, err error) error {
		if strings.HasSuffix(path, ".md") {
			pageNames = append(pageNames, path[len(w.path)-1:len(path)-3])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return pageNames, nil
}
