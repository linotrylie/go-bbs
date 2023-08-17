package utils

import (
	"GoFreeBns/utils/sliceutil"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

const RunModeDebug = "debug"     // 开发环境
const runModeTest = "test"       // 测试环境
const runModeRelease = "release" // 生产环境

// OsIsWindows 是否windows运行系统
func OsIsWindows() bool {
	return runtime.GOOS == "windows"
}

// RunModeIsDebug 是否开发环境
// debug 或 测试 都视为开发环境
// 注意:在测试文件中调用,都是 debug
func RunModeIsDebug() bool {
	return gin.Mode() == RunModeDebug || gin.Mode() == runModeTest
}

// RunModeIsRelease 是否生产环境
// 注意: 在测试文件中调用,都是 debug
func RunModeIsRelease() bool {
	return gin.Mode() == runModeRelease
}

func PanicInfo() string {
	maxCallerDepth := 100
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("file: %s\nline: %d\nfunction: %s\n", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}

	runtimePanicTypes := []string{
		"runtime.panicdivide",
		"runtime.sigpanic",
		"runtime.goPanicIndex",
		"runtime.panicdottypeE",
		"runtime.closechan",
		"runtime.mapassign_faststr",
		"runtime.deferreturn",
	}

	callersLen := len(callers)
	for index, data := range callers {
		if sliceutil.ContainsAny(data, runtimePanicTypes) {
			if index+1 <= callersLen {
				target := callers[index+1]
				if !strings.Contains(target, "runtime.") {
					return target
				}
			}
		}
	}

	// 没有找到位置,返回所有信息
	marshal, err := json.Marshal(callers)
	if err != nil {
		return ""
	}
	return string(marshal)
}
