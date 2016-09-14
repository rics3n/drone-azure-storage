package main

import (
	"github.com/drone/drone-plugin-go/plugin"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCommandBuildCorrectly(t *testing.T) {
	vargs := AzureBlobxfer{}
	vargs.StorageAccountKey = "xyzabc"
	vargs.Container = "my-container"
	vargs.StorageAccountName = "my-storage-account"
	vargs.Source = "__source__"
	vargs.Dest = "__dest__"
	w := plugin.Workspace{Path: "/test/path"}
	if !reflect.DeepEqual(command(vargs, w).Args, []string{
		"blobxfer",
		"my-storage-account",
		"my-container",
		filepath.Join(w.Path, vargs.Source),
		"--upload",
		"--remoteresource",
		vargs.Dest,

	}) {
		t.Error("command not composed correctly")
	}
}
