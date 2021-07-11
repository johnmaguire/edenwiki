package git

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWiki(t *testing.T) {
	name, err := ioutil.TempDir("", "wiki-test-")
	require.NoError(t, err)
	defer os.RemoveAll(name)

	w, err := CreateWiki(name)
	require.NoError(t, err)

	newPageContents := []byte("This is a new page's contents.")
	err = w.SetPage("A New Page", newPageContents)
	require.NoError(t, err)

	buf, err := w.GetPage("A New Page")
	require.NoError(t, err)
	assert.Equal(t, newPageContents, buf)

	err = w.SetPage("An Updated Page", newPageContents)
	assert.NoError(t, err)

	buf, err = w.GetPage("An Updated Page")
	require.NoError(t, err)
	assert.Equal(t, newPageContents, buf)

	updatedPageContents := []byte("This is an updated page's contents.")
	err = w.SetPage("An Updated Page", updatedPageContents)
	assert.NoError(t, err)

	buf, err = w.GetPage("An Updated Page")
	require.NoError(t, err)
	assert.Equal(t, updatedPageContents, buf)
}
