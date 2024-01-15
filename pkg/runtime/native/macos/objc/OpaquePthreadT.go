//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_t
*/

type OpaquePthreadT struct {
  Sig int64 `v8:"sig"`
  CleanupStack DarwinPthreadHandlerRec `v8:"cleanupStack"`
  Opaque [8176]byte `v8:"opaque"`
}
