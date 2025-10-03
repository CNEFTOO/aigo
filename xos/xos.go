package xos

import "runtime"

func GetPlatform() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "darwin":
		return "MacOSX"
	case "freebsd":
		return "OpenBSD"
	case "netbsd":
		return "NetBSD"
	case "Solaris":
		return "solaris"
	case "Android":
		return "android"
	case "IOS":
		return "ios"
	}
	return runtime.GOOS
}

func GetPlatformArch() string {
	return runtime.GOARCH
}

func GetOSXCPUNumber() int {
	return runtime.NumCPU()
}

func GetSystemVersion() string {
	return runtime.Version()
}
