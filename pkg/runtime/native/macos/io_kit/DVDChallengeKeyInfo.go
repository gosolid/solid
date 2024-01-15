//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.DVDChallengeKeyInfo
*/

type DVDChallengeKeyInfo struct {
  DataLength [2]byte `v8:"dataLength"`
  Reserved [2]byte `v8:"reserved"`
  ChallengeKeyValue [10]byte `v8:"challengeKeyValue"`
  Reserved2 [2]byte `v8:"reserved2"`
}
