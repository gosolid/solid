//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc._opaque_pthread_cond_t
*/

type OpaquePthreadCondT struct {
  Sig int64 `v8:"sig"`
  Opaque [40]byte `v8:"opaque"`
}
