// Code generated by go-bindata.
// sources:
// html/templates/apps.html
// html/templates/libs.html
// html/templates/main.html
// DO NOT EDIT!

package templates

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _htmlTemplatesAppsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x24\x12\xd7\x2d\x6c\xa1\xe5\xf8\xb3\x6d\xe9\x9e\xee\x85\x13\x9b\xc4\xe0\xd8\xa9\x3d\x50\x38\x94\xef\x7e\x4a\xa0\xb4\xbb\x6a\xa5\xbb\x7b\x03\x9e\x67\x3c\xf6\x6f\x46\x8f\x1c\xef\x37\xaa\x22\x58\x67\x0c\x25\x90\x8a\xc0\xf2\x8a\x3f\x24\x88\x8c\x7d\x9b\x49\xbb\x10\x18\xa1\x81\x75\xe0\xa5\x0c\x08\x8a\x12\xa2\x0d\x03\xdf\x5e\xc0\xd4\xb9\xb0\xf7\xba\x24\x29\xf3\xed\x25\x67\xcf\x99\xd2\x60\xa3\x48\x49\x60\x12\x7c\xfb\x99\x53\x48\x7c\xca\x96\x3c\x62\x4e\x19\x7c\x46\x5c\x72\xe0\x44\x38\x26\x22\x82\xf9\xd5\xcf\xc8\x24\x9a\xcb\xb9\x03\xca\x99\x72\xf0\xa5\x2a\xcf\x15\x5c\xce\x91\x66\xc2\xb7\x0d\xac\x05\x33\x09\x63\x60\xa3\x44\xb3\xa9\x6f\x27\x00\x99\xf9\x82\x71\x4a\x56\x11\x95\x6e\xa8\x14\x18\xd0\x24\x2b\x82\x48\xa5\x78\x2f\xe0\xba\x5b\x71\x2b\x38\x32\xe6\x55\x73\x53\x2e\xdd\xc8\x18\x1b\x71\x09\x2c\xd6\x1c\xd6\xbe\x6d\x12\x52\xbb\xa8\x3b\x1d\x79\x56\xbb\xa8\xaf\x9e\x86\x55\xa2\x1e\x26\xad\x93\xca\xd9\xc5\x68\x32\x58\x0d\xe2\xc6\x74\x5d\xbf\x7e\x58\x7e\xef\x27\x95\xab\xd3\x46\x6d\x92\xb6\xa3\xae\x18\xb7\x9e\x79\x27\x6e\xb7\x1e\x30\x6d\xf1\x71\xa3\x3b\x49\x6d\x14\x69\x65\x8c\xd2\x3c\xe6\xd2\xb7\x89\x54\x72\x9d\xaa\x85\x29\x06\x89\xb7\x93\xf4\x42\x45\xd7\x81\xe5\x01\x09\x05\x43\x91\x20\xc6\xf8\xf6\x36\x28\x7f\x9d\x50\x69\xca\x34\xa3\xbb\xd0\xa4\x45\x31\x14\xc5\xfb\xdd\x45\xe0\x50\xa2\xe7\xe5\xa4\x40\x07\x1e\x24\x81\xe0\xa1\x26\x9a\x33\xe3\x61\x48\x82\xcd\x46\x13\x19\x33\xe4\xb6\xb2\xec\x9e\xb3\xe7\x3c\xf7\x4a\xd5\xed\x93\x94\xe5\xf9\x6e\x0f\x93\xb4\x5c\xeb\x82\x0f\x76\x80\xb0\x25\xdc\x6c\x1c\x74\x48\xb6\xd5\xe8\x8b\xbf\x3f\x09\xe5\x79\x99\xdb\x9e\x7f\x28\x78\x58\x66\x6f\x78\xf8\x92\x2d\x88\x0a\x2e\xfa\x53\x7b\x3b\x60\x2f\x0c\x36\x9b\xb2\xac\x44\x41\xc5\xfd\x61\x80\x3c\x93\x12\x21\xf6\xa9\x6f\x5c\x82\x29\x73\x5b\xdd\xc3\x50\xb8\xf0\xcd\xbd\x24\xcb\x8a\x7b\xf7\x84\x79\x6e\x1d\x1c\xbc\x30\xdf\x6c\xa1\xb8\xa4\x6c\x55\x0a\x05\x9e\xf9\xe9\x56\xeb\xe0\x2d\xe1\x66\xb3\x2b\x73\x2f\x0b\x21\xcf\x6d\x44\x09\x10\x07\x54\x1c\x0b\xe6\xdb\xa0\x94\x00\x9e\xd9\x08\x38\x14\xf1\xeb\xfe\xef\x85\x90\xe7\x76\xf0\x2a\xdd\x33\x6d\xb8\x92\xe5\x64\x5f\xa8\x99\xa4\xe5\x68\xca\x59\xff\xa2\xec\x1c\x81\xcb\x31\x05\x96\x67\x22\xcd\x33\x40\x46\x47\xaf\x56\x8f\x14\x65\xee\xec\x69\xc1\xf4\xba\xb4\xf8\x76\xe9\xd4\xdc\x53\xb7\xea\x1a\xc1\xd3\xd2\xd6\xb3\x77\x5d\xdd\xeb\xd6\xd4\xe9\xd7\x1e\x5c\xcf\x97\x8f\xd7\xbd\xda\xdd\x55\xff\xef\xf4\xf6\xbc\x77\x39\x1f\x69\xac\xaf\x9a\x78\x98\xc5\x0d\xd2\xfa\xd1\xe9\x3e\xb7\xbf\xde\xde\xf7\x5b\xb8\x93\x75\xda\xed\x66\x2d\x99\x64\x9d\xb3\xde\xbc\xff\xb1\xab\x3d\xbc\x65\xfd\x08\x9a\xca\x99\x71\x23\xa1\x16\x74\x2a\x88\x66\x25\x39\x99\x91\x15\x16\x3c\x34\x38\x53\x59\xc6\xb4\x3b\x33\xb8\xea\x56\x4f\xdd\x26\x5e\xa4\xf4\x45\xfc\xb8\x9b\x56\xd6\x0f\xe3\xa4\xf9\xfb\xc9\x63\x75\xd8\x83\x65\x6d\x24\xcf\x1f\x6a\x69\x3c\x58\x25\x77\xcd\x1e\x1e\x47\x43\xd3\x1a\x9c\x27\x77\x3c\x9c\xd4\x9a\xb3\xf3\x29\x99\xb7\x07\x66\xbe\x9c\x2c\xcc\x72\x4a\x2a\x61\x7d\xf8\xbf\xbb\xf9\xb7\xaf\xcd\xec\xd7\xc7\xe6\xfd\x3e\xba\x3f\x46\x8d\x71\xc6\x66\x49\xfd\xae\x72\x4a\x2f\x66\x7f\x40\x63\x79\x73\xf5\x6d\xca\x70\x77\xd8\xe1\xa3\xd1\x78\x38\x5c\x8d\xa7\xed\x87\x8c\x57\x6f\x9f\x16\xf7\xb4\xb5\x9e\xdd\x11\x7d\x76\x72\xde\x18\xdc\x5f\xa6\x8f\xe2\x3f\xf4\x11\x58\x87\x47\xd3\x85\x8c\x80\x2b\x89\x8e\x8e\xd1\xc6\x42\xe8\xf0\xe8\xd3\x9f\xef\x5a\xfc\xaf\x4f\xc7\xee\x6e\x7d\x74\x6c\xe5\xc7\xd6\x9b\xe3\xf6\x16\xc5\x2f\x9e\xdd\x7e\x36\xac\x7f\x02\x00\x00\xff\xff\x95\x2a\x20\x56\x48\x06\x00\x00")

func htmlTemplatesAppsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesAppsHtml,
		"html/templates/apps.html",
	)
}

func htmlTemplatesAppsHtml() (*asset, error) {
	bytes, err := htmlTemplatesAppsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/apps.html", size: 1608, mode: os.FileMode(420), modTime: time.Unix(1525722917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesLibsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x04\x89\xeb\x16\xb6\xd0\x72\xfc\xd9\xb6\x74\x4f\xf7\xc2\x89\x4d\x62\x70\xec\xd4\x1e\x28\x1c\xca\x77\x3f\x25\xa1\xb4\xb7\x6a\xa5\xbb\x7d\x03\x99\x67\x3c\x93\xdf\x8c\x1e\x39\xde\x6f\x54\x85\xb0\x49\x19\x8a\x21\x11\x2d\xcb\xcb\xff\x90\x20\x32\xf2\x6d\x26\xed\x5c\x60\x84\xb6\xac\x03\x2f\x61\x40\x50\x18\x13\x6d\x18\xf8\xf6\x12\x66\xce\x85\xbd\xd7\x25\x49\x98\x6f\xaf\x38\x7b\x4e\x95\x06\x1b\x85\x4a\x02\x93\xe0\xdb\xcf\x9c\x42\xec\x53\xb6\xe2\x21\x73\x8a\xe0\x33\xe2\x92\x03\x27\xc2\x31\x21\x11\xcc\xaf\x7e\x46\x26\xd6\x5c\x2e\x1c\x50\xce\x8c\x83\x2f\x55\xd1\x57\x70\xb9\x40\x9a\x09\xdf\x36\xb0\x11\xcc\xc4\x8c\x81\x8d\x62\xcd\x66\xbe\x1d\x03\xa4\xe6\x0b\xc6\x09\x59\x87\x54\xba\x81\x52\x60\x40\x93\x34\x0f\x42\x95\xe0\xbd\x80\xeb\x6e\xc5\xad\xe0\xd0\x98\x57\xcd\x4d\xb8\x74\x43\x63\x6c\xc4\x25\xb0\x48\x73\xd8\xf8\xb6\x89\x49\xed\xa2\xee\x74\xe5\x59\xed\xa2\xbe\x7e\x1a\x55\x89\x7a\x98\xb6\x4f\x2a\x67\x17\xe3\xe9\x70\x3d\x8c\x1a\xb3\x4d\xfd\xfa\x61\xf5\x7d\x10\x57\xae\x4e\x1b\xb5\x69\xd2\x09\x7b\x62\xd2\x7e\xe6\xdd\xa8\xd3\x7e\xc0\xb4\xcd\x27\x8d\xde\x34\xb1\x51\xa8\x95\x31\x4a\xf3\x88\x4b\xdf\x26\x52\xc9\x4d\xa2\x96\x26\x5f\x24\x2e\x37\xe9\x05\x8a\x6e\x5a\x96\x07\x24\x10\x0c\x85\x82\x18\xe3\xdb\x65\x50\xfc\x3a\x81\xd2\x94\x69\x46\x77\xa1\x49\xf2\x62\xc8\x8b\xf7\xa7\xf3\xc0\xa1\x44\x2f\x8a\x4d\x81\x6e\x79\x10\xb7\x04\x0f\x34\xd1\x9c\x19\x0f\x43\xdc\xda\x6e\x35\x91\x11\x43\xee\x0d\x0f\xee\x39\x7b\xce\x32\xaf\x50\xdd\x01\x49\x58\x96\x21\xe4\x99\x84\x08\xd1\xda\x6e\x91\xfb\x8d\x4b\x30\x28\xcb\x3c\x5c\x6a\xbb\x06\x4c\xd2\x5c\x03\x9d\xc3\xc3\x8e\x1e\x4a\xfc\xed\xd6\x41\x87\xa2\x6c\x8d\xbe\xf8\xfb\xd7\xa0\x2c\x2b\x72\xe5\xcb\x0f\x49\x9a\x16\xd9\x76\x9a\xbe\x64\x73\xdc\x1c\x9a\xfe\x6b\xf6\xdd\x34\x5e\x90\x03\xe5\x65\x05\x67\xc1\x14\xe4\x3c\xb9\x05\xdf\xf4\x15\x3c\xc8\xfb\xee\x09\xb2\xcc\x3a\x28\xf2\x79\xe9\x4d\x99\xe4\x92\xb2\x75\xd9\xeb\x86\x07\xa6\x38\xfc\xd2\xd5\x3a\x78\x4b\xb0\xdd\xee\xca\xdc\xcb\x5c\xc8\x32\x1b\x51\x02\xc4\x01\x15\x45\x82\xf9\x36\x28\x25\x80\xa7\x36\x02\x0e\x79\xfc\x7a\xfe\x7b\x2e\x64\x99\xdd\x7a\x95\xee\x99\x36\x5c\xc9\x62\x73\x2f\xd4\x4c\xd2\x62\xf4\x62\x97\x3f\x29\x3b\x3b\xe0\x62\x0d\x2d\xcb\x33\xa1\xe6\x29\x20\xa3\xc3\x57\x9f\x87\x8a\x32\x77\xfe\xb4\x64\x7a\x53\xf8\xbb\x7c\x74\x6a\xee\xa9\x5b\x75\x8d\xe0\x49\xe1\xe9\xf9\xbb\x96\xee\xf7\x6a\xea\xf4\x6b\x1f\xae\x17\xab\xc7\xeb\x7e\xed\xee\x6a\xf0\x77\x72\x7b\xde\xbf\x5c\x8c\x35\xd6\x57\x4d\x3c\x4a\xa3\x06\x69\xff\xe8\xf6\x9e\x3b\x5f\x6f\xef\x07\x6d\xdc\x4d\xbb\x9d\x4e\xb3\x16\x4f\xd3\xee\x59\x7f\x31\xf8\xd8\xd2\x1e\x2e\x59\x3f\x82\xa6\x72\x6e\xdc\x50\xa8\x25\x9d\x09\xa2\x59\x41\x4e\xe6\x64\x8d\x05\x0f\x0c\x4e\x55\x9a\x32\xed\xce\x0d\xae\xba\xd5\x53\xb7\x89\x97\x09\x7d\x11\x3f\x9e\xa6\x9d\x0e\x82\x28\x6e\xfe\x7e\xf2\x58\x1d\xf5\x61\x55\x1b\xcb\xf3\x87\x5a\x12\x0d\xd7\xf1\x5d\xb3\x8f\x27\xe1\xc8\xb4\x87\xe7\xf1\x1d\x0f\xa6\xb5\xe6\xfc\x7c\x46\x16\x9d\xa1\x59\xac\xa6\x4b\xb3\x9a\x91\x4a\x50\x1f\xfd\xf2\x34\xff\xf5\xaa\x99\xff\x7c\xd3\xbc\x3f\x47\xef\xc7\xb8\x31\x49\xd9\x3c\xae\xdf\x55\x4e\xe9\xc5\xfc\x0f\x68\xac\x6e\xae\xbe\xcd\x18\xee\x8d\xba\x7c\x3c\x9e\x8c\x46\xeb\xc9\xac\xf3\x90\xf2\xea\xed\xd3\xf2\x9e\xb6\x37\xf3\x3b\xa2\xcf\x4e\xce\x1b\xc3\xfb\xcb\xe4\x51\xfc\x8f\x39\x5a\xd6\xe1\xd1\x6c\x29\x43\xe0\x4a\xa2\xa3\x63\xb4\xb5\x10\x3a\x3c\xfa\xf4\xe7\xbb\x16\xff\xeb\xd3\xb1\xbb\x7b\x3e\x3a\xb6\xb2\x63\xeb\x4d\xbb\xbd\x45\xf1\x8b\x67\xcb\x6f\x86\xf5\x4f\x00\x00\x00\xff\xff\x8e\x65\x9b\xee\x45\x06\x00\x00")

func htmlTemplatesLibsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesLibsHtml,
		"html/templates/libs.html",
	)
}

func htmlTemplatesLibsHtml() (*asset, error) {
	bytes, err := htmlTemplatesLibsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/libs.html", size: 1605, mode: os.FileMode(420), modTime: time.Unix(1525722917, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesMainHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\xeb\x8e\xdb\x36\xf6\xff\x3c\x7e\x8a\x13\x06\x68\x6d\x8c\x25\xcd\x34\x93\x4b\x3d\x92\xff\xff\x36\xc9\x74\x93\x26\x9d\x34\x69\x2e\x6d\x51\xa0\x94\x78\x2c\x71\x2c\x91\x0c\x49\xd9\x71\x0d\x03\xfb\x34\xfb\x60\xfb\x24\x0b\xea\x66\xf9\x32\x59\x64\x17\xd8\x2f\x33\xe4\xe1\x8f\x87\xe7\x77\x2e\xe4\x91\xc3\x3b\x4f\xae\x1f\xff\xf2\xeb\xab\xa7\x90\xd9\x22\x9f\x0e\x42\xf7\x0f\x72\x2a\xd2\x88\xa0\x20\xd3\x01\x40\x98\x21\x65\x6e\x00\x10\x16\x68\x29\x24\x19\xd5\x06\x6d\x44\xde\xfe\x72\xe5\x3d\x22\xfd\x25\x41\x0b\x8c\xc8\x82\xe3\x52\x49\x6d\x09\x24\x52\x58\x14\x36\x22\x4b\xce\x6c\x16\x31\x5c\xf0\x04\xbd\x6a\x32\x06\x2e\xb8\xe5\x34\xf7\x4c\x42\x73\x8c\xce\xfd\xb3\x1d\x55\x99\xb5\xca\xc3\x8f\x25\x5f\x44\xe4\x83\xf7\xf6\x3b\xef\xb1\x2c\x14\xb5\x3c\xce\xb1\xa7\x97\x63\x84\x2c\xc5\x76\xa7\xe5\x36\xc7\xe9\x13\x54\xa0\x29\xa3\x3a\x0c\x6a\x41\xbd\x68\xec\xaa\x1d\x03\xf8\x4c\x5a\x8b\x6c\x6d\xf1\x93\xf5\x18\x26\x52\x53\xcb\xa5\x98\x80\x90\x02\x2f\x21\x96\x9a\xa1\xf6\x62\x69\xad\x2c\x26\xe7\xea\x13\xd4\x78\xb8\x7b\x76\xf6\x30\x9e\xcd\x2e\x37\xbb\x7a\x26\x99\x5c\xa0\x86\x35\xfc\x07\xfa\xee\x3f\x88\xef\x5d\x42\xad\x30\x0c\x7a\x56\x86\x39\x17\x73\xd0\x98\x47\xa4\x92\x9a\x0c\xd1\x12\xc8\x34\xce\x22\xe2\xfc\x63\x26\x41\x60\x2c\x4d\xe6\x8a\xda\xcc\x8f\xa5\xb4\xc6\x6a\xaa\x12\x26\xfc\x44\x16\x41\x27\x08\x2e\xfc\x73\xff\x2c\x48\x8c\xd9\xca\xfc\x82\x0b\x3f\x31\x86\x00\x17\x16\x53\xcd\xed\x2a\x22\x26\xa3\xf7\x1e\x5d\x78\xdf\xa6\xef\x7e\xbe\x60\xbf\x5e\x2d\x97\xef\xdf\xdc\x3c\x7b\xf2\x9b\x78\xf1\xf4\xbd\xf8\xf4\xf8\x06\xdf\xbc\xbf\x52\xd9\x73\xbe\xfc\xe1\xd5\x07\x7d\x7e\xc3\xd8\xb3\xec\x1a\x53\x5e\x9e\x5f\x2d\xaf\xef\x7f\x7c\xfd\xc3\xe2\xea\xc3\x35\x7b\xfe\xdb\x05\x81\x44\x4b\x63\xa4\xe6\x29\x17\x11\xa1\x42\x8a\x55\x21\x4b\x43\xbe\x80\x57\x69\xd0\x9f\x49\x61\xe9\x12\x8d\x2c\xb0\x22\xa4\x31\x47\x6a\xd0\x04\x8b\xfb\xfe\x99\x7f\x5e\x33\xa2\x79\x7e\x1b\x8f\x53\x76\xf6\xea\xd1\x3d\xf1\xed\x9c\xfe\xfc\xf2\xf1\xf2\xe6\xd1\xd5\xc5\xeb\xe7\xdf\x3f\x78\x60\xff\x7a\xb6\xbc\xfe\xb1\xd0\x2c\xbe\x78\x70\xaa\xa4\x7e\x12\x5c\x2f\xf4\xf3\xd3\x7b\x0f\xdf\x7f\x7c\xf6\xf2\xe1\x5b\xf9\xbd\x5d\xfe\xed\xfa\xc1\x4f\x79\xfa\x59\x1e\x61\x50\x17\x46\x18\x4b\xb6\x9a\x0e\x06\x21\xe3\x0b\xe0\x2c\x22\x54\xa9\x1a\xc0\x50\x79\x55\x1e\x06\xd3\x41\x18\x30\xbe\x70\x28\x93\x68\xae\x2c\x18\x9d\x6c\xb9\xba\x80\xdd\x18\x86\x39\x5f\x68\x5f\xa0\x0d\x84\x2a\x82\x45\x89\x01\xe3\xc6\xba\x81\x7f\x63\xc8\x34\x0c\xea\xbd\xd3\x41\x78\xc7\xf3\xe0\x8b\x34\xf5\x76\x83\xe7\x4d\x8f\x9b\x51\x0a\x35\x4f\x2b\x47\xd3\x4f\x5c\x9a\xfa\xf4\x6a\x58\x25\xcb\xae\x0d\x1d\x95\xe9\x60\xf0\xae\xac\xe2\xa3\xa4\x40\x61\x87\x5f\x77\xbc\xbf\x1e\xc3\x7a\x00\xc0\xa8\xa5\x13\x98\x95\x22\x71\x25\x01\xc3\x51\x25\x05\xd0\x68\x4b\x2d\x9a\x49\x0b\x5b\x6f\xc6\xcd\x3c\xa7\xc6\x5e\xeb\x74\x02\x84\xb4\x22\xb9\x3b\xcd\x25\x65\x5c\xa4\x13\x98\xd1\xdc\x60\x2b\x45\xad\xa5\xee\xc3\x66\x3c\xb7\xa8\xbf\x53\xea\x50\xf8\x82\xc7\x7d\x21\x7e\xa2\x85\xca\xd1\x4c\xe0\x77\xd2\xb1\x20\x63\x20\x19\x2a\xcb\xa5\x1b\xbd\xa4\xc6\xa2\x2e\xb8\x60\xc6\x4d\xcb\x18\xb5\x97\x56\x2b\x37\x14\x53\xd4\x56\xd3\x84\x8b\x94\xfc\x51\xeb\x74\x65\x5d\x31\x4a\x34\x52\x57\xf0\x15\xfb\x93\x13\x3e\x83\xa1\xcd\xb8\xf1\xa5\x4e\xe1\x4e\x04\x84\xd4\xf2\x93\x4a\x38\x43\x9b\x64\x4f\xa8\xa5\xc3\xd1\xe0\xe4\xa4\xd3\x51\xa0\xcd\x24\x33\x93\xc6\x65\x1d\x6a\xeb\x52\x80\x4e\x71\xe3\x9c\xed\x4a\xeb\xf0\x66\xba\xd9\xdf\xe0\x2c\x89\x5a\x4b\xda\x2d\xd5\x4a\xe5\x51\x88\x80\x3c\x2d\x94\x5d\x55\xb7\x3c\xf9\xbc\xd2\xbe\x05\x10\x81\xd5\x25\x0e\x8e\x29\x24\x3b\xf8\x3a\xe0\x0e\xdf\xd8\xd3\xac\x56\x39\xd8\x1d\xe8\xa7\x68\x87\x24\xa0\x8a\x07\xff\x57\x3f\x38\xa7\x2d\x7e\xb4\x05\xd9\x0c\xc5\x50\xa3\x51\x52\x18\x84\x68\xda\xe3\xd4\x1c\xe7\xf2\x0d\x22\x68\x31\xd5\xfc\xb2\x07\xea\x1c\xe3\x16\x7c\xaa\x94\xf1\x73\x14\xa9\xcd\x9c\x97\xce\x46\x3b\x0a\xf7\x79\xfd\x24\x41\xa3\x92\x06\x96\xdc\x66\xc0\x50\x19\xd2\x43\x6f\xba\xf1\xa6\x67\x71\x42\x6d\x92\x0d\x1b\x15\xbb\xf6\x26\x52\x18\x99\xa3\x9f\xcb\xb4\x06\x8c\xf6\xc9\x74\x27\xbf\x6e\xe8\x4c\x80\xc0\x69\x5d\x0b\xfe\x0e\x45\x38\x05\x32\x06\x63\xa9\x2d\x0d\x24\x92\x1d\x47\xd6\xeb\x47\xed\x9c\x71\x41\xf3\x7c\x35\x1c\x8e\x9c\x99\x7b\x91\xae\x2a\xb1\x06\x37\x85\x5c\x17\x5a\xbf\xfe\x65\x7c\x83\x89\x35\x63\x30\x65\x6c\xac\xde\xcd\xdd\x5a\x76\x98\x88\xcd\x65\xd1\xec\xbd\xdc\xcb\xb7\xdd\x55\xbf\x3e\x73\xe8\xe6\xce\x46\x19\xdf\xf8\x2e\x53\x7c\x2e\x92\xbc\x64\x68\x9a\x53\x46\xa3\xcb\x43\x4b\x7f\xc4\x95\xf9\xdf\x58\x7b\x5d\xad\xfa\x73\x5c\x99\xf6\x94\x1d\x37\x57\x1c\xe6\xb8\x72\x14\xe6\xb8\x3a\xb4\x7e\x0b\xd6\xc8\xca\x04\x87\x4e\xcd\xd8\x61\x47\xfb\x19\x24\xe3\x9b\xdf\xe7\xb8\xfa\x03\xa2\xd6\xa8\x6a\xda\xcf\xf7\xad\xcd\x5b\xe9\x66\x0c\xeb\x4d\xdf\x4b\xd5\x1f\x8b\x85\xca\xa9\xc5\x09\xfc\x59\xbd\x7a\x49\x4e\x8d\x89\x48\xe1\x7d\x43\xda\xfe\x2a\x9c\x49\x5d\xb4\x0b\x6e\xec\x71\x91\x73\x81\x50\xc4\xde\x37\x04\x16\x9e\x14\x13\x53\xc6\x05\xb7\xbe\xd2\xb8\xa8\x9a\xb9\xee\x3e\xeb\xb4\xb8\x47\x74\xab\x9f\x0b\x55\x5a\x2f\xd5\xb2\x54\x50\xe8\xde\x61\x9f\x01\x7a\x4a\xa3\x42\xc1\x76\xb0\xae\x19\x54\x54\x1c\x83\xbb\xfe\x8d\x4c\x53\x6e\xb3\x32\xae\x9e\xc2\x30\x70\xd0\x9d\x93\xea\xa7\xbc\x27\xa8\x14\x80\x5d\x29\x8c\x48\xa5\x60\x87\xb8\x6b\x56\xb5\xcc\x1d\xe9\x42\x32\xd7\xf5\x48\x9d\x12\x50\x39\x4d\x30\x93\x39\x43\x5d\x49\xa8\xe0\x7f\x55\x3d\x23\x48\x5d\x5d\x1e\x04\x68\x69\xe5\x4c\x26\xa5\xe9\xf9\xa3\x3e\xfc\xe4\xe4\xe4\x24\x8c\x4b\x6b\x65\xc7\x22\xb6\x02\x62\x2b\x3c\xa5\x79\x41\xf5\x8a\xc0\x84\x71\x43\xe3\x1c\x59\x44\x9a\xfa\x24\xd3\x90\xb7\x70\xd6\x06\x64\x46\x0d\xcc\xa8\x67\x14\x17\x02\xb5\x1b\xaa\x32\x37\xe8\xcc\xe5\xb3\xfe\xd6\x80\x4f\xe1\x4d\x42\x45\x18\xd4\x07\x77\xa1\x0e\x1c\xcd\x6e\xd6\x73\x4e\x58\xe6\xed\x71\x39\x37\xb6\x39\xb1\xd5\xdc\x3e\xb6\xed\xad\x3a\x85\xb3\xdd\x88\xe6\xfc\xc8\x6e\x8f\x5b\x2c\x9c\x8a\x99\xd4\x9d\x0e\xe0\xa2\x7b\xbb\xf7\x23\x4d\x9b\xc6\xf2\x6e\x17\x95\xba\xf5\x26\xf0\xff\x49\xce\x93\xf9\x36\xff\xa4\x4e\xa3\x46\xcb\x9e\x8d\xee\x85\x76\x41\x9b\xae\xd7\x8d\x64\xb3\x09\x03\xba\x9f\x53\x58\x5f\x19\x0b\x0f\x73\x83\xbb\xd8\x66\x6d\x37\x8f\x72\xbe\xf5\x61\x99\x6f\xc7\x7d\x17\xf6\xd2\x9a\xe6\xa8\x2d\x54\x7f\x3d\x46\x45\x8a\xba\x33\xd3\x5d\xde\x3d\xe2\xeb\x75\x7d\x9f\xc3\x66\x73\x54\xe9\x9e\xa5\xbb\xd5\xb6\x17\xf7\x9e\xc5\x2f\x6a\xd9\x3f\xff\xfe\x8f\xfd\x7c\xbc\xc5\x07\x95\xa6\xee\xf9\x84\xaf\xbe\x82\x83\xb7\xf4\x20\xea\x3d\xc2\xd6\xa5\xaf\xd7\xbc\x49\x7c\x81\x5e\x71\x50\xc7\x15\x64\x07\x0f\xf5\xae\xfa\x8b\x0b\x59\x33\x35\xc5\xde\x4e\xb7\xd7\x35\xf0\xdd\x5e\x37\xf1\x18\xd5\xf3\x03\xa0\x83\xea\x43\xa1\x13\xb3\x63\xe2\xa3\x37\xc2\x4e\xb9\xd7\x77\x3b\xd8\x0c\xc1\xf9\xa2\x77\x35\x74\xbd\x6a\x1b\xda\x7f\xe7\xbd\xf3\x23\xe6\x7e\xa1\x09\x39\x8f\x0f\x4d\x78\xc1\xe3\xa3\x9a\xc3\xe0\x38\xe9\xd0\x66\x6d\x51\xe6\x3c\x1e\x03\x77\x35\xd9\xbc\x61\x95\xd9\xee\x94\xf1\xb6\xed\x1e\xb9\x13\x63\x2e\xd8\x64\x8e\xab\x6a\x4f\xf5\x44\xdf\xc2\xc6\xa5\x74\x0b\xd9\x6c\x8e\x9b\x95\x1d\x89\x5b\x70\x18\x38\x87\xec\x7e\xd1\xe8\xdb\x5f\x7f\xc7\x1d\xf2\xd2\x2d\x2f\xaa\xd4\x11\x5e\x2e\x1c\xe3\xed\x37\xc6\x1e\x2f\xaa\xd4\xed\xbc\x42\xcb\x76\x53\xbd\x4e\xbf\xf5\xba\xdd\xe5\x6e\x8e\xdb\xdc\xcd\x6e\x73\xb7\x6b\x60\x86\x4e\xc3\x7f\xe5\xf1\xda\xdf\x0b\xd4\x86\x4b\x71\xc4\xe5\xc7\x0d\xbb\xc5\xe1\x87\xae\x0d\x83\x8a\xf1\x67\xde\xd5\x83\x3b\x73\x4f\x50\xc3\xff\x1c\x6c\x46\x83\x81\xeb\x8f\xad\x2b\x25\x88\x40\xe0\x12\xde\x95\x38\x74\xdd\x0f\xe6\x13\x20\x77\xdd\x07\xf9\x78\xe0\xda\x98\x41\xef\xfb\x35\x0c\x6a\xab\xc2\xa0\xfe\xe9\xeb\x5f\x01\x00\x00\xff\xff\xa2\xaa\x7d\xb8\x0b\x13\x00\x00")

func htmlTemplatesMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesMainHtml,
		"html/templates/main.html",
	)
}

func htmlTemplatesMainHtml() (*asset, error) {
	bytes, err := htmlTemplatesMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/main.html", size: 4875, mode: os.FileMode(420), modTime: time.Unix(1525769621, 0)}
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
	"html/templates/apps.html": htmlTemplatesAppsHtml,
	"html/templates/libs.html": htmlTemplatesLibsHtml,
	"html/templates/main.html": htmlTemplatesMainHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"apps.html": &bintree{htmlTemplatesAppsHtml, map[string]*bintree{}},
			"libs.html": &bintree{htmlTemplatesLibsHtml, map[string]*bintree{}},
			"main.html": &bintree{htmlTemplatesMainHtml, map[string]*bintree{}},
		}},
	}},
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

