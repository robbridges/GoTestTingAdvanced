package sub_test

import (
	"GoTestingAdvanced/sub"
	"os/exec"
	"strings"
	"testing"
)

var (
	testHasGit bool
)

func init() {
	if _, err := exec.LookPath("git"); err == nil {
		testHasGit = true
	}
}

func TestGitStatus(t *testing.T) {
	if !testHasGit {
		t.Skip("Git not found on local pc")
	}
	want := "On branch master"
	got, err := sub.GitStatus()
	if err != nil {
		t.Fatalf("Gitstatus() err = %s; want nil", err)
	}
	if !strings.Contains(got, want) {
		t.Errorf("GitStatus() = %s; want prefix %s", got, want)
	}
}
