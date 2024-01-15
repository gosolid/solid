//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_mutex_t
*/

type OpaquePthreadMutexT struct {
  Sig int64 `v8:"sig"`
  Opaque [56]byte `v8:"opaque"`
}
