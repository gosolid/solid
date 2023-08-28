//js:package fs

package fs

var _ DirEntry = &dirEntry{}
var _ direntryv8 = &dirEntry{}

type dirEntry struct {
	isFile      bool
	isDirectory bool
	name        string
}

//js:method
func (e *dirEntry) IsDirectory() bool {
	return e.isDirectory
}

//js:method
func (e *dirEntry) IsFile() bool {
	return e.isFile
}

//js:get name
func (e *dirEntry) Name() string {
	return e.name
}
