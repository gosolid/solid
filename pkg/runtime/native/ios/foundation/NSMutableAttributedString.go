//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSMutableAttributedString : Foundation.NSAttributedString
*/

type NSMutableAttributedString struct {
  *NSAttributedString

}

func (r *NSMutableAttributedString) ReplaceCharactersInRange() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSMutableAttributedString) SetAttributes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

