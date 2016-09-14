package main

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/drone/drone-plugin-go/plugin"
)

func TestCommandBuildCorrectly(t *testing.T) {
	vargs := AzureBlobxfer{}
	vargs.StorageAccountKey = "xyzabc"
	vargs.Container = "my-container"
	vargs.StorageAccountName = "my-storage-account"
	vargs.Source = "__source__"
	w := plugin.Workspace{Path: "/test/path"}
	if !reflect.DeepEqual(command(vargs, w).Args, []string{
		"blobxfer",
		"--strip-components",
		"6",
		"my-storage-account",
		"my-container",
		filepath.Join(w.Path, vargs.Source),
	}) {
		t.Error("command not composed correctly")
	}
}
