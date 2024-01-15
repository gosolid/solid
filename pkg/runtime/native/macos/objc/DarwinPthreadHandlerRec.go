//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.__darwin_pthread_handler_rec
*/

type DarwinPthreadHandlerRec struct {
  Routine func(...any) (any, error) `v8:"routine"`
  Arg void `v8:"arg"`
  Next DarwinPthreadHandlerRec `v8:"next"`
}
