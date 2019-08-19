package main

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	_ "gocloud.dev/blob/fileblob"
)

const file = "test.txt"

func TestRootCommand_Valid(t *testing.T) {
	dir, clean := newTempDir()
	defer clean()
	ctx := context.Background()

	upload(ctx, "file:///"+dir, file)
	if uploaded, err := fileExists(dir + "/" + file); err != nil && uploaded {
		t.Fatal("FAIL: Validation Failed without 'Not Exist' error: ", err)
	} else if !uploaded {
		t.Fatal("FAIL: File was not found in specified location: ", err)
	}
}

func fileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	return !os.IsNotExist(err), err
}

func newTempDir() (string, func()) {
	dir, err := ioutil.TempDir("", "upload-test-dir")
	if err != nil {
		panic(err)
	}

	return dir, func() {
		os.RemoveAll(dir)
	}
}
