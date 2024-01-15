//js:package native/macos/foundation
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
  *NSSecureCoding
}

func (r *NSISO8601DateFormatter) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) StringFromDate() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) DateFromString() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) TimeZone() (*NSTimeZone, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) SetTimeZone() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) FormatOptions() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSISO8601DateFormatter) SetFormatOptions() error {
  return fmt.Errorf("unimplemented")
}

