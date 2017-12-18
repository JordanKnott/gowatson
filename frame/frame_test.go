package frame

import (
	"testing"
	. "gopkg.in/go-playground/assert.v1"
)

func TestAddFrame(t *testing.T) {
	AddFrame(Frame{ID: "1", Project: "ABC", Start: "0", Stop: "15", Tags: []string{}, UpdatedAt: "0"})

	frame, err := GetFrame("1")
	Equal(t, err, nil)
	Equal(t, frame, Frame{ID: "1", Project: "ABC", Start: "0", Stop: "15", Tags: []string{}, UpdatedAt: "0"})
}

func TestClearFrames(t *testing.T) {
	Equal(t, len(GetFrames()), 1)
	ClearFrames()
	Equal(t, len(GetFrames()), 0)
}

func TestGetInvalidFrame(t *testing.T) {
	_, err := GetFrame("1")
	NotEqual(t, err, nil)
}

func TestDeleteFrame(t *testing.T) {
	AddFrame(Frame{ID: "1", Project: "ABC", Start: "0", Stop: "15", Tags: []string{}, UpdatedAt: "0"})
	Equal(t, len(GetFrames()), 1)
	DeleteFrame("1")
	Equal(t, len(GetFrames()), 0)
}

func TestSaveFrames(t *testing.T) {
	AddFrame(Frame{ID: "1", Project: "ABC", Start: "0", Stop: "15", Tags: []string{}, UpdatedAt: "0"})

	savedFrames := SerializeFrames()
	expectedJson := `[{"ID":"1","Project":"ABC","Start":"0","Stop":"15","Tags":[],"Message":"","UpdatedAt":"0"}]`
	Equal(t, string(savedFrames[:]), expectedJson)
}

func TestLoadFrames(t *testing.T) {
	loadedJsonRaw := []byte(`[{"ID":"1","Project":"ZFG","Start":"0","Stop":"15","Tags":[],"Message":"","UpdatedAt":"0"}]`)
	UnserializeFrames(loadedJsonRaw)

	Equal(t, len(GetFrames()), 1)
	frame, err := GetFrame("ZFG")
	NotEqual(t, err, nil)
	Equal(t, frame.ID, "1")
	Equal(t, frame.Project, "ZFG")
	Equal(t, frame.Start, "0")
	Equal(t, frame.Stop, "15")
	Equal(t, frame.Tags, []string{})
	Equal(t, frame.Message, "")
	Equal(t, frame.UpdatedAt, "0")

}