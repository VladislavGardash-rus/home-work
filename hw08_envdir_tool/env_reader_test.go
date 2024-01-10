//nolint:all
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Success case: file with empty first string", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "new_dir")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		err = ioutil.WriteFile(filepath.Join(dir, "file_name"), []byte("\n"), 0o666)
		require.NoError(t, err)

		expectEnv := Environment{
			"file_name": EnvValue{Value: "", NeedRemove: false},
		}

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("Success case: file name with low case", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "new_dir")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		err = ioutil.WriteFile(filepath.Join(dir, "file_name"), []byte("file_name"), 0o666)
		require.NoError(t, err)

		expectEnv := Environment{
			"file_name": EnvValue{Value: "file_name", NeedRemove: false},
		}

		env, err := ReadDir(dir)
		require.NoError(t, err)
		require.Equal(t, env, expectEnv)
	})

	t.Run("Success case: testdata dir", func(t *testing.T) {
		environments := Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"HELLO": EnvValue{Value: `"hello"`, NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}

		env, err := ReadDir("testdata/env")
		require.NoError(t, err)
		require.Equal(t, env, environments)
	})
}
