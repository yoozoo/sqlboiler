// Code generated by go-bindata. DO NOT EDIT.
// sources:
// override/templates/17_upsert.go.tpl (6.883kB)
// override/templates/singleton/mysql_upsert.go.tpl (1.13kB)
// override/templates_test/singleton/mysql_main_test.go.tpl (5.129kB)
// override/templates_test/singleton/mysql_suites_test.go.tpl (255B)
// override/templates_test/upsert.go.tpl (1.845kB)

package driver

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\x48\x0e\x7f\xb6\x3f\x05\xd7\xe8\xee\xca\x07\x55\xed\x01\x87\x7b\xc8\x21\x0f\xcd\x9f\x76\x73\x4d\xda\x24\x6e\x2e\xc0\x05\x41\x31\x91\x28\x67\x90\xf1\x8c\x3a\x1a\x25\xf5\xe9\xf4\xdd\x0f\xe4\x48\x96\xe4\xd8\x8e\xfb\x27\x87\x7d\x8a\x35\x43\x91\x1c\xfe\xc8\x1f\x39\x4a\x59\xbe\x84\x17\x42\x49\x91\xc3\xce\x2e\x44\x6f\xe8\x17\xe6\xd1\x27\x71\xa3\x10\xfc\x9f\xe8\x83\x98\x61\x55\x0d\x59\x34\x8f\x6f\x71\x26\xfc\x36\xbd\xd0\x4a\xc0\x7f\x21\x9a\xb4\xbb\xfc\x82\x4c\x21\x7a\x93\x24\xef\x94\xb9\x11\x0a\x5e\x56\xd5\xf0\xd5\x2b\xb8\xc8\x72\xb4\xee\x1d\x08\xe7\x70\x96\xb9\x1c\x84\x06\xa9\x69\x2d\x04\xa1\x13\x48\x0c\xf2\x5a\x91\x25\xc2\x21\x18\x0b\x72\xaa\x8d\x45\x30\x1a\x62\xa3\x53\x25\x63\x17\x0d\xd3\x42\xc7\x10\x18\xf8\x4b\x59\x7a\xff\xa3\x8b\x6c\x22\xf5\xb4\x50\xc2\x56\xd5\xb8\xb1\x12\xb0\x13\xda\x38\x88\x3e\x98\x7d\xa3\x1d\x7e\x75\x55\x15\xbb\xaf\xa4\x8a\x1e\xa2\x7a\x31\x84\xb2\x44\x9d\x90\x93\xb5\xe5\x7d\xa3\x8a\x99\xce\xc3\xda\xb9\xfa\x11\x6e\x8c\x54\x51\xfd\x30\x06\xb4\xd6\x58\x28\x87\x03\x8b\xae\xb0\x1a\x4c\xe4\x0d\x7b\xbb\x5d\x9b\xfc\xde\x3b\x74\x07\x7b\xc1\xb8\x2c\x51\xe5\xc8\x7e\x84\xd0\x6c\xd4\x92\xf5\xbe\x4e\xaa\x2a\xdc\xe8\xc9\x78\x58\x0d\x87\x0b\xa7\x87\x3e\xdc\x14\xc0\x4e\xc8\xe9\xe7\xa9\xd0\x32\x5e\x0a\xfe\xe9\x8f\x45\x1f\x58\x67\x4e\x6b\x1c\x80\xad\xe1\x38\x7d\x6e\x3c\xca\xe1\x40\xa6\xe4\x14\x65\xe7\xff\x13\x8c\x7f\xb0\xd1\x5f\x76\x41\x4b\x45\x5e\x0c\x32\x0a\x51\xc0\xfa\x2e\xad\xc8\x0e\xad\x0d\xd0\xda\xf1\x78\x38\xa8\x56\x01\xb7\x06\xa9\x55\x40\x41\x91\x4b\x3d\xa5\x67\xfc\x8a\x71\xe1\x8c\xfd\x96\xc2\xe9\xa8\xce\xbe\x0f\xc5\xd3\xc7\xf1\x24\x47\x7c\xec\x0e\x6b\x97\x3a\x51\x7d\x0c\x6d\x2b\x5e\x2f\x75\xde\x7a\x3a\xd6\xdb\x43\xbe\x22\xcf\xba\x79\x45\x6e\x3c\x1f\xac\xf7\xc2\xc2\x6c\x3e\x39\x3b\x5e\x19\xcc\x0b\x2d\xbf\x14\x8d\x55\xd8\x85\xab\xeb\xdc\x59\xa9\xa7\x25\xf3\xac\x15\x7a\x8a\xf0\x42\x86\xf0\x22\x36\xaa\xc3\xb4\xcd\x0b\x64\x61\x40\x92\x32\x65\x91\xc8\xeb\xa3\xd5\x51\x59\xf2\x8a\xa7\xed\x51\xe8\xe5\x1a\xb7\xea\xdf\x15\x7b\xbb\xc8\x85\xe7\xc8\xb2\x09\x62\x0f\x29\x48\x4c\x5c\xcc\x50\x3b\xe1\xa4\xd1\x90\x1a\x0b\xb7\xe6\x01\x9c\x81\xcc\x9a\x0c\xad\x9a\x43\x91\x63\x1f\x0e\xb6\xd8\x43\x64\xdb\x24\xfd\x73\xe5\xe8\xa2\x4d\xc8\x14\x0c\xec\xb6\xe9\x54\xb7\x0d\xde\xcf\xa3\x0f\xf8\x10\x8c\xca\x32\x3a\xbd\x9b\x7a\xf4\x76\x40\x1b\x28\xcb\x5e\x23\xa6\x70\xdd\xcb\x04\x13\x0e\x61\xc1\xa7\x1d\x71\xfe\x79\xa4\x09\x48\x45\xd0\x8c\x9c\x9c\x61\xee\xc4\x2c\xfb\xec\xa5\x3e\xdf\xa2\xca\xd0\x8e\x20\x82\xca\x4b\xb7\x35\xf2\x87\x31\x77\x75\x5a\x75\xab\x29\x31\x7b\x98\x1a\x8b\x3e\xa8\x2c\xb4\x75\x69\x3d\x2e\x9e\xf6\xb4\xe4\xee\xa0\xcd\xc5\xe1\x40\xff\xe7\x00\x53\x51\x28\xc7\x83\xc8\x97\x02\xad\xc4\x3c\xfa\x60\xf4\xbf\xd1\x9a\x7a\x6b\x82\x04\x6b\x0d\xfa\x81\x79\xd0\x2d\xec\x75\xa4\x2f\xa5\xbb\xad\x85\x43\x30\x63\x52\xeb\x0b\xe3\x09\xad\x5b\xd6\x29\xeb\xe4\x00\x29\xd4\xc1\x42\xf7\x98\x10\x7d\xbd\x0e\xcf\x58\x68\x0a\x96\x87\x00\x1e\xa4\xbb\x05\x01\x8e\x27\x28\x77\x2b\x1c\xd4\xfb\x4d\xed\x50\x1d\x09\x28\x58\x33\xc4\x6c\xb7\x41\xf7\xd5\x2b\xd8\x2b\xa4\x4a\x20\x16\xf1\x2d\xc2\x1d\xce\x41\xea\x97\x4a\x6a\x84\x62\xaa\xa4\x9a\xc3\x4b\x98\xcd\xf3\x2f\x0a\xee\x73\xc8\xe8\x6f\x66\xcd\x8d\xc2\x59\x3e\x1c\xdc\x14\x29\x85\x20\x77\x76\x26\xf4\x54\x21\x35\xb9\xbd\x22\x4d\xd1\x06\x63\xde\x8d\x2e\xad\x74\x38\x61\x12\x0a\x72\x67\x63\xa3\xef\xa3\x23\x67\x44\xd0\xcb\xf3\xe8\xbd\xd4\x09\xd1\x1d\x25\xdf\xe7\x10\x62\xd2\xea\xe9\xaa\x2f\xb7\x6f\x54\xce\x21\x59\xd6\x1d\xf3\x69\xda\xe5\xbd\xb9\xc3\xe0\xf7\xe8\xf7\xa7\xdc\xe8\xd3\xc0\x7a\x37\xfa\x72\xdf\xe3\xc6\x63\x9d\x9d\xec\xfc\x09\xba\x9a\x94\xdc\xa0\x8a\xb0\xdd\xd9\x05\xda\xad\x37\xc6\xc3\x41\x0b\xde\x69\xd1\x80\x77\x53\xa4\x63\x2e\xe5\x95\x65\xe1\xcb\x76\x9f\xd2\xe5\xa4\x70\xd1\xf9\xb1\x89\xef\x48\x13\x27\x50\xe8\xf3\x28\x21\x43\x4f\xbf\x7f\x75\x87\xf3\xeb\xad\x0d\x5d\x68\xe5\x4d\x0d\x07\xd4\x07\x89\x07\xb8\x26\x7c\xf5\xfc\x52\x1b\xa6\x00\x34\xc3\xa7\x45\x47\x8e\xf4\xd1\x3b\xea\x3c\x51\x9d\x0e\x07\x83\x75\x1e\x34\x25\xfa\xb4\x48\x97\x24\xb6\x93\x36\x85\xeb\xbe\xd0\x66\x03\x3d\x8e\x87\x83\x41\xdd\x0c\x77\x76\x97\x8a\xe0\xa2\xf3\xf4\xe3\xfe\x9f\x5a\x39\x13\x76\xfe\x1e\xe7\x1d\x61\x0a\x71\xc3\x48\xde\x78\x87\x8e\x9e\xee\x2f\x85\xf6\x4c\x64\x1a\x82\x5a\xea\x36\x21\xc4\xa6\x50\x09\xf3\xfd\x0d\x93\x4f\x7d\x56\x4f\x4d\xa0\x64\xce\xdd\x87\x09\x8a\xcc\x41\x97\x64\x26\x34\x49\xcf\x32\x85\xd4\xf7\x03\x8b\x2e\x6c\xd3\x9f\x5e\xe2\x3c\x88\x88\x97\xe7\xb0\xeb\xf5\xfb\x4c\x3a\xa3\xa5\x13\x62\xe5\x20\x91\x42\x61\xec\x42\x18\x2d\xb9\x36\x6a\x5a\x70\xd3\x7b\x5b\x8d\x16\xbd\x06\xd8\x85\x74\xe6\xa2\x49\x66\xa5\x76\x29\x87\x7f\x34\x39\x3c\x3e\xdc\xff\x04\xbf\xe6\xf0\xf6\xfc\xe3\x09\x9d\xf7\xf8\xac\xaa\x96\x74\x97\x65\x74\x7e\x56\x55\x70\xf9\xc7\xe1\xf9\x21\xfc\x9a\x8f\x18\x17\x3f\xa2\xe5\xd1\x3f\x8d\xd4\x41\x7b\xca\xa3\x04\xb5\x3b\x2b\x8c\xc3\x89\x92\x31\x36\x1e\x47\xc7\x67\x21\x34\xbf\xcf\xcf\x38\xc5\xc7\x21\x8c\xc2\xd1\xb8\xd1\x56\x2b\xb8\xbc\x45\x8b\xfb\x4a\x14\x39\x32\x3e\xe4\xd0\xc8\x1f\xf8\xdc\xff\x7c\xdd\x0d\xdc\x02\x76\x7f\xd8\x7b\xa1\x0a\x3c\x11\x59\x26\xf5\x34\xe4\x52\x6b\x5b\xdd\x9e\xd4\x49\xbd\xb5\xae\x75\x7e\x9a\x67\x18\xae\x23\x80\x85\xda\x36\xc2\xf5\x78\xd0\x69\xeb\xbd\xbe\x4e\xec\xd5\xe4\x23\x1d\x98\x04\xeb\x64\x5c\x60\xf3\xdc\xce\x92\x5d\x32\xb8\xc2\xd5\xbe\xaf\xec\x6c\xe5\xbb\x2b\x87\x91\x69\x1a\x53\x86\xec\x48\x27\xd2\x62\x4c\x79\xeb\x17\xfe\x45\x12\x1f\xd3\xc0\x50\xe3\xb9\x17\xaa\x37\x54\xf0\x66\xfe\xd6\x9a\x59\x73\x04\x56\x58\x93\x6c\x0f\xa4\xb1\x27\x45\xef\x49\x0e\x57\xd7\x52\x3b\xb4\xa9\x88\xb1\xac\x16\xd3\xc5\x72\xb0\x3a\x81\x6c\x5e\x6c\x8d\x9f\x3a\xbb\xde\x74\x47\x87\x3f\xa9\x4c\xfd\x78\x7a\x80\x37\xc5\xf4\xc4\x24\xc8\x5a\xa9\x50\xde\x72\xa1\x28\x1d\xb4\xfb\xdc\x9c\x6c\xa3\x8b\x4b\x75\xfc\xb4\x34\x45\x67\x31\x93\xbe\x88\x85\x3e\x16\xb9\xf3\x6c\x7e\x74\xd0\xbd\xcf\x2c\xed\xd4\xf7\x1a\xbe\xd5\xac\xda\x1a\x2c\x8d\xf5\x7e\xd5\x62\xce\x13\x5f\x3d\xb6\xd2\xf0\xc9\x43\x7e\xd0\x71\xda\xfb\x14\x45\xd1\x98\xb5\xd0\xe4\xbf\xf9\xe5\xda\x42\xc0\x93\xed\x06\x45\xf5\xc5\xaa\xa7\x73\xb5\x9b\x9f\x9b\x84\xff\x36\x07\x1f\xbf\xf6\xed\xae\x35\x83\xf6\x8a\x92\xe8\xb7\x08\xba\xd4\xd2\x8d\xd6\xb3\xcf\xa6\x46\x41\x93\xcd\x32\x23\x2f\x20\x5f\x0b\x20\x25\xbe\xa2\xd5\x03\x90\xda\xfd\xfd\x6f\x3d\xe7\x68\xd3\x4f\xbe\x27\x22\x83\xab\xeb\xa2\x16\xa1\xf5\x86\xfe\x78\xa0\xeb\x97\xcc\x86\x9a\x59\x74\xc2\xa9\x71\x06\x78\x3e\xa9\xef\x3a\x4f\x7a\xea\xbd\x6c\x62\xef\xb3\x24\xea\x88\x25\x34\x48\xad\x0d\xe7\xa1\xb5\x93\xb9\x8e\xdf\x0a\xa9\xda\x32\x30\x8a\xbf\x94\xf2\x98\x93\xe0\xd7\xa6\x08\x4e\xdf\xe3\x7c\x71\x4b\x7e\xdd\x42\xb6\x74\xf7\xe7\xcf\x52\xdc\x74\x17\x9a\x7a\xa2\x9f\xa4\x53\x7e\x9a\xab\xd9\x71\x49\x9a\x64\x4d\xe4\xfd\xf0\xb2\x55\x05\x3c\xfa\xc5\x46\x45\xc4\xac\x55\x15\xf8\x53\xfb\x93\xd5\x38\x31\xef\xfc\xf6\xdb\xfa\x08\xff\x95\x76\x97\x77\xae\x5e\x5f\xd3\xde\x66\xaa\xbe\x1a\xb5\x61\xa9\xaa\xd1\xf5\x7a\xa8\x7a\x97\xc5\x45\x8e\x3c\x5b\x07\xe9\x4e\x29\x3f\xa1\x64\x2c\x3a\x2b\xf1\x1e\x9b\x7b\x1d\xf3\x73\xbe\xa6\x84\x80\x8e\xdb\x4b\xf7\x4d\x5d\x66\x9b\x6e\x15\xb6\x55\x35\xfe\x31\xfe\x6f\x06\xab\x2d\x5a\x40\xf7\x04\x9e\x92\x16\x05\xb7\x4c\x8c\x1d\x7a\x63\xed\xe7\xe6\x21\xe8\xdb\x7b\xac\x2e\x9a\xc4\x82\x27\x0c\x6a\x85\x5e\x7f\x97\x34\x57\xa8\x5c\xc1\x9a\xdf\xaa\xbe\x21\xd4\x9f\x90\x12\x99\xc9\x0a\xfe\x4c\x93\xf8\xbb\xc4\xe6\x9c\xa0\xd8\x75\x4b\x62\xe7\xd1\x3d\x6a\xbb\x8b\x59\x73\x01\xdc\x42\x9c\x2f\x7c\xb0\xeb\x23\xb5\xb5\x81\xc5\xc5\x6f\xb0\xe1\x0b\xd3\xe2\x9f\x25\x89\x79\x93\x3a\xb4\xdf\xf5\x75\xa9\xa6\x84\x4e\x1f\x67\xa5\x9a\x08\xb7\xfb\x95\xf3\x7f\x01\x00\x00\xff\xff\x45\x71\x40\x0b\xe3\x1a\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x69, 0xaa, 0x63, 0x2a, 0x19, 0xe1, 0x48, 0x1b, 0x3f, 0x19, 0xbb, 0xeb, 0xb4, 0x32, 0xfc, 0xc3, 0xbe, 0x87, 0x1c, 0xd8, 0xb3, 0xe6, 0x80, 0xb9, 0xf5, 0xff, 0xa6, 0x4, 0x15, 0x2b, 0x10, 0x27}}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8e\xd3\x30\x10\xc7\xcf\xf6\x53\x0c\x91\x56\x8d\x25\x2b\xcb\x5e\x91\x7a\xd8\xa5\x65\x15\x28\xfd\x2e\x08\x21\x0e\x6e\x3d\x6e\x2d\xa5\x4e\xf1\x47\xa1\x42\x7d\x77\xe4\x24\x6d\xb3\x4b\x41\x1c\xf6\x92\x8c\x3d\x33\x7f\xcf\x6f\x66\x6e\x6f\x61\x19\x74\x21\x17\x3b\x87\xd6\x4f\x02\xda\xc3\xc7\xc3\x6c\x32\xa8\x6f\x1d\x08\x88\x07\xe7\x85\xc7\x2d\x1a\x0f\xce\x5b\x6d\xd6\x10\x5c\xfc\xfa\x0d\x42\xa8\x12\x7b\xc2\x0b\xd8\xd9\x72\xaf\x25\xca\x8c\xaa\x60\x56\xd7\x75\x53\xa9\x05\x48\xab\xf7\x68\x5d\xd6\xd3\xa2\xc0\x95\xe7\xe0\xc5\xb2\xc0\xa1\xd8\x62\xa3\xcf\x21\xec\xa4\xf0\xc8\xe1\xc7\x46\x7b\x2c\xb4\xf3\xf0\xf5\x5b\xed\x63\xa7\x1a\x7e\x51\x72\xf1\x76\xe3\xed\x56\x98\x75\x81\x59\x2e\xd1\xf8\x49\x28\x3d\xce\x0a\xbd\xc2\xf8\x64\x36\x98\x70\x88\xff\xe9\xa4\xa5\xc9\x28\xb9\xbc\x7c\x5d\xe1\x8f\xe4\x73\x02\xa3\x94\x2c\x83\x82\x37\xed\xc4\x47\xf4\x0f\x41\x29\xb4\x29\xa3\x44\xa2\x42\xdb\x72\x8e\xc3\xc9\xb9\x0c\x2a\xa6\xef\x85\x85\x55\x59\x84\xad\x71\x0d\x14\x25\x5a\x41\x81\x26\xbd\xd4\x08\xaf\xba\xf0\x3a\xc2\x92\x53\x68\xb7\x09\x76\xd9\xfb\x52\xb7\x42\x39\x24\x3c\x61\x94\x1c\xe9\x59\xa6\x6e\x23\x83\xee\x49\x43\x6d\x7d\xf6\x6e\x67\xb5\xf1\x2a\xa5\x84\x44\x02\x1e\xff\x49\x3e\x9c\xf5\xa7\x73\xc8\x1f\x87\xa3\x69\x1f\xf2\xe1\x7c\x04\x37\x0e\xd2\x1b\xc7\xe0\xd3\xfd\x60\xd1\x9f\x55\x76\x52\x05\x9f\x7b\x50\x9d\x9a\xb2\x2a\xbb\x05\x5b\x88\x15\x6e\xca\x42\xa2\x75\x55\x13\x17\x0e\x73\x23\xf1\x67\xdb\xc1\x9f\xb1\x72\xb8\xe3\x70\xc7\xa2\x14\xa3\x84\x58\xf4\xc1\x1a\x58\x06\x95\xcd\x2a\xe2\xb4\xa1\x7b\x46\xd1\x40\x9c\x19\xfe\x52\x3c\x8c\x86\xd0\x5b\x8c\x07\xf9\xdb\xfb\x79\x1f\x3e\xf4\xbf\xc0\x62\xdc\x8b\x66\x45\xf5\x04\xaa\xc5\xf4\x62\x48\x71\xe2\xaa\xb4\xa0\x39\xec\xe3\xd6\x58\x61\xd6\xd8\x2c\x7a\x35\x1b\xad\x40\x5f\xa6\x1d\xa9\xb2\xcf\x56\x7b\x7c\x38\x78\x4c\x3b\xbc\x13\x5b\x72\xa4\x84\x7c\x8f\x8b\x29\x9f\x2e\xde\x3f\x36\x76\xcf\x68\x4b\xac\x69\x64\xad\x71\xcd\x93\x40\xb7\x69\x5a\x9a\xfc\x67\x66\x5d\x20\xeb\x34\xd3\xb9\x36\xb6\x23\xfd\x1d\x00\x00\xff\xff\x4c\x0d\x4e\x35\x6a\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3e, 0xcc, 0xa0, 0x4c, 0x62, 0x37, 0xa8, 0x14, 0xaa, 0x37, 0x45, 0xfd, 0x34, 0x76, 0xe0, 0x86, 0x6, 0x53, 0x4, 0xc, 0xf3, 0x3, 0xab, 0xc9, 0xcc, 0x8f, 0xb8, 0xe0, 0xcd, 0xca, 0x43, 0x5}}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xe3\x36\x16\x7e\x96\x7e\xc5\xa9\x81\xe9\x4a\x53\x85\x19\xa0\xc0\x3e\xa4\x10\x82\x89\xe3\x14\x41\x27\x37\x3b\xbb\xc5\xa2\x29\x5a\x46\xa2\x13\x22\x12\xa9\x90\x54\x1c\x6f\x90\xff\xbe\x38\xa4\x2e\x94\x63\x79\x33\xbb\xf3\xd8\xa7\x44\xe2\xc7\x73\xf9\xce\xe1\xe1\x27\x3f\x51\x05\xea\xee\xf9\x6c\xbd\xb8\xfa\xf2\xc0\xd6\x90\x82\x62\x77\xec\xb9\x22\x67\xb5\x36\x53\x59\x56\xbc\x60\xd1\x9f\xd1\x61\x19\x47\x51\x72\x23\xe2\xc3\x1b\xfd\xc3\xf4\xe2\x7c\x71\x3d\xff\x7c\x7a\x7e\x4d\x3e\x1e\x9e\x5c\xcc\x67\xa7\x3f\x9f\xc3\x2f\xb3\x7f\x91\x8f\x87\x37\x22\xfe\xe1\xcf\x38\x0c\xcd\xba\x62\x50\xae\xf5\x63\x71\xcd\xb4\x61\x0a\xb4\x51\x75\x66\xe0\x25\x0c\xf2\xdb\xa9\x14\x02\x3e\xea\xc7\x82\x1c\x1f\x85\xf8\xe2\x9c\x96\x0c\x10\xc2\xc5\x5d\x18\xdc\x4b\x6d\x00\xfa\xe7\x5a\x33\xe5\x3f\x57\x54\x6b\xff\x59\xeb\xa2\x94\x39\xeb\xd7\xa5\xb2\xfb\xb9\x30\x61\x18\xc8\xca\x70\x29\x4e\x78\xd1\x01\xc2\xc0\x30\x6d\x8e\x8f\xac\xd7\xce\xc8\x03\xaf\x16\x57\x5f\xa6\x65\x0e\xb7\x52\x16\xe1\x6b\x18\x2e\x6b\x91\x01\x17\xdc\x44\xb1\x8b\xfb\x8c\x72\x01\x29\x7c\xef\xe5\xf5\xf2\xda\x21\xa3\x12\x3e\x7a\x2b\x31\x68\x66\xea\x2a\x8a\x81\x29\x25\x15\x5a\x40\xae\x99\x52\xee\x45\x18\x06\x4f\xbc\x62\x8a\x2c\x98\x39\x66\x4b\x5a\x17\x26\x9a\xd8\xfd\xa4\x49\x68\x92\xc0\xc4\xa8\x9a\x4d\xe2\x71\x28\xe6\x3a\x49\xe0\xc7\x1f\x3f\xfd\x3d\x0e\xc3\xa0\x24\x0d\x99\x29\xb8\x1d\x3f\x33\xb3\xb0\x19\xb6\x1b\xf2\x5b\x41\x4b\x6b\xb2\x24\x96\xe8\x51\x24\xae\x3a\x9c\x2d\xc0\x28\x0e\x57\x1d\xce\x16\x66\x14\x87\xab\x0d\x0e\x0b\xe4\xe1\x4e\xc5\x30\x1f\x0b\x6a\xab\x3a\x6a\xaf\x65\xc9\xa2\xbd\x8a\x8e\x6e\x40\x8c\x9f\xbe\x57\x72\x6f\xcf\x91\x94\x45\xe7\xe2\x81\x57\xfa\xb1\xc8\xca\x7c\x82\xec\x62\xed\x52\x78\xa2\x05\x25\x47\xec\x8e\x8b\x7f\xd2\x82\xe7\x14\xdb\x2b\x8a\x49\xf3\xc0\xa2\x30\x08\x2c\xc4\x39\x3f\x97\x66\x56\x56\x66\x1d\x39\x1a\x13\x18\xb0\x96\x8c\x82\x91\xfd\x0e\xec\x4a\xd1\x81\xcf\xa5\x89\xec\x3f\xb3\xc7\x9a\x16\x3a\x72\x8c\x26\xf0\xa9\xdb\xe0\x68\xdc\x61\xde\xb5\x49\x87\x6f\x69\x19\xdf\xd0\xb0\xdd\xed\xe8\xd8\x4f\xc2\x20\x26\xd3\x7b\x96\x3d\x44\xc8\x11\x5f\xda\x16\xff\x2e\x05\xc1\x0b\x6c\xfa\x40\x31\x53\x2b\x81\x6f\xc3\xe0\x35\x0c\x83\xfd\x7d\x98\x2a\x46\x0d\x03\x0a\x8a\x8a\x5c\x96\xfc\xdf\x2c\x87\xfc\x16\x30\x06\x62\x4d\x14\x4c\x44\x7e\x51\x63\x48\x53\xf8\x64\xcd\x6d\xd4\xba\xb3\x40\x16\x86\xde\x16\xcc\x2d\x74\x19\xc6\xce\x67\x13\x55\x0a\x25\x29\xe9\x03\xbb\xe8\x66\x42\x14\xff\x34\x1e\xaf\x54\x9a\xfc\xaa\x68\x15\x31\x85\x85\xcb\x64\x5d\xe4\xe2\x6f\x06\xd0\x04\xb8\xb9\x02\x4b\x5e\xd8\x76\x6a\xbc\x7c\x37\x68\x2b\x34\xe7\xb9\xce\x95\xac\xae\x6d\xf0\x5b\xdc\x0e\x78\x0a\x5e\x87\x3b\x33\x4b\xd8\xbb\xf7\x86\x41\x90\xd7\x65\x85\x21\x1c\xa4\xc0\x9e\x59\x46\xa6\xb2\x2c\xa9\xc8\x9b\xce\xc6\xd5\x49\x82\x21\xb9\x71\xa2\x1d\x17\x09\x4c\xf6\xf6\x84\xdc\xcb\xa9\xa1\x6e\xb9\x25\x31\x70\x11\x8c\x5b\x1c\xb3\x86\xa6\x6e\xa9\x66\x76\xdd\x2b\x28\xc6\xa8\x12\x58\xa1\x39\x2e\xc9\x25\xaf\x58\x14\xf7\x71\x2f\x4c\x8e\x39\x1e\xa4\xf0\xfd\xed\xda\x30\x4d\x8e\xea\xe5\xd2\x8e\x5b\x2f\x94\x71\x50\x6f\x88\x2c\x4c\x2e\x6b\x1c\x37\xab\xe1\x4b\x47\xed\xc0\x5d\xe8\x1b\x47\x8c\x1d\xf7\x82\xad\x4e\x7e\x61\xeb\x63\xa6\x8d\x92\x6b\xa6\x22\xef\xba\x4c\x40\xc5\x9b\x9b\x9c\xe1\x8d\x20\x43\xbf\x9e\x7d\x14\x54\x99\xdd\xe5\xdc\x68\xc1\x25\xe5\x05\xcb\xc1\x48\xd0\xb8\x17\xba\x62\x42\xe6\xaa\x81\xad\x38\x6c\x1e\x3f\xb6\x6f\xe2\x6e\xc3\xd5\xb6\xc4\x7e\xa5\x7c\xab\xa3\x65\x69\xc8\xa5\xe2\xc2\x14\x02\x3d\xc4\x9b\xef\x06\xd5\x68\x66\x50\x14\xc7\xef\x8c\x71\x45\xb9\x81\xa5\x54\xa3\xac\x84\x41\xf0\x07\x36\x02\x99\x16\x52\xb3\x28\x86\xfd\x7d\xf8\xbc\x44\x75\xd2\x9e\x16\xae\x21\x97\x82\x25\x90\x21\x02\xcc\x3d\x83\x95\xe2\x86\x01\x13\x39\xc8\xa5\x7d\x51\xf1\x8a\x85\xdb\x19\xfe\x5f\xf3\xde\x68\x96\xff\x33\xf3\xcd\x5e\xc0\xc4\x1b\x23\x82\x17\x3b\xf4\x8a\x2e\xce\x64\xce\x22\x4f\x4c\xc5\xcd\x5f\x4c\x43\xaf\xb8\xc9\xee\xc1\xae\xbe\x84\x41\x46\x35\x6b\xf4\xc9\x41\x3f\x35\x27\xf3\xd9\xd5\x3f\x4e\xe7\xb3\xe3\x49\x8b\x58\xd2\x42\x0f\x21\xc7\xa7\x8b\xcf\x47\x5f\x2c\xa4\x19\x18\xfe\xea\xe5\x7c\x76\x32\x9b\x3b\x0b\x3b\xc4\xd5\x70\xd4\x78\x61\x36\x76\x90\xde\x45\x85\xfc\x2e\x23\x1c\x43\x0d\x7c\x0f\xe7\x75\xfa\x41\xdb\x71\xd4\x4b\xc3\x78\xdc\xd1\xe6\x7d\xd1\xcb\x39\x53\x56\x09\x34\x03\x88\xcb\xda\xf0\x82\x5c\xb3\xb2\xb2\xb0\x09\x8a\x37\x67\xbf\xbd\x21\x76\xdd\x8c\xa3\x95\x75\x9d\xb1\xf5\xb2\xd1\xd7\xd3\x4b\x74\x6d\x09\x0e\x83\x3f\x92\xa6\x1d\xa5\xc6\x93\x6e\x1a\x0d\xe1\x1c\x4b\x4d\x4e\x35\xde\xe6\xcf\x5c\x1b\xdb\x82\xee\x6e\xb2\x36\x52\xc0\x2a\x86\xc1\x2b\xb0\x42\x33\xf8\x8a\x38\xed\x8d\x08\x42\x1a\x9c\x0f\x06\xca\x4e\x33\x62\x80\x58\x81\x93\xaa\xe9\x70\xcb\xd5\xe4\xb7\xac\xe0\x4c\x98\xdf\x11\xd2\x2f\x2f\x9b\x55\xdc\x9c\x7e\xd0\x37\xc2\x16\xa7\x09\xfe\x2d\x0c\xb5\x4d\xfa\x21\x6f\x60\xf8\xb4\x15\x86\x02\xab\xb7\x86\x4f\xb1\x27\x2d\x50\x8c\xc6\x98\xa3\x13\x15\x5b\xbc\x50\xad\x57\x52\xe5\xbd\x09\xbb\x05\x53\xdb\x82\xd6\xba\xd8\xc3\x83\xd1\xa3\xbb\xc3\xd4\x2a\xa5\xd8\xb9\x77\x94\x0f\x7d\x76\xfc\x54\x4a\x1a\x99\xc9\x22\x35\x59\xb5\x8b\xc6\x6e\xc0\xfd\xc5\xe4\x57\x30\xe9\x1f\x78\x6c\xfa\xb2\x22\x56\x2b\xc6\xfd\x7c\xc4\x77\xcd\xe5\x30\x3e\x11\x86\x62\xac\x9f\x07\x38\x7a\xf1\x3c\xfa\x93\xa7\x39\xbf\xad\x0a\x82\x0f\xfa\xa7\x37\x4a\xa8\x75\x5e\x12\x55\x8b\x69\x99\x47\xfa\xb1\x68\x75\xf6\x64\x47\x1c\xbe\x9c\xdc\x1d\x05\x22\xfb\x18\xf0\x80\xe3\x1c\xd0\xdf\x34\x1a\xc3\xa8\xca\xe5\x4a\xf8\xb1\xf0\xa5\xd5\x90\xf6\x7b\xff\xed\x3c\x69\x97\x3a\xc6\xff\xab\x88\x3e\xf8\x7a\x15\xed\x5d\x7e\x52\x93\x39\x2b\xe5\x13\xb6\xd2\xbb\x46\x7f\x4b\x00\x0a\xc1\xa4\xbd\x55\x9b\xab\x26\x01\xaa\xee\x34\x10\x42\xda\x9b\xb2\xcb\xda\x2e\xa4\x40\xab\x8a\x89\x3c\xfa\xed\x77\x07\x78\xd9\x94\xc7\xaf\xce\x04\x21\x04\x1b\x30\xdb\xa2\xac\x1b\x8f\x1e\x0e\x61\x9d\x30\x75\x76\x35\x39\x67\xab\x39\xa3\x39\x53\x2e\x52\xb4\xa6\x9d\xe8\xdd\x26\x9f\xf5\xb8\xb2\xce\x7c\xb9\xec\x4c\x74\x2f\xdd\xdd\xe2\x36\x87\x5e\x3d\x70\x79\x5e\x8b\xb7\xa5\xf0\xf5\x4d\x7b\xa1\xa9\x5a\x08\x2e\xee\x0e\x26\x1d\x9b\x2e\xb7\x78\x08\x77\xae\x7d\x15\xb4\xb1\xba\xa1\x91\x36\xbf\x30\xdf\x23\x76\x32\x29\xb0\x55\xa3\xe6\x67\xa8\xc4\x95\x2f\x1e\xef\xda\x8d\xa6\x4d\xac\x79\xeb\x6e\xf8\xb3\x4e\xd0\x23\x1a\xce\x1e\x0b\x72\x51\x31\xd1\x7f\x28\xe5\x8a\x3f\x31\x45\xec\x47\xc4\x51\xcd\x8b\xfc\xaa\x66\x6a\xdd\x24\xd4\xfe\x4e\xe0\xc6\xe4\xf0\x74\xb6\xd3\xbc\x1d\xd7\xcd\x78\xf4\x86\xe2\xb0\x06\x3d\x11\xc9\x1b\x76\x86\x89\xbc\x86\xff\x09\x00\x00\xff\xff\xdd\x6c\xf4\x1c\x09\x14\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xa, 0x66, 0xf8, 0x27, 0x50, 0x1f, 0xc8, 0x5, 0xd5, 0x88, 0xc1, 0x79, 0xe3, 0x28, 0xff, 0xb6, 0x1b, 0xf9, 0xfa, 0x1, 0x4b, 0xf0, 0xeb, 0x10, 0x5, 0xf8, 0xdc, 0xfe, 0xba, 0x82, 0x80}}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x10, 0xc4, 0x71, 0xaf, 0xd9, 0x16, 0x41, 0x8b, 0x4b, 0xfc, 0xe8, 0xba, 0xfd, 0xfa, 0x4d, 0x2c, 0x1, 0xd1, 0x0, 0xe1, 0xb0, 0x78, 0xee, 0x7f, 0xd0, 0x65, 0xf3, 0xa1, 0x43, 0xba, 0x3c, 0xe7}}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\xcd\x6e\xdb\x3c\x10\x3c\x8b\x4f\xb1\x9f\xf1\xb5\xa0\x0a\x85\x69\xaf\x29\x7c\xc8\xdf\x21\x68\x6b\xb8\xb1\x7c\x2e\x18\x69\xe5\x10\xa6\x49\x95\x5c\xd5\x76\x05\xbe\x7b\x41\xc9\x76\x9c\xd8\x69\x73\x68\x0f\x39\xf8\x87\x8b\xd9\x9d\x99\x5d\x2e\xdb\xf6\x04\xfe\x97\x5a\x49\x0f\x67\x43\x10\xe7\xf1\x1f\x7a\x91\xcb\x3b\x8d\xd0\xff\x88\x91\x5c\x60\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xd3\x7a\xac\x1b\x27\x75\x08\xd3\xda\xa3\x23\x4e\xf0\x2e\x02\x94\x99\x89\x3c\x85\x96\x25\x24\xc6\xd2\x49\xad\x51\xf3\x94\xb1\x44\x55\xa0\xd1\xf0\x5d\x81\x2b\xbb\x34\x13\x65\x66\x8d\x96\x2e\x84\x4b\xab\x9b\x85\xf1\x29\x0c\x87\xbf\x83\x8d\x9d\x5a\x48\xb7\xfe\x84\xeb\x5d\x42\xcb\x92\x84\xc4\x64\xae\x6a\x3e\x88\xdf\xb5\x32\x33\xa0\xce\xc3\x52\xd1\x3d\x58\xa3\xd7\x50\xf7\x79\x30\xc7\x35\x14\x7d\xe6\x20\x65\x49\xd8\xc9\x5a\xac\x27\x5f\x3f\xef\x99\x7b\xa0\x9c\x1a\xf5\xbd\xc1\x7d\x7d\xef\xff\xc8\x69\x2c\x34\x5d\xda\x96\x0c\xc8\x42\x61\x4d\xa5\x55\x41\x60\x4d\xcf\xcd\x12\x8f\x58\xc6\xde\x3b\x69\x4a\xbb\x50\x3f\x51\x8c\x70\x39\x41\x2c\x79\xca\x92\x1f\xd2\x01\xba\xee\x63\x1d\x4b\x4e\x4f\xe1\x9c\x08\x17\x35\x01\xdd\x23\xdc\x8c\x26\xd7\xb7\x39\x78\x55\x22\xd8\x0a\xa4\x81\xe9\x38\x46\x58\x62\x63\xc5\xa3\x56\xda\xde\x6f\x2c\xba\xcf\x39\x21\xd7\x14\xc4\xa3\x98\x0c\xde\xda\x0c\x9e\x69\xfe\xd5\x45\xbe\xae\xd1\x67\x50\x49\xed\x31\xfd\xd8\x15\xfa\x6f\x08\x46\xe9\x4d\x47\xae\xa3\xd4\x8a\x0f\xa6\xa6\xeb\x05\xd9\x07\x96\xe3\x8a\xc0\x77\xdc\x67\xf0\xc6\x0f\xb2\x58\x6f\xd3\x98\xb6\x55\x15\x18\x4b\x20\x46\xf6\xd2\x1a\xc2\x15\x85\x50\xd0\x2a\x5a\x2b\xfa\xb3\xb8\x90\xc5\x7c\xe6\x6c\x63\x4a\x9e\xb6\x2d\x9a\x32\x04\x96\xf4\x90\x2f\x8d\xa7\x7c\xc5\xbb\x2a\xfb\x15\x0e\x02\x77\x56\x69\x71\x81\x33\x65\xba\x1a\xda\xe3\x7e\x2c\x5f\xf1\x82\x56\x59\x34\xb8\x65\x78\x11\x28\x65\x49\x89\x15\x3a\x88\x6b\xc3\x53\x68\xe1\x1b\x0c\x81\x56\xe2\xd6\x6a\x7d\x27\x8b\x39\x4f\x21\xc4\x11\xef\x86\x61\xc5\x66\x8b\x9e\x33\x1e\x87\x82\xa6\x84\x93\x10\x20\x9e\x3a\xfe\x1b\x53\xa1\xe3\xe9\xe3\xd3\xcb\xe6\xd2\x74\x74\xc7\x87\x72\x30\x8d\xc2\x36\x86\xba\xc0\x93\xab\xb5\x7d\x02\x78\x2a\x2e\x23\xe6\x85\xf2\x1f\x9c\x1f\xaa\xe4\x5b\xda\x08\xe9\x88\x23\xe8\xc3\x23\xc8\x60\x29\x4d\x5c\x23\x04\x87\x85\x75\x65\x06\x33\x4b\x67\x83\xac\xc7\x6f\x44\x3f\xd9\x97\xe9\xf8\xea\x3c\xbf\x3e\xb6\x2f\x7f\x6d\x23\x9e\x85\x1d\xbc\x5a\x42\x88\x7f\xba\x3e\xaf\xef\x5e\xbd\x92\x6b\x15\xd8\xaf\x00\x00\x00\xff\xff\x05\xd5\xa1\x1c\x35\x07\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x92, 0xf5, 0x7f, 0x80, 0xcc, 0x1f, 0xaf, 0x82, 0xb2, 0x12, 0x19, 0x6, 0xd4, 0x69, 0xc3, 0x57, 0xe4, 0x53, 0x13, 0x49, 0x4c, 0x87, 0xe7, 0xd9, 0xe0, 0xe1, 0xe6, 0xf5, 0x6c, 0x40, 0x20}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"templates/17_upsert.go.tpl": templates17_upsertGoTpl,

	"templates/singleton/mysql_upsert.go.tpl": templatesSingletonMysql_upsertGoTpl,

	"templates_test/singleton/mysql_main_test.go.tpl": templates_testSingletonMysql_main_testGoTpl,

	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,

	"templates_test/upsert.go.tpl": templates_testUpsertGoTpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl":   &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
