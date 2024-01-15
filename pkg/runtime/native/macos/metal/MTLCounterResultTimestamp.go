//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLCounterResultTimestamp
*/

type MTLCounterResultTimestamp struct {
  Timestamp uint64 `v8:"timestamp"`
}
