//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSCopying
*/

type NSCopying struct {

}

func (r *NSCopying) CopyWithZone() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

