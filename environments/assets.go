// Code generated by go-bindata.
// sources:
// assets/environment-template.yml
// DO NOT EDIT!

package environments

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

var _assetsEnvironmentTemplateYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5a\x7b\x4f\x1b\xb9\x16\xff\x9f\x4f\x61\x72\x2b\x65\xbb\x62\xf2\x98\x00\xa5\x23\xb1\xab\x10\xd2\x12\x95\xd0\x28\x93\x52\x69\x69\x55\x39\x1e\x27\xb1\x3a\x63\xcf\xb5\x3d\x40\xb6\xb7\xdf\xfd\xca\xf6\xbc\x1f\x59\xe8\xe3\xd2\x4b\x85\x4a\xec\xe3\xe3\xdf\xf9\xf9\x9c\x63\xfb\x38\x96\x65\xed\x0d\xdf\xbb\x0b\x1c\x84\x3e\x94\xf8\x15\xe3\x01\x94\xd7\x98\x0b\xc2\xa8\x03\xda\x76\xaf\xdf\xb3\x7a\x2f\xad\xde\xcb\xf6\xde\x0c\x72\x18\x60\x89\xb9\x70\xf6\x00\x98\x50\x21\x21\x45\x78\xb1\x0d\xb1\xfa\x0c\x80\xfe\x0b\xb8\x92\x13\xba\xd6\x0d\xe7\x58\x20\x4e\x42\xa9\x55\x25\xf2\x40\x6e\x43\x0c\x24\x03\x91\xc0\x9d\x58\x6c\x05\x23\x5f\x3a\x40\xda\x9d\x80\x20\xce\xf2\xda\x31\x85\x14\x6d\x9d\x5d\xfa\x8c\x48\xac\x12\xac\x18\x07\xd7\xb3\x51\x3d\xa2\xa1\xef\xb3\x3b\xec\x5d\x43\x3f\xc2\xc2\x28\xb5\x80\x67\xe6\x4f\x3f\x79\x04\x41\x89\xbd\x22\xb6\x4c\xe8\x1c\x0b\xc2\xb1\x37\x82\x21\x44\x44\x6e\xf3\xb6\x5f\x45\xc1\x12\xf3\xe2\xc0\x76\xbf\x5d\x05\x6f\x04\x01\x5b\x01\x12\x9b\x21\x14\x7e\x1f\x46\x14\x6d\x00\xa1\x60\xcb\x22\x0e\xc6\x23\x17\x20\x3f\x12\x52\xeb\x9c\xc2\x7b\x97\xfc\x8d\xff\x71\x3e\xbb\x66\xbe\x29\xbc\x27\x41\x14\x00\x5a\x37\xef\x06\x4a\x80\x20\x05\x4b\x1c\x03\xc0\x5e\x03\x84\x37\x78\x7b\x05\x83\x07\x2d\x77\x2c\xaa\xac\x82\x42\x30\x44\xa0\xc4\xe0\x8e\xc8\x0d\xb8\x63\xfc\x33\xe6\x19\x80\x0e\x00\x97\x18\xde\x62\xb0\xf4\x21\xfd\xac\x06\x78\x44\xc0\xa5\x8f\x81\xeb\x5e\x00\x88\x10\x16\xa2\xe4\x28\x6d\x65\xa2\xeb\x5e\xe8\xe5\xac\xf1\x0d\x37\x5a\x52\x2c\xc1\x8a\xb3\x00\xdc\x6d\x08\xda\x68\x18\x4a\xb8\xa2\xb3\x62\xc5\x94\xd0\x4b\x4c\xd7\x72\xe3\x80\xf6\x4b\x43\xe5\x14\xde\xa7\x4d\xfd\x93\x76\x11\x4b\xaf\xa3\xff\x75\x7b\x79\x07\x9b\x41\x29\x31\xa7\x0e\x68\xfd\xf6\xe1\x83\xf7\xa5\x7f\x30\xf8\xfa\xfc\xc3\x87\xce\x43\x3e\x74\xe3\x3f\xed\xaf\xcf\x5b\x5a\xe5\x88\x51\x21\x39\x24\x54\x16\x6c\x6c\x07\x91\x90\x6a\xcd\x20\xb8\x85\x3e\xf1\xc0\x68\x72\x3e\x07\x4b\x9f\xa1\xcf\x0e\xb8\xef\xe8\x7f\xdd\xfb\x4e\x7b\x6f\x8a\x25\xf4\xa0\x84\x8a\xa7\xe1\x7b\xd7\x71\x46\x3e\x8b\x3c\x13\xe8\x4a\x93\x33\xa1\x12\xf3\x15\x44\xf1\xba\xa6\x61\xfe\x9a\xb3\x28\x8c\xa3\x44\x45\xc6\x25\x5c\x62\x3f\xf9\xa8\x7e\xbc\x84\x83\x56\x1a\x8c\x23\x46\x57\x64\x1d\x71\xad\xba\x95\xca\x16\x53\x47\xf2\x63\x15\x92\x48\x7d\x87\x09\xee\x42\x5f\xec\x5a\x85\xb6\xc4\x19\x1e\x02\x76\x18\x49\x06\x5c\x04\x7d\x42\xd7\x8f\x05\x5c\x0a\xfe\x42\x5f\x1c\xa0\x45\x12\x35\x8e\x54\x49\x35\x69\x36\xf0\x98\x24\x49\x13\x90\x7f\xb6\xca\xe3\xf3\x69\xb1\x49\x45\x21\x2f\xa6\x2a\x0a\x31\x5c\x1c\xfa\x06\x6b\xe9\x35\x87\x54\xe6\x02\x05\xfc\x66\x22\x53\x65\x56\xca\x28\x7e\xfe\x10\x5d\x99\x37\xd6\xa9\x4c\x35\xd4\x26\xd3\xa2\xa6\x58\x24\x9f\x8a\xd2\xe4\x01\x10\x8b\xa8\x4c\xb5\x15\x52\x64\x51\x4b\x92\x01\x77\x6a\x19\x31\xea\x11\xe5\x08\x7a\xc1\x2e\xa0\x28\x18\xd8\x7a\x45\x1d\xe7\x8a\xc9\x56\x16\x12\xba\x69\xfc\xef\x08\xfa\xa2\xe5\x80\x9b\xfd\x39\x5e\x25\xa4\x1c\x80\x76\xfb\xe3\xde\x14\x86\x21\xa1\x6b\x11\x07\xdf\x1c\xaf\x09\xa3\x0b\x36\x9c\x4e\x8c\x92\x48\x58\x18\x0a\x69\xf5\x13\x9d\xc3\xe9\x64\x72\xee\x00\x18\x10\xcb\x5e\x0e\x96\xc7\xbd\xc3\x7e\x22\x78\x87\x85\xb4\xec\x1a\x41\x88\x8e\x4f\x5e\xd8\xc8\xec\x58\x38\x32\x82\x75\x1a\x7b\x03\x7b\x70\xb2\x7c\x61\x52\x15\x0c\x2d\xca\xb8\xdc\x34\xce\xbf\x5a\xda\xab\xbe\xfd\xf2\x28\x91\x16\x2c\x8a\xa5\xeb\x40\x1c\x0e\x8e\x0e\x5f\xf4\xed\x5e\x01\x6d\x9d\xda\xe5\x0a\xf7\x5e\x1e\x79\xab\xaa\xda\x3a\x69\xf4\xe2\x64\x75\x38\x80\x87\x89\x6d\x08\x53\xc9\xa1\x5f\x2b\x8b\xfb\xf8\x78\x75\x72\xa2\x78\x30\x3b\x80\x09\x6e\x23\x79\x3d\x1b\x25\x43\x94\x6f\x3a\xa0\x9f\x64\xee\xfe\xb1\x09\xd9\x68\xe9\x13\x34\xfc\xab\xdf\x20\x66\x1f\x16\xc4\xec\x1a\xb1\x7e\x55\x6c\x50\x23\x66\x1b\xb1\x39\x16\x2c\xe2\xc8\x1c\x44\x52\x78\x66\x43\xd2\x89\x7a\x3c\xb2\x1d\x27\x39\xcb\xcc\x38\x0b\x31\x97\x04\xa7\xc9\x64\x44\x3c\x7e\x66\xf2\xfd\xfe\x2b\x42\xbd\x09\x9d\xc2\x10\xdc\x14\x6c\x3f\x50\x8a\x0f\xcc\xde\xf0\x31\x1e\x37\xa6\x6a\x77\x3d\xa7\xc2\x8d\xc2\x90\x71\xb5\x9f\x4a\x1e\xe1\x76\xb9\xfb\x82\x09\x49\x61\x80\x45\x49\xa0\x9c\x84\x80\x76\xfb\x6a\xc6\x36\x14\x18\x34\x29\xad\x65\xf3\x4c\x77\x83\x85\xd7\x21\x9a\x78\xb1\xfe\x84\x87\x87\xda\x9d\x2e\x67\xc9\xfa\x29\x0c\x4d\xd7\x24\x7c\x4b\x2f\x75\x7a\x75\x80\x32\x2f\xf1\xa7\x5b\x48\x7c\xb8\x24\x3e\x91\xdb\xbf\x18\xc5\x0e\xd8\x77\xb1\x8f\x91\x04\x37\xa0\x77\x00\xf6\x5f\x2b\x63\x84\x0e\x6e\x33\x60\x01\xd7\xb9\x4d\xf2\x0d\xde\x3a\xe0\x0a\x4b\x75\xca\x49\xd3\x90\x3e\x6f\x3a\x31\xa4\x0a\x33\xf6\x13\x31\x63\xff\x40\x66\xfa\x3f\x85\x99\xc1\x13\x31\x33\xf8\x81\xcc\xd8\x3f\x88\x19\x7d\x4a\xa3\x58\xbe\x86\x12\xdf\xc1\x6d\x3d\x33\x25\xa1\x06\x8a\xbe\x61\xf6\xeb\xd9\xe8\x41\x00\xae\x67\xa3\xb8\x7f\x28\x25\x44\x9b\x00\xd3\xc7\x2d\x54\x69\x96\x54\xa2\x6a\x99\xc1\x36\x67\x91\xc4\x0b\x95\xb1\xea\x01\x65\xfd\x8f\x82\xf1\xcd\xae\xab\xe7\xdb\x01\x25\xbe\x3c\x84\x98\x7a\xe2\x2d\x75\x6a\x88\x6d\xc0\x99\x19\x92\xc2\x2d\x33\x90\x9d\xa9\x24\xa1\xfa\x40\x9b\x73\xfb\xe2\x3d\x05\x80\x87\x12\x9c\xe6\xef\x6c\x9e\x61\x7c\xa1\x53\xb7\x87\x1d\x01\x5a\x3b\xa0\xc1\x3a\x33\xa2\x64\x59\x3a\xf5\xa3\x28\x28\x25\xd7\xa7\x83\x6d\x7f\x07\xec\xc1\xd3\xc1\x1e\x7c\x03\xec\x38\x36\x86\xc8\xaf\x87\x98\xf5\xff\xec\x28\x9c\xd0\x25\x8b\xa8\x37\x0e\x37\x38\xc0\x1c\xfa\x33\xc6\x65\x19\xe3\x98\x4a\xde\x90\xbf\x4a\x42\x0d\x68\x33\xa9\x12\x35\x25\x3b\x01\x98\x47\x3e\x36\x55\x1a\x07\xb4\xfb\xbd\x41\x72\x86\x9a\x71\x26\x19\x62\xbe\x03\xda\xc7\xed\x9c\xec\x10\x99\xdb\x3d\xcc\x5d\x66\xc7\x6b\x8e\x85\x3a\x84\xad\xa0\x2f\xd2\x53\xd8\x8e\xd0\x56\x36\xcf\x21\x5d\xe7\x2e\x42\xaf\x38\x0b\x34\x02\xfb\xb0\x9d\x36\x2e\x98\x9a\xfe\xe8\x68\x70\xd4\xce\x98\x73\xdd\x8b\x5f\x87\xaf\xc3\x9f\xc2\x97\x06\x50\xaa\x18\x34\x73\x66\xdb\x25\xc6\x4c\x43\x4c\xd7\x85\x94\xe1\xaf\xc3\xd7\xd1\x13\xfb\xd7\x49\xaf\xc4\x95\x69\x78\x1b\x49\x4d\xd6\xaf\x43\x54\xef\xbb\x88\xca\xdf\x86\xbe\x89\xa7\x32\x4d\x69\x10\x96\x36\xbe\xb2\x31\x0f\xdc\x0b\x6a\x07\x7c\xcf\xce\xfb\x30\x9e\x4b\xdb\xdf\x53\x83\xb7\xbf\x03\xfc\xe0\xa9\xc1\x0f\x1e\x05\x7e\x8c\xc4\xc8\x14\xb5\x6a\xc0\xe9\x92\x6f\x52\xc2\x1f\x8f\xdc\x61\x24\x59\x5c\xfe\xd4\x25\xde\xca\x90\x9c\x40\xe1\x83\x96\x2e\x1f\x60\x6f\x3e\x36\x6d\xe9\xb3\x91\xba\x0b\x4d\x3c\x4c\x25\x59\x91\x04\x9a\xda\xc9\x77\x78\x5a\x6d\xa7\xbd\xab\x33\x61\xca\xdc\xcb\x0a\xf5\x5c\x5d\xbb\x33\x63\x46\x8c\x4a\x48\x28\xe6\x49\x9d\x42\x24\x77\x3b\x42\x75\xcd\x30\x7d\xa8\xc9\xca\x88\x66\x64\xbe\xac\x5b\x2d\x58\x1a\x99\xba\xb2\xf0\x88\x63\x0d\x62\xc6\x7c\x92\x95\x6a\x93\xa2\x8f\x4b\xd6\x14\xe6\x0a\xd4\x0b\x12\x60\x16\x49\x07\xcc\x16\xfd\xa3\xa9\x6e\x7e\x17\x7a\x50\xe2\xe2\xf0\xdc\x6a\xcc\x99\xaf\xfe\x33\x52\x99\xa2\x29\xa1\xa9\x89\x13\xea\x62\x7e\x4b\x50\xc1\x3a\x6d\xdf\x19\x94\x68\x53\xb6\x1b\x80\x19\x8c\x04\x56\x50\xf2\x38\xd4\xcf\x7b\x48\xe4\x5b\x5a\x04\x9f\xcf\x84\x55\x7a\x77\x7b\x55\xcd\x62\x69\xf9\xfc\x63\x85\xb6\xb7\xfe\xc1\x82\xc8\xcc\x60\x94\x2b\xf2\x25\x2d\x41\x00\xa9\x57\x28\xdd\x03\xd0\xeb\x7f\x82\x9e\xf7\x29\x29\xfa\x7e\x92\xec\x13\xca\xc7\x4c\x65\xbc\xba\xc5\x47\x4b\xf0\x9f\x52\x2f\x00\xff\xda\xef\x2e\x09\xed\x2e\xa1\xd8\x54\xfa\x30\xda\x30\x15\x64\x9f\x46\x97\xef\xdc\xc5\x78\x7e\xfa\xec\x4b\x16\x9c\x5f\x01\xf8\xe3\x0f\xd0\xc5\x12\x75\x31\x12\xea\xb7\x63\xd0\xe7\xd4\xac\x88\x8f\x4b\xc8\x5b\x7a\x04\x5a\x51\xf5\x6b\x6d\xa2\x50\x8f\x6a\x55\x61\x53\x89\xa9\x6c\x84\x7d\x13\x40\x42\x3f\x56\x9a\x85\x84\xe8\xf3\xe9\xb3\x2f\x9a\x6a\x57\x7d\x98\x78\x5f\x2b\x52\x5c\x57\xae\x13\x31\x53\xc7\x2e\x4b\x05\xcc\x53\xfe\xd4\xeb\xf5\x0e\x7b\xb9\x1d\xce\xfc\xb0\x3b\xaa\xf6\x5e\xce\x98\x2c\xf5\xac\x75\x12\xaa\xf6\x64\x66\x6f\x18\xfb\x2c\x3a\x9e\x36\x1f\x46\x92\x59\x1c\xfb\x0c\x7a\x98\x7f\x23\x11\x15\x3d\x96\x9a\xa1\x4a\x8d\xe4\x64\xbd\xc6\x5c\x9c\x86\x4c\xc8\x4e\xa4\x23\xad\x22\x14\x42\xb9\x39\x4d\x6b\xb9\x9d\x6a\x24\x74\x12\xa7\xee\x34\x7a\x73\x45\x29\xd4\xc7\x8f\xd3\x2e\x0b\x65\x17\xde\x09\xed\x6f\x0a\x35\xa1\x44\x02\xeb\x16\x58\x96\x5e\x36\x90\x5f\x36\x95\xed\xbe\x02\xcb\xe2\x31\x96\x9a\xa0\xd4\xbd\x6a\xe9\xc0\xce\x85\x04\x80\x47\x14\x8a\xd3\xd2\x92\x08\x93\x4c\x4a\xde\x29\xb6\xe2\x96\x14\x22\x32\x5e\x05\xe3\xab\xe5\x66\x00\xb0\x2e\x30\x7b\xa5\x73\x54\xbe\x5f\x44\x1c\xcf\x23\x4a\x55\xaa\x68\x92\xaa\x89\x13\x60\xde\x64\xea\xa3\x65\xa7\xe4\x3f\x38\x58\xc3\x0e\x37\x09\xe0\xda\x5c\x8e\xf3\x25\xc5\xe2\x33\xcf\x81\xd9\x20\x5a\x39\xb6\x5b\x07\xe6\xd1\x22\x2d\x2f\xba\x18\x45\x9c\xc8\x6d\xfc\xe0\x0a\x6e\xcc\x98\x0b\x26\xa4\xfb\x3a\x95\x2a\xbc\x1d\x96\x6a\xee\xd9\xf3\xe9\x04\x06\x49\xeb\x8c\x33\x45\x52\x2c\x3b\x1e\xd9\xa5\x8e\x78\x44\xf2\xc6\x05\xf6\x27\x2b\x70\x93\x7b\xf5\x8a\xa1\x17\x3f\x19\x43\xae\x98\xbe\x77\xb7\x12\x6c\xef\x04\xe6\xe7\xb9\xb4\x0d\xc0\x2b\xea\x38\x67\x50\xe0\xe3\xc3\xfc\x1a\xd5\x04\x64\x2e\x99\x02\xeb\xbe\x18\x5e\xdb\x28\x30\x8f\x74\xbe\x0f\xac\x2d\x80\x77\xc2\x52\x2b\xb4\x64\x4c\x0a\xc9\x61\x58\x10\x7e\x92\x58\xa9\x4c\x2a\xf4\xd6\x08\x2c\x0c\x9e\xfd\xf9\xb0\x99\x6b\x4e\x64\x3b\xa6\xae\x2e\x63\x65\xa3\x9d\x0c\xa7\x2a\xab\x54\xd7\xba\xea\xc1\x33\x28\x37\x0e\x68\x75\x93\xe8\x98\xb3\x5c\x50\x59\xa9\xe3\xa8\x66\x33\xb7\xfa\xab\x7e\xc2\x58\xa6\x6e\x96\xa1\x10\x51\x80\x95\x80\x39\xcc\x9c\x33\x14\x05\x2a\x41\xa7\x54\xba\x12\x4a\x5c\x6c\xb2\xc0\x78\xb5\xc2\x48\x3a\x20\x7f\x4d\x37\x13\x10\x8a\x48\x08\xfd\x62\xf4\x27\x47\x9d\xbd\x62\x90\x63\x64\x77\x60\x00\xff\x66\x14\xde\xa9\xed\x36\xc8\xf5\xc7\x97\xbc\xbd\xbc\xbc\x90\xc2\xc9\x00\x37\xf0\xa4\xed\x20\x79\xaa\x8c\x65\x26\x90\x30\x12\x56\x9c\x2b\xb3\x93\x55\x83\xe5\xb5\xb6\xef\xb2\xbe\x0e\xb5\xb1\x53\x38\xfa\xc8\x89\xb3\xc3\x7e\xb9\x5f\xb9\x91\xea\xaa\x38\x7b\x8d\xec\x39\xe6\x8f\x91\x26\x02\xb1\x5b\xcc\x67\xcc\xf7\xc7\xd4\x0b\x19\xa1\xb2\x46\xcc\x8d\x96\x01\x91\xbf\x57\x7a\xb8\x53\x6d\x13\x8e\x52\x56\x68\x4e\x76\x59\x07\xb4\x7e\x57\x4b\x61\x32\x64\xc3\x7d\x2c\x9f\x54\x9b\xae\x28\xb5\x55\x47\x3d\xa2\xf0\x5d\x9b\xf1\xc8\xd5\x73\xa5\x99\x1a\x64\x5a\x4b\xe9\x7b\x42\x4d\x91\x20\x75\x8c\x49\x98\x95\x18\x24\x0a\x0b\x95\x80\x99\x79\x7d\x2d\x56\x98\x6a\x1a\x47\xc4\xe3\x93\xb0\x5a\xb8\x1a\xfb\xcb\xff\x89\xf9\x97\x67\x3f\xc9\xf2\x62\xbd\xa8\xa6\x31\xb1\x3c\x5f\x5a\x19\x5f\x9e\xd9\x6a\x35\xe6\x51\xd3\xa3\x4f\x1d\xae\xa6\x3d\xbc\x16\x64\x0e\x62\x0a\x26\xc5\x97\x16\x6b\x34\x05\xe6\x42\x54\x98\x50\x71\xfa\x1a\xcb\xa1\x94\x66\x85\x3a\x71\x73\x9e\xe0\xbc\x90\x71\xe3\x9c\x94\x6a\xb0\xc7\x97\x67\xff\x0f\x16\x56\xc0\xd7\x9a\x58\xe6\x61\x8c\xc4\xd8\x5f\x56\x6d\xf3\xa1\x90\x04\x5d\x32\xe8\x9d\x41\x1f\x52\x44\xe8\xfa\xda\x76\x9c\xac\x21\xce\x6b\x4d\xd5\x94\xf2\xe6\xf5\xa3\x8b\x0c\xa5\xa3\x5a\x69\xa7\x54\x46\xa6\xc6\x5d\xaa\xe4\x49\xeb\x2a\x32\x4d\x46\xc6\x03\x1a\x0c\xcc\x73\x30\xe4\x34\x39\xd6\xe9\xb9\x62\x91\xf8\x4b\x89\x66\x83\xc8\xa1\x33\x93\xaf\x18\xbf\x83\xdc\xcb\x02\x0e\xf2\x35\x96\xda\x92\xb2\xbe\x58\x51\x4e\x22\xdd\xf8\x4a\x21\x9a\xf9\xd6\xc5\x62\x31\x4b\x8d\xaf\x2a\x78\x30\x0d\xe5\x49\x6b\x4e\x2d\x09\x88\x1d\x30\xea\xf2\xdb\xdb\x48\x86\x91\x71\x11\x75\x34\x7d\xc7\xe3\x13\x44\xfc\x86\xa4\x0f\xa7\x1b\x29\x43\xa7\xdb\xd5\xb7\xf6\xb1\xbf\xec\x9c\x5f\xb9\xfa\xc4\xd6\xdd\x03\xe5\x6f\x99\xaa\xb4\xf8\x6e\x7e\x59\x59\x70\x45\x66\x41\x6f\xc6\x6b\x61\x89\x0b\xca\x86\x9c\x02\xb6\x02\x72\x83\xb5\xde\x44\xb0\x53\x53\xdc\x2b\xa9\xcd\x6f\xf7\xc5\x2f\x19\xc3\x00\xa7\x3a\xb3\xaf\xbd\x75\xe2\xe7\x7c\xaf\xaa\x2d\xd9\x04\x0a\x6a\x16\x1b\x0c\x88\x97\x28\xba\x0d\xd1\xde\x7f\x03\x00\x00\xff\xff\xc8\x38\x05\x7c\x34\x2e\x00\x00")

func assetsEnvironmentTemplateYmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsEnvironmentTemplateYml,
		"assets/environment-template.yml",
	)
}

func assetsEnvironmentTemplateYml() (*asset, error) {
	bytes, err := assetsEnvironmentTemplateYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/environment-template.yml", size: 11828, mode: os.FileMode(420), modTime: time.Unix(1483945855, 0)}
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
	"assets/environment-template.yml": assetsEnvironmentTemplateYml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"environment-template.yml": &bintree{assetsEnvironmentTemplateYml, map[string]*bintree{}},
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

