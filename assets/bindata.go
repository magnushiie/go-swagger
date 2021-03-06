package assets

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

var __2_0_schema_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x6d\x73\xdb\xb8\xf1\x7f\x7f\x9f\x02\xa3\xcb\x8c\x2f\x13\x5b\xca\xe5\xff\x7f\xd3\x74\x3a\x37\xee\x39\xbd\xba\x4d\x6a\x4f\x9c\xb4\x2f\x6c\x65\x06\x22\x21\x0b\x77\x14\xa9\x10\xa4\x6d\x5d\xea\xef\xde\xc5\x03\x29\x82\x04\x48\xf0\x41\x8e\x7d\xe1\xcd\x9c\xed\x48\xc0\x62\x77\xb1\xd8\xfd\xed\x02\x04\xbf\x7c\x87\xd0\x24\xa1\x49\x40\x26\xaf\xd1\xe4\x18\xfd\xe3\xe2\xec\x5f\xe8\xc2\x5b\x91\x35\x46\xcb\x28\x46\x17\xb7\xf8\xfa\x9a\xc4\xe8\xd5\xf4\x25\x3a\x3e\x3f\x9d\x4e\x0e\x79\x07\xea\xf3\xd6\xab\x24\xd9\xbc\x9e\xcd\x98\x6c\x32\xa5\xd1\xec\xe6\xd5\x8c\x89\xbe\xd3\x5f\x59\x14\x7e\x2f\x1b\x3f\x93\x1f\x15\x7a\xf0\x2f\x8f\x54\xc3\x28\xbe\x9e\xf9\x31\x5e\x26\x47\x2f\xff\x5f\x75\x56\xfd\x92\xed\x46\x30\x15\x2d\x7e\x25\x5e\x22\x3f\x8b\xc9\xe7\x94\xc6\x84\x0f\x7f\x39\x51\x03\x4f\x0e\x81\xa1\x70\x19\xf1\xdf\x1b\x9c\xac\xd8\x64\x2e\xda\x62\xdf\xa7\x09\x8d\x42\x1c\x9c\xc7\xd1\x86\xc4\x09\x25\x0c\xfa\x2d\x71\xc0\x88\x68\x00\x8d\x13\x12\x87\xda\xb7\x5c\x1f\xf0\xd5\xa7\xbb\xa3\xfc\x1f\x5c\x84\x98\x2c\x39\x2b\xdf\xcf\x7c\xb2\xa4\xa1\x20\xcb\x66\x37\x24\xf4\xa3\xf8\xcd\x5d\x42\x42\x06\x1f\x4c\x44\xeb\x7b\xf8\x79\x2f\xc9\x1b\xe8\x66\x2c\x17\x68\x67\x62\xb2\x24\xa6\xe1\xb5\x10\x53\x7c\x4e\xc2\x74\x2d\xc4\x04\xd5\x4b\x89\xc4\xc7\x3e\x61\x5e\x4c\x37\x9c\x03\xde\xeb\xc3\x8a\xe4\x73\x74\x43\x62\xce\x07\x8a\x96\x28\x59\x51\x86\xfc\xc8\x4b\xd7\x24\x4c\xa6\x8a\x33\x49\x43\xea\xaa\x51\x38\xd1\x4a\xeb\xb7\x8a\x58\xe2\xc2\x38\x98\xcd\x1a\xf3\x96\x93\x34\xa6\xbb\x8f\x95\xb6\xf9\xe7\x9f\x2e\x3f\x7d\xb9\x9f\xa1\xd7\x57\xf0\xdf\xfc\xc5\x0f\x3f\xbd\x86\xbf\xfc\x17\xcf\x7f\x7a\x36\xa9\x13\x73\x99\x06\xc1\x16\x7d\x4e\x71\x40\x97\x94\xf8\xe8\xe3\xfb\x53\x94\x44\x20\x2a\x41\x9c\x35\x29\x36\x91\x56\xaa\x31\xbe\xc0\x8c\x9c\x83\x61\xb8\x30\x5f\xe4\x72\x56\xcb\x0e\xa7\x8a\xb8\xbd\x65\x4c\xf0\x81\xd1\x9b\x3b\xbc\xde\x04\xe4\x35\x3a\x98\xe1\x0d\x3d\x28\x71\x22\xec\xbb\x60\x0f\x56\xed\xab\x86\x6f\x29\xa8\x5c\xa3\xe0\xc1\xb7\x69\x89\x44\x89\xb9\x63\x14\x50\xa9\x8e\x77\xa7\xef\xde\x20\x2e\x29\x43\xd8\xf3\xc8\x26\x01\xad\x2d\xb6\x3b\x2d\x1d\xd6\x33\xb1\x26\x3e\xc5\x1f\xa0\x7b\x95\x0d\xb0\x6d\x3f\xf5\xda\xb3\xa1\x86\x46\x1e\x0e\x91\xa2\xd1\x8b\x0d\xb1\xdc\x1b\xb5\x29\x9b\x69\x3d\x0b\x5f\x37\xf7\x2f\x36\x2e\x8d\x1f\xe3\x35\x01\x83\x71\x62\x42\xb5\x3d\xb1\x51\x8b\x09\xdb\xc0\x87\x2e\xf6\x91\x35\xb5\xd2\x62\xc4\x83\xd5\x97\x6c\x1d\x4c\x2d\x6b\x69\xec\x7f\xd2\x46\x4f\xa6\x4e\x1a\xd5\x04\x5f\x33\xd3\x2a\xc4\x71\x8c\xb7\x3b\x3b\xa0\x09\x59\x17\xdb\x59\x07\x04\x7a\x13\xd5\xe6\x3e\xef\x9d\x86\xf4\x73\x4a\x4e\x15\x8d\x24\x4e\x89\xc6\x03\xb9\xe3\x0b\x1c\x07\x27\x91\xe7\x20\x92\xd6\xba\xe4\xe0\x4d\x36\x54\xf1\xae\x86\x28\x66\x5a\x2d\xbf\x90\x90\xc4\x38\x40\xbc\x3b\x77\x9f\xdc\x8f\xe3\x45\x94\x26\x86\xd5\xaa\x05\x41\xe5\xf4\x79\xf0\x93\x61\x7c\x17\x2a\x9a\x02\x60\xb6\x84\x2c\x41\x50\x7c\xad\x07\xc2\x1a\x45\x19\x83\x61\xa6\x2f\x7d\x82\x0c\x81\x51\xaa\x4a\xc1\x10\x6d\x34\x8b\xa3\x36\xe9\xf0\x18\xc9\xa9\x47\x38\xf4\xc1\xbb\x10\x8f\x82\x87\x16\x44\xab\xb1\x41\xe3\x08\x48\x65\x6a\xec\x33\x3a\x03\xec\x12\x26\xd4\xcb\x03\x31\x44\xf0\x05\xc4\xe5\xc6\xc1\x75\x4a\xdd\x19\x08\xa2\x90\xe3\x80\xc2\xe7\xc5\xb1\xd1\xc5\x2a\x4a\x03\x88\x00\x04\xf9\x74\xb9\x24\x31\x40\x03\xb4\x8c\xa3\xb5\x68\x21\xf4\x34\x45\xe8\x17\x9a\xac\xd2\xc5\xd1\x32\xc0\x37\x11\xd8\x18\x5a\xe3\xf8\x37\x3f\xba\x0d\x11\x00\x0a\x1c\x04\xd1\x2d\xf1\x2d\x52\x80\x19\xad\xd9\xd9\xf2\x82\xc4\x37\xd4\xeb\x33\x8f\x3c\xba\x0a\x62\x9c\x7b\x26\xc9\x09\x30\x5a\xaf\x45\x08\x8d\x09\xf6\x12\x37\x73\xcd\x1a\x1b\x29\x05\x30\x20\x38\x57\x37\x4a\x59\xe3\xaa\xc1\x97\x03\x77\x89\x3b\x57\xd7\xf0\xb3\xec\xa9\xb9\x86\x4c\x1b\x30\x31\x60\x6b\x9a\x85\xb5\x5c\xfe\x96\xb5\x18\x42\xac\xea\x39\x85\xd4\x07\x03\xa3\xcb\x2d\xb4\x45\x9c\x5c\xc6\xa5\xd2\x04\x82\x71\x21\x0f\x98\x41\x02\x80\x43\xfa\xbb\x90\xcb\x32\xb3\x69\x1c\xf4\xe4\xe5\xe3\xfb\xb7\x68\x13\x51\xe0\x07\x98\x51\x78\xcd\xab\xea\x75\xaa\x13\xd2\x91\xac\x91\x35\x58\xf2\xb4\x2f\x73\x82\x06\x82\xe9\x82\xa8\xce\x9c\xb4\x64\xe1\x52\x32\xd3\x60\x88\x55\xe3\xb6\x1a\xa2\x16\x68\x84\x45\xb4\x0d\x2e\xfb\xb3\xae\xa2\x45\x29\x91\x04\xc6\x9c\xa2\xd3\xe4\x80\x21\x12\x7a\x51\x1a\xe3\x6b\xf0\x60\x30\xdd\x29\xe3\x41\x01\x9d\x5d\x00\xf2\x8c\xd6\x10\xf2\xe8\x22\xc8\xbb\x3d\xa8\xd1\xe5\x63\x3a\x19\x9a\x71\x02\x2b\xa0\xd7\xd1\x8f\xbc\x27\x01\x08\x7e\x23\x93\x16\x96\x31\x44\x43\x9f\xde\x50\x1f\xb2\x2a\xd0\x99\x2f\xd8\x65\x53\x04\xec\x6f\xd1\x3a\x05\xfc\x0e\xd1\x22\xce\x3a\xaa\x2e\x07\x59\x42\x75\x30\xad\x24\x4e\x7b\xc4\x11\x85\xb9\x81\xd4\xcc\x89\x18\x97\x94\x03\xc1\x3a\x34\x52\x67\xc8\x2e\x29\x83\x4d\xfb\x16\xba\x8d\x98\x56\x55\x4b\x2a\x7c\x96\x66\xf3\x2c\x04\xf3\x8f\xd1\x1a\x82\xb4\x2c\xda\xc8\xf1\x99\x8a\xfe\x0b\x61\x73\x30\x59\x92\x1c\x83\x79\xe4\x9f\xa8\xf4\xd1\x57\x10\x49\x24\x60\x7a\x4e\x68\xce\x6f\x2c\x69\xc0\xd0\xb2\xe7\xe3\xb5\x17\x3f\x26\x80\xf8\x18\x44\x1c\xe1\x22\x99\x88\x90\x85\xf4\xcc\x98\x68\x3d\x90\x54\xd9\x70\xfb\x15\xca\x96\xd7\xb4\x94\x46\x77\xe0\x25\x06\xab\xd9\x49\x36\x6a\x5e\x6c\x12\x5f\x5a\x42\x08\x77\xa9\xf3\xc6\xe0\xe0\x8e\x86\x87\xf0\xdb\xad\x5c\x2f\x91\xd5\x1d\x27\xfd\x36\x39\xc4\x4b\x7c\xf4\xfb\xcb\xa3\x3f\x1d\xcd\x5f\xcc\xd4\x9f\x57\x57\x47\x2f\xe6\x2f\x9e\xf1\x76\x3d\xdc\xd4\x9a\xae\xc9\x07\xc9\x53\xab\x6a\xd7\xe5\xd5\x15\xdb\xf1\xf1\xe7\xab\xab\xe9\x5f\xae\xae\x66\x9c\x9f\xba\x2a\x58\x5e\xdb\xc9\x62\xf1\xdf\x3f\x7c\x38\x47\x6b\x00\x32\x10\x7b\x4b\xde\x84\xb3\x8d\x4b\xd3\xea\x86\x3d\x76\x95\x91\x27\x9c\xdd\xea\xb5\x0f\x4d\x78\xbd\xfe\x21\xbe\xaa\xd6\x40\x4c\x33\x59\xf8\xf2\x5e\xeb\x6e\x29\x82\x94\x1a\x4e\x20\x18\x40\x82\xb7\xed\x95\x75\x2e\x62\x4a\x20\x4f\x93\x94\x32\x33\xc8\xe7\xfa\xab\xa5\xbc\x39\x07\x87\xe8\xba\x73\x4e\x6b\x71\xaa\xe2\x3b\xd7\x82\x51\x99\x66\xce\xd8\xa9\xdf\x4b\xf4\x25\xe8\x3d\xf4\x83\xad\x06\x85\x77\x6b\xcc\x38\xb6\xa1\x6e\x6b\xd1\x6b\xcb\xda\x6d\x8d\x46\x0c\xf5\xdb\x32\x5b\x86\xaa\x76\x27\xb6\x14\x9d\xa1\xd8\x32\x96\x77\x6b\x48\xee\xda\xdb\x69\x9a\x8a\xbc\x35\x24\x77\xcd\xcd\xeb\xb7\xb2\xa1\x50\x43\xab\xb2\xa9\x50\xa6\xe6\x73\xb8\xe1\xe1\x84\x58\x0d\x73\x11\x45\x01\xc1\x61\xd9\x32\x97\x38\x0d\x12\x2d\x1a\x55\x18\xad\xd6\xa3\xeb\x38\xd5\x6a\xd2\x82\x96\x35\x15\x12\xf8\x7e\x28\xbc\xf3\x88\xe2\x85\x22\xdc\x1a\xff\x5c\x13\xc7\x12\x58\x93\xab\x48\x07\xa2\xa3\x6f\x1b\x76\x27\xe4\x93\x00\xd6\xd6\x20\xa4\xa2\x4d\x19\xf4\x77\xa7\xb5\x22\xb8\xb2\x5c\xba\x29\x0a\x27\xde\x6a\x20\x4a\x03\xf9\x2d\xe3\xa2\x33\x6e\x53\x39\xd7\x20\x64\xdf\x3c\x5b\xe5\xb1\x8b\x09\xdf\x4d\x20\x48\x43\x24\x5f\xf0\x4a\xcd\x16\xdd\xe0\x80\xfa\x12\x4c\x32\xc8\x29\x52\x68\x13\xf9\x22\x3b\x3a\x50\xee\xa6\x58\x7c\x58\x53\x7d\xc9\xfe\x38\xec\xaa\xff\xe1\x12\x50\xf1\xfc\xcb\xff\xdd\x3f\x7f\xf6\xdf\x4f\x3f\xa8\xf1\x9f\x3f\x6b\xe7\xc1\xff\x8d\x83\x94\x58\xca\x19\x7b\x70\x2b\x61\x94\x94\xf0\xa7\x79\x86\x1c\x75\xd4\xa8\x25\xa3\x18\xed\x05\xd9\x89\xd2\x64\x7e\x52\x9f\x05\x13\x8c\x42\x72\xc6\x87\xba\xec\x90\x90\x37\x27\xf1\xfc\xcc\xca\x7b\x22\x76\x4d\xbc\x5d\xc7\xb9\x91\xb5\xd6\xf9\x4d\x71\x95\xec\x3d\x41\xce\x8f\xe3\xb8\x42\x06\x6c\x77\x7c\xce\x2e\x26\x6b\x6c\x41\xd9\x95\xb4\xba\x86\x54\xde\xba\xc6\xfa\x9d\xb3\xe5\xaa\x10\x83\x97\x7f\xe4\x10\x13\xa3\x31\xab\xef\x86\x82\x2f\x9a\x51\x09\x4a\xcd\xd6\x94\xe8\xb5\x02\x8d\x09\x53\x1e\x92\x9f\x4b\xca\xbe\x04\x5f\x23\xf6\x39\xe5\x39\xac\x84\xa8\x23\x59\x39\x54\xcc\xd2\xdb\xb9\x71\xee\xf3\xfa\x4b\x6b\x3b\x36\x65\xc8\xb6\x00\x17\xd3\x35\xe5\x25\x6c\x26\x33\x62\x23\x3d\x2f\x0a\x02\x50\x39\x74\xf8\x9b\x91\x27\xdb\x46\x62\xa9\x97\x05\xbb\x64\x10\xd9\x81\x64\xd6\xd8\x48\x69\x8d\xef\xe8\x3a\x5d\xbb\x51\xca\x1a\x5b\x56\x9d\x17\xa4\x0c\x94\xf2\xae\x0d\xc9\x4a\x2f\x33\x97\xd0\xde\x9d\x4b\xd5\xb8\x81\xcb\x36\x24\x2b\xbd\x6c\xba\x7c\x4b\xc2\xeb\xc4\x11\x75\xed\x9a\xdb\x64\x6e\x45\x2d\x6f\x6e\x43\x83\xaa\x34\xe7\xb6\xcf\x21\x1a\xdb\xa4\x3c\x75\x5f\x2a\x79\x6b\x9b\x8c\x6d\x68\x65\xad\x8d\xb4\xf4\x12\x95\x03\xb9\x62\x07\xb3\xad\x84\xce\xf6\x11\x5a\x6d\x02\x56\x1e\x85\xf0\x72\x56\x49\xbe\x2c\x32\xee\xda\x5b\x56\x7e\xfb\x60\x6d\x0c\x15\x65\xc8\x54\x73\xf2\x0e\xf0\xb3\xf2\xf6\x5b\x8e\x9e\x63\xb1\x01\x79\x0b\xf8\x1a\xdd\x1d\xf1\x9a\x97\x00\xd7\xcd\x27\x05\x78\xd5\xd0\xd0\xc6\x7a\xb2\x6a\x11\xf9\xdb\xf3\x7c\x07\xa7\xd3\xfe\xb2\x08\x21\xfc\xa7\x02\x1e\xf3\xc7\x98\x99\x0f\x55\xbb\x94\x85\x53\x43\xe9\x32\xcf\xc7\xf8\x46\x2c\xe5\xa9\x0f\x3f\xb7\x23\x4e\x04\x50\x48\x94\x14\xfe\xe1\xad\x53\xd6\xef\xcc\xce\xc0\x1b\xf1\x3b\xc6\x2d\x21\xbb\x8f\xc2\x4e\x38\x61\x70\x29\xaa\xe6\x17\x44\x1e\x36\x2b\xcd\x82\x57\xb8\x79\x5a\x40\x48\xc1\x10\xdb\x94\xbd\x6c\xec\xdd\xae\x88\xc8\x65\x21\x5b\x85\x34\x4c\x1e\xc6\xce\xd9\xe3\x93\x92\x8d\xc7\x5b\xc8\x5a\x04\x0e\xa6\x1d\x8a\x6a\x7d\x90\x7c\x5f\xcc\x9c\xaf\xf4\x8b\x74\x71\x51\x66\xe4\xc9\x55\xdb\x9e\xa8\x05\x7c\xd5\x05\xa5\x92\x17\xf3\x92\x1a\x9d\x64\x57\x27\x39\x68\x42\x56\xc8\xc2\x0a\xb9\xd9\x98\x90\x8d\x09\x59\x1d\x97\x63\x42\x56\xa5\x36\x26\x64\x4f\x38\x21\x33\x66\x54\xc0\x7b\xbc\x1d\x61\xcc\xb7\x0e\x63\x84\x19\x8c\x28\x66\x44\x31\xb2\xcd\xe3\x44\x31\xff\x81\x79\x7f\xc7\x3d\xdb\x08\x67\xd0\x08\x67\x46\x38\x33\xc2\x19\x54\x86\x33\xdc\xe5\x9d\xe0\x04\x8f\x88\xe6\x5b\x47\x34\x99\x25\x8c\xa0\xe6\xa9\x82\x1a\xf8\x63\x49\xf9\xc3\xdc\x23\xb8\x19\xc1\x4d\x3d\x97\x23\xb8\xa9\x52\x1b\xc1\xcd\x1f\x0d\xdc\xf0\x93\xde\x23\xb0\xc9\x03\x4a\xe1\x33\xbe\x0d\x3f\xdf\x2f\xfa\x79\x7c\x08\x87\x9b\xc3\x88\x6e\x9e\x2a\xba\xf9\x36\x50\xcd\x08\x66\xd0\x08\x66\x46\x30\x33\x82\x19\x54\x06\x33\x61\x14\xfe\x75\xc0\xb3\x72\xa5\x83\xde\xee\xcf\x41\x58\x0f\xf2\xe4\x20\x06\x35\xd2\xa8\xa9\x3a\xb5\xa0\x62\xdb\x8a\x6b\x41\xc2\x82\x10\x33\x0a\xf3\x12\x9c\x34\xa8\xde\x5d\x71\xfa\x49\xc7\x16\x4c\x56\x66\xde\xcc\x5e\xe5\x60\x97\xeb\xf3\x55\xc7\x00\x58\x24\xb2\xa2\xac\x78\x13\x19\x20\x00\xac\xdd\xed\x2a\xe9\x3c\xe8\xfd\x2d\xce\xc8\xb8\xeb\x93\x8f\xdd\x01\x84\xf1\xe6\xb9\x4c\x3e\xc7\x4b\x6b\x67\x3b\x71\x66\x92\x5e\x5b\x78\xda\x63\xc0\x22\xd5\x2e\xb0\xa3\xd3\x90\x75\xd8\xa4\xd1\x77\x76\x18\xb1\xe9\xa8\x75\x03\x20\xea\x32\xe2\x10\xa8\xa9\xc3\xb8\x83\x40\xab\x2e\xf2\x0e\x81\xbf\x7a\xc9\xdb\x0b\xa4\xb9\x8e\xac\x85\x8d\x88\x89\x7c\xe0\x54\x25\x29\xdd\x00\xdd\x00\x23\x9f\xc8\xf5\xf4\xb2\x13\x08\xec\xa0\xf3\x5e\x48\x71\x9f\x9a\xde\xf7\xc0\xf5\x8a\x76\xc0\xa2\x1d\x94\xdd\x08\x58\x41\xe1\xf6\x07\x6b\x1f\x42\xeb\x0f\x32\x7a\xbd\xea\x6d\x15\xb1\x3e\x0c\xc8\xa8\x7f\x2c\xca\x0e\xad\x33\x85\x2e\x9e\xcc\x9a\x4e\x98\xaa\x2b\x7d\x20\x06\x27\x67\x1c\xc8\x58\x0d\xc1\xe1\xb6\x8c\x6d\x35\x0e\xea\x1f\x16\x28\x81\xdb\xa2\x3c\x86\xbb\x8b\xf4\x65\xfc\x63\xe9\x2b\xf3\xc5\x46\x2d\x58\x41\x85\x47\xc4\x39\x72\xb6\xec\xcd\x7e\xb9\x37\x6a\x07\x07\x41\x15\x94\xd4\xdc\xc3\x64\x93\xc4\x76\x41\x93\xb3\x42\xcd\x6e\xde\xba\x0a\xad\x8f\xef\x23\x87\x27\x93\xdb\x71\xd6\x4e\xa1\x3e\xe5\xd0\x13\xd4\x84\x93\x28\xee\x02\xbe\x63\xc8\x42\xcf\xc2\xc0\x7a\x09\x55\xe7\x3b\x67\xee\xd6\x95\x9b\xe0\xcc\x3a\xe0\x0d\x2d\x90\x67\xf8\xeb\x9f\x54\x81\xb8\xa8\xd0\xd2\x56\x47\xa9\x44\xe9\x90\x05\xf6\xba\x8c\x75\x7c\x00\x5c\x75\x1d\xcb\xbe\x63\xd9\xb7\x49\x9b\x63\xd9\x77\x2c\xfb\x76\x2d\xfb\x9a\xde\xc9\xd1\xf7\x0d\x18\x19\xcd\xf7\x12\x3b\xf3\x1b\x58\x77\xf5\xc8\xef\x8c\x8a\xae\x3c\xb7\x6d\xa2\xd1\x3d\xe8\x54\x23\x4b\x05\x56\x99\x3d\xb7\xd3\xe5\xae\xb6\x6b\x2d\x75\x4d\xeb\x91\x7f\x3f\x61\xd3\x61\x73\xd6\x28\x82\xb8\xf1\x6a\x83\x1b\xdf\x90\x60\xc1\x87\x30\xf7\x77\x5d\x7a\x82\x83\x88\xe9\x22\xad\x5e\xa0\xd6\x1b\x66\xdd\xc6\x78\xb3\x19\xea\xca\x40\xe3\x74\xf2\xd7\xbc\x0c\x35\x9d\x75\x77\xda\x0f\x3d\xd5\x3d\x2f\x6a\x1a\x0c\x7d\x56\x1c\xc2\xc3\x15\xfb\x8d\x8e\xe6\x64\x47\x62\x18\x47\x63\xd8\xb9\xb1\xf2\xbc\xc0\x8c\x7a\xc7\x69\xb2\xe2\x6f\xa7\x90\xc7\x4e\x2e\xaa\xd7\x4d\x1e\xba\x88\x8f\x37\xf4\x9f\x64\xdb\xb5\x77\x84\x81\x89\x57\xa7\x90\x0c\x50\x8f\x26\xfd\xa8\x9c\x63\xc6\x6e\xa3\xd8\xef\x47\xe5\x78\xc3\x79\xe9\xa5\x12\x45\xc8\xf3\x08\x63\x3f\x47\x3e\x31\xd0\x99\x1b\x6d\xa3\x6e\x5a\xf6\xb2\xf6\xf7\x7d\x3d\x96\x10\xa8\xf7\x19\xa5\xc7\xb7\x88\x4b\x46\xbf\xbf\xb9\x39\x44\xc5\x7d\xee\x3d\xce\x94\x94\xc8\x32\x55\x5d\x7d\x7f\xbb\xe3\x71\xe5\x3b\x15\x0e\xd1\x40\x8f\x25\x3e\x3e\xfb\xb1\xb8\xbd\xbd\xda\xd1\x32\x88\x6e\x45\x1d\x02\x86\x8e\x62\xf5\x1a\x9d\x8f\x4e\x6f\x23\xe8\x6c\x55\x52\x4e\x5b\xd1\x83\x33\xd4\x85\x2a\x55\x7a\xb3\xd0\x65\x1e\x48\xe1\x7a\xaf\xad\x60\xf0\x42\xf6\x30\x63\xc6\xb2\xba\x5a\x70\xec\xf0\xee\xa4\xa7\x6f\xc5\x95\xb0\xfb\x20\x56\x9c\x44\xbf\x91\xa7\x69\xbd\x1b\xa5\xaf\x07\xb1\xde\x5c\x4d\xa3\xd5\xea\x56\x6b\x82\x79\xa3\xe1\x36\x40\x84\x5c\x65\xa3\xed\x7e\x55\xdb\xad\x66\x16\x5f\x07\x39\x3c\x6d\x73\xce\xb5\xf8\x47\xc0\x11\xe3\x62\x41\xe6\xc5\x72\x51\x9e\xc5\x01\x6a\xb9\xba\xc8\xfa\xa8\xfa\x2b\x55\x06\xac\x72\xe7\xef\xb4\xaa\xe8\xb7\xa1\xb4\x5d\x7a\xb3\x41\x33\x4b\x86\x47\x68\x76\x34\x50\x48\x88\x2f\xdf\x28\xc9\x60\x3e\x10\x56\xaf\x29\x90\xaf\x9f\x09\x02\xfd\x8d\x04\x95\x63\x03\x86\x3b\x54\x75\x0f\x64\xd0\x47\x9b\xca\x56\xe5\xed\x79\xae\x45\x1b\xe3\xdd\xf6\xc5\x4a\x8d\xf3\x4e\x42\xe1\x0d\x33\x9d\x74\x9d\xc4\x38\x64\xc0\x08\xbf\xbf\x36\x89\xbc\x28\x30\xbf\xda\xd6\xa0\x28\xeb\x62\x2f\xe4\xd6\x49\xb2\xe1\x7e\x9b\xff\x66\xfc\x8f\x5b\xf5\x93\x4d\x5a\x4b\x5a\xb3\x57\x6c\x63\x65\xc7\x88\xc7\x6e\xc4\x1d\xb7\xf2\x57\x22\x7f\x6d\xe8\x46\x7b\xbf\xd9\xae\x44\x2e\xda\xd7\x8e\xbe\x7b\x80\x7a\x20\x36\x0e\xd5\x3e\x98\x23\x3f\xe5\xf3\xd3\x03\x9d\x9d\xce\xc8\x9b\x3d\xf1\xa0\xe7\xa5\x77\x43\x95\xf7\xe9\x07\x3b\x23\x9d\x7b\x49\xd3\x06\xe3\x90\xe7\xa2\xf3\x81\x2a\xdb\xfa\x83\x9d\x85\xce\x86\xa8\x39\x42\x30\xfc\xf9\xe7\x5c\xae\xca\x41\x80\xc1\xce\x3c\x57\xe5\x1a\x74\x2c\xf3\xa9\x84\xc2\x7c\x55\xce\x0f\x0c\x7f\xf6\xb3\xa0\xc5\xbd\x8e\xa6\x9f\xf5\xdc\x05\xe4\xf2\xa9\x86\xc1\xce\x30\x17\xd4\x58\x39\x38\xb5\x4f\x2d\xee\x73\x30\xb3\x12\xcd\x07\x27\x06\x3d\x9f\x9c\x2f\x84\x70\x38\xe3\x0f\xcb\x06\xaf\xe3\x8e\x7d\x6f\xd8\xb7\x7a\xac\x48\xc7\xb7\xdf\xf1\xff\xef\xff\x17\x00\x00\xff\xff\xf5\xe4\x24\xa1\x2c\x8b\x00\x00")

func _2_0_schema_json_bytes() ([]byte, error) {
	return bindata_read(
		__2_0_schema_json,
		"2.0/schema.json",
	)
}

func _2_0_schema_json() (*asset, error) {
	bytes, err := _2_0_schema_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "2.0/schema.json", size: 35628, mode: os.FileMode(420), modTime: time.Unix(1421575197, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _jsonschema_draft_04_json = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x3b\x6f\xdb\x30\x10\xde\xfd\x2b\x04\xa5\x63\x52\xb9\x40\xa7\x6c\x45\xbb\x18\x68\xd1\x0c\xdd\x0c\x0f\xb4\x75\xb2\x19\x50\xa4\x42\x51\x85\x0d\x43\xff\xbd\xa4\xa8\x07\x29\x91\x92\x2d\xbb\x48\xb4\xc4\xe1\xbd\xbe\x3b\xde\x8b\xe7\x45\x20\xbf\x10\xc7\xe1\x73\x10\x1e\x84\xc8\x9e\xa3\xe8\x35\x67\xf4\x29\xdf\x1d\x20\x45\x9f\x19\xdf\x47\x31\x47\x89\x78\x5a\x7e\x8d\xf4\xd9\x43\xf8\xa8\x85\x3e\xe9\xff\x67\x48\xc6\x90\xef\x38\xce\x04\x66\x54\x49\x7f\x67\x1c\x02\xcd\x12\xa4\x20\x50\xad\xa2\xe3\x4e\x30\xc5\x8a\x39\x97\xdc\x1a\x71\x45\xd0\x6c\xdf\x38\x47\x27\x8b\x50\x11\xc5\x29\x03\xa5\x1c\x55\xe4\x47\x9b\x98\x62\xba\x12\x90\x2a\x7d\x5f\x7a\x24\x5c\x9f\x9f\xa5\x83\x1c\x12\xa5\xe2\x21\x0c\xca\x96\xa9\xec\xf8\xc3\x8c\xe5\x12\xd7\x5f\x58\x51\x01\x7b\xe0\x7e\x10\xb8\x66\x18\xc2\xc0\x69\x91\x4a\x8e\xe5\x25\xfa\x7f\x40\x82\x0a\x22\x96\x43\x3b\x88\x90\xdf\x0a\xea\xda\x82\x1d\x19\x91\x8b\xfa\x58\xa5\x21\xc5\x1c\x6b\x9d\x0a\x42\x50\x06\x1b\x27\x8c\x1c\xa7\x19\x81\x3f\xd2\x97\x7c\x68\x1a\x68\xe5\xc0\xba\x8d\x74\x10\x6e\x19\x23\x80\xa8\xfa\xd9\x3a\x1e\x84\xb4\x20\x44\xff\x4d\xb7\xfa\x84\x6d\x5f\x61\x27\xd4\xaf\x5c\x70\x4c\xf7\xa1\xcf\x7e\x45\x9d\x73\xcf\xc6\x65\x36\x7c\x8d\xa9\xf2\xf2\x94\x28\x28\x7e\x2b\xa0\xa1\x0a\x5e\x40\x07\x73\x61\x80\x6d\x6d\x34\x8e\xe9\xd3\x8c\xb3\x0c\xb8\xc0\xbd\xe8\xe9\xa2\xf3\x78\x53\xa3\xec\x01\x49\x18\x4f\x91\xba\xab\xb0\xe0\x38\x74\xc6\xaa\x2b\xca\x7b\x6b\x16\x58\x10\x98\xd4\xeb\x14\xb5\xeb\x7d\x96\x82\x26\x4b\xcf\xe6\x71\x2a\xcf\xb0\x4c\xcd\x2a\xf7\x3d\x6a\x9b\x74\xf3\x56\x5e\x8f\x02\xc7\x1d\x29\x72\x59\x28\xbf\x5a\x16\xfb\xc6\x4d\xfb\xe8\x58\xb3\x8c\x1b\x77\x0a\x77\x86\xa6\xb4\xb4\xf5\x64\x93\xbb\xa0\x24\x88\xe4\x1e\x84\xad\x13\x37\x21\x9c\xd2\x72\x0b\x42\x74\xfc\x09\x74\x2f\x0e\xbd\x9e\x3b\xd5\xbc\x2c\x1f\xaf\xd6\xd0\xb6\x52\xbb\xdf\x22\x21\x80\x4f\xe7\xa8\xb7\x78\xb8\xd4\x7d\x74\x07\x13\xc5\x71\x05\x05\x91\xa6\x91\xf4\x7b\x38\x3d\xe9\x1e\x6e\x1d\xab\xef\x3c\x0c\x74\xbf\x7d\xd5\x6c\xce\x89\xa5\xbe\x8d\xf7\x66\xce\xee\xd1\x86\x67\x80\x34\xad\x8f\xc3\xb3\xae\xc6\x1c\xe3\xb7\xc2\x96\xd9\xb4\x72\x0c\xf0\xab\x92\xe9\x5a\x05\xee\x5c\xb2\x87\xc6\x7f\xa9\x9b\x17\x6b\xb0\xcc\x75\x77\x96\x16\xb7\xcf\x1c\xde\x0a\xcc\x21\x1e\x53\x64\x0e\x73\x4f\x81\xbc\xb8\x07\xa6\xe6\xfa\x50\x55\xe2\x5b\x4d\xad\x4b\xb6\xb6\x81\x49\x77\xc7\xca\x68\x1a\x90\x67\xd7\x78\x3f\x3c\xba\xa3\x8e\xdd\xe8\x7b\xc0\x8a\x21\x03\x1a\x03\xdd\xdd\x11\xd1\x20\xd3\x46\x72\x55\x7d\x93\x0d\xb3\xcf\x34\x52\x46\x03\xd9\x8d\x75\xe2\x0e\x42\xbd\xb9\xdf\xe9\xdd\x34\xb6\x24\x9b\x5b\xa4\x56\x3f\x6b\xac\xd8\x01\x30\x1e\x25\xce\x3a\x77\xc6\x73\xd4\xbd\x96\xc9\xf5\x06\xbc\xca\xf8\x44\xb0\x2e\x09\x5a\xf3\xf5\x3a\x94\x7b\xb7\xa8\x9f\x7f\x17\x8e\x58\x53\xb2\x0e\xfc\xf5\x92\x8c\xc2\x4c\x49\xca\x84\xe7\x7d\x5d\xb6\x2f\x7e\x4f\x79\xba\x96\xe6\x75\xb7\x87\x9b\x0d\xdc\xb5\xbd\xae\xbb\x85\xb8\x8e\x64\x67\xd1\xe8\x18\xe5\xe2\x5f\x00\x00\x00\xff\xff\x4e\x9b\x8d\xdf\x17\x11\x00\x00")

func jsonschema_draft_04_json_bytes() ([]byte, error) {
	return bindata_read(
		_jsonschema_draft_04_json,
		"jsonschema-draft-04.json",
	)
}

func jsonschema_draft_04_json() (*asset, error) {
	bytes, err := jsonschema_draft_04_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "jsonschema-draft-04.json", size: 4375, mode: os.FileMode(420), modTime: time.Unix(1421575197, 0)}
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
	"2.0/schema.json": _2_0_schema_json,
	"jsonschema-draft-04.json": jsonschema_draft_04_json,
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
	"2.0": &_bintree_t{nil, map[string]*_bintree_t{
		"schema.json": &_bintree_t{_2_0_schema_json, map[string]*_bintree_t{
		}},
	}},
	"jsonschema-draft-04.json": &_bintree_t{jsonschema_draft_04_json, map[string]*_bintree_t{
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

