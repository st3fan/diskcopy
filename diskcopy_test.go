package diskcopy_test

import (
	"io"
	"os"
	"testing"

	"github.com/st3fan/diskcopy"
)

func Test_Open(t *testing.T) {
	file, err := os.Open("testdata/VideoWorks Disk 1.image")
	if err != nil {
		t.Fail()
	}

	image, err := diskcopy.NewImage(file)
	if err != nil {
		t.Fail()
	}

	if _, err := image.Seek(0x446d4-84, io.SeekStart); err != nil {
		t.Fail()
	}

	walkTheDog := make([]byte, 12)
	if _, err := image.Read(walkTheDog); err != nil {
		t.Fail()
	}

	if string(walkTheDog) != "WALK THE DOG" {
		t.Fail()
	}

	defer image.Close()
}
