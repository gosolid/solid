//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_services"
)

/*
struct ApplicationServices.ICFileSpec
*/

type ICFileSpec struct {
  VolName [32]byte `v8:"volName"`
  VolCreationDate int `v8:"volCreationDate"`
  Fss core_services.FSSpec `v8:"fss"`
  Alias core_services.AliasRecord `v8:"alias"`
}
