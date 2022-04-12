// Copyright 2022 Robert S. Muhlestein.
// SPDX-License-Identifier: Apache-2.0

package compfile_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/compfile"
)

func ExampleFile_nil() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil))
	fmt.Println(compfile.New().Complete(nil, ""))
	//Output:
	// [bar/ blah/ come/ foo.go]
	// [bar/ blah/ come/ foo.go]
}

func ExampleFile_pre_File() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "fo"))
	//Output:
	// [foo.go]
}

func ExampleFile_pre_Dir_Only() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "bar"))
	fmt.Println(compfile.New().Complete(nil, "bar/"))
	//Output:
	// [bar/foo.go bar/other]
	// [bar/foo.go bar/other]
}

func ExampleFile_pre_Dir_or_Files() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "b"))
	//Output:
	// [bar/ blah/]
}

func ExampleFile_pre_Dir_Specific() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "blah"))
	//Output:
	// [blah/file1 blah/file2]
}

func ExampleFile_pre_Dir_Recurse() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "com"))
	fmt.Println(compfile.New().Complete(nil, "come/"))
	//Output:
	// [come/one]
	// [come/one]
}

func ExampleFile_dir_File() {
	os.Chdir("testdata/file")
	defer os.Chdir("../..")
	fmt.Println(compfile.New().Complete(nil, "bar/fo"))
	fmt.Println(compfile.New().Complete(nil, "bar/foo"))
	//Output:
	// [bar/foo.go]
	// [bar/foo.go]
}
