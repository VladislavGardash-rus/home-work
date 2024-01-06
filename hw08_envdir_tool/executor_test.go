package main

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("Success execute command", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "test")
		require.NoError(t, err)
		defer os.RemoveAll(dir)

		err = os.Mkdir(filepath.Join(dir, "envs"), 0777)
		require.NoError(t, err)

		err = ioutil.WriteFile(filepath.Join(dir, "envs", "BAR"), []byte("bar"), 0666)
		require.NoError(t, err)

		err = ioutil.WriteFile(filepath.Join(dir, "time_test.sh"), []byte("#!/usr/bin/env bash\n\necho -e \"BAR is (${BAR})\narguments are $*\""), 0666)
		require.NoError(t, err)

		err = os.Chmod(filepath.Join(dir, "time_test.sh"), 0777)
		require.NoError(t, err)

		env, err := ReadDir(filepath.Join(dir, "envs"))
		require.NoError(t, err)

		returnCode := RunCmd([]string{filepath.Join(dir, "time_test.sh"), "any"}, env)
		require.Equal(t, 0, returnCode)
	})

	t.Run("Fail execute command", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)

		returnCode := RunCmd([]string{"cat", "."}, env)
		require.Equal(t, 111, returnCode)
	})
}
