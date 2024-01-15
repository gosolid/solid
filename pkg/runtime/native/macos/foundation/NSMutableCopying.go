//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSMutableCopying
*/

type NSMutableCopying struct {

}

func (r *NSMutableCopying) MutableCopyWithZone() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

