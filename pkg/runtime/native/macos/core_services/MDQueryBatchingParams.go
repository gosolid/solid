//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.MDQueryBatchingParams
*/

type MDQueryBatchingParams struct {
  FirstMaxNum uint64 `v8:"firstMaxNum"`
  FirstMaxMs uint64 `v8:"firstMaxMs"`
  ProgressMaxNum uint64 `v8:"progressMaxNum"`
  ProgressMaxMs uint64 `v8:"progressMaxMs"`
  UpdateMaxNum uint64 `v8:"updateMaxNum"`
  UpdateMaxMs uint64 `v8:"updateMaxMs"`
}
