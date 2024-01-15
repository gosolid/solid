//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSSpellServerDelegate
*/

type NSSpellServerDelegate struct {

}

func (r *NSSpellServerDelegate) SpellServer() (NSRange, error) {
  return NSRange{}, fmt.Errorf("unimplemented")
}

