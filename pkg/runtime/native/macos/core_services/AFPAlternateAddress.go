//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AFPAlternateAddress
*/

type AFPAlternateAddress struct {
  FVersion byte `v8:"fVersion"`
  FAddressCount byte `v8:"fAddressCount"`
  FAddressList [1]byte `v8:"fAddressList"`
}
