package backend

/*
import (
	"github.com/yuin/gopher-lua"
	"watson/frame"
	"encoding/json"
)

func CallFunction(name string, L *lua.LState) {
	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal(name),
		NRet: 1,
		Protect: true,
	}); err != nil {
		panic(err)
	}
}

func PluginInit(L *lua.LState) {
	CallFunction("init", L)
}

func PluginCleanup(L *lua.LState) {
	CallFunction("cleanup", L)
}

const luaFrameTypeName = "frame"
func RegisterFrameType(L *lua.LState) {
	mt := L.NewTypeMetatable(luaFrameTypeName)
	L.SetGlobal("frame", mt)
	L.SetField(mt, "new", L.NewFunction(newFrame))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), frameMethods))
}

func newFrame(L *lua.LState) int {
	id := L.CheckString(1)
	project := L.CheckString(2)
	start := L.CheckString(3)
	stop := L.CheckString(3)
	tags := L.CheckString(3)
	tags := L.CheckString(3)
	person := &frame.Frame{ID: id, Name: name}
	ud := L.NewUserData()
	ud.Value = person
	L.SetMetatable(ud, L.GetTypeMetatable(luaFrameTypeName))
	L.Push(ud)
	return 1
}

func checkFrame(L *lua.LState) *frame.Frame {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*frame.Frame); ok {
		return v
	}
	L.ArgError(1, "frame expected")
	return nil
}

var frameMethods = map[string]lua.LGFunction{
	"id" : frameGetSetID,
	"name" : frameGetSetName,
}

func frameGetSetID(L *lua.LState) int {
	f := checkFrame(L)
	if L.GetTop() == 2 {
		f.ID = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.ID))
	return 1
}

func frameGetSetName(L *lua.LState) int {
	f := checkFrame(L)
	if L.GetTop() == 2 {
		f.Name = L.CheckString(2)
		return 0
	}
	L.Push(lua.LString(f.Name))
	return 1
}

func PluginSync(L *lua.LState, frames []frame.Frame) string {

	b, err := json.Marshal(frames)
	if err != nil {
		panic(err)
	}

	if err := L.CallByParam(lua.P{
		Fn: L.GetGlobal("sync"),
		NRet: 1,
		Protect: true,
	}, lua.LString(b)); err != nil {
		panic(err)
	}

	ret := L.Get(-1)
	L.Pop(1)
	if str, ok := ret.(lua.LString); ok {
		return string(str)
	} else {
		return ""
	}
}
*/