package models

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_app_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x5a\x7b\x6f\xdc\xb8\x11\xff\x3f\x9f\x82\x10\xee\x8f\xd6\x58\xaf\x1f\xc1\x5d\x5b\xa1\x2d\x60\xaf\x9d\x8b\x7b\x4e\xce\xd8\x75\x72\x40\x83\xa0\x90\xb5\x5c\xaf\xba\x12\xa9\x92\x94\x2f\x1b\x63\xbf\x7b\x87\x0f\x49\x7c\x69\x1f\xbe\xf4\x50\xdc\x2b\x2b\x0e\x87\xc3\x79\xfe\x66\x78\xcf\xcf\x68\x8e\x17\x05\xc1\x28\xc9\xea\x3a\x41\x9b\xcd\x2b\x84\x9e\xe1\x1f\x84\x92\x8b\x5f\x66\xf7\xb8\xaa\xcb\x4c\xe0\x37\x94\x55\x99\xf8\x88\x19\x2f\x28\x49\x50\x8a\x92\xf3\xd3\xb3\xd3\xe3\xd3\xbf\xc0\xdf\xc9\x48\x93\x4f\x28\x99\x17\x02\xd6\x79\x92\x1a\x16\xc0\xea\x19\x09\xc3\x03\x25\x35\xa3\x39\xe6\xfc\x38\xef\x29\xd1\x58\x1f\xa9\x38\x5c\x96\x19\x59\x4d\xca\x86\x0b\xcc\x24\x0f\x94\xbc\x21\x69\x7a\xfd\x9f\x26\x2b\x25\xcf\x4f\xf2\xcb\x14\x2f\xe0\x8f\x49\x4b\x85\x36\x23\x94\x24\xe8\x33\xd2\x4c\x36\x46\x96\xbb\x8c\x65\x15\x06\x02\x2e\x85\x8d\x0a\xf3\x90\xc1\x69\x39\x66\xc7\xb5\xa4\x75\x25\x89\x4a\xed\xd0\xb5\x22\x5b\xd2\x9a\x4f\xf0\xf1\x7e\x5d\x63\xa5\xa4\x99\x60\x05\x79\x34\x0a\x52\x4b\x57\x78\x91\x35\xa5\x50\xab\xee\x77\x9e\xb3\xa2\x16\xad\x7a\x13\xb3\xb4\x19\x75\x27\xd5\x4d\xe4\x14\x20\x7d\xdf\x54\x0f\x20\x41\xe4\x10\x65\xa6\xd3\xa1\x63\xa4\x16\xef\x3e\x20\xbe\xcc\x18\xe6\x88\x2e\x10\xce\xf2\x25\x32\xb7\x0d\xcf\xbf\x26\x4f\x05\xa3\xa4\xc2\x44\xc4\xe5\x18\xbe\xec\x96\xbb\x46\xaf\xfa\x13\x5e\xff\xaf\x8f\x98\xe2\x12\x67\x1c\xff\x0e\x76\x9b\xe2\x9a\xf2\x42\x50\x16\xbb\xd3\x6f\x3b\x6c\x46\x1b\x96\x63\x94\xd3\x39\x46\xac\x3f\x26\x10\x61\xd6\x3c\x10\x2c\xf8\xc0\xf9\xb7\x05\x17\x7f\x85\x58\x87\x48\x9b\x9c\xa7\xa9\x26\x4e\xd3\x9b\xf9\xdf\x5f\x22\xd3\xc7\xbb\x09\xe2\xfa\x3c\xb4\xa0\x0c\x89\x65\xc1\x91\x4c\x2d\x81\x54\x6d\x36\x71\xa4\xb2\xec\x29\xa3\x8f\x8b\x6d\xde\x4b\xc9\x13\xfd\x02\x37\x57\xa6\x44\x4f\x86\xdf\x68\xd0\x6f\x42\x11\xee\x26\x03\x4a\xe9\xf5\x01\x34\x52\x19\x2f\xd5\x45\x54\x07\x4e\xae\x9a\x62\xae\xec\x68\xdb\x27\x99\x40\x5a\xa1\xd5\x3d\xad\x8b\x7c\x4a\xcb\x98\x9f\xb6\x42\xde\x5c\xbc\x4b\x53\x45\x63\x49\x72\xc7\x68\x8d\x99\x28\xb0\x6b\x74\x99\xd4\x39\x6f\x2a\x2c\xe9\xef\x68\x59\xe4\xeb\x2b\x9a\x37\x41\x4c\x7b\xf6\x91\xc9\xfe\xfc\x18\xf2\xfd\xd9\x9f\xac\x43\xb4\x6b\x09\xb0\x92\xd9\xff\xc9\x59\x42\x1e\x3f\x45\x7e\xbd\x58\xe0\x5c\x59\xf7\xa2\x2c\xe9\xaf\x1e\x37\x23\x7a\x41\xf2\xa2\xce\x4a\x5d\x01\x66\x98\x3d\x15\x39\x56\xe9\x1f\x5c\xa2\x7a\x98\x67\xe3\xac\xca\xbe\x52\x92\xfd\xca\xc7\x39\xad\x54\xf2\x8f\xf0\xb9\xc8\x8d\x9f\xc0\x3e\x2e\x78\xda\x5f\x1c\x76\x78\xe4\x1b\xe7\xb7\xbd\xea\x70\x86\xb2\x22\x96\x52\xf8\x93\xc4\xfd\x2c\x35\xa9\x75\xed\xea\xc0\xd7\x80\xa6\x5c\xbf\x87\xda\xa4\x74\x30\xaf\x0a\x02\xd1\xc7\x32\x88\xdb\x40\x17\xc9\x0e\x03\x29\x9a\x7d\x8c\xa4\x08\x1d\x43\x49\xc5\x06\xa6\xb0\x54\x96\x1c\xc9\x9f\xad\x63\xea\x0f\x68\xb3\x43\x6d\xf6\xaf\x9e\x72\x13\x16\xb2\xde\xb5\xb7\xb8\xf5\xad\x32\x75\x9a\xbe\x69\x88\x96\x6a\x2f\xef\x9e\x40\x2a\x0c\x3d\x79\xf6\xfa\xb2\xc9\x57\x58\xf4\x98\xe2\x1f\xb4\x30\xae\x71\x0c\x37\x85\xff\xe4\x2a\x97\xc0\x9f\x7b\x88\xa1\xc4\x98\xe2\x47\x15\xcd\x70\xf9\xd0\xcf\x80\xb1\x29\x55\x3e\x57\xcd\xd4\x64\xa6\x13\x87\x6d\x07\xa2\x24\x72\x39\x59\x28\x60\x05\xbf\xc7\x5f\x8b\x3a\xd1\x87\x0c\xba\xdf\xdb\x8c\xcc\x4b\x05\x36\xda\x48\xc0\x5f\x00\x7c\x10\x88\x15\x87\xee\x1d\xae\xa0\x0e\xcc\x8a\xaf\x4a\x9d\x67\xe7\x7f\x76\x97\xdb\x84\xa2\x85\xfe\x11\x8b\x0b\xa1\xbd\x22\xc8\x3a\xd2\x27\x18\x09\x22\x2c\x99\x36\x44\x14\xda\x87\x09\x68\xfc\xdf\xdc\x3d\xe0\x1e\xd6\x68\xa3\x7c\xeb\xf5\x69\x32\xec\x0a\x71\x3c\xc6\xba\x7c\xb8\x13\x92\x1d\x40\xca\x75\x1e\xf1\xf0\x9b\x43\xca\x71\xde\xb0\x42\xac\x93\x01\x56\x5c\x86\x50\xb7\xd8\xa6\xef\x9f\x1b\x51\x37\x62\x10\xf3\x76\xd7\xa2\x86\x6e\xa7\xa4\x2e\x61\x57\xc3\xb1\x10\x50\xc3\xbc\x22\xfe\x31\x2b\x1b\x63\x4b\xe3\x5e\x1d\x5d\xaf\xee\x57\xed\xbf\x37\xaf\xe0\x40\x4c\xe6\x8a\xaf\x85\xfc\x63\xb8\x5c\x37\x02\xcf\x88\x65\xe4\x11\xa3\xef\x56\x28\xfd\x1b\x1a\x5f\x13\xc1\x54\xf6\xe2\xed\x1d\x34\x66\x07\xba\xa6\x86\x90\x94\x74\x9b\x4d\x9f\xb2\xb7\x21\xf8\xf8\x9e\x1e\xcf\x8f\xf4\xf9\x46\xdc\xed\x82\xb7\xd0\xdc\x13\x1a\x2b\xa1\x3b\x51\xfb\x13\xf1\x58\x5e\x02\x16\x26\xb4\xaa\x20\xa8\xec\xca\x3b\x84\xcb\x9c\xaa\x0f\xac\x72\xbd\x55\x32\x33\x5c\x80\x9f\x4d\x1d\x03\x86\xad\xd3\x44\x44\x81\x0d\x05\xc3\xf3\x09\x6d\x9c\x5c\xdf\xcb\xe3\xe1\x7c\x47\x9e\xb3\xe1\x83\xef\x97\x18\x11\xb5\x55\x62\xfc\x82\x80\x17\x43\xf0\xaa\x9c\xa3\x50\xbf\x80\x75\xa3\x47\x24\x28\x02\x67\x04\x48\x29\x6f\xb3\xc2\xb8\x46\xac\x21\x04\xb4\x80\x28\x41\x6b\x08\x35\x94\x9b\x7e\x67\xd7\x6d\x6e\xaa\xec\x11\x1f\xac\xd6\xdf\xa0\x3e\x9d\xf3\x82\x13\xb7\xe8\x4d\x96\xcc\xef\x7f\x88\x1f\x09\x6b\xef\x2e\xa5\x76\xa6\x17\xef\xa4\x56\x20\xd1\x80\x83\xe2\x9d\x52\x58\xae\xff\xed\x2f\xbe\x6f\x34\x74\x09\x44\x39\x7e\xf2\x13\x2c\xf2\xa2\x4f\x1c\x91\xa4\xd1\x92\xe8\x2c\x31\xda\xc1\xdf\xca\xba\x03\x27\x38\x95\xdc\xac\x42\x67\x21\x18\xce\xaa\x76\x5c\x10\x2d\xe2\xc9\x0c\xba\xd1\x2e\x06\xce\xfa\xcc\x65\xee\x5f\x2c\xd0\xf8\x6d\xc6\xef\xb4\x24\x56\x1a\xba\xa5\x8f\xfc\x03\x77\x1a\xf1\x08\x4c\x56\x14\x9d\xaa\x07\x60\x44\x8f\xf2\x34\x28\x38\x71\xa0\x47\x1c\xea\x79\x80\xc3\x85\x79\x52\x36\x0f\xa0\x5b\x54\x5b\xf0\xdd\x9e\xe8\x6e\x2b\x08\x8f\xc1\xf0\xbd\x80\xb8\x07\xa1\x57\xc6\x8a\x77\x8d\x98\xe2\x9c\xb2\x39\x58\xff\x73\x74\x97\x05\x1b\x3f\x0d\x21\xa3\x8c\x91\x14\xf0\x7b\xda\x72\x3d\x82\xbf\xb8\x72\x8f\x93\x10\x83\xc1\xf5\xf2\x95\x52\xa6\xaa\x0e\xc7\x47\x06\x27\x05\x78\xd4\x47\xa4\xc8\xa3\x70\x90\xd5\x2b\x9f\xc6\x2d\xeb\xd2\x68\x17\xb9\x1a\x85\x6c\x75\x29\x4d\x23\x61\xe0\x4e\xbf\x82\xe4\x50\xa8\xb6\xc6\x4a\xd7\xc6\x7a\x0d\x57\x5c\x41\xdf\x4f\x4e\x1f\x27\x1d\xd6\xb8\x51\xaf\x94\xce\xd7\xfd\x12\x7f\x48\x92\xe8\xf1\x50\x5b\x34\x87\x62\x6b\xa8\x9e\xc6\x33\xe0\xf5\x64\x76\x9f\xf1\xd5\x95\x3c\xae\x10\x91\xfe\xbe\x06\xa1\xf8\xcf\xca\x1b\x9c\x56\x60\xd4\xf5\x7a\x2a\x56\x3e\x47\xda\x78\x4d\x2e\xfb\x72\xff\x0c\x8b\xd8\x8a\x99\xb3\xf1\xe9\x7e\x6d\x83\x39\xf8\x9e\xae\x30\xd9\x89\x8c\x07\x51\x71\x6f\xa8\x58\x8b\xb1\xdd\xa9\x41\x93\xbd\x0e\x93\xb0\xd9\xb0\x07\x56\x1d\xa3\xf6\x9b\x47\xea\x4d\xea\x3a\x72\xfb\xbb\xb7\xa5\x6b\x63\xda\x82\x80\xd7\x3e\x89\xd4\xb8\xc1\x70\xd2\xb1\x00\x08\xfe\x4b\xc0\x27\xf0\x87\x5e\xf0\xad\x7d\x5f\xdc\x59\xc2\x8a\xe9\x79\x89\xab\xfa\xad\x4e\xd1\xf2\xfa\xbf\xf0\x06\x67\x8e\x1d\x4e\xad\x6d\x52\x1f\xfa\x45\x30\x72\x1c\x26\xfe\x9e\x0e\x18\x46\xf5\x36\x31\xc3\x10\xf5\xfd\xb9\x6b\x42\xbb\xc6\xa5\x0f\x7e\x8f\xf6\x96\x66\xf3\x4b\xd3\x47\x79\x2e\x58\xc2\x52\xdb\x62\xf1\x3d\x7d\xb1\x4f\x8c\xbb\xd3\x65\xf0\x4a\xe0\xb7\x18\x52\x0a\x27\x2d\xf6\x6b\xc5\x08\x7d\x57\x53\x26\xe4\xb2\x26\x1c\xdf\xc1\x4f\xee\xf4\x7f\xdf\x01\x63\xf8\x04\x24\x7f\xe0\x75\x59\x08\xb3\x25\x49\x93\x3f\xc6\x13\xad\xe2\x63\xee\x28\xd9\xc9\xcc\x4d\xe6\xf8\x4b\xcb\xe9\x14\xbe\xb7\xca\x7a\xd1\xc4\x39\xc6\xf0\xb0\x91\xf7\x81\xd2\xbe\xa5\x3c\xf6\xa4\x70\x88\xa4\xfa\x9c\x29\xb4\x2d\xb4\xe2\xa0\xf9\xc3\x44\x7e\x91\x43\xb8\x28\xdb\xaa\x9f\xb6\x89\x93\xd6\x12\xde\x1d\x6d\x0c\xee\x27\x97\xce\x78\x10\x95\x57\xef\x67\x3a\x44\xbd\x07\xae\x6d\x0e\xd8\x0a\x63\x94\xe2\x8b\x14\xd9\xee\xbb\xe5\x90\x6b\x8e\x3d\xb7\xfc\x66\xae\xe9\xab\x24\xcc\x2b\x07\xb1\x1e\x1a\xa1\xd9\x76\x0e\x7f\xbf\xc8\x0b\xfc\x5e\x68\x97\x1f\xcc\xcc\x8c\xe9\x47\x46\x9b\x7a\x10\x59\xea\x87\x18\x87\x74\x27\xba\x54\x64\x6e\x0f\x1b\x14\x02\xb4\x47\x21\x38\x7e\xe8\xd4\xe8\x15\x81\xc4\x91\xe8\x86\x3c\x32\x8d\x8e\x55\x3a\x2e\xf4\x4f\xe5\x86\x76\x63\x90\x7c\xac\xf3\x9b\xb9\x63\x56\xf9\xea\x12\x83\xae\x96\x9a\x86\x35\x53\x66\x5c\x14\x79\x5f\x11\xe0\xdc\x34\xb5\x0b\xc4\x1e\x38\xbc\x7f\x0e\xeb\xab\x8f\xf9\xe6\xdc\x77\x42\x09\xc1\xaa\x07\xba\x62\x59\x21\xe7\x20\xba\x65\xd3\x3b\xaf\x49\xf6\x50\x62\x79\x35\xc1\x1a\x3c\xb2\xc7\x9e\x3f\x9c\x0e\xf0\xb1\xa7\x78\x28\xb9\x99\x97\x78\x78\x13\xa3\x9c\xff\x93\x12\xdc\x1e\xd0\x2f\xbd\xc5\x59\x29\x96\x93\x25\xce\x57\x3e\x76\xd1\x4b\xeb\xfb\x25\x18\x63\x49\xcb\xb9\x6a\x1d\xdd\xe9\xec\x0d\x01\x08\xf2\xa4\x1a\x91\xef\xbd\x0a\xcf\x1e\xf5\x90\xfc\x19\x2d\x0a\xc6\x45\x2e\x4f\x50\x16\x1d\x18\xef\xbe\x76\xbe\x7f\x20\xcb\xe8\xe9\x7d\x1d\xb6\xee\x20\x5f\x1d\x31\xb1\xea\x79\xd9\x7e\x08\x5d\xe8\xf6\x72\x42\xe9\xaa\xc0\x33\xb0\xbc\xea\x18\x79\x67\x87\x4f\xcf\x7e\xb7\x9d\x2d\x14\xe4\x90\xf8\xd5\xe1\x61\xb9\x48\xd0\x51\x85\x71\x30\xe0\xf4\xfe\x10\x33\x1e\xd6\x56\x6b\x7a\x40\x67\xe6\x8d\x9f\x1d\x40\x14\x9f\xac\xf8\x4f\x7f\x03\x33\x95\xbd\x9e\xfc\x06\x87\x08\x5e\x96\xee\x67\x02\xfe\x14\xc2\x7b\x64\x73\x96\xbd\x49\xc1\x8e\xf1\x83\xfb\x06\x18\x4c\x3b\xfa\x17\xc1\xa0\xff\x4f\x70\xce\xbd\xe7\xc1\x97\x4c\x04\xe2\x0d\x84\x35\x82\x09\x20\x4f\xf8\x32\x18\x1f\x15\x39\xc1\xea\x3a\xae\x6d\xef\xf0\x79\x71\xcb\x43\xed\xb7\x7f\x83\x1d\x34\xb2\x5a\xc5\x3a\x07\xf7\xd0\x5b\xe6\x60\x5d\x79\x1e\xf0\x51\x7c\x94\x34\xb0\x87\xe1\x47\x19\xf6\xec\x46\x8d\xa5\xa1\x8c\xbe\x61\xb4\x8a\x66\xf3\xdd\xdc\xa6\x3e\xaf\x5f\x0a\xb1\xdc\x83\x57\x7e\xbe\x53\x78\x20\xb9\x68\xc4\x92\xb2\xe2\x2b\x8e\x16\xc1\x60\x57\x64\x34\xe6\x0c\xc6\x62\xc7\x1c\x45\xd8\x78\x5f\xb6\x3c\x54\x07\xdd\x8e\x5e\xdd\x3a\xd1\xb5\x5f\xb4\xc2\x17\x26\x37\xd9\xcc\x5e\xa7\xa9\x79\x46\x35\xd9\xe6\x0a\x97\x58\xfa\x49\x97\x8c\xe1\x86\x02\xea\xe4\x8e\x6c\xa4\x06\x66\x50\x12\x05\xa3\x6a\x18\x06\xc1\xfe\x24\xc5\xe8\x02\xe7\x3e\x7b\xf4\x82\xa6\x1d\x50\x24\x7c\x0d\x16\xae\x24\x32\x6e\xd1\x62\xfb\x72\xeb\xe0\xbc\x8e\x5e\xfe\x4f\x17\xa3\x18\xb4\x0c\x32\x7e\x4c\x6d\x96\xd6\xfe\x1b\x00\x00\xff\xff\xf0\xcc\x12\x81\x28\x27\x00\x00")

func templates_app_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_app_tmpl,
		"templates/app.tmpl",
	)
}

func templates_app_tmpl() (*asset, error) {
	bytes, err := templates_app_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/app.tmpl", size: 10024, mode: os.FileMode(420), modTime: time.Unix(1442322257, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_service_mysql_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\xdd\x6f\xda\x3a\x14\x7f\xe7\xaf\xb0\xfc\x74\xaf\xc4\xe5\x52\x2a\x5d\xe9\x46\xd3\x24\x0a\xb4\x8a\xb4\x76\xa8\xb0\xee\x61\xea\x83\xb1\x0f\xc8\x5a\xb0\x3d\xdb\xe9\xc4\x2a\xfe\xf7\xd9\x21\x09\x36\x49\xe8\x87\xb4\x7e\xa0\x90\x73\x7e\xe7\x77\xbe\xed\xe7\x67\xc4\x60\xcd\x05\x20\x6c\x40\x3f\x71\x0a\x18\xed\xf7\x3d\x84\x9e\xdd\x3f\x42\x78\xfc\x75\xb1\x84\xad\xca\x88\x85\x6b\xa9\xb7\xc4\x3e\x80\x36\x5c\x0a\x8c\x12\x84\x47\xc3\x8b\xe1\x3f\xc3\xff\xdd\x1f\xee\x1f\xd4\xe7\x44\x93\x2d\x58\xa7\x83\x93\xd2\x84\x37\x92\x65\x92\x3a\x0b\x6c\x61\xa5\x26\x1b\x08\x64\x4e\xba\xdc\x29\x28\xcc\xdd\xe5\xdb\x15\xe8\xd2\x54\x21\x9a\xc2\x9a\xe4\x99\x2d\xa4\x17\xc3\x58\x62\xa8\xe6\xca\x56\xae\xd4\x14\xc8\x1c\x38\x90\xe1\xbf\x00\xfd\x75\x73\xf5\x37\x2e\x51\xfb\x0a\x8e\xa7\xc4\x92\x15\x31\x5d\x7e\x2c\xac\xe6\x62\xd3\xe5\x07\x51\xea\x9c\x23\xa5\x2a\x62\x25\x07\x12\x2e\x21\x4d\x17\x52\x61\x2c\x11\x14\x0a\xd2\xf7\xb8\xc1\x56\x03\x3b\x1a\x6c\x39\xd5\xf2\x9c\x3b\x15\x0f\xa2\x19\x31\x06\xad\xa5\x0e\x3c\x93\x0c\x4c\xd3\xb5\xb9\x53\xfc\x29\x35\x7b\x83\x5b\x31\xe7\xc2\x35\x12\x68\xa4\x2a\x3b\x0d\x86\x45\xbe\x12\x60\x4d\x0b\x81\x43\x7f\xe2\xc6\x7e\x70\x6d\x97\x24\xb3\xc9\x28\x49\x0e\xba\x49\x92\xb2\x8f\x5d\x9c\x0e\xf4\x30\x9f\x20\x53\x5a\x6d\xd0\x7d\x71\x9d\x5d\x54\xe1\x0f\x94\xbb\x8c\x35\xaf\x28\x1a\xe4\x0f\x8a\xb6\xc7\x79\x0c\xd1\x39\xef\xe3\x3b\x1f\x5e\x6d\xb9\x17\xd8\xc7\x9f\x73\xab\xf2\x28\x93\x78\x2e\xb5\xbd\xbc\x1c\xfe\xb7\xa4\x6a\xcc\x98\xf6\x22\x67\x80\x64\x39\x1c\x1e\xaf\x45\x92\xdc\x80\x1d\x5b\xeb\xbe\x7f\x3b\x76\x08\xee\x23\x3c\x13\x4c\x49\x2e\xec\xc0\x23\xc1\x18\x8c\x1e\xd1\x3e\x6c\x8d\xa3\x6d\xff\xf8\x3e\xdb\x05\xf2\xc4\xf0\x4c\x3c\xdd\xee\xcc\x8f\x2c\x9c\xcc\xc8\xf2\x3d\xac\x7d\x22\x6a\x79\x2b\x3a\xec\xdc\x36\x74\x2d\x6f\x45\x87\x6d\xd2\x86\xae\xe5\x1e\x1d\x55\xe1\x1e\x8c\xcc\x35\x85\xa8\x0e\x0b\xa0\xb9\xe6\x76\x77\xa3\x65\xae\x5e\x6a\x81\x58\x39\x68\x84\xb9\x96\x0a\xb4\xe5\x10\x4f\x8b\x93\x14\xaa\x27\x7d\xb2\xf5\x71\xa0\x6a\x91\xf7\x43\xf5\x88\x21\x15\x9b\xa2\xbc\xae\x48\x81\x0e\xf2\xc1\xa6\xca\x51\x5a\x49\x65\xe6\x0d\x5a\xaa\x7c\xed\xae\xb5\xdc\x96\x05\xc7\xbe\xfe\xfe\xdd\x52\x9e\xbe\x99\x70\xa6\x53\x1f\x2a\x1e\x0e\x8a\xdf\x7f\x87\xb8\xcc\xd4\xe1\xe7\x31\xf2\xc8\x4d\x46\xca\xa2\x0c\xfb\x59\x09\x00\xfb\x8e\xc5\xf1\x52\x46\xef\xa7\xee\x63\x7a\x15\x2a\xbf\x2a\xa3\x11\xe4\x0d\x99\x2d\x40\x29\x33\x51\x2c\xd5\x8e\x3b\x1b\x4f\x3d\x1f\x2f\x06\x73\x9c\xa4\xd7\x44\xd2\x76\xe0\xd6\x9e\x35\x84\x47\x87\xca\x34\x54\x64\x13\x7f\x66\x44\xd8\xe8\xd8\xea\xc4\xa5\x0c\x84\xe5\x6b\x0e\x3a\x26\xf6\xf1\x2c\x2c\xa1\xdf\xef\x0e\x63\x74\x02\xbf\xab\x87\xaf\x39\xed\xfd\xce\x42\x35\x50\x61\xdd\x4f\x80\x33\xb1\x71\x57\x9d\xba\x9e\x71\x1d\x6f\x89\x71\x17\x97\x78\x0b\x34\x47\xbf\x03\x12\xaf\x9e\xe6\xbe\x89\x60\xf1\xe0\x44\x92\x7c\x95\x71\x9a\xed\xc6\xd4\x6d\x13\xc3\x57\x59\xe1\xec\x9a\x64\xe6\xb4\xe9\x0e\xb5\xab\x5a\x65\xa3\x46\xb1\xdc\x9d\x18\xd1\xc4\x17\xb3\x1e\x26\x29\x5a\x38\x6e\xa3\x3d\x36\xdb\xb4\x57\x7d\xee\x7b\xee\x9a\x08\x82\xf9\x9b\xe1\xef\x00\x00\x00\xff\xff\x17\x8f\xc1\x85\x31\x0a\x00\x00")

func templates_service_mysql_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_service_mysql_tmpl,
		"templates/service/mysql.tmpl",
	)
}

func templates_service_mysql_tmpl() (*asset, error) {
	bytes, err := templates_service_mysql_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/service/mysql.tmpl", size: 2609, mode: os.FileMode(420), modTime: time.Unix(1442615805, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_service_postgres_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\xdd\x6e\xda\x30\x14\xbe\xe7\x29\x2c\x5f\x6d\x12\x63\x94\x6d\x17\x8b\xa6\x49\x14\x68\x15\x69\xeb\x50\x61\xdd\xc5\xd4\x0b\x63\x1f\x90\xb5\x60\x5b\xb6\xd3\xa9\xab\x78\xf7\xd9\x21\x09\x76\x42\xa0\xff\x28\xe4\x9c\xef\x3b\xff\xc7\x7e\x7a\x42\x0c\xd6\x5c\x00\xc2\x06\xf4\x03\xa7\x80\xd1\x6e\xd7\x43\xe8\xc9\xfd\x23\x84\xc7\xbf\x16\x4b\xd8\xaa\x8c\x58\xb8\x92\x7a\x4b\xec\x1d\x68\xc3\xa5\xc0\x28\x41\x78\x34\xbc\x18\xbe\x1b\x7e\x76\x7f\xb8\xbf\x57\x9f\x13\x4d\xb6\x60\x9d\x0e\x4e\x4a\x0a\x4f\x92\x65\x92\x3a\x06\xb6\xb0\x52\x93\x0d\x04\x32\x27\x5d\x3e\x2a\x28\xe8\x6e\xf2\xed\x0a\x74\x49\x55\x88\xa6\xb0\x26\x79\x66\x0b\xe9\xc5\x30\x96\x18\xaa\xb9\xb2\x95\x2b\xb5\x09\x64\xf6\x36\x90\xe1\xff\x00\xbd\xb9\xbe\x7c\x8b\x4b\xd4\xae\x82\xe3\x29\xb1\x64\x45\x4c\x97\x1f\x0b\xab\xb9\xd8\x74\xf9\x41\x94\x3a\xe5\x48\xa9\x8a\x58\x69\x03\x09\x97\x90\xb6\x0b\xa9\x30\x96\x08\x0a\x85\xd1\xd7\xb8\xc1\x56\x03\x3b\x1a\x6c\x39\xd5\xf2\x94\x3b\x95\x1d\x44\x33\x62\x0c\x5a\x4b\x1d\x78\x26\x19\x98\xb6\x6b\x73\xa7\xf8\x57\x6a\xf6\x02\xb7\x62\x9b\x0b\xd7\x48\xa0\x91\xaa\x78\x5a\x16\x16\xf9\x4a\x80\x35\x47\x0c\x38\xf4\x37\x6e\xec\x17\xd7\x76\x49\x32\x9b\x8c\x92\x64\xaf\x9b\x24\x29\xfb\xda\x65\xd3\x81\xee\xe6\x13\x64\x4a\xd6\x96\xb9\x9f\xae\xb3\x8b\x2a\xbc\x26\xcf\x4a\x1a\xbb\xd1\x2e\x4f\xe7\x03\xce\x2b\x3b\x2d\x0f\xee\x14\x3d\x1e\xec\x21\x4e\x17\x81\x0f\xf2\x74\x8c\x35\x73\x2f\xe0\xc7\x3f\x72\xab\xf2\x28\x9d\x78\x2e\xb5\xfd\xf4\xf1\xc3\x68\x49\xd5\x98\x31\xed\x45\x8e\x80\x64\x39\xec\x1f\xaf\x44\x92\x5c\x83\x1d\x5b\xeb\xbe\xff\x3e\xb4\x09\xee\x23\x3c\x13\x4c\x49\x2e\xec\xc0\x23\xc1\x18\x8c\xee\xd1\x2e\xec\x8f\x03\xb7\x7f\x7c\x1d\x77\x81\x6c\x10\xcf\xc4\xc3\xbc\xcc\x76\x38\xa1\x11\xf9\x2d\xac\x7d\x2e\x6a\x79\x17\x41\xd8\xc4\xc7\x08\x6a\x79\x17\x41\xd8\x34\xc7\x08\x6a\xb9\x27\x88\xca\x71\x0b\x46\xe6\x9a\x42\x54\x90\x05\xd0\x5c\x73\xfb\x78\xad\x65\xae\xce\xf5\x42\xac\x1c\x74\xc4\x5c\x4b\x05\xda\x72\x88\x67\xc7\x49\x0a\xd5\x46\xc3\x54\xad\x8b\xaa\xcd\xde\x0f\x11\x91\x91\x54\x6c\x8a\x52\xbb\x82\x05\x3a\xc8\xc7\x9b\x2a\x67\xd5\x4a\x2a\x33\xcf\x69\xa9\xf2\x75\xbc\xd2\x72\x5b\x16\x1f\xfb\x5e\xf0\xef\x96\xb2\xf9\x66\xc2\x99\x4e\x7d\xb4\x78\x38\x28\x7e\xdf\x0f\x71\x99\xac\xfd\xcf\x7d\xe4\x91\x9b\x92\x94\x45\x49\xf6\x73\x13\x00\x76\x1d\x9b\xe4\x5c\x52\x6f\xa7\xee\x63\x7a\x19\x2a\x3f\x2b\xa9\x11\xe4\x65\xc9\x2d\x70\x29\x33\x51\x38\xd5\xde\x3b\x19\x52\x3d\x2e\x67\xe3\x39\x0c\xd6\x73\x82\x39\x76\x08\xd7\x9e\xb5\x84\x07\x87\xca\x4c\x54\xc6\x26\xfe\x1c\x89\xb0\xd1\x51\xd6\x89\x4b\x19\x08\xcb\xd7\x1c\x74\x6c\xd8\xc7\xb3\xb0\x84\xfe\xb9\xd9\x0f\x53\x03\x7e\x53\x8f\x60\x7b\xf2\xfb\x9d\xb5\x6a\xa1\xc2\xd2\x37\x80\x33\xb1\x71\xd7\x1f\x7c\x7c\xd5\x3b\xf9\x77\x62\xdc\x7d\x26\x5e\x07\xed\x1d\xd0\x01\x89\xd7\x50\x7b\xf7\x44\xb0\x78\x7c\x22\x49\xbe\xca\x38\xcd\x1e\xc7\xd4\xad\x15\xc3\x57\x59\xe1\xef\x9a\x64\xa6\xd9\x77\xfb\xf2\x55\xdd\xb2\x51\x0d\x22\x77\x86\x44\x73\x5f\x4c\x7c\x98\xa7\x68\xf3\xb8\xd5\x76\xdf\xee\xd4\x5e\xf5\xb9\xeb\xb9\xdb\x23\x08\xe6\x2f\x8c\xff\x03\x00\x00\xff\xff\xde\x67\x65\x1d\x48\x0a\x00\x00")

func templates_service_postgres_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_service_postgres_tmpl,
		"templates/service/postgres.tmpl",
	)
}

func templates_service_postgres_tmpl() (*asset, error) {
	bytes, err := templates_service_postgres_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/service/postgres.tmpl", size: 2632, mode: os.FileMode(420), modTime: time.Unix(1442322257, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_service_redis_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xd1\x6f\xda\x3e\x10\x7e\xe7\xaf\xb0\xfc\xf4\xfb\x49\x8c\x41\x27\x75\x6a\x34\x4d\x42\x0c\xaa\x48\x5b\x87\x80\x76\x0f\x55\x1f\x8c\x7d\x50\x6b\x89\x1d\xd9\x4e\xb7\xaa\xe2\x7f\xdf\xd9\x09\x10\x27\x1d\x6d\xa5\xb5\x80\x22\xdf\xdd\xf7\xdd\x7d\x77\xb6\xf3\xf4\x44\x04\x6c\xa4\x02\x42\x2d\x98\x07\xc9\x81\x92\xdd\xae\x47\xc8\x13\x7e\x09\xa1\xe3\x1f\xcb\x15\xe4\x45\xc6\x1c\xcc\xb4\xc9\x99\xbb\x01\x63\xa5\x56\x94\x24\x84\x9e\x0d\x47\xc3\x77\xc3\x0b\xfc\xd0\x7e\xe5\x3e\x67\x86\xe5\xe0\xd0\x87\x26\x35\x04\xae\x7e\x61\x8e\xad\x99\x85\xc6\x1a\xae\xae\x1e\x0b\x08\x30\x4b\x67\xa4\xda\xd6\x10\x55\x00\x6c\x58\x99\xb9\x60\x1d\xc6\x06\xcb\x8d\x2c\xdc\x3e\x83\xda\x91\x88\x9a\x81\x48\x25\xe0\x37\xad\x03\x76\xfb\x48\x9a\x2a\xeb\x98\xe2\x10\x38\xbb\x59\x9c\x4c\x02\x8d\x9c\xf1\x7b\x18\xb8\xb3\x41\x2e\xb9\xd1\x7f\x4b\x08\x1d\x57\xf7\x40\x1c\x22\x12\xbd\xc1\x54\x2a\x4e\xe2\x34\x29\xb1\xf8\x4e\x52\x73\x66\xed\x2f\x6d\xc4\x1b\x64\x89\xab\xbf\x56\x88\x2b\xc8\x7f\x48\xb0\x06\x62\x20\xd7\x0f\x20\xfe\xef\x12\x2d\xcb\xb5\x02\x67\x9f\x2f\xfc\xab\xb4\xee\x13\xb6\x39\x49\xa6\x93\xb3\x24\xa9\x7c\x93\x24\x15\x9f\x4f\xd4\x79\x33\x9f\x10\x5b\xa3\x76\xe8\x6e\x0a\xfe\x3c\xd5\x91\x05\xe3\x3d\xc5\x69\x86\x03\x72\xaf\x81\x4f\xbf\x97\xae\x28\xa3\x62\xe8\x5c\x1b\x77\xfe\xe1\xe3\xc5\x8a\x17\x63\x21\x8c\x37\x21\x00\xcb\x4a\xa8\x1e\x67\x2a\x49\x2e\xc1\x8d\x9d\x6f\xe6\x2d\xa1\x0b\x28\x32\xc9\x99\xa7\xba\x34\xba\x2c\x68\x1f\x31\x8c\xcc\x99\x79\x9c\x2a\x31\xd7\x52\xb9\x81\x07\x02\x6b\x29\xb9\x23\xbb\x66\xcf\x8e\x54\xfe\xf1\x9f\x50\x05\xa0\x16\xcf\x54\x3d\x2c\x40\x48\xdb\xdc\x3a\x11\xd1\x02\x36\x5e\xa6\x83\xdd\x47\x47\x3a\x2d\xc0\xea\xd2\x70\x88\x94\x5a\x02\x2f\x8d\x74\x8f\x55\x2e\x2f\x34\x29\x76\x6e\xb4\x6a\x6e\x74\x01\xc6\x49\x88\x47\x0a\x2d\xc1\xb5\xd5\x49\xe3\xeb\x20\xfb\xd3\xa5\xdf\x74\x8f\x18\x52\xb5\x0d\x8a\xa3\x6e\x0d\x1f\xe2\x8b\x4d\x0b\xa4\x74\x9a\xeb\xcc\x03\x3a\x1e\x64\x9c\x19\x9d\xd7\x3d\xa0\xbe\x25\x7e\x6d\xa5\xdb\x2b\x13\x29\x4c\xea\x4b\xa5\xa3\xe1\x20\xfc\xbf\x1f\x9d\xd3\x5a\xab\xea\xef\x2e\xca\x09\xa7\x37\x15\x91\xc6\x7e\x9e\x1b\x01\xbb\xce\xc0\x4f\xfc\xf1\x50\x6d\x9c\x17\x85\xcd\x98\x75\x32\x04\xec\xf7\xda\x1b\xe4\x6d\x29\xbb\xa8\x94\x0d\x28\x64\xdb\x82\x39\xec\xfb\x54\xd8\xa8\x9c\xfd\x69\x70\xb2\xa4\xce\xdc\xbe\xb2\xa4\xee\xbc\xbf\xa6\xae\x71\xe9\x34\xde\x2d\x92\xcf\x98\xcc\xf0\x10\x33\x53\xc5\xd6\x19\xf8\x36\x6c\x58\x66\xa1\xdf\x76\xfe\x26\x95\x36\xf5\x55\x74\x5d\x6c\x0d\x13\x3e\x1f\x67\xca\xd8\x35\x24\x75\xa5\xc5\xe1\xdc\x3f\x88\x10\x5d\x08\xc7\xba\x9f\x6d\xe7\x15\x5e\x69\x51\x70\xa7\xdf\x2d\x80\xa9\xda\xe2\x8d\x7a\x98\xfd\xb8\x29\x57\x65\x1e\xe2\x27\x59\x69\xeb\x8b\x92\x8e\x62\x97\x78\x86\x9b\x96\xb6\xbc\xf1\x40\x1c\x33\x0c\x8d\x59\x3a\xc6\x7f\x86\xe4\x5b\xf9\xc5\xdb\x2e\x8c\xc7\x6d\x73\x40\xa2\x7d\x8f\x07\xcb\x5d\x77\x50\x7a\xfb\xdf\x5d\x0f\x5f\x21\x40\x09\xff\xd6\xf0\x27\x00\x00\xff\xff\x64\xe5\x91\xd9\x4d\x08\x00\x00")

func templates_service_redis_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_service_redis_tmpl,
		"templates/service/redis.tmpl",
	)
}

func templates_service_redis_tmpl() (*asset, error) {
	bytes, err := templates_service_redis_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/service/redis.tmpl", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1442322257, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"templates/app.tmpl": templates_app_tmpl,
	"templates/service/mysql.tmpl": templates_service_mysql_tmpl,
	"templates/service/postgres.tmpl": templates_service_postgres_tmpl,
	"templates/service/redis.tmpl": templates_service_redis_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"app.tmpl": &_bintree_t{templates_app_tmpl, map[string]*_bintree_t{
		}},
		"service": &_bintree_t{nil, map[string]*_bintree_t{
			"mysql.tmpl": &_bintree_t{templates_service_mysql_tmpl, map[string]*_bintree_t{
			}},
			"postgres.tmpl": &_bintree_t{templates_service_postgres_tmpl, map[string]*_bintree_t{
			}},
			"redis.tmpl": &_bintree_t{templates_service_redis_tmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

