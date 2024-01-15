//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.RuleBasedTrslRecord
*/

type RuleBasedTrslRecord struct {
  SourceType int16 `v8:"sourceType"`
  TargetType int16 `v8:"targetType"`
  FormatNumber int16 `v8:"formatNumber"`
  PropertyFlag int16 `v8:"propertyFlag"`
  NumberOfRules int16 `v8:"numberOfRules"`
}
