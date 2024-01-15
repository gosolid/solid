//js:package native/ios/foundation
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

func (r *NSKeyedUnarchiver) SetDecodingFailurePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeDoubleForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeFloatForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) RequiresSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodingFailurePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedArrayOfObjectsOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedObjectOfClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveObjectWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveTopLevelObjectWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedDictionaryWithKeysOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) InitForReadingWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeBoolForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeIntForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) FinishDecoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeInt64ForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedArrayOfObjectsOfClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) ClassForClassName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeInt32ForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeBytesForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) InitForReadingFromData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchiveObjectWithFile() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) DecodeObjectForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) SetRequiresSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedDictionaryWithKeysOfClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) ContainsValueForKey() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedUnarchiver) UnarchivedObjectOfClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

