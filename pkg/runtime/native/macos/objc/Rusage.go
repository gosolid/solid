//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.rusage
*/

type Rusage struct {
  RuUtime Timeval `v8:"ruUtime"`
  RuStime Timeval `v8:"ruStime"`
  RuMaxrss int64 `v8:"ruMaxrss"`
  RuIxrss int64 `v8:"ruIxrss"`
  RuIdrss int64 `v8:"ruIdrss"`
  RuIsrss int64 `v8:"ruIsrss"`
  RuMinflt int64 `v8:"ruMinflt"`
  RuMajflt int64 `v8:"ruMajflt"`
  RuNswap int64 `v8:"ruNswap"`
  RuInblock int64 `v8:"ruInblock"`
  RuOublock int64 `v8:"ruOublock"`
  RuMsgsnd int64 `v8:"ruMsgsnd"`
  RuMsgrcv int64 `v8:"ruMsgrcv"`
  RuNsignals int64 `v8:"ruNsignals"`
  RuNvcsw int64 `v8:"ruNvcsw"`
  RuNivcsw int64 `v8:"ruNivcsw"`
}
