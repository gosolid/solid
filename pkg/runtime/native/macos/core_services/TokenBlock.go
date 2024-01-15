//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TokenBlock
*/

type TokenBlock struct {
  Source *byte `v8:"source"`
  SourceLength int64 `v8:"sourceLength"`
  TokenList *byte `v8:"tokenList"`
  TokenLength int64 `v8:"tokenLength"`
  TokenCount int64 `v8:"tokenCount"`
  StringList *byte `v8:"stringList"`
  StringLength int64 `v8:"stringLength"`
  StringCount int64 `v8:"stringCount"`
  DoString byte `v8:"doString"`
  DoAppend byte `v8:"doAppend"`
  DoAlphanumeric byte `v8:"doAlphanumeric"`
  DoNest byte `v8:"doNest"`
  LeftDelims [2]int16 `v8:"leftDelims"`
  RightDelims [2]int16 `v8:"rightDelims"`
  LeftComment [4]int16 `v8:"leftComment"`
  RightComment [4]int16 `v8:"rightComment"`
  EscapeCode int16 `v8:"escapeCode"`
  DecimalCode int16 `v8:"decimalCode"`
  ItlResource **byte `v8:"itlResource"`
  Reserved [8]int64 `v8:"reserved"`
}
