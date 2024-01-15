//js:package native/macos/io-kit
package io_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
struct IOKit.fdisk_part
*/

type FdiskPart struct {
  Bootid byte `v8:"bootid"`
  Beghead byte `v8:"beghead"`
  Begsect byte `v8:"begsect"`
  Begcyl byte `v8:"begcyl"`
  Systid byte `v8:"systid"`
  Endhead byte `v8:"endhead"`
  Endsect byte `v8:"endsect"`
  Endcyl byte `v8:"endcyl"`
  Relsect uint `v8:"relsect"`
  Numsect uint `v8:"numsect"`
}
