//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct IOKit.DriverServiceInfo
*/

type DriverServiceInfo struct {
  ServiceCategory uint `v8:"serviceCategory"`
  ServiceType uint `v8:"serviceType"`
  ServiceVersion objc.NumVersion `v8:"serviceVersion"`
}
