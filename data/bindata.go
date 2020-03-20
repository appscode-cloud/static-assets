// Package data Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// clouds.json
// customers.json
// press.json
// testimonials.json
package data

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

var _cloudsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd3\xb1\x6a\xc3\x30\x10\x06\xe0\x3d\x4f\x71\x78\x0e\xf1\x9e\x2d\x6d\xa1\x4b\x4b\x0a\x19\x3a\x2b\xe7\xab\x6a\x2c\xeb\x84\xa4\x34\x34\x25\xef\x5e\x24\x05\xa7\x95\x34\x6a\x0b\xa7\xcb\xff\x7f\x18\xee\x67\x05\xd0\x89\xb3\xeb\xb6\x10\x7e\x02\x74\x13\x7d\x77\xdb\x34\x5b\xa7\x89\x62\xc9\xcb\x3b\x40\xe7\x2c\x86\x8d\x4f\xef\x8d\xdb\xf6\x3d\x0e\x7a\x23\x8c\x71\xc8\x03\x6d\x90\xe7\x7e\x9c\x85\x24\xd7\xa3\xe2\xd3\xe0\x7a\x71\x76\x1b\xa3\xe5\x2d\x2c\xb4\x29\x1f\xfe\xbe\x9b\xc5\x85\x35\xbc\xd3\x11\x0e\x64\xbf\x46\x24\xd7\xc5\x95\xeb\x0a\xe0\xba\x8e\xae\xcb\xc9\x52\x29\x8b\xd3\x36\xb6\x10\x55\xd5\xbd\x8e\x68\xd9\xf1\x87\x87\x5d\xac\xcb\x64\x12\x4b\x57\x98\x35\x51\x49\x34\x55\xd3\x33\xb3\x54\x04\x8f\x61\x2b\x07\x0d\xa3\x1c\xbd\x50\x8c\x24\x74\x21\xfb\xf7\xd8\x84\xf8\x37\xb1\x6a\x7d\x4a\x0b\xfb\x58\x99\x59\xd5\xa8\x79\x28\xbf\xdf\x6d\xdc\xc4\x97\xb2\xaa\xb2\x97\x54\x93\x99\xd8\x90\x76\x5e\xe0\x54\xb0\xee\x2f\x4d\x64\x4b\x5c\x15\xb7\x37\xa4\x0f\xb1\x2c\xf3\x19\x81\x13\xf9\x02\x77\x1b\x37\x91\xa5\xac\x2a\xeb\x2d\xd5\x64\xa6\xa3\xb0\x34\x93\x17\xaa\x60\xdd\x5f\x9a\xc8\x96\xb8\x2a\xee\x61\x29\xcb\x6f\x82\x71\x22\x5b\x5e\x43\x1a\xb7\xb9\x83\x98\x55\xbf\x80\x54\x73\x37\xad\xae\xbf\x01\x00\x00\xff\xff\x80\x8b\x61\x4c\x6d\x05\x00\x00")

func cloudsJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudsJson,
		"clouds.json",
	)
}

func cloudsJson() (*asset, error) {
	bytes, err := cloudsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "clouds.json", size: 1389, mode: os.FileMode(420), modTime: time.Unix(1567469222, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _customersJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xc1\xea\xa3\x30\x10\x87\xef\x7d\x8a\xe0\x79\x6b\xee\xde\x5a\xba\xec\x16\x3c\x14\xf7\xb6\x97\x25\xea\x54\x03\x31\x09\xc9\xa8\x74\x97\xbe\xfb\x92\xd4\x50\x41\x6b\xfd\xd7\x93\x90\xdf\xcc\xe7\x37\x03\xf3\x6f\x47\x48\xc4\x25\x82\xb1\x37\x8b\xd0\xd8\x28\x21\xee\x8d\x90\x48\xb2\x06\xa2\x84\x44\x67\x97\xfe\x1a\xd2\x6f\x8f\xac\x87\xdc\x72\xf4\x71\x8d\xa8\x6d\x42\x69\xdf\xf7\xf1\x18\x14\x17\xaa\x09\xe5\x42\x55\x6a\x5c\x5b\x94\x32\x66\x5a\xdb\x42\x95\xe0\xea\x28\x6f\x58\x05\x96\x16\xad\x45\xd5\x80\xb1\x74\x4c\xa2\xae\x3d\xd6\xb2\x0a\x38\x5b\xab\x3e\x4a\x08\x9a\x16\x76\x84\xdc\xdd\x6b\x04\x42\xd7\x60\x38\x9b\xfa\x7f\x0f\xc9\xb2\x7b\x00\x6c\xf3\x0e\x94\x55\xce\x7f\x21\x37\x33\xc2\xbf\xfd\xf3\xb2\xad\x6f\xdd\xa6\xea\x11\xab\x3c\xb5\xe1\x05\xd4\x6d\x9e\x0b\x98\xda\x5e\x5c\xf8\xd3\x87\xe4\xf0\x63\x41\x7b\x44\xd9\x26\x3e\x02\x3d\xf4\x6d\xb7\xa8\x6f\x5a\x59\xf5\xec\x36\x55\xcf\x86\xe0\xb5\xf3\xd0\xba\xcd\x77\x80\xac\x5a\x75\x09\x9d\x12\x2d\x72\x25\x67\x2e\xf1\xf4\x0c\x63\x09\xb8\xe0\x5d\xce\x57\x7e\xdd\x7d\x04\xa2\xa3\xdf\x1f\x45\x0b\xfb\x34\x9b\x99\xe6\xca\x84\x7d\x8e\x23\x78\x07\x86\x35\x7a\x3a\x4b\xca\x3b\xc8\x5c\xf2\x7a\x88\xd0\xbc\x6d\xfb\x81\xb2\x6a\xfd\x39\x37\x25\x9f\xca\x1e\xcf\xd9\xe9\x4c\x2e\x78\x23\x29\x96\x6f\x2e\xd3\x23\x9c\x4a\xcc\xda\xcf\xad\x3d\x85\xfa\xff\xee\x53\x55\xa9\x3f\x19\x20\x97\xec\x9d\x3f\xb3\xc8\x0c\xce\x5c\xe9\x61\x08\xe6\xdd\x13\x4a\x87\xce\xbd\x16\x0c\xaf\xca\x34\xb1\x32\xd5\xe7\xfa\x03\x2d\x7c\x17\xb4\x77\xf7\xff\x01\x00\x00\xff\xff\xd8\x50\x11\x81\x83\x06\x00\x00")

func customersJsonBytes() ([]byte, error) {
	return bindataRead(
		_customersJson,
		"customers.json",
	)
}

func customersJson() (*asset, error) {
	bytes, err := customersJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "customers.json", size: 1667, mode: os.FileMode(420), modTime: time.Unix(1571828234, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pressJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x4f\x6f\x1b\x47\x0c\xc5\xef\xfe\x14\x84\xcf\xfa\xd3\xd8\x69\x8d\x06\x28\x8a\x56\x4e\x90\x20\x85\x5b\xb8\x2e\x8c\xa2\xe8\x81\x9a\xa1\xb4\x53\xcd\x0c\x07\x24\x57\x5b\xa1\xc8\x77\x2f\xb8\x6b\x39\x42\x61\x07\x39\xa4\x47\x69\x67\xc8\x37\xef\xfd\xc8\x3f\xce\x00\xfe\x39\x03\x00\x38\xb7\x64\x99\xce\x5f\xc1\xf9\x25\x70\xa3\x0a\xca\xbd\x04\x82\x26\xfc\x17\x05\x53\xb0\x0e\x0d\x0a\xee\x08\xde\xf7\x6b\x92\x4a\x46\x0a\x84\x9a\x48\xce\x67\x53\x09\xed\x4b\x41\x39\x78\x91\x37\xc2\x05\x42\xee\xd5\x48\x40\x0d\x8d\xa0\x60\xc5\x2d\x15\xaa\x06\xc6\xa0\x15\x9b\x76\x6c\x0a\x58\x23\x5c\xdf\xce\x20\x70\x69\x58\x13\x57\x30\xe6\xac\xb0\xf1\x12\x6f\xa9\x59\xe2\xd9\xd8\x33\xce\xc6\xb3\xef\xfb\x75\x0a\x2c\x15\x30\x15\xaf\xb4\x49\x39\x83\x75\x04\x5b\x6c\x0a\xa9\x9e\xe8\x5b\x1c\x95\xe5\x54\x77\x2e\xab\x33\x6b\xfa\x6a\xb9\x1c\x86\x61\x91\xea\x86\x07\x96\x1c\x17\x81\xcb\x12\xc5\x52\xc8\xb4\xbc\xbc\x78\xf1\xf2\x9b\x8b\xab\xe5\xe5\xdc\x5d\x98\x4f\x2e\xcc\x8f\x2e\xcc\xdd\x85\xb9\xbb\x30\xdf\x3d\x76\x99\x4f\x2e\x2c\x3a\x2b\xf9\xd8\xd0\x52\xa1\x3b\xbe\x25\x8c\xde\xf6\x6b\x28\xa9\xaa\x8b\x1d\xff\x79\x38\x93\x0a\x6e\xdd\xf1\x29\x00\xf7\x4f\xc2\xa9\xc8\x10\xeb\x02\x5b\xd3\xc0\x91\x46\x8d\xe3\x05\x5d\x36\x21\xd5\xe5\x8b\x45\xab\xdb\xf3\xf1\xea\x87\x33\x80\x0f\xb3\x27\xa2\xbc\xed\x6b\x4d\x75\x0b\xbf\x73\x2f\x70\x8d\x86\x6b\x54\x82\x9f\x4f\x2d\x82\xfb\x64\xdd\xf8\xfb\xfa\xc7\x27\x72\x3c\x39\x98\x14\x10\xa2\xa4\xbd\x57\xdc\xb0\xc3\x91\xea\x68\xbc\x50\xc5\xa4\x8a\x35\x10\xa0\x70\x5f\x23\x44\x6a\x99\x0f\x7e\xd2\x23\x93\x07\x1d\xd8\x5a\x4e\x01\x2d\x71\xd5\x05\xbc\xe5\x81\xf6\x24\xb3\x89\x0c\xff\xee\xc5\xe2\x51\x66\xc6\x03\x89\x77\x55\xf3\x84\x11\x94\x1a\x8a\x93\x14\xb8\x06\x92\xba\x80\xbb\x8e\x1e\xa4\x1f\x39\x85\x01\x15\x82\x10\x1a\x45\x40\x57\x3c\xe0\x01\x78\xe3\xdf\xf7\x29\x8e\x22\x40\x53\x69\x99\xa0\x50\xe8\xb0\x26\x2d\xfe\x9a\x47\x8d\x07\xf7\x4a\x8d\x05\xb7\x04\x7a\x50\xa3\x72\x7c\xa7\x62\x21\x68\x19\x6d\xc3\x52\xbc\xfa\x78\xf6\xe4\x51\x9f\xe4\xcd\x1f\x46\x75\x9b\x2a\x91\xa4\xba\x6d\x1c\x03\xaa\x8d\xc1\x3a\x4b\x71\x7d\x8a\xd4\xd1\x85\x39\xb5\xa4\x1c\x69\x7e\xf5\xed\xf2\x19\xb6\xbe\x7a\x84\xeb\xa7\xa4\x46\xf5\xcb\xe0\x75\xf1\x19\x78\xad\x6e\x56\x6f\xe0\x75\xd5\x5e\x48\x4f\x91\x7a\x57\x8d\xc4\xe7\x47\x70\x9d\x72\xb2\x03\x0c\x0e\x19\xc2\x8a\xc4\xe0\x17\xe1\xad\x60\x79\x02\x36\x8f\x73\x95\xb9\x8f\x70\x83\x96\xf6\x04\x2b\x2e\xad\x37\x4f\xe5\x8d\x53\x35\x7a\x0c\x1d\xfa\x94\x9b\x70\xec\x03\xc5\x31\x98\x4a\x03\x04\x12\x4b\x9b\x87\x20\x3c\x6e\x6f\xe2\xae\x74\x94\x1b\xd0\xa8\x12\xc2\x58\x7d\x62\x81\x64\xda\x3d\x6c\x9d\xef\x28\x92\x7d\x0a\xa4\xb0\xa7\x1a\x59\x14\x84\x0a\xa6\x3a\x6d\x25\x4b\xeb\x4c\xd3\x23\xb8\x12\x60\x9d\xee\xa0\x8e\xdd\x1b\xb7\x3e\xa3\x38\x96\x86\xa9\x92\x00\x4b\xe8\x48\x4d\x1e\xb5\x8c\x68\x1a\xee\x48\xa1\xe3\x1c\x1d\x48\xbf\x99\x6a\xec\xd5\xe4\xf0\x2c\x37\xd6\x51\xa5\x41\x0d\xc3\x6e\x91\x78\x19\x6a\xd8\xcc\x1f\xdf\xae\xf3\xf0\x1c\x13\xff\xc3\xbe\xb9\xfc\x0c\x20\x4e\x18\x38\x99\x0a\x68\x18\x76\xd3\x80\x8f\x3b\x7d\x93\x44\x0d\x9a\xa4\x1a\x52\xcb\xa4\x4f\x90\xf0\x96\x72\xf1\xd1\x77\x93\x0a\xfb\xe9\x07\x8f\x7d\x9a\x7d\x5a\x3f\x96\x3c\xdd\x29\xe3\xa7\x93\xf5\x0f\xef\x7c\x70\xbd\x0e\xe6\xdd\x0c\xee\xb0\x60\x86\x5f\xb1\xc3\xd9\x04\x94\x47\xb5\x81\x1f\x5a\xd3\x15\x47\x02\xfa\xbb\x65\x96\x31\xa3\x01\x46\x09\x17\x10\xb0\xc2\x9a\xa6\x95\x91\x36\x89\xe2\x27\x67\xfc\xc0\xbd\xf5\xeb\xc9\xbd\x01\x2d\x74\xdf\xef\xbf\xbb\xba\xb9\xfe\xed\xea\xfd\xee\x9e\xf3\xeb\xa7\xd3\x7a\xf1\x71\x82\xef\xfd\xce\x97\xc9\xeb\xe5\x7f\xf2\x3a\xfb\xf3\xec\xdf\x00\x00\x00\xff\xff\x47\x51\x88\xb1\xee\x07\x00\x00")

func pressJsonBytes() ([]byte, error) {
	return bindataRead(
		_pressJson,
		"press.json",
	)
}

func pressJson() (*asset, error) {
	bytes, err := pressJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "press.json", size: 2030, mode: os.FileMode(420), modTime: time.Unix(1582350867, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testimonialsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xc1\x6e\xdc\xc8\x11\xbd\xfb\x2b\x2a\x3a\xc4\x09\x30\xa2\x2e\x01\x02\x2c\x36\x06\x04\xcb\x6b\x1b\xd6\xae\x05\x2b\x88\x0f\x71\x20\x14\xbb\x6b\xc8\xb6\x9a\x5d\x8d\xaa\xe6\x70\xa9\x20\xff\x93\x4b\x4e\xf9\x04\xff\x58\x50\xcd\xa1\x34\xd0\x7a\x2e\x0a\x74\x31\xc6\x6c\xb2\xfa\x55\xbd\x57\xaf\x4a\x7f\x7f\x01\xf0\xcf\x17\x00\x00\x27\x09\x07\x3a\xf9\x01\x4e\x2e\x47\x87\xf0\x09\x77\x78\x77\xc7\x91\x4f\x36\xcb\xe9\x57\x6e\xed\xf0\x4a\xd8\x8f\xae\xc0\xcf\x98\xb0\x23\x59\x4f\x69\xc0\x10\xed\x7c\x7d\x90\x85\xb7\x21\xd6\x78\x7d\x29\x59\x7f\x38\x3b\x73\x3e\x35\x98\xb3\x3a\xf6\xd4\x38\x1e\xce\xc2\x80\x1d\xe9\x99\x1b\xb5\xf0\x40\xa2\x67\x99\x44\x39\xe9\x59\x1c\x1d\xde\xc8\x8a\xa0\xc9\xa9\x5b\xc3\x3a\x1e\x32\xa6\xf9\x03\xcd\x16\x39\xa4\x42\xa2\xb3\x16\x1a\xf4\xe0\x62\x03\xb8\x7f\xc3\x71\xd2\x31\x96\xf0\x10\xa1\x90\x96\x30\x70\x0a\x58\xf1\xbe\xb7\x10\xd7\x4b\x08\x98\x50\xc1\x53\x0c\x5d\x5f\xc8\x43\x61\xa0\xd4\x61\x47\x30\x85\xd2\xc3\x79\xce\xfa\x9a\x3d\x41\x48\x50\x7a\xaa\x2f\x3a\x2c\xb4\x81\x99\x0a\x6c\xc7\xe4\x71\xa0\x54\x30\x42\x41\xbd\x05\xde\x82\x8e\x39\xb3\xd8\xe5\xe0\x47\xc1\x36\xd2\x06\x12\xa7\x53\xca\x3d\x0d\x24\x18\x61\x62\xb9\x8d\x8c\x5e\x97\x2b\x3e\x8c\x2d\x49\xa2\x42\xda\xc0\x67\x82\x44\xe4\x0d\x47\x4f\xd0\x92\x96\xd3\x2c\x94\x51\xc8\x6f\x60\xe0\xfa\x5f\xde\x06\x17\x28\x15\xf0\x58\xb0\x45\x25\xe0\x4c\x82\x85\x05\x1e\x12\x5f\xf1\x86\xe4\x47\x2d\x32\x37\xf0\x36\xec\x28\xdd\xe7\xf3\x52\x21\x93\x0f\x9d\x10\x19\xe8\xfb\x50\xed\x18\xa2\xb7\xef\xd7\x98\xba\xd9\xe7\xed\x82\x06\x4e\xb5\x5a\x84\x3a\x37\xf0\x0b\x43\x09\x03\xd5\x27\x13\xaa\x15\x0f\x93\x07\x8c\x11\xb8\xfd\x4a\xae\x84\x1d\x29\x08\xa1\xeb\xc9\x1b\x1e\x4c\x80\x03\xde\x85\xd4\xc5\x19\xb4\x67\x29\x90\x49\x02\x7b\x43\x60\xa1\x1a\x78\x0f\x13\x8f\xd1\x83\x90\xe3\x61\xa0\xe4\x1f\x08\x38\xc8\x6d\xcb\x02\x98\xe6\x83\xca\x81\x50\x44\x43\x60\xb5\x6d\x4e\x5e\x00\xfc\x6b\xf3\x5b\x8d\xff\x8c\x12\x18\x3e\xf0\xdd\x57\xbc\x7d\xa4\xf0\x6b\x8e\x63\x09\x9c\x14\xce\xc5\xf5\xa1\x90\x2b\xcf\xa3\xf2\xc1\x30\xdc\xdc\x56\x0c\xc7\x35\x4e\x31\xf7\x24\x01\xbf\xaf\xef\x1d\xcf\x87\x6d\xf8\x48\xdc\x7f\x5b\x4e\x61\x40\x93\x6d\x01\x0d\x43\x8e\x54\xb9\xa1\xed\xaa\x1d\x2b\xe1\xa8\xa6\xf6\x2c\x6c\xd9\xd6\xf3\x90\x42\x09\x58\x08\x78\x14\x68\x51\x08\x06\x32\x6d\x1f\x14\x7a\x15\x6f\x03\xef\x8b\xc2\x98\x3c\x49\x9c\x8d\x94\x42\xae\x4f\x1c\xb9\x9b\x97\xab\x7e\x2d\x94\x34\xec\x08\x2e\xff\xb4\xf6\x04\x60\xe4\xd4\x2d\xaa\x57\xc2\x21\x92\x2a\x5c\x5f\x5f\x82\xb5\x74\x27\x68\x04\x40\x50\x98\x7a\x2c\x0b\xfc\x51\xc1\xf5\xcc\x4a\xb0\x66\xc5\x3b\xfb\xa7\xf4\x24\xda\xdc\x3f\x2c\x84\x83\x7d\x88\x51\x19\x76\x24\x33\x08\x69\xe6\xe5\xfe\xa9\xa7\x64\x75\x70\x3c\x50\xcd\x78\x8f\xa6\x81\xb7\x42\x58\x60\x5f\xdb\xdf\x1d\x57\x4d\x1a\x29\xc2\x27\x1c\xbe\xfd\x47\xe8\x0e\x2e\xbf\xfd\x37\xd3\xdd\x23\xfd\x5c\xd0\xee\x63\x56\x78\x93\xba\x90\xe8\xb9\x1c\x72\xa8\x48\x6e\x64\x8f\xe4\x26\x56\x24\xc7\x55\x94\x25\x38\xea\xc7\xb6\x8d\xf4\xff\x08\x29\x68\xb5\x00\x42\x0d\xa4\x05\x26\x9c\xad\x8a\xa3\x52\x7d\xbc\x45\x5d\xc4\x23\x14\x83\xd9\x1d\xbc\x3b\xbf\x12\xfe\x75\x06\xd4\xaa\xa3\x90\x3a\x31\xa2\x1d\xa7\x22\x1c\x23\x49\x03\xe7\x05\xae\x0c\xdc\xbb\x0a\x6e\x63\xf4\xec\x6f\x71\x2c\x89\xe4\x54\x0b\xa7\xea\x4c\x55\x89\x71\xa4\x33\xb3\xaa\x04\x9e\x72\xe4\xd9\x2c\x57\x8f\xb6\xf9\xa7\xe0\x7a\x12\xb8\x44\x09\xbb\xf0\xed\xdf\x42\xdf\xe1\xea\xea\x1a\xae\x33\xb9\x80\x31\xe8\x33\x75\xba\x54\x18\x37\xf1\x1e\xc6\x71\xa2\x3c\xed\x56\xff\x79\x0a\x51\xef\x6b\x57\x2e\x7d\x85\xb0\xa5\xe9\xb0\x61\x5d\x1c\xd5\x06\x66\xe5\x68\xa2\xca\xdb\xca\xec\x9e\xa1\x2c\xb4\x25\x91\xea\xd2\xbf\xe5\xea\x33\x99\x89\xc7\x38\x43\x0c\xb7\xb4\x4a\xa1\x92\xe3\x38\x6d\x43\x37\x2e\x9d\xdb\xc0\x05\xbb\xb1\x4e\xc3\xb5\x91\xb3\x50\x29\x33\x74\xcc\xbe\x81\x73\x6b\x4e\xfb\x7a\x5c\x3e\x7e\x87\x8b\x4c\x82\x42\x18\xac\x21\xf1\xc1\x97\x5a\x72\x68\xaf\x85\x52\x53\xd3\x15\xc1\x44\x31\x2e\x79\xb6\x5c\x7a\xb3\x16\xcb\xea\xf2\xcf\x60\x96\x04\x2d\x46\x4c\x2e\xa4\xae\x81\x8f\x0f\xea\xf9\xeb\xeb\x2b\x50\x92\x5d\x70\xa4\x1b\xf8\x8c\xf3\x2d\xfc\xc2\xd3\xa6\xfa\x85\xc9\xb5\x70\x8d\xa8\xc5\x42\x95\x9e\x47\xc5\xe4\xd5\xbe\x36\x26\x83\x16\xf3\x4b\xc7\x29\xd9\x48\xb3\x09\x51\x0d\x46\x07\xe6\xd2\xc7\x19\xb0\xd4\xa4\x14\x07\x5a\x86\xd8\x31\x59\x5e\x60\x0a\x14\xe1\x2d\xcb\x10\xe7\xc7\xe3\x87\x52\x60\x81\x0b\xda\x51\xb4\xc9\xfb\x3c\x8a\xf4\x15\xc1\x4d\x57\x11\x1c\x57\x63\x1b\xc4\x87\xa7\xe8\xf0\x91\x52\x46\xb5\xb9\xb0\xd7\xda\x32\x2f\xb4\x08\xda\xa2\xb5\x65\x99\x50\xfc\x5e\x93\x31\x9e\xfa\xbd\x74\xc8\xef\x45\x55\x8f\x6c\x36\xfc\x81\x74\x69\xd4\x38\xc3\x25\x95\x97\x66\xb2\x4e\xe6\x5c\xfe\x08\x3d\xea\x32\x26\x8c\xe6\x21\xac\x13\xc4\x96\xb0\x3d\xdf\x46\xee\x41\x2f\x20\xb4\x42\x74\x47\x0d\xbc\x41\xd7\xc3\x80\x5f\x59\x8c\xce\xba\xd9\x58\xb4\xd6\x3c\x06\x17\x86\x27\x8a\x36\x32\x60\xcc\x1e\x0b\x1d\x1d\x0f\x17\x75\xa9\xf8\x49\xc8\x3f\x14\x6d\xcf\xea\x4f\x5c\xe7\x23\xfc\x1e\x5e\xbf\xf9\xf8\x5c\x94\xda\x3a\xb1\xad\xb7\x1f\x67\x14\xb5\xa0\x94\x27\x0d\x81\xcf\xf4\x52\x1e\x51\x69\xb6\x91\x51\x8a\x15\xda\x76\x52\xf8\x11\xa1\x17\xda\xfe\xe5\xcb\x7d\x02\x5d\x28\xfd\xd8\x56\xe8\xfb\xbb\x4f\x73\x44\xa3\xfd\xfe\xc1\x97\x93\x57\xe7\xcb\xaf\x43\x53\xff\xf1\x0c\x5f\x2d\xab\x88\x31\xcd\xe2\x7a\x32\xcd\xd4\xe5\xd2\xfc\xa3\x9d\x97\x89\x20\xf6\x4e\xa2\xe9\x69\x77\x9f\xde\xde\x6b\xe2\x74\x5d\x74\xbf\x9c\xbc\xfa\xb8\xff\x69\x20\xaa\xef\xd9\x06\xd4\x45\xac\x7f\x15\x68\xcf\x93\xc3\x3a\xf4\xb0\x2c\x4d\xfe\xe2\x1f\xff\x0b\x00\x00\xff\xff\x22\x68\x97\xfc\x49\x0d\x00\x00")

func testimonialsJsonBytes() ([]byte, error) {
	return bindataRead(
		_testimonialsJson,
		"testimonials.json",
	)
}

func testimonialsJson() (*asset, error) {
	bytes, err := testimonialsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "testimonials.json", size: 3401, mode: os.FileMode(420), modTime: time.Unix(1582350867, 0)}
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
	"clouds.json":       cloudsJson,
	"customers.json":    customersJson,
	"press.json":        pressJson,
	"testimonials.json": testimonialsJson,
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
	"clouds.json":       {cloudsJson, map[string]*bintree{}},
	"customers.json":    {customersJson, map[string]*bintree{}},
	"press.json":        {pressJson, map[string]*bintree{}},
	"testimonials.json": {testimonialsJson, map[string]*bintree{}},
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
