//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSUnarchiver : Foundation.NSCoder
*/

type NSUnarchiver struct {
  *NSCoder

}

func (r *NSUnarchiver) SetObjectZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) DecodeClassName() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) ClassNameDecodedForArchiveClassName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) IsAtEnd() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) InitForReadingWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) ObjectZone() (*foundation.NSZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) UnarchiveObjectWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) UnarchiveObjectWithFile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) ReplaceObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUnarchiver) SystemVersion() (uint, error) {
  return 0, fmt.Errorf("unimplemented")
}

