//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.OSClassDescription
*/

type OSClassDescription struct {
  DescriptionSize uint `v8:"descriptionSize"`
  Name [96]byte `v8:"name"`
  SuperName [96]byte `v8:"superName"`
  MethodOptionsSize uint `v8:"methodOptionsSize"`
  MethodOptionsOffset uint `v8:"methodOptionsOffset"`
  MetaMethodOptionsSize uint `v8:"metaMethodOptionsSize"`
  MetaMethodOptionsOffset uint `v8:"metaMethodOptionsOffset"`
  QueueNamesSize uint `v8:"queueNamesSize"`
  QueueNamesOffset uint `v8:"queueNamesOffset"`
  MethodNamesSize uint `v8:"methodNamesSize"`
  MethodNamesOffset uint `v8:"methodNamesOffset"`
  MetaMethodNamesSize uint `v8:"metaMethodNamesSize"`
  MetaMethodNamesOffset uint `v8:"metaMethodNamesOffset"`
  Flags uint64 `v8:"flags"`
  Resv1 [8]uint64 `v8:"resv1"`
  MethodOptions [0]uint64 `v8:"methodOptions"`
  MetaMethodOptions [0]uint64 `v8:"metaMethodOptions"`
  DispatchNames [0]byte `v8:"dispatchNames"`
  MethodNames [0]byte `v8:"methodNames"`
  MetaMethodNames [0]byte `v8:"metaMethodNames"`
}
