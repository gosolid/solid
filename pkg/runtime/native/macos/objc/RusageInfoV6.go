//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

/*
struct objc.rusage_info_v6
*/

type RusageInfoV6 struct {
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
  RiCpuTimeQosDefault uint64 `v8:"riCpuTimeQosDefault"`
  RiCpuTimeQosMaintenance uint64 `v8:"riCpuTimeQosMaintenance"`
  RiCpuTimeQosBackground uint64 `v8:"riCpuTimeQosBackground"`
  RiCpuTimeQosUtility uint64 `v8:"riCpuTimeQosUtility"`
  RiCpuTimeQosLegacy uint64 `v8:"riCpuTimeQosLegacy"`
  RiCpuTimeQosUserInitiated uint64 `v8:"riCpuTimeQosUserInitiated"`
  RiCpuTimeQosUserInteractive uint64 `v8:"riCpuTimeQosUserInteractive"`
  RiBilledSystemTime uint64 `v8:"riBilledSystemTime"`
  RiServicedSystemTime uint64 `v8:"riServicedSystemTime"`
  RiLogicalWrites uint64 `v8:"riLogicalWrites"`
  RiLifetimeMaxPhysFootprint uint64 `v8:"riLifetimeMaxPhysFootprint"`
  RiInstructions uint64 `v8:"riInstructions"`
  RiCycles uint64 `v8:"riCycles"`
  RiBilledEnergy uint64 `v8:"riBilledEnergy"`
  RiServicedEnergy uint64 `v8:"riServicedEnergy"`
  RiIntervalMaxPhysFootprint uint64 `v8:"riIntervalMaxPhysFootprint"`
  RiRunnableTime uint64 `v8:"riRunnableTime"`
  RiFlags uint64 `v8:"riFlags"`
  RiUserPtime uint64 `v8:"riUserPtime"`
  RiSystemPtime uint64 `v8:"riSystemPtime"`
  RiPinstructions uint64 `v8:"riPinstructions"`
  RiPcycles uint64 `v8:"riPcycles"`
  RiEnergyNj uint64 `v8:"riEnergyNj"`
  RiPenergyNj uint64 `v8:"riPenergyNj"`
  RiReserved [14]uint64 `v8:"riReserved"`
}
