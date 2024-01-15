//js:package native/macos/cf-network
package cf_network

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CFNetwork.CFNetServiceClientContext
*/

type CFNetServiceClientContext struct {
  Version int64 `v8:"version"`
  Info void `v8:"info"`
  Retain *func(...any) (any, error) `v8:"retain"`
  Release *func(...any) (any, error) `v8:"release"`
  CopyDescription *func(...any) (any, error) `v8:"copyDescription"`
}
