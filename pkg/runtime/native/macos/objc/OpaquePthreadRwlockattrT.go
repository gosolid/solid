//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_rwlockattr_t
*/

type OpaquePthreadRwlockattrT struct {
  Sig int64 `v8:"sig"`
  Opaque [16]byte `v8:"opaque"`
}
