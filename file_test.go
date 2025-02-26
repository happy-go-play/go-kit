package gokit

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func CreateTempFile(t *testing.T, dir, pattern string) (*os.File, func()) {
	file, err := os.CreateTemp(dir, pattern)
	require.NoError(t, err)
	return file, func() {
		require.NoError(t, os.Remove(file.Name()))
	}
}

func CreateTempDir(t *testing.T) (string, func()) {
	dir, err := os.MkdirTemp("", "testdir")
	require.NoError(t, err)
	return dir, func() {
		require.NoError(t, os.RemoveAll(dir))
	}
}

func TestCreateDirIfNotExist(t *testing.T) {
	dir, cleanup := CreateTempDir(t)
	defer cleanup()

	err := CreateDirIfNotExist(dir)
	require.NoError(t, err)

	nonExistentDir := dir + "/subdir"
	err = CreateDirIfNotExist(nonExistentDir)
	require.NoError(t, err)
	require.DirExists(t, nonExistentDir)
}

func TestCreateDirIfNotExist_Error(t *testing.T) {
	err := CreateDirIfNotExist("/invalid/dir")
	require.Error(t, err)
}

func TestAppendToFile(t *testing.T) {
	file, cleanup := CreateTempFile(t, "", "testfile")
	defer cleanup()

	data := []byte("test data")
	err := AppendToFile(file.Name(), data)
	require.NoError(t, err)

	content, err := os.ReadFile(file.Name())
	require.NoError(t, err)
	require.Equal(t, append(data, '\n'), content)
}

func TestAppendToFile_Error(t *testing.T) {
	err := AppendToFile("/invalid/file", []byte("data"))
	require.Error(t, err)
}
