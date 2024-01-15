//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation._predicateFlags
*/

type PredicateFlags struct {
  EvaluationBlocked uint `v8:"evaluationBlocked"`
  ReservedPredicateFlags uint `v8:"reservedPredicateFlags"`
}
