//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableString : Foundation.NSString
*/

type NSMutableString struct {
  *NSString

}

func (r *NSMutableString) ReplaceCharactersInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

