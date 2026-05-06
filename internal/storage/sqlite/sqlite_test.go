package sqlite

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"url-shortener/internal/storage"
)

func TestSaveAndGetURL(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test*.db")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	s, err := New(tmpFile.Name())
	require.NoError(t, err)

	id, err := s.SaveURL("https://google.com", "google")
	assert.NoError(t, err)
	assert.Greater(t, id, int64(0))

	url, err := s.GetURL("google")
	assert.NoError(t, err)
	assert.Equal(t, "https://google.com", url)

	_, err = s.GetURL("notexist")
	assert.ErrorIs(t, err, storage.ErrURLNotFound)
}

func TestDeleteURL(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test*.db")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	s, err := New(tmpFile.Name())
	require.NoError(t, err)

	_, err = s.SaveURL("https://example.com", "example")
	require.NoError(t, err)

	err = s.DeleteURL("example")
	assert.NoError(t, err)

	_, err = s.GetURL("example")
	assert.ErrorIs(t, err, storage.ErrURLNotFound)
}
