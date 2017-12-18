package main

import (
	"github.com/yuin/gopher-lua"
	"github.com/cjoudrey/gluahttp"
	"layeh.com/gopher-json"
	"net/http"
	"watson/frame"
	"fmt"
	"io/ioutil"
	"strings"
	"path/filepath"
	"watson/config"
)

func BackendStuff(filename string) {
	L := lua.NewState()
	defer L.Close()

	// Register User Types
	// backend.RegisterFrameType(L)

	// Register modules
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
	L.PreloadModule("json", json.Loader)

	// Load file and execute it
	if err := L.DoFile(filename); err != nil {
		panic(err)
	}

	// Run the plugin init function
	// backend.PluginInit(L)

	/*
	frames := make([]frame.Frame, 3)
	frames[0] = frame.Frame{ID: "1", Name: "Frame One"}
	frames[1] = frame.Frame{ID: "2", Name: "Frame Two"}
	frames[2] = frame.Frame{ID: "3", Name: "Frame Three"}
	*/

	// Sync frames
	//fmt.Printf(backend.PluginSync(L, frames))
}

func LoadDriver() {
	config.LoadConfig()
	driver := config.GetBackendDriver()
	println("Driver is " + driver)

	files, err := ioutil.ReadDir("./plugins")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		name := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		if driver == name {
			fmt.Println("Found driver! " + name)
		}
	}
}

func main() {


	frame.AddFrame(frame.Frame{ID: "1", Project: "ABC", Start: "0", Stop: "15", Tags: []string {}, UpdatedAt: "0"})
	f, _ := frame.GetFrame("2")

}
