//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.rusage_info_v2
*/

type RusageInfoV2 struct {
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
  RiChildUserTime uint64 `v8:"riChildUserTime"`
  RiChildSystemTime uint64 `v8:"riChildSystemTime"`
  RiChildPkgIdleWkups uint64 `v8:"riChildPkgIdleWkups"`
  RiChildInterruptWkups uint64 `v8:"riChildInterruptWkups"`
  RiChildPageins uint64 `v8:"riChildPageins"`
  RiChildElapsedAbstime uint64 `v8:"riChildElapsedAbstime"`
  RiDiskioBytesread uint64 `v8:"riDiskioBytesread"`
  RiDiskioByteswritten uint64 `v8:"riDiskioByteswritten"`
}
