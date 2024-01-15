//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSKeyedUnarchiver : Foundation.NSCoder
*/

type NSKeyedUnarchiver struct {
  *NSCoder

}

func (r *NSKeyedUnarchiver) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeIntForKey() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetDecodingFailurePolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) FinishDecoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedObjectOfClass() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedDictionaryWithKeysOfClass() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveTopLevelObjectWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) InitForReadingFromData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodingFailurePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeInt32ForKey() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeFloatForKey() (float32, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeDoubleForKey() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) RequiresSecureCoding() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedArrayOfObjectsOfClass() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedObjectOfClasses() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) InitForReadingWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetDelegate() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveObjectWithFile() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetClass() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) ClassForClassName() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveObjectWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeObjectForKey() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeInt64ForKey() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeBoolForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetRequiresSecureCoding() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeBytesForKey() (byte, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedArrayOfObjectsOfClasses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedDictionaryWithKeysOfClasses() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) ContainsValueForKey() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

