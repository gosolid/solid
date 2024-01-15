//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_condattr_t
*/

type OpaquePthreadCondattrT struct {
  Sig int64 `v8:"sig"`
  Opaque [8]byte `v8:"opaque"`
}
