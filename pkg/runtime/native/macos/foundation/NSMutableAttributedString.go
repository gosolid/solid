//js:package native/macos/foundation
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

func (r *NSMutableAttributedString) ReplaceCharactersInRange() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSMutableAttributedString) SetAttributes() error {
  return fmt.Errorf("unimplemented")
}

