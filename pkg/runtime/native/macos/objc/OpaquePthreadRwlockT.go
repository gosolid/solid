//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_rwlock_t
*/

type OpaquePthreadRwlockT struct {
  Sig int64 `v8:"sig"`
  Opaque [192]byte `v8:"opaque"`
}
