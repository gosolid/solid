//js:package native/macos/objc
package objc

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface objc.OS_object : objc.NSObject
*/

type OS_object struct {
  *NSObject

}

func (r *OS_object) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

