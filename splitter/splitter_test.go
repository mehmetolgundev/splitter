package splitter

import "testing"

func TestSplit(t *testing.T) {

	chanResult := Split("../test/testfile", 1)
	if chanResult == nil {
		t.Fatal("channel not found")
	}
	result := <-chanResult
	if result.Err != nil {
		t.Fatal(result.Err)

	}
	if len(result.Files) != 8 {
		t.Fatalf("file count isn't 8. count:%d", len(result.Files))
	}

}
