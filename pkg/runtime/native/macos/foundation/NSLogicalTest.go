//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSLogicalTest : Foundation.NSScriptWhoseTest
*/

type NSLogicalTest struct {
  *NSScriptWhoseTest

}

func (r *NSLogicalTest) InitOrTestWithTests() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLogicalTest) InitNotTestWithTest() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSLogicalTest) InitAndTestWithTests() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

