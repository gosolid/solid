//go:build darwin && !ios

package main

/*
#cgo darwin CFLAGS: -I/usr/local/include/chromium/third_party/boringssl/src/include/ -I/usr/local/include/chromium -I/usr/local/include/chromium/third_party/abseil-cpp -I/usr/local/include/chromium/third_party/skia -x objective-c
#cgo darwin CXXFLAGS: -I/usr/local/include/chromium/third_party/boringssl/src/include/ -I/usr/local/include/chromium -I/usr/local/include/chromium/third_party/abseil-cpp -I/usr/local/include/chromium/third_party/skia -x objective-c++ -std=c++20
#cgo darwin LDFLAGS: -L/usr/local/lib/chromium -lchromium -lcxx_cppdeps -lstdc++ -F/usr/local/lib/chromium -framework AuthenticationServices -framework DiskArbitration -framework ForceFeedback -framework CoreMIDI -framework Accessibility -framework CoreLocation -framework SafariServices -lcups -lbsm -framework CoreWLAN -framework ApplicationServices -framework SystemConfiguration -framework UserNotifications -framework ScreenTime -framework SecurityInterface -framework GameController -framework MediaPlayer -framework ImageCaptureCore -framework Metal -framework MetalKit -framework AppKit -framework CoreFoundation -framework QuartzCore -framework CoreText -framework IOSurface -framework AVFoundation -framework Cocoa -framework OpenGL -framework CoreVideo -framework MediaAccessibility -framework CoreGraphics -framework AudioUnit -framework AVFAudio -framework PDFKit -framework CoreServices -framework VideoToolbox -framework CoreMedia -framework IOKit -framework OpenDirectory -framework CoreBluetooth -framework IOBluetooth -framework CryptoTokenKit -framework ScreenCaptureKit -framework UniformTypeIdentifiers -framework LocalAuthentication -framework CoreAudio -framework Accelerate -framework AudioToolbox -framework Security -framework Carbon

int macosmain();
*/
import "C"

import (
	"os"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	// startWorker()

	exitCode := C.macosmain()
	os.Exit(int(exitCode))
}
