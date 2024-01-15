//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSWhoseSpecifier : Foundation.NSScriptObjectSpecifier
*/

type NSWhoseSpecifier struct {
  *NSScriptObjectSpecifier

}

func (r *NSWhoseSpecifier) StartSubelementIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) SetStartSubelementIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) SetStartSubelementIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) SetEndSubelementIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) InitWithContainerClassDescription() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) Test() (*NSScriptWhoseTest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) StartSubelementIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) EndSubelementIdentifier() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) EndSubelementIndex() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) SetEndSubelementIndex() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSWhoseSpecifier) SetTest() error {
  return fmt.Errorf("unimplemented")
}

