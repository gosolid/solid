//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.sigevent
*/

type Sigevent struct {
  SigevNotify int `v8:"sigevNotify"`
  SigevSigno int `v8:"sigevSigno"`
  SigevValue void `v8:"sigevValue"`
  SigevNotifyFunction func(...any) (any, error) `v8:"sigevNotifyFunction"`
  SigevNotifyAttributes OpaquePthreadAttrT `v8:"sigevNotifyAttributes"`
}
