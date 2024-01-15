//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.rusage_info_v0
*/

type RusageInfoV0 struct {
  RiUuid [16]byte `v8:"riUuid"`
  RiUserTime uint64 `v8:"riUserTime"`
  RiSystemTime uint64 `v8:"riSystemTime"`
  RiPkgIdleWkups uint64 `v8:"riPkgIdleWkups"`
  RiInterruptWkups uint64 `v8:"riInterruptWkups"`
  RiPageins uint64 `v8:"riPageins"`
  RiWiredSize uint64 `v8:"riWiredSize"`
  RiResidentSize uint64 `v8:"riResidentSize"`
  RiPhysFootprint uint64 `v8:"riPhysFootprint"`
  RiProcStartAbstime uint64 `v8:"riProcStartAbstime"`
  RiProcExitAbstime uint64 `v8:"riProcExitAbstime"`
}
