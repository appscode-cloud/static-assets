// Package hugo Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// config.json
// params.dev.json
// params.json
package hugo

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _configJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x92\xbd\x6e\xe3\x30\x10\x84\x7b\x3d\x05\xc1\xda\x77\xba\x4b\xe9\x67\x70\x9a\x00\xa9\x02\x17\x2b\x6a\x25\x11\x12\x7f\x40\x2e\x6d\x24\x81\xdf\x3d\x58\xfe\x48\x4e\x65\xcf\xee\xcc\xf2\xf3\xc0\xdf\x9d\x10\x12\x62\x44\x8a\xf2\x2c\x58\x09\x21\x03\x7a\xf7\xfe\x76\x91\x67\x21\x17\x22\x1f\xcf\x7d\x3f\x6b\x5a\xd2\xf0\x57\x39\xd3\x83\xf7\x51\xb9\x11\xfb\x48\x40\x5a\xfd\xa9\xe9\x53\xc9\xde\x30\x44\xed\x2c\x67\x0d\x44\xc2\xd0\x16\xa3\x0e\xc7\x13\xac\x81\x80\x5d\xf9\xf3\xd4\xa6\xda\xc0\x8c\xec\x93\xe5\x7a\x5f\xae\xf7\x75\xbe\xfb\x26\xbd\xfd\xb2\x15\x9d\xb7\x8f\x4e\x88\x07\x1b\xa5\xda\x5c\x1a\xd9\xf5\x51\x10\xe0\xbe\x63\xc2\x57\x0a\xd8\xc4\xac\xf0\xa0\x9c\x35\xc1\xe6\x14\x82\x6d\xb3\x4d\x5b\x37\xee\x0e\xe7\xd1\x46\x02\xb5\xb6\x81\x07\xb5\x22\x35\x35\x40\x40\x83\x04\xdb\x7e\xd1\xa9\x15\x03\xa3\x5d\x33\x94\x01\x6d\x5f\xd1\xa6\xa3\x6e\xf5\x8f\x11\xe5\x9a\x06\x1c\x07\x79\xca\xbf\x29\x2e\xfc\x85\x47\x37\x48\x1b\xb1\x98\x13\x84\x51\x5e\xeb\x59\xf5\x3f\x87\x6e\xee\x13\x66\x2e\x59\xc8\x78\xd7\x53\x36\x46\x84\xa0\x96\x4d\xcf\x0b\xb5\x23\x4f\xb9\x97\x9c\xf3\x0b\x04\x53\x72\xbc\x9f\x5c\x30\xf2\xba\x17\x37\x39\x47\x18\x2e\xda\xae\x4f\xed\x35\xbe\xa2\x2a\xe3\xb1\xaa\x9c\xf5\x5f\xd0\xb0\x6a\xc1\x19\xbd\xf5\x55\x9f\xe6\x46\xba\x47\xf7\x13\x00\x00\xff\xff\x84\x98\x6a\xcb\x82\x02\x00\x00")

func configJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJson,
		"config.json",
	)
}

func configJson() (*asset, error) {
	bytes, err := configJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.json", size: 642, mode: os.FileMode(420), modTime: time.Unix(1572861521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _paramsDevJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd2\x41\x6f\x82\x30\x18\x06\xe0\xbb\xbf\xc2\x70\xda\xb2\x51\xa7\x73\xc6\x78\xda\xb2\xcb\x76\xdc\x2f\x30\x6d\xf9\x84\xc6\xc2\xd7\xb4\x5f\x21\x66\xd9\x7f\x5f\xa4\x42\xb0\x12\xbc\xbe\x7d\x9f\xbc\x90\xf6\x77\x36\x9f\x27\x19\x96\x5c\x55\x7b\x6e\x8c\x93\x98\x41\xb2\x9b\x27\x05\x91\x71\xbb\xc5\xa2\xcb\xd2\xfa\x35\x2d\x7c\x8e\xac\x01\xc1\xb8\x31\xc9\xf3\x00\x1e\xbd\x80\x4c\x0c\x59\x48\xd2\x7a\x35\x8d\x6a\xee\x35\xc5\xae\x0d\xd3\x7a\x39\x4d\x0f\x68\xcb\x58\x9e\xb3\x49\x68\x0a\x6e\x4b\xb0\x43\x77\x89\xee\xee\x49\x15\xaf\x49\x75\x17\xb9\x42\x81\xce\x62\x18\xd2\x49\x2c\x4e\x04\xc2\x2b\x9d\x81\x75\x43\x7e\xce\x59\x7f\x30\x14\x25\xb7\x47\x20\xa3\xb9\xbc\xba\xbf\x41\xcc\x46\xb0\xd0\x98\xef\xbd\xd5\x57\x1b\x1a\xf3\xb1\xae\xd3\x5c\x1e\xe3\x72\x1b\xb2\xee\x95\x30\x89\x65\x28\xe7\x8a\x0a\x2f\xe2\x76\x48\xcf\xad\xfe\x61\x85\x3a\x35\x8a\x08\x6c\xdc\xbf\xc4\x2d\xf8\x30\xc6\x7d\x62\x06\x5f\x3f\x81\x1c\xb8\x04\x81\x78\xf3\x45\x4d\xd3\xb0\xee\x6c\x64\xe9\x84\x9e\xbc\x80\x31\x75\x39\x6a\x91\xec\xf7\xbe\x2b\x19\xa4\xc4\x8a\xb8\xa4\x3d\x94\x5c\xb5\xd6\x79\x63\xd0\xd2\xfb\xed\xdf\x77\x55\x53\x60\xd5\x5e\xc7\xd3\xf2\x61\xf3\xf6\xf2\xb8\x5a\x2f\xd3\xed\x7a\xbb\x49\x66\x7f\xff\x01\x00\x00\xff\xff\x78\xac\xdc\xe8\x79\x03\x00\x00")

func paramsDevJsonBytes() ([]byte, error) {
	return bindataRead(
		_paramsDevJson,
		"params.dev.json",
	)
}

func paramsDevJson() (*asset, error) {
	bytes, err := paramsDevJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "params.dev.json", size: 889, mode: os.FileMode(420), modTime: time.Unix(1572861521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _paramsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\xc1\x6e\x02\x21\x10\x40\xef\x7e\x85\xd9\x53\x9b\xb6\x58\x1b\x6b\x8c\xa7\x36\xbd\xb4\xc7\x7e\x81\x01\x76\x74\x89\xec\x0e\x81\xa1\x1b\xd3\xf4\xdf\x1b\x41\x37\x08\x5c\xdf\xbc\xc7\xa0\xec\xef\x6c\x3e\x6f\x5a\xec\xb9\x1a\x76\xdc\x18\x27\xb1\x85\x66\x3b\x6f\x3a\x22\xe3\xb6\x8b\xc5\x95\x31\x89\x7d\xf3\x98\xc8\x47\x2f\xa0\x15\xa9\x1a\x49\x55\xfc\xe1\x5e\x53\xee\x06\x58\xd5\xf7\x68\xfb\xdc\x3e\xb3\x42\x36\x1d\xb7\x3d\xd8\xd4\xbd\x20\xa6\xb0\x38\x56\xaa\xfc\x50\x26\x55\x61\xb9\x4e\x81\x6e\x73\x33\xd2\xe2\x02\xe2\x44\x20\xbc\xd2\x2d\x58\x97\x26\x67\xce\xa6\x41\x5a\xf4\xdc\x1e\x81\x8c\xe6\xf2\xe6\x7f\x4e\x30\xab\xc4\x42\xe3\x61\xe7\xad\xbe\xd9\xa1\xf1\x50\x73\x9d\xe6\xf2\x98\xcb\x01\xb2\xf2\x35\x0f\x8a\x3a\x2f\x72\x3b\xd2\xb3\x35\x7d\x00\x51\xa7\x51\x11\x81\xcd\xfd\x0b\x0e\xc1\xbb\x31\xee\x03\x5b\xf8\xfc\x8e\xc9\x9e\x4b\x10\x88\xc5\x8d\xc6\x71\x64\xd7\x59\x65\xd3\x09\x3d\x79\x01\xb5\xea\x32\x0a\x91\x9c\xf6\x7d\x0d\x32\x96\x12\x07\xe2\x92\x76\xd0\x73\x15\x5a\xe7\x8d\x41\x4b\x6f\xe5\xaf\xbf\xaa\xa6\xc3\x21\x3c\xc7\xc3\xf2\x6e\xfd\xfa\x7c\xff\xb2\x5a\x3e\x6d\x56\x9b\x75\x33\xfb\xfb\x0f\x00\x00\xff\xff\x9a\x2b\x9a\x44\x21\x03\x00\x00")

func paramsJsonBytes() ([]byte, error) {
	return bindataRead(
		_paramsJson,
		"params.json",
	)
}

func paramsJson() (*asset, error) {
	bytes, err := paramsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "params.json", size: 801, mode: os.FileMode(420), modTime: time.Unix(1572861521, 0)}
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
	"config.json":     configJson,
	"params.dev.json": paramsDevJson,
	"params.json":     paramsJson,
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
	"config.json":     &bintree{configJson, map[string]*bintree{}},
	"params.dev.json": &bintree{paramsDevJson, map[string]*bintree{}},
	"params.json":     &bintree{paramsJson, map[string]*bintree{}},
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
