package home

import (
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDir(t *testing.T) {
	require.NotEmpty(t, Dir())
}

func TestShort(t *testing.T) {
	d := filepath.Join(Dir(), "documents", "file.txt")
	require.Equal(t, filepath.FromSlash("~/documents/file.txt"), Short(d))
	ad := filepath.FromSlash("/absolute/path/file.txt")
	require.Equal(t, ad, Short(ad))
}

func TestLong(t *testing.T) {
	d := filepath.FromSlash("~/documents/file.txt")
	require.Equal(t, filepath.Join(Dir(), "documents", "file.txt"), Long(d))
	require.Equal(t, Dir(), Long("~"))
	ad := filepath.FromSlash("/absolute/path/file.txt")
	require.Equal(t, ad, Long(ad))
}

func TestLongWithEnv(t *testing.T) {
	t.Setenv("CRUSH_TEST_HOME", filepath.FromSlash("/tmp/crush"))
	require.Equal(t, filepath.FromSlash("/tmp/crush/file.txt"), Long("$CRUSH_TEST_HOME/file.txt"))
}

func TestLongWithUser(t *testing.T) {
	u, err := user.Current()
	require.NoError(t, err)

	expected := filepath.Join(u.HomeDir, "docs", "file.txt")
	require.Equal(t, expected, Long("~"+u.Username+"/docs/file.txt"))
}
