//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct CoreServices.LSLaunchURLSpec
*/

type LSLaunchURLSpec struct {
  AppURL *core_foundation.CFURL `v8:"appURL"`
  ItemURLs *core_foundation.CFArray `v8:"itemURLs"`
  PassThruParams AEDesc `v8:"passThruParams"`
  LaunchFlags int `v8:"launchFlags"`
  AsyncRefCon void `v8:"asyncRefCon"`
}
