//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSKeyedArchiver : Foundation.NSCoder
*/

type NSKeyedArchiver struct {
  *NSCoder

}

func (r *NSKeyedArchiver) RequiresSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) ArchiveRootObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) ClassNameForClass() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt32() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeBytes() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) OutputFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) ArchivedDataWithRootObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetClassName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt64() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeDouble() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeFloat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetOutputFormat() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodedData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) InitRequiringSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) FinishEncoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeConditionalObject() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeInt() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) SetRequiresSecureCoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) InitForWritingWithMutableData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSKeyedArchiver) EncodeBool() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

