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

This is your home page. Feel free to edit it however you please.

You can see all the [current pages here](/page/).`
)

// ErrPageNotExists is returned when a page does not exist.
var ErrPageNotExists = errors.New("page does not exist")

// Wiki encapsulates a Git-backed wiki.
type Wiki struct {
	path string
	fs   billy.Filesystem
	r    *git.Repository
}

// CreateWiki creates a new wiki on-disk at the specified path if one does not
// already exists and returns a new Wiki object for it.
func CreateWiki(path string) (*Wiki, error) {
	w := Wiki{path: path}

	// Create repository & git directories
	err := os.MkdirAll(path+"/.git", 0755)
	switch {
	case errors.Is(err, fs.ErrExist):
		// already exists, do nothing
	case err != nil:
		return nil, err
	}

	// Initialize or open Git repository to return with Wiki
	gitfs := osfs.New(path + "/.git")
	s := filesystem.NewStorage(gitfs, cache.NewObjectLRUDefault())

	w.fs = osfs.New(path)
	w.r, err = git.Init(s, w.fs) // Init if it doesn't exist

	switch {
	case errors.Is(err, git.ErrRepositoryAlreadyExists):
		// Open the repository if it does exists
		w.r, err = git.Open(s, w.fs)
		if err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	}

	// Create the home page if it doesn't yet exist
	homeFileName := homePageName + ".md"
	_, err = w.fs.Stat(homeFileName)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		// Create the home page
		err := w.SetPage(homePageName, []byte(homePageInitialContent))
		if err != nil {
			return nil, err
		}
	case err != nil:
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
