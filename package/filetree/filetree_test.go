package filetree

import (
	"testing"
)

// Expected structure of the test directory
var expectedStructure = map[string][]string{
	"testdata": {
		"file1.txt",
		"file2.log",
		"dir1",
	},
	"testdata/dir1": {
		"file3.txt",
		"file4.md",
		"subdir",
	},
	"testdata/dir1/subdir": {
		"file5.txt",
		"file6.json",
	},
}

// TestBuildTree validates that the generated tree matches the expected structure
func TestBuildTree(t *testing.T) {
	testPath := "testdata"
	hOpt := HiddenOptions{
		All:   true,
		Files: true,
		Dirs:  true,
	}

	tree, err := BuildTree(testPath, &hOpt)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}

	// Validate the root directory
	if tree.Root == nil {
		t.Fatal("Expected root node, got nil")
	}
	if tree.Root.Name != "testdata" {
		t.Errorf("Expected root node name 'testdata', got '%s'", tree.Root.Name)
	}

	// Recursively validate the structure
	validateTree(t, tree.Root, "testdata")
}

// validateTree recursively checks if the file tree matches the expected structure
func validateTree(t *testing.T, node *Node, path string) {
	expectedFiles, exists := expectedStructure[path]
	if !exists {
		t.Fatalf("Unexpected path in tree: %s", path)
	}

	// Track found files
	foundFiles := map[string]bool{}
	for _, child := range node.Children {
		foundFiles[child.Name] = true

		// If it's a directory, check its children
		if child.IsDir {
			validateTree(t, child, path+"/"+child.Name)
		}
	}

	// Ensure all expected files exist
	for _, expectedFile := range expectedFiles {
		if !foundFiles[expectedFile] {
			t.Errorf("Missing file or directory: %s/%s", path, expectedFile)
		}
	}
}
