//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSSecureCoding
*/

type NSSecureCoding struct {

}

func (r *NSSecureCoding) SupportsSecureCoding() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

