package integration

import (
	"testing"

	"github.com/bitrise-core/bitrise-plugins-init/version"
	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

var versionStr = version.VERSION

func Test_VersionTest(t *testing.T) {
	t.Log("version flag")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("")
		require.NoError(t, err)

		cmd := command.New(binPath(), "--version")
		cmd.SetDir(tmpDir)
		out, err := cmd.RunAndReturnTrimmedCombinedOutput()
		require.NoError(t, err, out)
		require.Equal(t, versionStr, out)
	}

	t.Log("version flag")
	{
		tmpDir, err := pathutil.NormalizedOSTempDirPath("")
		require.NoError(t, err)

		cmd := command.New(binPath(), "-v")
		cmd.SetDir(tmpDir)
		out, err := cmd.RunAndReturnTrimmedCombinedOutput()
		require.NoError(t, err, out)
		require.Equal(t, versionStr, out)
	}
}