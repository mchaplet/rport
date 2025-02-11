// Code generated for package api_token by go-bindata DO NOT EDIT. (@generated)
// sources:
// 001_init.down.sql
// 001_init.up.sql
// 002_plural_and_name.down.sql
// 002_plural_and_name.up.sql
package api_token

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\x2c\xc8\x8c\x2f\xc9\xcf\x4e\xcd\xb3\x06\x04\x00\x00\xff\xff\x8d\x12\x02\x7d\x15\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 21, mode: os.FileMode(420), modTime: time.Unix(1677581889, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x6a\xc3\x30\x10\x45\xf7\x39\xc5\xef\x2a\x31\xf8\x06\xa5\x0b\x55\x9e\x12\x11\xd9\x0e\xea\x88\x34\x2b\x23\xd2\x29\x98\xd2\x44\xc8\x2e\xe4\xf8\xa5\x91\xa9\x5d\xc8\x52\x3c\xbd\xe1\x3f\xed\x48\x31\x81\xd5\xb3\x25\x84\xd8\x77\xe3\xe5\x53\xce\xd8\xac\x00\xe0\x7b\x90\x74\x0e\x5f\x02\xa6\x37\x46\xd3\x32\x1a\x6f\x2d\xf4\x96\xf4\x0e\x9b\x3f\xfa\xf0\x84\xf5\xba\x28\x6f\x4a\x4c\xf2\xd1\x5f\xef\x0b\x13\x5b\x7e\x3f\x25\x09\xa3\xbc\x77\x61\x44\xa5\x98\xd8\xd4\x34\x6b\x15\xbd\x28\x6f\x19\xda\x3b\x47\x0d\x77\xbf\xf4\x95\x55\xbd\x2f\x71\x93\xe5\x1a\xfb\x24\xc3\x52\xce\x57\x87\xd3\x25\xe6\xd1\xf9\x9d\x9b\xfe\x6d\xca\x60\xef\x4c\xad\xdc\x11\x3b\x3a\xce\x3d\xe5\x14\x51\xac\x0a\x1c\x0c\x6f\x5b\xcf\x70\xed\xc1\x54\x8f\x3f\x01\x00\x00\xff\xff\x2c\xf2\x75\x27\x2d\x01\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 301, mode: os.FileMode(420), modTime: time.Unix(1677581757, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_plural_and_nameDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\x48\x2c\xc8\x8c\x2f\xc9\xcf\x4e\xcd\x2b\x8e\x2f\xcd\xcb\x2c\x2c\x4d\x8d\xcf\x4b\xcc\x4d\xb5\xe6\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x45\x52\xa3\x90\x52\x94\x5f\xa0\xe0\xec\xef\x13\xea\xeb\xa7\x80\x57\x61\x90\xab\x9f\xa3\xaf\xab\x42\x88\x3f\x42\xd0\x9a\x0b\x10\x00\x00\xff\xff\x93\xa1\xea\x76\x78\x00\x00\x00")

func _002_plural_and_nameDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_plural_and_nameDownSql,
		"002_plural_and_name.down.sql",
	)
}

func _002_plural_and_nameDownSql() (*asset, error) {
	bytes, err := _002_plural_and_nameDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_plural_and_name.down.sql", size: 120, mode: os.FileMode(420), modTime: time.Unix(1677582037, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __002_plural_and_nameUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\xbb\x0a\x02\x31\x10\x45\xfb\x7c\xc5\x2d\x15\xfc\x83\x54\x51\xa7\x58\x88\x13\x0c\x13\xd8\x2e\x04\x76\x8a\x45\x8c\x8f\x98\xff\x97\xb5\x71\x8b\x6d\xef\x3d\xe7\x38\x2f\x14\x21\xee\xe8\x09\xe5\x39\xe7\xcf\xe3\xa6\x15\x91\xd8\x5d\x08\x12\xfe\x5b\xb3\x66\x93\x6d\x28\xd3\x84\x5a\xee\x0a\xa1\x51\xc0\x41\xc0\xc9\x7b\x6b\x4e\x91\x9c\x10\x12\x0f\xd7\x44\x18\xf8\x4c\xe3\x4a\xcb\xbd\xce\xaf\xae\x79\x31\x0d\x00\x04\x5e\x47\x77\xbd\xe9\x7b\xf9\x0e\xbf\xf6\xde\x7e\x03\x00\x00\xff\xff\xcc\x54\x2b\x61\xa9\x00\x00\x00")

func _002_plural_and_nameUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__002_plural_and_nameUpSql,
		"002_plural_and_name.up.sql",
	)
}

func _002_plural_and_nameUpSql() (*asset, error) {
	bytes, err := _002_plural_and_nameUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "002_plural_and_name.up.sql", size: 169, mode: os.FileMode(420), modTime: time.Unix(1677581822, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"001_init.down.sql":            _001_initDownSql,
	"001_init.up.sql":              _001_initUpSql,
	"002_plural_and_name.down.sql": _002_plural_and_nameDownSql,
	"002_plural_and_name.up.sql":   _002_plural_and_nameUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"001_init.down.sql":            &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":              &bintree{_001_initUpSql, map[string]*bintree{}},
	"002_plural_and_name.down.sql": &bintree{_002_plural_and_nameDownSql, map[string]*bintree{}},
	"002_plural_and_name.up.sql":   &bintree{_002_plural_and_nameUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
