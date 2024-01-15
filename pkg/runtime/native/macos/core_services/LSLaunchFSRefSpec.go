//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.LSLaunchFSRefSpec
*/

type LSLaunchFSRefSpec struct {
  AppRef FSRef `v8:"appRef"`
  NumDocs uint64 `v8:"numDocs"`
  ItemRefs FSRef `v8:"itemRefs"`
  PassThruParams AEDesc `v8:"passThruParams"`
  LaunchFlags int `v8:"launchFlags"`
  AsyncRefCon void `v8:"asyncRefCon"`
}
