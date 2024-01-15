//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSSpecifierTest : Foundation.NSScriptWhoseTest
*/

type NSSpecifierTest struct {
  *NSScriptWhoseTest

}

func (r *NSSpecifierTest) InitWithObjectSpecifier() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpecifierTest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSSpecifierTest) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

