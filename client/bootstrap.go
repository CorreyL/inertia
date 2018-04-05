// Code generated by go-bindata.
// sources:
// client/bootstrap/daemon-down.sh
// client/bootstrap/daemon-up.sh
// client/bootstrap/docker.sh
// client/bootstrap/keygen.sh
// client/bootstrap/token.sh
// DO NOT EDIT!

package client

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

var _clientBootstrapDaemonDownSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xb1\x6e\xc2\x30\x10\xc6\xf1\xfd\x9e\xe2\x2b\xb0\x1a\x9e\x20\x43\x2a\x22\x54\xa9\xb8\x12\x52\x87\x4e\x60\xe2\x73\x38\xd1\x9c\xa9\x7d\xa8\xaf\x5f\x45\x74\x88\xd8\x6e\xf8\xdf\x4f\xdf\xf2\x65\x73\x16\xdd\xd4\x0b\xd1\x12\xaf\xa1\x4a\x8f\xda\x17\xb9\x19\x52\x2e\x38\x17\xd1\x41\x74\x40\xcc\xbf\x0a\xbb\x30\x62\xe0\x31\xeb\x9a\xa8\xb2\xc1\x31\xd1\xb6\xed\xf6\x1f\xfe\xe8\xdb\x7d\xd7\x88\x72\x31\x09\xee\x11\x4d\xe2\x8e\xed\xff\x05\x7d\x56\x0b\x53\x81\xa0\x11\x16\xae\x0c\xb1\x07\x2c\x69\x3a\xa5\xa2\xdc\x55\x45\x87\x35\xb5\xef\x87\xae\xdd\x7e\x1d\x0f\x9f\xde\xbf\xf9\x5d\x73\xaa\xf7\x98\x11\x73\x7f\xe5\x82\x5b\x85\xfb\x81\x73\x49\xbe\x8d\x0b\x16\x1a\x46\x6e\x56\xb3\x21\x8b\x13\xcd\xfb\x32\xc2\x25\xac\x9e\x4c\xfa\x0b\x00\x00\xff\xff\x62\xed\x0f\x4a\xfb\x00\x00\x00")

func clientBootstrapDaemonDownShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonDownSh,
		"client/bootstrap/daemon-down.sh",
	)
}

func clientBootstrapDaemonDownSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonDownShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-down.sh", size: 251, mode: os.FileMode(420), modTime: time.Unix(1522901431, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDaemonUpSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x5d\x6f\xe3\x36\x10\x7c\xd7\xaf\x98\xc8\x4e\xd3\x02\x95\xd4\x5c\x5f\x0e\x3e\xe8\x00\xf7\x22\x5c\x82\x26\x76\x60\xe7\x5a\x14\x69\xe0\xa3\xa5\x75\xc4\x5a\x22\x75\xe4\xca\x6e\xfb\xeb\x0b\x52\x72\xfc\x91\x00\x7d\x92\x4d\xee\x0e\x67\x67\x67\x77\x70\x96\x2c\xa5\x4a\x6c\x19\x04\x03\xfc\x22\xac\xcc\x61\x73\x23\x1b\xc6\x4a\x1b\x58\x62\x96\xea\x19\x6d\x83\x1b\x45\x86\xa5\x80\xa1\x6f\xad\x34\x54\x93\x62\x8b\xef\x0b\x69\x28\x67\x6d\x24\xd9\x1f\x41\x9c\xff\x10\x0c\x20\x54\x81\xa5\x91\xca\x25\x72\x49\x28\x04\xd5\x5a\x41\xab\x4a\x2a\x8a\x83\xc0\x12\x23\x22\xf7\xde\x17\x4b\x06\xc2\x3c\xb7\x1e\x2d\x0e\xae\xc6\xd9\xdd\x74\xb2\x98\x65\xb7\xd9\x78\x9e\xa5\xe7\x8f\x97\x4f\x76\x77\x78\x3f\x9d\x3d\xa4\xe7\x8f\xef\x9e\x6c\x70\x3d\x9d\x3f\x2c\xc6\x57\x57\xb3\x6c\x3e\x4f\xcf\x1f\x7f\x7e\xb2\x0e\x6c\x47\x50\xd6\xe2\x99\x50\x10\x0b\x59\xed\x31\x27\xe3\xbb\x2c\x95\x5d\x48\xd4\x31\x0a\x6e\xee\xc6\x9f\xb3\xb4\x5d\xe6\x95\x68\x55\x5e\x36\xa2\x48\xfa\x88\xd1\xf0\x98\x8a\xc7\x67\x14\x9a\xac\xba\x60\xd4\x82\x99\x0c\xb6\xa5\x60\x34\xda\xf0\x61\x99\xa6\x55\x16\x5a\x41\x2a\x7f\x9a\x6b\xc5\xc2\xa1\x3a\x61\x2c\x2a\xad\x9e\xdd\x57\x32\xa4\x45\x2d\x9a\x86\x0a\xb0\xee\x43\x8d\x13\x13\x07\x05\xc7\xc1\xa7\xe9\xe4\x61\x7c\x33\xc9\x66\x9d\x00\xef\x7f\x7a\x7f\xe9\xc8\xcc\x89\x5d\x53\x0e\xe4\x8f\x83\x7a\x5d\x48\x83\xa8\xc1\xf0\x7a\x7a\x97\x25\x8d\xd1\x7f\x51\xce\xa7\xc7\xd6\x56\xa7\x47\x71\x5f\xb5\x03\xfe\x54\x52\xbe\x86\x5c\x41\x54\x86\x44\xf1\x8f\xab\xc7\x77\xd2\x75\x95\xc5\x9a\x50\xe8\xad\x02\xfd\x2d\xad\x77\x46\x57\x75\x1c\x8c\x6f\x67\xd9\xf8\xea\x8f\xc5\xec\xcb\x64\x72\x33\xf9\x9c\x7e\xb5\x6d\xa1\x51\xe8\x7c\x4d\x06\x8d\x45\xf4\x0d\x51\xb4\x92\x95\xd3\x2d\x54\xa2\xa6\x74\x78\xd0\x99\xf0\x6b\x20\x57\x78\xc4\x19\xa2\x7f\x11\x0e\x4f\xc0\x42\x3c\x7d\x70\x0a\xa9\x00\x00\x28\x2f\x35\xc2\xfb\xb6\x73\xe6\x0b\x91\x5d\xfb\xfb\x36\xb0\x86\xad\x88\x9a\xd0\xe7\x1c\x92\x31\x35\xa2\x15\x4e\xdf\xc0\x47\x24\x05\x6d\x12\xd5\x56\x15\xde\x7d\xfc\xee\x32\x58\xc9\x0f\x41\xc7\x2a\x3c\x71\x43\x88\xb3\x14\x21\x93\xe5\x63\x66\x03\x5c\xe9\xad\xaa\xb4\x28\xfc\x94\x90\x65\x2a\x76\x7c\xbc\x2b\xe3\x83\x02\x76\xa1\x8e\xfb\xd0\x3b\xf1\x35\xd5\xc6\x91\xe9\x2e\x5f\xf3\xa3\xca\x52\xff\xec\xad\x7b\xd2\xf1\xc1\xb2\x95\x55\x01\x76\xc6\xb4\xa5\x6e\xab\x02\xa5\xd8\x10\x96\x44\x0a\x36\x6f\x2e\x0a\x48\xc5\xba\x4f\x73\xa6\xfb\xed\x7e\x0e\xc1\x48\x3a\x9a\xd1\x2b\x9a\xb7\xff\x43\xd1\x57\x1b\xc9\x63\x80\xb7\xc4\x74\xe6\x9a\xb5\x6a\x3f\x10\xd8\x4a\x2e\x21\xf2\x9c\xac\xdd\x8d\x40\xa9\x2d\xef\x90\xad\xfb\xb0\xf7\x5d\x30\x80\xa1\x8a\x36\x42\x71\x1f\xb2\xf7\xbd\x4b\x15\x55\xa5\xb7\x7e\x61\xed\xd1\xdd\x2f\xa3\xab\xd8\x0f\x0b\x79\x74\xd7\xf0\xbb\xcc\xc7\xd5\xda\xbc\xec\x08\x37\xad\xa5\xde\x82\x4b\x69\xb1\xd5\x66\x6d\x47\xc1\x00\x25\x73\x63\x47\x49\xf2\x2c\xb9\x6c\x97\x71\xae\xeb\xe4\xad\x4d\x31\x28\xf5\x36\x92\x1c\xf9\xbc\xa0\xd3\x6c\xd6\x8f\xcc\xcb\xe2\xeb\x96\xc4\xf0\x60\xb0\xc3\xe0\xc8\x92\xad\x42\x54\x20\x8a\x4c\x8d\x3f\xbd\xc4\x51\xb3\x77\x9d\x8f\x1f\x85\xc3\xe3\x55\x10\xee\x22\x37\x48\x36\xc2\x24\xa6\x55\x49\x07\x17\x3b\xe5\x46\x6f\x1d\xee\x53\x42\x3f\xf9\xe1\x28\x11\x4d\x93\x78\x49\xfb\x2b\x82\xbb\x48\xfb\xfb\xfd\xe9\x7c\x7e\xbd\xf8\x75\x32\xfd\x7d\xb2\x70\xdb\x77\x9e\x5e\xbc\x64\x26\xb1\xb5\x65\xb2\x56\x7a\xab\x16\xee\xbf\xbd\xd8\x65\x45\x6e\xce\xf7\x75\xf8\x41\xef\xef\xc2\xde\x4f\x9e\xc8\x7e\x9b\x87\xaf\x9d\xf3\x5f\x00\x00\x00\xff\xff\x12\x2d\x88\x33\xa1\x06\x00\x00")

func clientBootstrapDaemonUpShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDaemonUpSh,
		"client/bootstrap/daemon-up.sh",
	)
}

func clientBootstrapDaemonUpSh() (*asset, error) {
	bytes, err := clientBootstrapDaemonUpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/daemon-up.sh", size: 1697, mode: os.FileMode(493), modTime: time.Unix(1522901431, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapDockerSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x5d\x6f\xd4\x30\x10\x7c\xf7\xaf\x18\x7a\xa7\x16\x24\x7c\xa1\xf7\x48\x55\xa4\x72\x77\x4f\x54\x3a\xe9\xae\x7d\x46\xae\xb3\x49\xac\x26\x76\xea\xdd\xb4\x04\xc4\x7f\x47\xf9\x84\x82\x80\xe6\xc9\xf1\xcc\xce\x8e\x67\x77\xf1\x2a\xb9\x73\x3e\xe1\x42\xa9\x05\x3e\x86\x20\x2c\xd1\xd4\x0c\x83\xca\xd8\xc2\x79\x42\x16\x22\xd2\x60\xef\x29\xc2\xf8\x74\x3c\x6a\x1b\xaa\x3a\x30\xad\x94\x62\x12\x68\x52\x6a\xbb\xdf\x7c\xda\x1d\x3e\x1f\xf7\xb7\x87\xcd\xee\x32\x27\x59\x0d\xd4\x95\x0d\xd5\x04\x6e\x77\xc7\x9b\xcb\xb3\x44\xaa\x3a\xc9\x49\xf4\x48\xe0\xe2\xac\xeb\x7e\xbc\x77\x35\x9c\x67\x31\x65\x69\xc4\x05\x0f\x97\x61\x3b\x74\x76\x0c\x53\x46\x32\x69\x3b\x31\x28\x5d\x29\x97\xa1\x30\x5c\x4c\xf6\xd6\x1f\x92\x94\x1e\x13\xdf\x94\xe5\x05\xa4\x20\xaf\x00\x80\xbe\x38\xc1\x3b\x95\xb9\x0b\xa5\x32\x12\x5b\x64\xae\xa4\xd7\x6f\xf0\xad\x47\x17\xb8\x8a\x39\xbf\x1f\xcf\xc0\xf2\x1c\x1c\x9a\x68\x09\xb7\x87\xeb\x9f\xb7\x6b\xa4\xc4\xe2\xfc\xe0\xab\x53\x58\xf5\xe0\xe4\xc0\x36\xb1\xfc\x4b\xff\xee\xe3\x26\x0d\x03\x47\x67\x7c\xbc\xee\xba\xe8\x80\x93\xe5\xfa\x64\xb0\x58\x4e\x3a\x4f\x39\xc9\xff\x74\x7a\x8e\xde\xf7\xe5\x58\x9e\x8f\x0a\x4c\x33\x2b\x92\x34\xd1\x63\x40\xba\x77\x7f\xef\xe2\xbd\xaa\xcc\xd7\xe0\xb1\xdb\x1c\x87\x08\xbd\x25\x46\xa4\x87\xc6\x45\x82\x6d\x58\x42\x35\x65\xdb\x05\x9b\x47\xaa\xa1\x1f\xa6\xb2\x84\xc4\x26\xdc\xb2\x50\xa5\x23\x95\x64\x98\x7e\xf1\xd6\xfb\x6a\x9b\x59\x00\xba\x1d\x87\xa2\x66\x67\x0b\xdc\xc4\x16\x12\x90\x86\x27\x5f\x06\x93\xa2\x61\xe7\xf3\x21\x96\x10\xfb\x67\xbd\x1d\x99\x77\x94\x85\x48\x88\xc4\x21\x4a\x47\x92\x30\x49\x4f\x25\x73\xfe\xf3\x4c\xb1\x7c\xb6\x82\xf3\x6f\xb7\x74\xbf\xe7\x58\x3c\x43\xff\x8c\xd0\xd4\xa2\xbb\x98\x9b\x3a\x35\x42\x38\x3d\x9d\x6f\xf4\xbc\x81\xbd\x8d\xb9\xe2\x45\x36\xfe\x69\x20\x73\x2a\x73\x4a\xf5\x59\x32\xc5\x47\x67\x69\xda\x6c\x16\x13\x45\xfd\x08\x00\x00\xff\xff\x55\xed\x48\x14\xac\x03\x00\x00")

func clientBootstrapDockerShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapDockerSh,
		"client/bootstrap/docker.sh",
	)
}

func clientBootstrapDockerSh() (*asset, error) {
	bytes, err := clientBootstrapDockerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/docker.sh", size: 940, mode: os.FileMode(493), modTime: time.Unix(1522901436, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapKeygenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xdc\x30\x10\x46\xef\xf3\x2b\xbe\xb0\x81\x5c\x2a\xef\xbd\xa1\x0b\x6d\x13\xda\x3d\x74\x13\x68\x7a\x2a\xc5\x68\xad\x71\x34\xd8\x91\x8c\x66\xdc\xad\x2f\xfd\xed\xc5\xeb\xa5\x24\x64\x2f\xd1\x55\x6f\x78\x4f\x9a\xd5\xc5\x7a\x2f\x69\xad\x91\x68\x85\xfb\x92\xc3\xd8\xb0\xc2\x63\x18\xf7\xbd\x34\x6e\x28\xf2\xdb\x1b\xa3\xe3\xc9\x0d\x5e\x0a\x7c\x0a\xc8\xa3\x0d\xa3\x29\x2c\xf2\x89\x9b\xef\x2b\x22\x65\x83\x63\xa2\xed\x4d\x7d\x73\xfb\xfd\x61\xbb\xfb\xf8\xb0\xbd\xdb\x7d\xb8\xfc\x7a\xf7\xed\x76\x5d\xa9\xc6\xb5\x84\xba\xa8\xaf\x25\x71\x31\xf1\x75\xe0\xa1\xcf\x13\xdd\xff\xf8\x54\xbf\x71\xa6\x1a\xc6\xfd\x9c\xfc\x39\x72\xd3\x41\x5a\x04\x56\x93\xe4\x4d\x72\x42\x2b\x3d\xc3\xf7\x85\x7d\x98\xc0\x7f\x44\x4d\xdf\x93\xb4\xf8\x09\xd7\xe2\xf2\xa5\x09\xbf\xae\xe7\x87\x24\x02\x80\x23\x73\x71\xa4\x5e\x37\xbd\x20\xe7\xb3\xc2\xb6\x7d\xf6\x01\x08\x99\x35\xd9\x22\x7c\x87\x27\xdf\x31\xc4\xaa\xff\xb8\x6a\x74\x1d\x4f\x8f\x9c\xe0\xa6\x73\x25\x9b\x73\xd6\xe3\x78\x2b\xd7\xc4\xbd\x32\x2d\xda\x2f\x9c\xb8\x9c\xd6\x82\x83\x58\x44\xca\x18\xbc\xea\x21\x97\xb0\x08\x9f\xcb\x5e\x9b\x9c\xa1\xa8\x87\xdb\xe1\xea\x8a\x5a\x21\x3a\xe1\xda\xf8\x84\x47\xb1\x38\xee\xab\x26\x3f\x61\xb3\xc1\xdf\x65\x09\x5d\xca\x87\x54\xc7\xac\xa6\x44\x8d\xb7\xb3\xa9\xff\x02\x00\x00\xff\xff\x37\x00\x91\x4b\x4d\x02\x00\x00")

func clientBootstrapKeygenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapKeygenSh,
		"client/bootstrap/keygen.sh",
	)
}

func clientBootstrapKeygenSh() (*asset, error) {
	bytes, err := clientBootstrapKeygenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/keygen.sh", size: 589, mode: os.FileMode(493), modTime: time.Unix(1522017202, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientBootstrapTokenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcf\x41\x4b\xfb\x40\x10\x05\xf0\xfb\x7e\x8a\xf7\xa7\xfd\xd3\x53\xb2\xf7\x42\x0e\x45\x82\x29\xd6\x46\x8c\xe2\x45\x08\xdb\x64\x6c\x42\xcc\x6c\x9c\xd9\x55\xfc\xf6\x12\x4d\xf5\x38\xf3\x9b\x79\xf0\x56\xff\xec\xa9\x67\xab\x9d\x31\x4a\x01\x09\x19\xb3\xc2\xa3\x92\xc0\xc9\x39\x8e\xc4\x21\x35\xf7\xf9\x21\xdf\x55\x79\xf6\x5f\x67\xbc\x26\x26\x71\x81\xe0\xd0\x3a\x1a\x3d\x23\xf8\x81\x18\x51\x7b\x3e\xe3\xea\xb0\xc7\x8b\x17\xec\xee\xf6\x10\x7a\x8b\xa4\x41\x53\xa3\xb1\xf5\x68\x7d\x33\x90\x40\x22\x23\x49\x64\xc4\xb3\x01\x80\xe4\x1d\xeb\xa2\xbc\xcd\xb7\xd6\x4d\x93\xed\xbc\x86\x0b\x10\xaa\xaa\xa8\x6f\x8e\xe5\xd3\xb1\x2e\xca\xea\xa1\xca\x36\xbf\x37\x36\x55\xed\xec\xc0\xfe\x83\xeb\x79\xd6\xcd\xdf\xd7\x9c\x96\x7d\x67\x5e\x76\x09\x71\x90\xcf\xc9\xf7\x1c\xb2\x9e\x49\x42\xef\x16\x8a\xa7\xe6\xd5\x45\x6e\xba\xc9\xb5\x76\xa1\xed\x7a\x29\xfc\x53\xcc\x7c\x05\x00\x00\xff\xff\xfc\x14\x01\x6b\x23\x01\x00\x00")

func clientBootstrapTokenShBytes() ([]byte, error) {
	return bindataRead(
		_clientBootstrapTokenSh,
		"client/bootstrap/token.sh",
	)
}

func clientBootstrapTokenSh() (*asset, error) {
	bytes, err := clientBootstrapTokenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/bootstrap/token.sh", size: 291, mode: os.FileMode(493), modTime: time.Unix(1522901431, 0)}
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
	"client/bootstrap/daemon-down.sh": clientBootstrapDaemonDownSh,
	"client/bootstrap/daemon-up.sh": clientBootstrapDaemonUpSh,
	"client/bootstrap/docker.sh": clientBootstrapDockerSh,
	"client/bootstrap/keygen.sh": clientBootstrapKeygenSh,
	"client/bootstrap/token.sh": clientBootstrapTokenSh,
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
	"client": &bintree{nil, map[string]*bintree{
		"bootstrap": &bintree{nil, map[string]*bintree{
			"daemon-down.sh": &bintree{clientBootstrapDaemonDownSh, map[string]*bintree{}},
			"daemon-up.sh": &bintree{clientBootstrapDaemonUpSh, map[string]*bintree{}},
			"docker.sh": &bintree{clientBootstrapDockerSh, map[string]*bintree{}},
			"keygen.sh": &bintree{clientBootstrapKeygenSh, map[string]*bintree{}},
			"token.sh": &bintree{clientBootstrapTokenSh, map[string]*bintree{}},
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

