//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSISO8601DateFormatter : Foundation.NSFormatter
*/

type NSISO8601DateFormatter struct {
  *NSFormatter

}

func (r *NSISO8601DateFormatter) StringFromDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) DateFromString() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) TimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) SetTimeZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) FormatOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) SetFormatOptions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

