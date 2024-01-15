//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.fd_set
*/

type FdSet struct {
  FdsBits [32]int `v8:"fdsBits"`
}
