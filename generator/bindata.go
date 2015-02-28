package generator

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

var _templates_model_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xc1\x4e\xc4\x20\x10\x86\xef\x3c\xc5\xa4\x59\x13\x4d\xb4\x7b\xdf\xc4\x93\x7a\xf0\x62\x4c\xdc\x07\xd8\xb1\x9d\xb6\xb8\x05\x2a\x4c\x77\xd3\x10\xde\x5d\xa0\x6d\xd6\xf5\xe2\x0d\xca\xf7\x4f\x7f\x3e\xbc\xaf\xa9\x91\x9a\xa0\x50\xa6\xa6\x7e\xb0\x66\x20\xcb\x53\x11\x82\xf0\x5e\x36\x50\x3e\x9b\xea\x83\xad\xd4\x6d\x08\xde\x5f\xef\x48\xd7\x19\x2b\xdf\x97\xd4\x1b\x2a\x0a\x01\x12\x87\x8c\xfb\x69\x48\xbb\xc3\x97\x33\x7a\x57\x24\x0c\x2d\xaa\x99\x29\x0e\x62\xcd\x8b\x01\xab\x23\xb6\x04\x99\xc8\xcb\xf4\x75\xbb\x85\x7d\x27\x1d\x34\xb2\x27\x38\xa3\x83\x96\x34\x59\x64\xaa\xe1\x73\x02\xee\x08\xdc\x19\xdb\x96\x2c\xb0\x31\x7d\x99\xf8\x97\x5a\x72\xac\x16\x0f\xd7\x9c\x92\x6d\xc7\x10\x2f\x75\x22\x68\x46\xce\xa3\x3a\xd2\x30\x99\x11\x2c\x3d\xd8\x51\x5f\x4d\x5a\x7f\x01\x95\x51\x0a\x75\x2d\x16\x07\xaf\x6a\x30\x96\x5d\x08\x72\x5e\xc0\xad\x80\x58\xd7\xa2\x8e\xb5\x37\x47\x9a\xee\x61\x73\xc2\x7e\x24\xd8\x3d\xfe\xa2\x33\x94\x8e\xb3\x93\x58\x43\x6a\x6e\xa0\xb8\xf9\x2e\x16\x7c\x41\x66\x0f\x77\x17\x23\xff\x98\x5f\x41\x8e\x82\x93\xb5\xa7\x1e\x9d\x5b\xdc\x3b\xb6\x63\xc5\xe0\xc5\x5a\x6f\x7d\x1c\x49\x2e\x27\x99\xd4\xd0\xa7\x3b\xfe\x79\x70\x28\x2f\x83\x45\x10\x3f\x01\x00\x00\xff\xff\x80\x30\xad\xe3\x18\x02\x00\x00")

func templates_model_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_model_gotmpl,
		"templates/model.gotmpl",
	)
}

func templates_model_gotmpl() (*asset, error) {
	bytes, err := templates_model_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/model.gotmpl", size: 536, mode: os.FileMode(420), modTime: time.Unix(1424678182, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_modelvalidator_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\xb8\x09\x19\x60\xaf\x89\xdc\x01\xc3\x1e\xd6\x65\x40\xd1\x79\x68\x80\x2e\x0d\x9a\x6e\x2f\xc3\x80\x32\x32\x25\xb3\x91\x48\x85\xa4\x9c\x78\x82\xfe\xfb\x8e\xa4\x28\x5b\xb2\xe4\x58\x0e\x06\xec\x49\xd2\xdd\xf1\xf8\xf1\xbb\xbb\x8f\x2a\xcb\x25\x8d\x19\xa7\x10\xe4\x92\x65\x4c\xb3\x35\x5d\x93\x94\x2d\x89\x16\x32\xa8\xaa\x49\x59\xb2\x18\xc2\x4f\xf4\xa1\x60\x92\x2e\xd1\x80\x9f\x54\x4a\xf8\xe9\x12\xea\x38\xda\x78\xa7\x65\x19\xde\x10\xbd\xaa\xaa\x73\x08\xf0\xfd\x83\x88\x88\x66\x82\x57\x55\x70\x0e\xf8\xfd\x27\x49\x0b\xba\x78\xca\x25\x55\xca\x9a\x67\x6f\x6c\xae\x6f\x2e\x81\xb3\x14\xca\x09\x80\xa4\xba\x90\xdc\x58\x27\x66\x6f\xca\x97\x0d\x86\xdf\x19\xff\x40\x79\x62\xd2\xf7\x81\x68\xdc\xa3\x51\x58\xeb\x4e\xf6\x71\xa8\xc8\xd3\x41\x54\xde\x7d\x22\xaa\x6d\xf6\x51\xa8\x70\x27\x4d\x25\xef\xc7\x54\x3b\x4f\x40\xf4\xc5\x2d\x71\xa9\xbf\x8c\xad\x1e\xcb\x8a\x6c\xb0\x76\xc6\x79\x10\x51\x9c\x0a\xa2\x7f\xfc\x61\xda\xdb\x47\xbe\x84\x6e\x0b\xfb\xb5\x78\x8a\xd2\x42\x61\x3b\x37\xe6\xb1\x75\x3d\x80\xd7\x39\x5f\x8a\xd7\x6f\xd1\xc1\xeb\xcd\xe3\xf0\x16\xa9\x66\x79\x4a\x3f\xc6\x03\x90\x1b\xff\x4b\x51\xef\x6c\x34\x0a\xe1\x82\x0f\xd1\x69\x3c\xa7\xcd\x87\xcb\x79\x34\x0c\xff\xf4\x92\x17\x15\x4a\x8b\x2c\x16\x32\x23\xba\xa5\x7a\x3d\x20\x7f\xb3\x51\xcf\xd0\x67\x0c\x2e\xd0\x7e\x2a\x2d\x19\x4f\x86\xc8\x74\xfb\xaa\xa3\xd1\x7b\xd4\x2a\x65\x51\x9f\x48\x5f\x53\xba\x54\xb7\xec\x1f\x6a\x2d\x08\x52\x92\xec\x9a\x64\xf8\x69\x8c\xe6\x30\x8c\x9b\xda\xa6\x94\xf7\x43\x9a\xed\xcf\xec\x95\xa6\x99\x1a\x1c\x5a\xeb\x7d\xae\x72\x1d\x1c\x7e\x54\xeb\xcc\x63\x87\xf2\x10\xa0\xda\x7b\x12\xa0\x26\xf3\x28\x40\x7f\x70\xf6\x50\xd0\x03\x98\x76\x02\xfe\xdb\xdb\xf1\x7f\x30\x5d\x06\xc6\x2d\xf6\x7b\x4a\x6f\xa3\x15\xcd\xc8\xad\xe9\x53\x40\xd7\x7c\x0e\xca\xda\x41\x59\x47\xef\x8e\x13\x1c\x07\x60\x06\xf9\xeb\x37\xf8\xfc\x19\x06\xdb\x14\xdd\xaf\x5e\x21\x90\xb2\x94\x84\x27\x14\x42\xcf\x3f\x60\x62\x7c\xcd\x53\x3c\xb6\xf9\x9f\x11\x39\x95\x7a\xb3\x9d\x14\x08\x77\x54\xc0\xbe\xa5\x8a\x3a\x7c\x5c\xe8\x7d\x8c\x37\x75\x06\xd7\x2b\x2f\xdc\xcf\xf1\xf3\x76\xb9\x64\x86\x78\x92\x6e\x93\x34\x07\xc7\x2d\xad\x15\xaf\xfc\xaa\x32\x24\x20\x0b\x76\x5a\x67\x70\xd1\x76\x1a\xc3\xf7\x26\xc2\x12\x01\x70\x14\x12\x80\x9d\x33\x23\x98\x41\x82\xe1\x97\xf6\x6e\x9d\xa2\x0b\xa9\xba\xe7\xb8\x16\xfa\x6d\x9a\x8a\x47\xfc\x07\x0c\xfa\x52\x06\x7b\x6d\x37\xeb\x15\xe6\xae\xd4\x89\xbb\xaf\x34\x6a\x4b\x33\x16\xcb\xc9\x36\x38\xa7\x81\xfa\x2b\xd1\xe4\xf3\x26\xa7\xad\x01\xe8\xc3\x61\x2c\x76\x2a\xa6\xa7\x8a\xef\x3e\xb7\x4d\x6d\xaf\xd4\x8d\xff\x85\xae\xaa\x76\x3d\xf6\xfe\xac\x7d\x6b\x60\x31\xc0\xad\x7d\x67\x0f\xe5\x2e\x0f\xfc\xc1\x6a\x67\xe8\xbf\xa8\x7a\x92\x08\xae\x09\xe2\xec\x2c\xef\xdc\x18\x7d\xeb\x30\x94\x3e\x7d\xb4\x8c\xb6\xd7\x76\x4b\x60\x16\x77\x0b\x96\x93\xe8\x9e\xe0\x64\x58\x95\xb1\xaf\x68\x34\x95\xfa\xbc\x62\x0a\x62\x86\x53\xf5\x48\x14\x24\x14\x91\x61\xd2\x25\xdc\x6d\x40\xaf\x70\xd4\x1e\x49\x92\x50\x09\x5a\x88\x34\x34\xf1\x0b\xd3\x55\x3c\x41\xa7\x5f\x97\xb1\x64\xa5\x01\x59\x5f\x53\x88\x0b\x6d\x53\xad\x28\x87\x8d\x28\xb0\x54\x17\xb2\xe0\xad\x4c\x7e\x0b\x88\x44\x96\x11\xbe\x9c\x4c\x58\x96\x0b\xa9\x61\x8a\xa5\x0d\x12\xa6\x57\xc5\x5d\x88\xbe\x79\x44\x54\x41\xd2\xaf\x2c\x9b\x27\xe2\xa2\x5e\x3d\x77\xcd\x1d\x1c\x13\x8a\x57\x7b\x9c\xe9\xa3\x42\xbd\x12\x07\x13\xab\x18\x4e\x44\xce\xee\xe9\xe6\x1c\xce\xd6\xa6\x43\x4d\xb7\x86\x57\x16\xa7\x97\x15\xe3\x36\x83\x57\xe2\xd1\xf1\xd6\x8e\x21\xf8\xf6\x21\xa8\xc3\xeb\x10\x47\xfd\xcc\xf2\xec\xdb\xba\x51\x7d\xe5\x28\xc4\x82\xbc\x2f\x90\x87\x77\x29\x51\xaa\x96\xb1\xb8\xe0\x11\x98\xc1\xff\x44\x23\x8a\x4d\x29\x9d\x1d\xbe\x43\xd3\x4e\xdc\x0c\xba\xb3\x02\xee\xcc\xb8\x2e\x61\xf8\xba\x99\x39\x31\xb0\x53\xe3\x26\xe0\x3d\x51\xf5\x22\x9c\x36\x77\x94\x35\x91\x58\x27\x05\x7f\xfd\x6d\x83\x77\x29\xf0\xfa\xca\xa8\x3f\xf5\x40\x8e\xd6\x4c\xb7\x51\x87\xfe\xc0\x7b\x72\x3d\x3c\xe1\x60\x01\x5d\x02\xc9\x73\xe4\x70\x8a\x1f\xe7\x26\x64\x66\x05\x72\xcb\xec\xf6\xcd\x41\x30\x72\x89\xb1\x46\x1d\x5f\x37\x79\x76\x35\xd1\x0c\x91\x50\x4c\xd3\x2d\xfe\x85\xf1\x98\x55\x61\x18\xee\xe7\xaf\x97\x23\x2e\x2b\x35\x70\x16\x79\xfa\x6d\x47\x34\xc5\x80\xdd\xcb\xa7\x45\xda\x00\x65\x87\x6a\xbc\xdd\xc4\xd4\xf8\x59\xfa\x0e\x17\xfd\xd9\x3b\x67\xff\x94\x6d\xe5\xf8\x37\x00\x00\xff\xff\x97\xa4\x77\x70\x7f\x10\x00\x00")

func templates_modelvalidator_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_modelvalidator_gotmpl,
		"templates/modelvalidator.gotmpl",
	)
}

func templates_modelvalidator_gotmpl() (*asset, error) {
	bytes, err := templates_modelvalidator_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/modelvalidator.gotmpl", size: 4223, mode: os.FileMode(420), modTime: time.Unix(1425156718, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_server_builder_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\x6d\x6f\xdb\xb6\x13\x7f\xfd\xf7\xa7\x38\x18\x2d\x20\x15\xaa\xf2\x7f\x1d\x20\x03\xb2\xa6\x43\xb3\x61\x5d\x91\x14\xdb\x8b\x20\x18\x18\xe9\x6c\x73\xd1\x53\x49\x2a\x99\x17\xf8\xbb\xef\xf8\x24\x51\x96\xec\xda\x4d\x8a\x19\x01\x22\x91\xc7\x7b\xbe\xdf\x1d\xd5\xb0\xec\x9e\x2d\x11\x9e\x9e\xd2\x4f\xf6\x71\xb3\x99\xcd\x4e\x4e\xe0\xf3\x8a\x4b\x58\xf0\x02\xe1\x91\x49\x58\x62\x85\x82\x29\xcc\xe1\x6e\x0d\x6a\x85\x20\x1f\xd9\x72\x89\x02\x54\x5d\x17\xa9\xa6\x7f\x9f\x73\xc5\xab\x25\x6d\xfa\x73\x25\x5f\xae\x14\x34\xa2\x7e\x40\x58\xb4\xca\xb0\x5a\x61\x05\xeb\xba\x05\x81\x6f\x45\x5b\x0d\x38\x79\x11\x90\xd5\x65\xc9\xaa\x7c\x36\xe3\x65\x53\x0b\x05\xd1\x0c\x60\x2e\x95\x20\xee\x72\xae\x9f\x2b\x54\x27\x2b\xa5\x1a\xf3\xb2\xe4\x6a\xd5\xde\xa5\x74\xe8\x24\x63\xb2\x65\xc5\x5f\xbc\x3c\x59\xd6\x6f\x1d\xdb\x43\x68\x4e\x64\x83\xd9\x61\x84\x4a\x2c\x4a\x75\x10\x69\xc9\xf3\xbc\xc0\x47\x26\xf0\x48\xf2\x13\x89\x59\x2b\xb8\x5a\xcf\x67\x74\xf0\xe9\x49\xb0\x8a\xe2\xf3\xea\x1e\xd7\x09\xbc\x7a\x60\x45\x8b\x70\x7a\x06\xe9\xa5\x71\x8e\xa4\x68\x69\x22\xbd\xbd\xd9\xd0\x03\xf9\x9b\x57\x6a\x01\xf3\xd7\x5f\xe6\x8e\xdc\x91\x60\x95\xd3\x53\x6c\x82\x4b\xd1\x3e\x6f\x9a\x8f\xac\xa4\xcd\xf3\x4f\x97\xfa\xfd\xb2\x5a\xd4\xe9\x05\xca\x4c\xf0\x46\xf1\xba\x22\x5a\xb5\x6e\x70\x44\x4a\x2e\x68\x33\x05\x4f\xc4\x53\xbb\x0d\xfa\xdf\x1b\xfd\x9e\x5e\xd4\x59\x5b\x62\xa5\x68\x3f\xab\x2b\x85\x7f\xab\x6e\xbf\xb7\x31\x7d\x67\xb7\x88\x68\x45\xa1\x2e\x50\x48\x47\x54\xb2\xe6\xc6\xc6\xfa\x56\x87\x38\xfd\x60\xb7\x89\x70\x51\x8b\x92\x29\x4f\x07\x36\x14\xe9\x15\x2e\x39\x3d\xae\x03\x57\x69\xe6\x92\x74\x20\xdf\x58\x53\xdf\x15\x4c\x4a\x6b\x81\xdb\x12\x94\x80\xfa\x9c\x96\xcb\xb4\x9e\x76\x91\x44\xd0\xeb\x9c\x8e\xfc\x8a\x39\x67\x9f\xc9\xfe\xcd\x66\x4e\x89\x5c\x22\x68\x67\x18\x21\x93\xec\x5c\x18\xbd\x68\x11\x78\x3c\x50\xec\x93\xa8\xf3\x36\x9b\x54\xcc\x6d\x0d\x15\x6b\xfc\xe2\xd1\x8a\x75\xec\xbc\x62\x7e\x61\x5a\xb1\x6b\x97\x72\x17\xb8\xe0\x15\xd7\xe1\xf7\x89\xc5\x17\x94\x6a\xf2\x47\x26\x79\x76\xde\xaa\xd5\x84\xe6\x7a\x79\xa0\xf5\xa2\xad\x32\xcd\x82\xaa\x9b\x29\x50\xec\x1e\x25\xb4\x12\x45\x45\xe4\x40\xe1\x84\x86\xce\x3e\xd6\x22\x37\x2f\x02\x55\x2b\x2a\x6b\x2d\xaf\x32\xde\xb0\x82\x04\x93\x14\x4e\xd8\x81\x42\x47\x9d\x36\x49\x06\xe5\x14\xcf\x98\x61\xfc\x48\xc5\x04\x77\x5a\x27\xb3\x33\xb2\xde\xa8\xa4\xd5\x88\x6c\x2a\x25\x60\xff\xc7\x10\x69\x9c\xf3\x72\x36\x9b\x04\x50\x88\x5a\xc4\xbd\x57\xbc\xc5\x94\xea\xbf\xe0\xfa\x39\x26\x33\xc2\xc6\x7b\x82\xbb\x6f\x35\x92\xec\x23\xb8\xad\x35\x03\x60\x0d\x07\xaa\x70\xad\x86\xd5\xc0\xc0\x2a\xcf\x89\x80\x5b\x14\xa5\x9d\xeb\xba\x15\x99\xaf\xf6\x7d\xee\x38\xc0\x0d\xd3\x69\xf2\x5b\xa3\x11\xda\x66\xc7\xc8\x29\xae\x52\x41\x22\x15\xa9\x56\xa9\xf6\xd4\xbe\xc6\x4d\x1a\x3b\x65\x3f\xb4\x84\xf1\xc1\xe9\x9e\xba\x4b\xbb\xae\x1f\x4d\xcb\x09\x3b\x56\x3a\x49\x62\x8d\x28\xe4\x3e\x16\xbb\x4e\x8d\x9c\x40\xf6\x5e\xa3\x78\xc0\xf7\xda\x53\x40\x3d\x2e\x63\x45\x41\xfe\x37\x2d\x8d\x42\x84\x7e\x5d\x60\x86\xfc\x01\xf3\x44\x9b\x2a\x50\x2f\x31\xc8\x71\xc1\xda\x42\x79\x4f\x58\x7e\x77\xad\x32\xcd\x30\xa3\xe3\xe4\x35\xfd\x2c\xa0\x7e\x74\xf9\xad\x1b\x29\xd1\x05\x42\xf5\xcf\x84\xd1\x20\xe3\x15\xca\x86\x22\x81\x7f\x50\xe1\xa2\x48\xe0\x8d\x5b\xfd\xd2\xa2\x54\x5d\x44\x6d\x33\xff\xc9\x61\xa7\xcf\x43\x1d\x04\x9f\xbf\x64\x83\x4d\x0b\x8f\xb0\x33\x2d\xc3\xa4\xc8\x95\x35\x45\xb8\x10\xbd\xd9\xea\x05\xb1\xe7\x1b\xc5\xdb\x70\x4c\xfd\xe1\x7f\x56\x18\x8c\xf8\xa4\x5e\x8e\xd5\xed\xca\xe9\x61\x79\x0d\x81\xb9\x95\xaa\x2e\x9d\x5e\x40\xcd\x8c\xe7\x4c\xd5\xe2\x08\x05\x87\xcc\x23\x03\x41\x1e\x13\x1c\x5b\xa7\xb9\xa5\x48\x7a\x29\x7e\xe3\x77\xbf\x10\x9b\xa6\xb7\xd3\x9c\xf4\x3c\xcf\x8d\x00\xcf\x39\xe0\xe5\xe3\xe0\x78\xa1\xdf\xc1\x30\x14\xae\xb4\x82\x8a\x0e\x6d\x39\xc2\x68\x2f\x85\xc2\x62\xb3\x52\xeb\xfd\xc0\x04\xb4\x55\x10\xf4\x9b\x5b\xeb\x88\xe9\xce\x49\xab\x54\x83\x63\x63\x77\xf4\xbf\xb3\x33\xa8\x78\x61\x04\xc1\x50\xcc\x19\x81\x57\x43\x45\x14\x85\xab\x89\xe9\x65\x13\x8c\xe6\x1a\x83\x76\xa0\x4f\xdf\x3d\x0f\x53\xae\xeb\x81\xcf\x55\xce\x33\xda\xa7\xdc\xae\x0e\x7a\x80\x9e\x06\x9e\x9f\xab\xa3\x66\xb2\x4f\xbf\x10\xba\x0f\x53\xcb\x83\xe4\x73\x35\x73\x7c\x46\xca\x59\x2d\x0a\xac\x06\xc7\x63\xf8\x01\xfe\xef\x84\x39\x00\xd1\x45\x68\x00\x70\x11\xcd\x4b\x2e\xa5\x86\xaa\xb0\x62\x4e\xe1\xb5\x9c\xfb\x16\x2f\xd3\x9f\x6b\x5e\x6d\x6b\x44\x7f\xb1\x95\x3f\xeb\xd8\x92\x51\x54\x95\x03\x58\x27\x0c\xa0\x6b\x88\xd2\xd0\x63\x0b\x27\x6c\x5c\x0c\x96\xe4\xab\x2a\x68\x6b\x3c\x3f\xa2\x28\x07\x52\xa2\x8e\xc9\xe5\x45\x37\x9b\x1c\x09\xed\xc6\x49\x3b\x31\xb6\x17\x67\x8d\x3c\xef\x87\x8b\x5a\xc8\xce\x50\x0d\x34\x6c\xb0\xd5\x35\x69\x3d\xcd\xf3\x05\xd7\xed\xc1\xe5\x36\xc8\x6c\x85\x84\x0e\x47\x58\x3d\x12\x1b\x39\x1e\xe1\xa0\x6f\xee\x0d\xbe\x80\xae\xcd\x7e\x3c\xd8\x77\x33\xec\x80\x99\x03\x63\x3d\x28\xec\xaa\x3d\x81\x52\xf7\x5c\xba\x2d\x95\x34\x91\x45\x5f\xe3\x18\xdb\x4b\x06\x58\x0c\xb7\x7a\xea\xc3\xb6\x82\xbc\xde\x36\x35\x25\xb5\xe8\x6c\x65\x48\xdd\xca\x01\x58\xa0\x7f\x74\xf9\x43\x53\x21\xd6\x49\xf3\xd3\x99\xbf\xd0\x4c\xcc\xda\xd6\x80\x1b\x2d\xe5\x96\xaa\xcd\xc7\x21\xed\x48\x22\x1b\x89\x36\x81\xa6\x1f\x71\xe9\xfe\x47\x23\x25\xcb\xf0\x69\xd3\xe7\xca\xee\x4c\x19\xe3\x88\xe1\x17\x6f\xe2\x1e\x46\x86\x1a\x86\xb3\xf1\x2e\x15\x7b\x1a\x17\x71\x63\xb0\x77\x2b\xdd\x36\x13\x9b\xef\x34\x23\xbf\xa4\xe6\xc4\x2e\x86\x6d\xcd\xc3\xb7\x8d\x43\x21\xc7\xd4\xaa\x1f\x0c\x8a\x43\x7c\xe8\xcf\xda\xfe\xed\xdb\xd4\xb0\x80\xfc\xfd\x71\xaa\x76\x4a\x7d\x5f\x33\x77\xb4\x63\xca\x26\x94\x13\x95\xfe\xca\x27\xbb\x9e\x3d\x59\x1e\x5d\x33\xee\x2b\x63\xd0\xcf\xbf\x5e\x0e\x9e\x83\xaf\x84\x3f\x13\x28\x55\x5f\x02\x81\x22\x83\x2a\x28\xd5\xb8\x06\x06\x92\x07\x3b\xe7\x45\x41\xe0\xc4\x69\x46\xf9\x87\x0c\x1c\x17\x46\x78\xc3\xed\xab\xc3\xe5\xd9\x36\x81\xce\xb9\x43\x87\x94\x89\x6c\x78\xc9\xdc\xf0\x53\xc2\x30\x37\xfc\x15\xfe\xe5\x72\x23\x94\x73\x70\x6e\x74\xb3\x90\xcf\x8d\xe1\x34\xf5\xf5\xd4\xf0\x0c\x5e\x20\x35\x06\x92\xff\xdb\xd4\x08\xbe\x8a\x7c\xcf\xd4\x70\x23\x50\x30\x5e\x84\xdf\xb6\xba\xcc\xe8\xee\xf4\xdf\x38\x62\xf4\x62\x26\xe7\x8b\x28\x14\x9a\xc0\x5d\x5d\x17\x76\x88\x70\x63\xd8\xd8\x59\xfe\xdb\x5c\xac\x87\x40\x3f\x95\x8d\xc9\x74\x93\x73\x6c\xdf\x31\xc2\xf7\xc8\x8f\x7a\x93\x53\x66\xf7\xc1\x6f\x30\x58\xf6\xde\xa3\xc6\xc0\xc8\xa7\x8e\xc5\x2a\x01\xea\x10\xa7\x53\xa1\xf4\x8c\x6e\x02\x63\x6f\xfb\x40\x98\x93\x3a\x00\x87\x3b\x70\x6c\x88\xf7\xcf\x58\xba\xff\xb8\x39\xb0\x62\x0f\x19\x04\x5f\x3f\x3f\xe2\xe3\x55\xdd\x2a\x76\x57\xa0\xfb\x10\x3a\xe1\x7c\x8d\x13\xc9\x98\x63\xa2\xc5\xf5\xb3\xb4\x46\xf9\xad\xd9\x7e\x9f\xcb\x47\xf5\x1d\xe6\x44\xbc\xe7\xc2\xb0\xf5\x49\x66\xaf\x98\x9b\x60\xbe\x71\x55\xd8\x7f\xa9\x21\xdb\x87\x35\x38\x61\xba\x73\xda\x84\xf5\xd3\x77\x94\x78\x58\x89\x87\x6b\xf6\x1d\x95\x99\xf8\x8a\xe6\x21\xe1\xdf\x00\x00\x00\xff\xff\x39\x4c\x24\xb1\x74\x19\x00\x00")

func templates_server_builder_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_server_builder_gotmpl,
		"templates/server/builder.gotmpl",
	)
}

func templates_server_builder_gotmpl() (*asset, error) {
	bytes, err := templates_server_builder_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/server/builder.gotmpl", size: 6516, mode: os.FileMode(420), modTime: time.Unix(1424678182, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_server_operation_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\xdd\x6f\xe3\x44\x10\x7f\xf7\x5f\x31\x44\xc7\x91\x54\xc6\x7d\x2f\xba\x07\x28\x88\xde\x03\x47\xd5\x56\xf0\x88\xb6\xeb\x49\xbc\xd4\xde\x75\xf7\x23\x69\x88\xfc\xbf\x33\xfb\x61\x27\x6e\xdc\x80\x28\x94\x93\x4e\xba\x75\x76\xe6\x37\xbf\xf9\xcd\xc7\xb6\x65\xfc\x81\xad\x10\x76\xbb\xe2\x3a\x1e\xbb\x2e\xcb\xce\xcf\xe1\xae\x12\x06\x96\xa2\x46\xd8\x30\x03\x2b\x94\xa8\x99\xc5\x12\xee\xb7\x60\x2b\x04\xb3\x61\xab\x15\x6a\xb0\x4a\xd5\x85\xb7\xff\xa1\x14\x56\xc8\x15\x5d\xf6\x7e\x8d\x58\x55\x16\x5a\xad\xd6\x08\x4b\x67\x03\x54\x85\x12\xb6\xca\x81\xc6\xaf\xb5\x93\x23\xa4\x3e\x04\x70\xd5\x34\x4c\x96\x59\x26\x9a\x56\x69\x0b\xf3\x0c\x60\x26\xd1\x9e\x57\xd6\xb6\xb3\x8c\xbe\x76\x3b\xcd\x24\x91\x7e\xf7\x80\xdb\x1c\xde\xad\x59\xed\x10\x2e\x3e\x40\xf1\x31\x78\x18\x4a\xc1\x1b\xf9\xeb\xae\xa3\x03\x91\x10\xd2\x2e\x61\xf6\xe5\xe3\x2c\x99\x27\x13\x94\x25\x9d\x16\x21\x63\x92\xe0\xb2\x66\xc6\x7c\x62\x0d\x5d\x5f\x11\x85\x9a\x68\x91\x23\xea\x25\xe3\x94\x83\xa2\x74\x2b\x66\x81\x33\x09\x55\xb8\x06\xc2\x12\xa5\xf7\xbc\x72\xc4\xf9\xc0\x1d\x5a\xa6\x59\x63\x32\xbb\x6d\xf1\x2f\x91\x77\x81\x8b\x58\x42\xf1\xad\xb3\x95\xd2\xe2\x0f\x2c\x03\x41\x62\x15\xad\xbd\x8c\xa5\xe3\x68\x80\x91\x76\xc6\xd5\x16\x88\x0d\xf1\x40\xad\x95\x2e\xe0\x47\xb4\x26\xa8\xa9\xf1\xd1\xa1\xb1\x29\x3c\x50\xe9\x36\x58\xd7\xfe\xff\xa0\x35\x72\xa7\x85\xdd\x06\x41\xb8\x68\x59\x3d\x44\xbe\x75\x9c\xe0\xcd\x4f\xaa\xc4\xba\xe7\x38\x8f\x57\xd7\x01\xab\xeb\xc6\x69\xc4\x5f\xf3\x5e\xc4\x33\xdf\x41\x3d\x6a\xd7\x2d\x20\x39\xdf\xa0\x75\x5a\x9a\x4b\xd5\xb4\x35\x3e\xfd\x7c\xff\x3b\x72\x1b\xac\x83\x17\x39\x8d\x03\xe7\x31\xa3\x45\xac\x4e\x6d\xf0\x5f\xe1\x12\x30\x0f\x0a\xbe\x07\x7f\xad\xc2\xaf\xd5\xaf\xa7\xf4\x7f\xe8\x35\x84\x9e\x94\x27\x9c\xe2\x2e\xf8\x84\x9b\xb1\x3f\x70\x8d\x34\xaa\x5e\x2b\x89\x1b\xf0\x83\x59\xf4\x6d\x1d\xc7\x04\x27\x87\x42\xb5\x7e\xc4\x85\x92\xd9\xd2\x49\x7e\x8c\x3b\xe7\xf6\x09\xce\x1a\x51\x12\xd2\x86\x69\x2c\x2e\x15\x0d\xc9\x93\xcd\xd3\xbc\xe9\xe9\x51\x5a\xc0\xd9\x33\x7e\x7e\xa2\x74\x50\x12\xde\x8f\xaf\x76\x09\xf2\x02\x28\x56\x9e\x8a\xaf\x2f\xfa\x00\x21\xe5\x28\xdd\xf7\x8a\xdf\x5a\xea\xa2\x55\x50\x6f\xf4\x15\xc5\x99\x98\x6d\x30\x56\x3b\x6e\x43\xfc\x14\x68\x2a\x9f\xa1\x6b\xfa\xf2\xa4\x66\x80\xa9\x32\xed\xcb\x71\x75\x4a\x04\x4f\x3c\xa8\x4a\x95\xa7\x2e\xe2\x28\xd6\xa8\x13\xab\x67\xf2\x2c\xe0\x16\xf5\x1a\xaf\xee\xee\xae\xe7\x3a\x95\xef\x06\x4d\xab\xa4\xc1\x5f\x69\x3f\xa0\xce\x41\xc3\x59\xfa\x3d\xb4\xfb\x22\x4a\xaa\x9c\xc5\x1c\x7e\xf3\xcb\xf6\x28\x4a\x9f\x5c\x71\xe3\xad\x3e\xca\xa5\x9a\x53\x5f\x4e\xae\x36\x17\xf6\x4f\xe8\xdd\xd3\x50\x83\xd3\xdc\x53\xf2\xb8\xbe\xd3\x09\xce\x7b\x7e\xf1\x01\xa4\xa8\x03\x31\x38\x45\x27\x64\x56\x52\xa6\x04\x91\x50\x68\x3d\xc4\x69\xcf\xfb\x9c\x08\x70\x11\x80\x62\xdb\xd0\xd1\xcf\xc1\xb0\x27\x23\x4d\x1f\xf8\x11\xf6\xbb\x05\x66\xc3\x12\xdf\x75\xb3\x3e\xaf\x7e\x10\xe3\x57\x31\x3f\x5a\x47\x34\x52\xc2\x7e\x65\x40\x3d\xc4\xe7\x92\xfe\xd1\x3c\xd5\xf5\x96\xe6\x69\x6c\xdb\xb7\xda\x68\x38\x53\xfa\x27\x85\xfb\x4e\xc8\xf2\x17\xff\x38\xa5\xfa\x0d\xfa\xe5\xcf\x3a\xef\xfd\x31\xc6\xb0\x1d\x42\x16\x24\x71\xbf\x28\xbe\x19\xc9\xee\xd3\xb8\xa7\x30\xfb\x4d\xf8\xdf\x54\xe1\xa5\xc7\x71\x6a\xef\x6a\xdc\xbf\x02\x9e\xeb\x94\xcd\x45\xba\x9f\x92\x2f\xcd\x52\xf1\xd2\x0e\x9d\x54\x6a\x88\x38\x74\x4b\x28\x31\xe3\xd6\x85\xa2\xa6\xbf\x15\x0e\x1e\x8f\x37\xec\xe1\xec\xf5\xb0\x2f\x08\x3d\xd1\x20\x07\xaf\xd0\xdf\x2c\xd0\x64\x13\xff\xa3\x2a\x0c\xaf\xd9\x67\x23\xfd\xdb\x2b\x9f\xde\xec\x3f\x03\x00\x00\xff\xff\x6a\x9d\x7c\x9d\xd5\x0b\x00\x00")

func templates_server_operation_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_server_operation_gotmpl,
		"templates/server/operation.gotmpl",
	)
}

func templates_server_operation_gotmpl() (*asset, error) {
	bytes, err := templates_server_operation_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/server/operation.gotmpl", size: 3029, mode: os.FileMode(420), modTime: time.Unix(1424678182, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_server_parameter_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x59\x5b\x73\x13\x3b\x12\x7e\xf7\xaf\xd0\xba\x80\xb2\x83\x99\xe4\x81\xda\x07\xd8\xec\x03\x21\x2c\xa9\x62\x59\x36\x81\xbc\x00\xb5\x28\x33\xb2\xad\x65\x6e\x91\x34\x49\x8c\xcb\xff\xfd\x74\xeb\x32\xa3\xb9\xc6\x36\x49\x9d\x3a\xe7\xe1\xe0\xd1\xa5\xf5\x75\xf7\xd7\x17\x29\xeb\x75\xc4\xe6\x3c\x65\x64\x9c\x0b\x9e\x70\xc5\x6f\xd8\x0d\x8d\x79\x44\x55\x26\xc6\x9b\xcd\x68\xbd\xe6\x73\x12\xfc\x9b\xa7\x1f\x58\xba\x50\x4b\x18\x81\x6f\x26\x04\x79\x75\x4c\xec\x42\x56\x4d\x4f\xd6\xeb\xe0\x13\xc5\x65\x33\x32\x86\xdf\x1f\xb2\x90\x2a\x9e\xa5\x9b\xcd\x78\x46\xe0\xfb\x92\xc6\x05\x3b\xbd\xcb\x05\x93\x52\x0f\xeb\x51\x4f\xfa\xf4\xb5\x16\xfe\xb7\x63\x92\xf2\x98\xac\x47\x84\x08\xa6\x0a\x91\xe2\xe8\x08\xd1\xb0\x34\xaa\x50\xd1\xbb\x41\x54\x6e\x7a\x4f\x54\x95\xf4\x9d\x50\xc1\x49\x8a\x89\xb4\x1b\x93\x9d\xdc\x03\xd1\x0f\xb3\xc5\x88\xfe\xb1\x9b\x9d\x78\xca\x93\x22\xe9\xf5\x1d\x4e\x0e\x22\x9a\xc7\x19\x55\x7f\x7f\x39\xe9\x42\x36\x75\x2e\x34\x47\xe8\xaf\xd3\xbb\x30\x2e\x24\x50\xa9\x1c\xde\xd5\xaf\x03\x78\xcd\xe4\xef\xe2\x75\x47\x34\xf0\xba\xe1\xdd\xf0\x16\xb1\xe2\x79\xcc\xfe\x33\xef\x81\x5c\xce\xff\x2e\x6a\xef\xa0\x9d\x10\x9e\xa6\x7d\xe6\xc4\x99\xfd\xe2\xc3\xc8\xdc\x1a\x86\xfe\xb7\xca\x36\x61\x21\x55\x96\xcc\x33\x91\x50\x55\x4b\x38\x1d\x18\xdf\xe9\x55\xf7\x58\x0f\x07\xcc\x42\xfd\x29\x95\xe0\xe9\xa2\xcf\x96\xe6\x5c\xb9\x1d\xf8\x0a\xb4\x8c\x79\xd8\x95\x1e\x3f\x32\x16\xc9\x0b\xfe\x8b\xe9\x11\xc0\x28\x68\xf2\x91\x26\xf0\x89\x83\xa8\x0b\x4f\xd1\xb3\x31\x4b\xbb\x11\x4d\xdb\x11\x7b\xa6\x58\x22\x7b\x43\x56\xcf\xde\xe7\xb7\x06\x0e\x17\xa8\x56\xf2\xae\x21\x39\x04\xc8\xce\xee\x05\xa8\x94\xbc\x13\xa0\x2f\x29\xbf\x2e\xd8\x00\x26\x6f\xc1\xce\xfc\xfe\x8b\xc7\x56\x2e\xb2\x9c\x09\xb5\xea\x60\xea\x99\xfc\xe4\xca\x3c\xee\x00\xeb\xe4\x31\x40\xed\xac\xfe\x24\xc0\x25\xbe\xaa\x67\xf2\x44\x87\xad\x89\x33\x28\x45\x75\x19\xdd\x31\xdd\x29\x26\x4b\x15\x05\xac\x0d\x01\x8d\xf8\xaa\xef\x6c\x28\x79\xc5\xd3\xa8\x04\x3d\xee\x5a\xa1\xa5\xe1\x32\xe6\x19\x00\x28\xc8\x52\x85\xcb\x82\x33\x98\xb9\xbb\xa4\x80\x21\x44\xb7\x15\x8a\xc7\xc1\x45\x1e\x73\xf5\x66\x65\x14\x34\xbe\xc3\xf5\xfe\xda\xaf\x5d\xa3\xdf\x8d\x77\x4f\xb2\x38\x66\x21\xfa\xb7\x4c\x45\x3a\xb4\x63\xc9\xba\x8e\x14\xf4\xb6\xd2\xcf\x9b\x94\xbf\x70\x56\x42\x88\x8c\x6e\xa8\x20\xb5\x39\xfd\xf9\x79\x95\xb3\xe6\xa6\x4b\x4b\xbb\xd3\x98\x25\x00\x0e\x25\xcc\x8b\x34\x9c\xd4\x16\x61\x22\xd2\x0c\x3b\x59\xf2\x38\x6a\xb3\xaf\x9a\x32\x47\x4c\xc9\x01\x90\x2d\x13\x32\xb0\xe2\x61\x95\x66\x62\x9d\x3b\x4d\xbe\x11\x23\x04\x20\x96\x9c\x05\x0a\x03\x67\x47\xc0\x8e\xba\x3e\x88\xf3\xe8\x75\x63\xec\x1f\xa4\x61\x8f\xc6\x82\xe7\xcf\x2d\x08\x70\x29\x08\xb4\x90\x5b\xf4\xac\x26\x6a\xac\x47\x1e\x98\x09\xe0\xe1\x0d\x20\x47\x1e\xde\xa0\x29\x66\x2e\x86\x4b\x33\x78\x2b\xea\x96\xd4\x3c\xf0\x08\x30\x05\x3c\x36\x07\x78\x01\xeb\x87\x2c\x5a\xf1\x2c\xd5\x46\x42\xe3\x4e\xca\x33\x06\x6b\x9a\xef\x0d\x93\x32\x86\x31\x80\x8d\x4b\x20\x46\x91\x5e\x8a\xd4\x15\x1a\xa2\x45\x3b\x13\xd5\x72\x11\x9e\x4a\x9a\x34\x3d\x26\x34\xcf\x81\xdc\xf5\x53\xc4\x8c\x68\x4b\x4f\xf5\x06\x13\x18\x5a\xdc\xde\x90\x07\xcc\xd1\x81\xba\x81\x7b\x37\xe4\xc3\xa7\x95\x09\x08\xb5\x22\x15\xc9\x6a\xe9\xae\x11\x3a\x7e\x8e\xf2\x83\xe6\xb7\x5d\xe8\xe1\x7e\x0c\x33\xb4\x0f\x71\x89\xac\x4c\xc4\x39\x0d\x7f\xd2\x05\x33\x75\x5f\xff\x84\xd9\xd1\xe1\x21\xf9\xbc\xe4\x92\xcc\x79\xcc\xc8\x2d\x95\x64\xc1\xc0\x2e\xa0\x50\x44\xae\x56\x44\x2d\x19\x91\xb7\x74\xb1\x80\xd8\x55\x59\x16\x07\xb8\xfe\x34\x82\xc8\x4d\x17\x30\xe9\xf6\x25\x7c\xb1\x54\x04\xd2\xce\x0d\x83\x1c\xa7\xb4\xa8\x25\x4b\xc9\x2a\x2b\x40\xaf\x17\xa2\x48\x6b\x92\xdc\x11\x24\xcc\x92\x84\xa6\xd1\x68\xc4\x93\x3c\x13\x8a\x4c\x40\xe9\x71\xca\xd4\xe1\x52\xa9\x7c\x8c\x1f\x0b\xae\x96\xc5\x55\x00\x0b\x0f\x43\x2a\x0b\x1a\xff\x9f\x27\x87\x8b\xec\x85\x15\xb5\xcd\x9a\x43\x13\xea\x5b\x2d\x4d\x78\x14\xc5\xec\x16\xaa\x89\x86\x80\xaa\xc0\x46\xed\x0a\x41\x53\x30\xde\x93\x9f\x6c\x35\x23\x4f\x74\xc8\x20\x17\x82\x33\x8d\x5c\x5a\x22\xe1\xb4\x4e\xda\x60\x0c\xc8\xec\x73\x32\x7e\x7a\x3d\xb6\xcb\xed\x12\xe3\x95\xa9\xb6\x3c\x3a\x2e\xa6\x52\x9a\x16\x4c\x77\x63\x12\x8c\xa2\xc9\x29\x09\x8d\x63\x6d\xb6\xab\xac\x48\x23\x92\x9b\x59\xcc\xd7\x38\x08\x5b\xdf\x17\x60\x3c\x6f\x3f\xc1\xac\xaf\x93\x15\xca\x56\xab\x9c\x87\x20\x42\x3b\x11\xf8\x0f\x3a\x91\xec\x4a\xd3\x3e\x22\x73\x91\x25\x84\x12\xd4\x31\x38\x67\xd0\x97\x49\x35\x82\x0d\xac\x1b\x11\xf4\xee\x45\xa8\x6c\x86\x37\x76\x30\xad\xa3\x74\xd9\xfb\x2d\x93\xa1\xe0\xb9\x49\x94\x46\xb1\xda\x90\x53\x5b\x73\xfa\x93\x2d\x4f\x16\x75\x55\x3e\x2b\xf3\x18\x62\xbe\x81\x38\xb4\xe8\xc0\x08\x6a\x49\x30\x30\xc1\x2e\x60\x0d\xd7\xd3\xc1\x17\x30\x4c\x2f\x99\x11\xae\x08\x40\x2f\x12\x18\x55\x4b\xaa\x90\x5e\x70\x39\xbb\x43\xa2\xa6\x0b\x49\x38\x7e\xe9\x52\x4c\x89\x0d\x5b\x7a\x15\xb3\x09\xa8\x37\x4f\x14\xd8\x61\xc1\xe1\xe7\x6a\x6a\x6a\x03\x56\x66\x26\xe6\x34\x64\x08\x05\xcd\x2e\xb5\x00\x93\x2e\x25\x1e\x76\xcb\xc1\x43\x05\xd8\x16\xb6\x51\x1d\x02\x09\x53\xcb\x2c\x22\x68\x77\x39\xc2\x6a\x4f\x30\x58\xcf\x59\xc8\xa0\xd4\x09\xab\xf0\x41\x97\x91\xa7\xbe\xb6\x13\x41\x0e\x7c\xdf\xcc\x88\xc8\x0a\x88\x97\x83\x8a\x9f\xd0\xa7\xab\x70\xc9\xa2\x73\x9c\x70\x90\xd1\x43\xd8\xa0\x40\x9d\x20\x5f\xbf\xeb\x31\x57\x95\x83\xf7\x54\xfe\xb7\x60\x62\xe5\x1c\x77\x2d\x75\xc7\x13\x7c\x39\xff\x10\xe8\x89\x49\x55\x02\x88\xdd\x80\x85\xdb\xad\xf7\xbc\xd3\xc5\x03\x77\x4e\x9a\xa9\x56\x43\x69\x7a\xcc\xea\xf4\xcd\xa6\x96\x4c\xeb\xe6\x09\xd0\xc9\x2d\x96\x4c\xae\x65\xf0\x2f\xa6\xaa\xee\x7d\x6a\x6d\x62\xef\x98\x1d\x57\x47\xa2\xcd\x50\x26\x4d\xf8\xd0\xdd\xc4\xb4\xac\x8e\xa5\xa6\xd0\x8e\x80\xcc\xbd\xa1\x19\x1c\xc6\x10\x8f\x09\xf2\x3d\xa3\x50\x96\xf6\x87\x19\x18\x01\x8f\x09\xb1\x24\x4c\xe5\xf6\x77\x50\x0d\xca\x21\xff\xc6\xd9\xbc\x81\x1a\x74\x65\xc7\x27\x34\x22\xdc\xed\x81\xed\xed\xe9\x3a\x00\x62\x7b\xf7\x91\xdd\x4e\x5e\x1e\x1d\x41\xe7\x26\x40\x3a\x16\x2d\x5d\xaf\xbe\x8d\xeb\x47\x7f\x1b\x93\x39\x85\x89\xe8\x15\x79\x7a\x33\x36\xea\x69\xfd\x88\xd6\xcd\x1c\xd2\xb6\x73\x3b\x97\x1d\xbb\x1a\x17\x20\xf0\xf5\x5b\xc8\x30\xaf\x48\x53\x6d\xa3\x68\x73\xdc\x8c\x6e\x6a\x56\xdd\xcf\xcd\x68\x37\xdd\x31\x3e\xac\x97\x6d\x2f\x55\xe6\x71\xcf\xed\x0f\x1e\xed\x1d\x77\xbf\x8e\x04\xd0\x77\xc3\x7b\x38\x4a\x63\xa9\xa9\xd3\xfa\x41\x74\xe9\xf3\xd1\x23\x2a\xe4\x7b\xaf\xac\x09\x67\xf2\x4d\x16\x39\x2f\xd5\xae\x29\xe6\x3c\xf0\x2b\x96\x53\xe1\x7e\x00\x70\xdc\x30\x23\xcf\xb6\x08\x86\xad\x51\xda\x40\x05\x18\x92\x9d\xe2\xe7\xa4\x11\x9e\xe3\xae\xab\x58\x7f\x98\xb6\x38\x89\xc5\xfb\x7f\x8d\xdb\x42\xbb\x2c\xeb\xfb\x7f\x6a\x1a\xe4\x7b\x03\xdd\x75\xed\x3d\xb7\x82\xb6\x08\x77\x4f\x98\xdc\xeb\xc9\x41\x6f\x9a\xff\xae\x20\x9b\xfd\xb4\x5f\x9b\x51\xf5\xff\x6d\xb2\x46\x4b\x97\x9d\xa0\x0d\x00\x2b\x21\xd8\x0c\x61\x79\xe6\x25\x0c\xfb\x0b\xd0\xe1\xc3\x2b\xec\x9e\x92\x7f\x92\xa3\xce\x7b\xf9\x09\xb4\x6e\x99\xe4\x8a\x55\xcf\x1c\x86\x1a\xb0\x2b\x08\x02\x47\xec\xfa\x5b\x06\xb4\xdd\x4f\x42\xd7\x58\xe9\xee\xbc\x6c\xb3\x88\x7e\x9d\x69\xf6\x2c\x7e\xc7\xe2\x47\x42\xf9\x8e\xe1\x3d\x54\x74\xbe\xb6\x0d\xf5\x78\x15\x94\xaa\xc7\xeb\x49\xd9\xf4\xd6\xbe\x8f\x97\x2f\xe1\xa4\xa7\x29\x2d\x5f\x59\x30\x33\x4d\x2c\xf4\xb2\x7b\x99\x12\xdd\x31\x72\xc1\x22\x9f\x04\xe5\x03\xa8\x9b\xbc\x28\x1f\xe3\x7b\x9f\x3a\x00\xd3\x96\x8f\x0c\x95\x83\xf5\x15\x7b\xe8\x05\xc7\x7b\xbb\x41\xf9\xfb\xbc\xd0\x0c\xbe\xcd\x94\xaf\x32\x56\xba\xbd\x70\xb4\x5f\xd5\x8e\x4d\x17\x5f\x7f\xf3\xe8\x59\x09\xb2\x3a\xf4\x84\xee\xb8\xa2\xa6\x1c\x0e\x39\x67\xff\xb6\xe7\x7b\xff\xee\x31\x68\xe9\x1a\xeb\xfb\x0a\xf1\xc3\x51\xf3\xeb\xf7\x1d\xc8\x29\xed\x1f\x58\x74\x84\xa3\x17\x4a\x8b\xd5\x98\xa9\x97\x1d\x1f\xf7\x44\xbf\x5b\x3a\xe0\xf0\x46\x69\xab\x8e\x79\xcb\xe6\xb4\x88\xd5\xa5\xb9\x72\x47\x6c\x7e\xe9\xee\xea\xdd\x0f\xca\xf5\xf5\x03\xcf\xc6\x9a\xac\x15\xee\x67\xcf\xb4\x8e\xee\x00\x3f\x95\xf5\x10\xc9\x2d\xad\xb7\x77\x3d\x96\x00\xe7\xfa\xa9\x74\xe8\x95\x6a\x88\xbc\xf5\xa7\x22\x9f\xbd\x17\x28\xe3\x4f\xa2\x70\x07\x87\xab\x3f\x47\x60\xde\xad\x47\x57\x0f\xde\x5d\x19\x7e\xaf\x0e\xc3\x49\x77\xf8\x7d\xbd\x33\x30\xfd\xbf\x94\x94\xff\xfe\x11\x00\x00\xff\xff\x1b\xb4\x80\xa7\xe2\x21\x00\x00")

func templates_server_parameter_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_server_parameter_gotmpl,
		"templates/server/parameter.gotmpl",
	)
}

func templates_server_parameter_gotmpl() (*asset, error) {
	bytes, err := templates_server_parameter_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/server/parameter.gotmpl", size: 8674, mode: os.FileMode(420), modTime: time.Unix(1425157101, 0)}
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
	"templates/model.gotmpl": templates_model_gotmpl,
	"templates/modelvalidator.gotmpl": templates_modelvalidator_gotmpl,
	"templates/server/builder.gotmpl": templates_server_builder_gotmpl,
	"templates/server/operation.gotmpl": templates_server_operation_gotmpl,
	"templates/server/parameter.gotmpl": templates_server_parameter_gotmpl,
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
		"model.gotmpl": &_bintree_t{templates_model_gotmpl, map[string]*_bintree_t{
		}},
		"modelvalidator.gotmpl": &_bintree_t{templates_modelvalidator_gotmpl, map[string]*_bintree_t{
		}},
		"server": &_bintree_t{nil, map[string]*_bintree_t{
			"builder.gotmpl": &_bintree_t{templates_server_builder_gotmpl, map[string]*_bintree_t{
			}},
			"operation.gotmpl": &_bintree_t{templates_server_operation_gotmpl, map[string]*_bintree_t{
			}},
			"parameter.gotmpl": &_bintree_t{templates_server_parameter_gotmpl, map[string]*_bintree_t{
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

