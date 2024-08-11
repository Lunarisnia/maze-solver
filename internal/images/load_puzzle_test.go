package images

import "testing"

func Test_LoadPuzzle(t *testing.T) {
	_, err := loadPuzzle("foo/bar")
	if err != nil {
		t.Fail()
	}
}
