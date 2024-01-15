//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSMessagePort : Foundation.NSPort
*/

type NSMessagePort struct {
  *NSPort

}

