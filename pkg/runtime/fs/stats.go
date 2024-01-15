//js:package fs

package fs

//go:generate go run github.com/grexie/isolates/codegen

import (
	"time"
)

var _ Stats = &stats{}
var _ statsv8 = &stats{}

//js:class Stats
//js:export Stats
type stats struct {
	isFile         bool
	isDirectory    bool
	isSymbolicLink bool
	contentType    string
	accessed       int64
	modified       int64
	changed        int64
	created        int64
	size           int64
	mode           FileMode
	uid            int
	gid            int
}

//js:method
func (s *stats) IsDirectory() bool {
	return s.isDirectory
}

//js:method
func (s *stats) IsFile() bool {
	return s.isFile
}

//js:method
func (s *stats) IsSymbolicLink() bool {
	return s.isSymbolicLink
}

//js:get contentType
func (s *stats) ContentType() string {
	return s.contentType
}

//js:get atimeMs
func (s *stats) AccessedMilli() int64 {
	return s.accessed
}

//js:get mtimeMs
func (s *stats) ModifiedMilli() int64 {
	return s.modified
}

//js:get ctimeMs
func (s *stats) ChangedMilli() int64 {
	return s.changed
}

//js:get birthtimeMs
func (s *stats) CreatedMilli() int64 {
	return s.created
}

//js:get atime
func (s *stats) Accessed() time.Time {
	return time.UnixMilli(s.accessed)
}

//js:get mtime
func (s *stats) Modified() time.Time {
	return time.UnixMilli(s.modified)
}

//js:get ctime
func (s *stats) Changed() time.Time {
	return time.UnixMilli(s.changed)
}

//js:get birthtime
func (s *stats) Created() time.Time {
	return time.UnixMilli(s.created)
}

//js:get size
func (s *stats) Size() int64 {
	return s.size
}

//js:get uid
func (s *stats) Uid() int {
	return s.uid
}

//js:get gid
func (s *stats) Gid() int {
	return s.gid
}

//js:get mode
func (s *stats) Mode() FileMode {
	return s.mode
}
