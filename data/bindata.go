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

var _customersJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xcf\x8a\xab\x30\x14\x87\xf7\x7d\x8a\xe0\xfa\x36\xd9\xbb\x6b\xe9\xe5\xde\x82\x8b\xe2\x72\x36\x43\xd4\x53\x0d\xc4\x24\xe4\x8f\xd2\x19\xfa\xee\x43\x52\x43\x05\xad\xed\xd4\x95\x98\xdf\x39\x5f\xbe\x73\x20\xdf\x1b\x84\x12\xe0\xaa\x01\xcd\x68\x92\x22\xff\x8f\x50\x22\x68\x0b\x49\x8a\x92\xbf\x31\xf9\x73\x3b\xef\xa1\x30\xcc\x86\xa8\xb1\x56\x99\x94\x90\xbe\xef\x71\x04\xe0\x52\xb6\xb1\x94\xcb\x5a\x8e\xeb\xca\x4a\x60\xaa\x94\x29\x65\x05\xbe\x8e\xb0\x96\xd6\x60\x48\xe9\x8c\x95\x2d\x68\x43\x22\x85\xf8\x56\xac\x44\x1d\x51\xa6\x91\x7d\x92\x22\xab\x1d\x6c\x10\xba\xfa\xd3\xe4\x0b\x0a\x3d\x23\xfc\x11\x8e\x97\x6d\x43\xeb\x3a\xd5\x80\x78\xc9\x53\x69\x56\x42\xe3\x8a\x82\xc3\xd4\xf6\xe4\xc3\xff\x21\x44\xbb\x7f\x0b\xda\x23\xca\x3a\xf1\x11\xe8\xa6\x6f\xba\x45\x7d\xed\x44\xdd\xd3\xcb\x54\x3d\x1f\x82\xc7\xce\x43\xeb\x3a\xdf\x01\xf2\xd2\xaa\x2b\xe8\x24\x77\x96\x49\x61\xa6\xbe\x87\x7b\x88\x05\xd8\x05\xef\x6a\xbe\xf2\xf7\xee\x23\x10\x19\x5d\xbf\xe7\x0e\xb6\x59\x3e\x33\xcd\x99\x72\x73\x1f\x87\xb3\x0e\x34\x6d\xd5\x74\x96\x8c\x75\x90\xfb\xe4\xf1\x10\xb1\x79\xdd\xf6\x23\xe5\xa5\xf5\x17\x4c\x57\x6c\x2a\xbb\x3f\xe6\x87\x23\x3a\xd9\x0b\xca\x6c\xf5\xe4\x65\x06\x84\x57\xc1\xd4\xbd\x6f\x1d\x28\x24\xdc\xbb\xcd\x64\x2d\x3f\x73\xb0\x4c\xd0\x67\xfe\xd4\x58\xaa\xed\xcc\x2b\xdd\x0d\xc1\xbc\x7b\x4a\xc8\xd0\xb9\x55\x9c\xda\xb3\xd4\x2d\x96\xba\x7e\x5f\x7f\xa0\xc5\xef\x82\xf6\xe6\xfa\x13\x00\x00\xff\xff\x62\xde\x69\xe3\xbf\x05\x00\x00")

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

	info := bindataFileInfo{name: "customers.json", size: 1471, mode: os.FileMode(420), modTime: time.Unix(1567469222, 0)}
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

	info := bindataFileInfo{name: "press.json", size: 2030, mode: os.FileMode(420), modTime: time.Unix(1567469222, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testimonialsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xcd\x6e\x1c\x37\x12\xbe\xeb\x29\x6a\x75\x58\xef\x02\x52\x0f\x76\xe3\x20\x80\xe1\x18\x12\x2c\xff\xc1\x4a\x24\x58\x41\x7c\x88\x03\xa1\x9a\x5d\xdd\x4d\x89\xcd\x22\xaa\xaa\x67\xdc\x0a\xf2\x3e\xb9\xe4\x94\x47\xf0\x8b\x05\x64\xcf\x48\x03\xd9\xf2\x41\xd1\x65\x30\x68\xb2\x7e\xc8\xef\xab\xef\xe3\x2f\x3b\x00\xbf\xed\x00\x00\xec\x46\x1c\x68\xf7\x09\xec\xfe\x80\xe2\x19\xde\xf2\xd5\x05\x5e\xee\xee\xcd\x6b\x17\x5c\xe7\xa5\x33\x0e\xa3\x79\x8e\x0a\x87\xe2\x7a\x6f\xe4\x6c\xb3\x83\x06\xf4\x21\xef\x19\x72\x78\x75\x59\xc2\x0f\x28\xa4\x9e\xc4\x63\xe5\x78\xd8\xec\x4c\xc2\xad\x0f\xa5\x54\x6f\x96\xf4\xc9\x62\xe1\x9a\x58\x61\x4a\xea\xb8\xa1\xbc\x75\xe1\x07\xec\x48\x17\x6e\x54\xe3\x81\x44\x17\x89\x44\x39\xea\xa2\x64\x3f\x9f\xb3\x57\x17\xa9\xdb\x24\x75\x3c\x24\x8c\xd3\x5b\x9a\x72\xde\x4d\xd9\xad\x92\xcd\xe8\x6c\xbd\xba\xe4\x09\x3b\x92\xcd\xa2\x91\x9a\x1f\x38\x7a\x2c\xfd\xff\x3c\xaf\xc2\x80\x0d\x81\x37\x50\x3f\xa4\x40\x80\xb1\x01\x6a\x5b\xef\x3c\x45\x83\x96\x05\x46\x05\x63\x48\xc2\xf9\x1a\xca\xba\x8f\xde\x3c\x1a\x01\x8f\x02\x35\x0a\xc1\x40\x86\x01\xde\x8e\x35\x49\x24\x23\x85\x15\xcb\x65\x60\x6c\x2a\x78\x63\x0a\x63\x6c\x48\xc2\xe4\x63\x07\x46\xae\x8f\x1c\xb8\x9b\xe6\x52\x1f\x8d\xa2\xfa\x25\xc1\xf1\x63\xd0\x31\x25\x16\x03\x0c\x1c\x3b\x58\x79\xeb\x41\x09\x87\x40\xaa\x70\x76\x76\x0c\x3e\x1a\x75\x82\x19\x19\xf0\x0a\xab\x1e\x6d\x6e\x7f\x54\x70\x3d\xb3\x12\x6c\x4e\xc5\xcb\xfc\x63\x3d\x89\x56\xd7\x1f\x8d\x70\xc8\x81\x18\x94\x61\x49\x32\x81\x90\x26\x9e\xeb\xaf\x7a\x8a\xf9\x1e\x1c\x0f\x54\x4e\xbc\xee\xa6\x82\x57\x42\x68\xb0\xbe\xdb\x7f\xed\xee\x00\xfc\xbe\xf7\x25\x3a\xc5\x91\x02\xbc\xc3\xe1\xd3\x9f\x42\x57\x70\xfc\xe9\xaf\x44\x57\xb7\x88\x75\x44\xcb\x93\xa4\xf0\x22\x76\x3e\xd2\x0d\x34\x5b\xa4\xca\x49\x2a\xc1\xc1\x0b\x5d\x55\x81\x13\x5d\x1d\x24\xf1\x8e\xfa\xb1\xae\x03\x3d\x20\xbf\x72\xa1\x73\x59\x77\x7b\x1e\x4a\xb7\xd5\x45\xa2\x3b\xa8\xb6\xd5\xc4\x3f\x61\x9b\x57\xb0\x9e\x80\x50\x3d\xa9\xc1\x0a\xa7\x7c\xd5\xa3\x52\xf9\xdc\xa2\xce\x0c\x13\x0a\x1e\xeb\x40\xf0\xfa\xf0\x54\xf8\xe3\x04\xa8\x85\x6c\x3e\x76\x92\xd9\xe0\x38\x9a\x70\x08\x24\x15\x1c\x1a\x9c\xe6\xe6\x5e\x97\xe6\xf6\x32\x86\xeb\x2a\x8e\x25\x92\xec\xab\x71\x24\xe0\x76\xa6\x6b\x18\x69\xd1\x09\x51\x84\x86\x52\xe0\x69\xa0\x68\x5a\xdd\x85\xea\x3b\xef\x7a\x12\x38\x46\xf1\x4b\xff\xe9\x0f\xa1\x2f\x00\x7a\x7a\x06\x67\x89\x9c\xc7\xe0\xf5\x73\x9d\x90\x50\x62\x49\xe8\xa0\xa1\xe5\x46\x58\xaa\x48\xf6\x20\x40\x4a\x69\xf0\x3c\x5c\x37\xf8\x15\x0c\xb7\xea\xdf\x07\xc3\x37\x65\xaa\xe7\xb9\x44\x68\x69\xb5\x3d\xf0\x2e\x8c\x6a\x24\x5a\xe0\x5b\x51\x81\x74\x03\xfa\x1a\xbc\x24\xd4\x92\x08\x35\x5f\x84\xf1\x3d\x81\x10\x86\x30\x41\xf0\x97\xb4\x61\x49\xc1\xcd\x71\x6c\x7d\x37\xce\x93\x5f\xc1\x11\xbb\x31\xa3\x76\x2d\x04\x49\xc8\x6c\x82\x8e\xb9\xa9\xe0\x30\x0f\x77\x8e\x1e\xe7\xe0\xd7\x38\x33\xc8\x2b\xf8\x21\x0f\x34\xde\xe8\x5a\x4d\x0e\xf3\x36\x6f\xe5\x68\xba\xe9\x60\x45\x21\xcc\xe7\xac\xd9\xfa\x2c\x4d\xf9\x54\xc7\xdf\x41\x96\x34\xa8\x31\x60\x74\x3e\x76\x15\x9c\xdc\x10\xeb\xa7\xe7\xa7\xa0\x24\x4b\xef\x48\xf7\xe0\x3d\x4e\x97\xf0\x23\xaf\xf6\x8a\xde\x64\x26\x1b\x97\x8c\x6a\x39\x95\xf5\x3c\x2a\xc6\x46\x73\x74\x86\xd2\xab\x65\xbd\x75\x1c\x23\xb9\xd9\x7a\x8a\x40\xe9\xc0\x6c\x7d\x98\x00\xad\x1c\x4a\x71\x20\x30\x3f\xd0\x9d\x8c\x3d\xc2\xe8\x29\xc0\x2b\x96\x21\x4c\xb7\x7d\x8d\xa2\x67\x81\x23\x5a\x52\x96\x96\xcf\xf5\xa7\x29\xc1\x07\xb5\x97\xc6\x67\xde\x55\x38\x3e\x08\x4b\xe7\xbc\xe7\x5d\x69\xaa\x4a\xf1\x0e\x82\x96\xba\xf7\xa1\xe6\x2d\xf2\x8c\x9a\xad\x66\x4d\xbf\xd9\x82\xd4\x04\x7d\xd7\x5b\xcb\xb2\x42\x69\xd6\x34\x0d\x61\xbf\x59\xb3\x89\x9a\x35\xcf\xca\x52\xb6\x9b\xff\x90\xce\x63\x1d\x26\x38\x26\x7b\x94\x75\xdb\xc9\x94\xec\xbf\xd0\xa3\xce\xce\x93\x91\x1f\xfc\xc6\x94\xb8\xbd\xa6\x40\xc6\x7b\x6b\x3c\x10\x6a\x21\xba\xa2\x0a\x5e\xa0\xeb\x61\xc0\x0b\x96\x8c\xb0\xe6\xb0\x9c\xad\xce\x8a\x84\x33\xe8\x2b\x0a\xd9\x85\x60\x4c\x0d\x1a\xdd\xe9\x38\x47\xe5\x01\xf3\x52\xa8\xb9\xb9\xb4\x35\xd0\x2f\xb9\x58\x2e\xfc\x1b\x9e\xbf\x38\xf9\x1c\x65\x69\xbf\x7d\xfc\xcd\xff\xff\x77\xd0\xe5\x2f\x0f\xe6\x29\x4d\x79\xb3\xb4\xa5\x9f\xaf\xa8\x10\xaa\xa1\xd8\xbd\x5c\xe4\x3d\x3d\x92\x5b\xe8\x66\x71\x49\x28\x96\xef\x7e\x60\x35\x78\x8a\xd0\x0b\xb5\xdf\x7f\xb8\x3e\x41\xe7\xad\x1f\xeb\xd2\xfb\xba\xf6\x7e\x0a\x98\x99\x70\xfd\xe1\xc3\xee\xb3\xc3\xf9\xdf\xb6\x2b\x3c\x5d\xe0\xb3\xf9\xc1\x93\xc1\x67\x71\x3d\x65\x1a\x65\xaa\x94\x27\x44\x3d\xcd\x96\x22\x79\x4f\xa4\xd5\xfd\x6a\xef\x5f\x5e\xd3\x64\x3f\x0f\x25\x1a\xcb\x87\xdd\x67\x27\xeb\xbf\xb9\x89\xa2\x8e\xf9\x9d\xd5\x05\x6c\xca\xdb\xa4\xe7\x95\xc3\xe2\x9a\x68\xb3\x14\xec\xfc\xfa\x77\x00\x00\x00\xff\xff\x6a\x92\xf8\x07\xe4\x0a\x00\x00")

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

	info := bindataFileInfo{name: "testimonials.json", size: 2788, mode: os.FileMode(420), modTime: time.Unix(1567469222, 0)}
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
	"clouds.json":       &bintree{cloudsJson, map[string]*bintree{}},
	"customers.json":    &bintree{customersJson, map[string]*bintree{}},
	"press.json":        &bintree{pressJson, map[string]*bintree{}},
	"testimonials.json": &bintree{testimonialsJson, map[string]*bintree{}},
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
