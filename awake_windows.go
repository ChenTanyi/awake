package awake

import (
	"syscall"
	"time"
)

const (
	interval = 30 * time.Second
	// https://docs.microsoft.com/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate
	ES_AWAYMODE_REQUIRED = 0x00000040
	ES_CONTINUOUS        = 0x80000000
	ES_DISPLAY_REQUIRED  = 0x00000002
	ES_SYSTEM_REQUIRED   = 0x00000001
)

var (
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
)

func Awake() {
	ticker := time.NewTicker(interval)
	SetThreadExecutionState := kernel32.NewProc("SetThreadExecutionState")

	for range ticker.C {
		SetThreadExecutionState.Call(uintptr(ES_DISPLAY_REQUIRED))
	}
}

func AwakeTemperory(duration time.Duration) {
	ticker := time.NewTicker(interval)
	timer := time.NewTimer(duration)
	SetThreadExecutionState := kernel32.NewProc("SetThreadExecutionState")

	for {
		select {
		case <-ticker.C:
			SetThreadExecutionState.Call(uintptr(ES_DISPLAY_REQUIRED))
		case <-timer.C:
			return
		}
	}
}
