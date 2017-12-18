package frame

import (
	"errors"
	"encoding/json"
	"watson/common"
)

type Frame struct {
	ID        string
	Project   string
	Start     string
	Stop      string
	Tags      []string
	Message   string
	UpdatedAt string
}

var frames []Frame

func AddFrame(frame Frame) {
	frames = append(frames, frame)
}

func DeleteFrame(ID string) bool {
	for index, frame := range frames {
		if frame.ID == ID {
			frames = append(frames[:index], frames[index + 1:]...)
			return true
		}
	}
	return false
}

func ClearFrames() {
	frames = []Frame{}
}

func GetFrame(ID string) (Frame, error) {
	var frame Frame

	for _, f := range frames {
		if f.ID == ID {
			return f, nil
		}
	}
	return frame, errors.New("Invalid ID")
}

func GetFrames() []Frame {
	return frames
}

func SerializeFrames() []byte {
	framesJson, err := json.Marshal(frames)
	common.PanicIfError(err)
	return framesJson
}

func SerializeFramesPretty() []byte {
	framesJson, err := json.MarshalIndent(frames, "", "  ")
	common.PanicIfError(err)
	return framesJson
}

func UnserializeFrames(framesJson []byte) {
	emptyFrames := make([]Frame, 0)
	json.Unmarshal(framesJson, &emptyFrames)
	frames = emptyFrames
	PrintFrames()
}

func PrintFrames() {
	for _, frame := range frames {
		print(frame.Project)
	}
}