package examples

import (
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)


func TestAccRepoTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "repo", "ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccRepoPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "repo", "python"),
		})

	integration.ProgramTest(t, &test)
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	checkBitbucketUsername(t)
	checkBitbucketPassword(t)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
	}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/bitbucket",
		},
	})

	return baseJS
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePy
}

func checkBitbucketUsername(t *testing.T) {
	username := os.Getenv("BITBUCKET_USERNAME")
	if username == "" {
		t.Skipf("Skipping test due to missing BITBUCKET_USERNAME environment variable")
	}
}

func checkBitbucketPassword(t *testing.T) {
	username := os.Getenv("BITBUCKET_PASSWORD")
	if username == "" {
		t.Skipf("Skipping test due to missing BITBUCKET_PASSWORD environment variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}
