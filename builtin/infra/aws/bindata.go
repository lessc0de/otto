// Code generated by go-bindata.
// sources:
// data/vpc-public-private/main.tf
// data/vpc-public-private/outputs.tf
// DO NOT EDIT!

package aws

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
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
	name string
	size int64
	mode os.FileMode
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

var _dataVpcPublicPrivateMainTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x91\x3d\x4f\xc3\x30\x10\x86\x77\xff\x8a\x57\x2e\x23\x0a\xb4\x42\x6c\x4c\x0c\x4c\x7c\x0c\x88\xb5\x72\xec\x53\x63\xe1\xda\x95\x7d\x69\x14\x55\xf9\xef\xd8\xf9\x60\x00\x51\x90\x38\x6f\xe7\x57\x77\xcf\xa3\x5b\xe1\x81\x3c\x45\xc5\x64\x50\xf7\x78\x66\x0e\x97\x30\x01\x3e\x30\xc8\x58\xc6\x5e\xf9\x56\x39\xd7\x57\x42\xac\xf0\xa8\xac\xc7\xdb\xcb\x3d\xb8\x51\x8c\xce\x3a\x07\x1d\x3c\x97\x2e\x1d\x29\xf6\xdc\x58\xbf\xab\x44\xa4\x14\xda\xa8\x09\x52\x75\x69\x7b\x3c\x68\x09\xb9\xcf\x21\x89\x93\x40\x2e\x6d\x4d\xdc\xd6\x2e\xe8\x77\xdc\x41\xae\xaf\xab\xf1\x5d\xad\x6f\xa5\x18\xff\x59\xed\xd2\x1c\x2d\xf5\xa4\xf6\x54\x82\x21\xc3\xc9\xb1\x3b\x88\xa1\xe0\xbc\x36\x84\x43\x5b\x3b\xab\x91\xda\xda\x13\xc3\x26\x74\x0d\x45\xc2\x82\x90\x0a\xa0\x27\x5d\xfc\x38\x64\x6e\x82\xf5\x4c\xb1\x84\x47\xfe\x5d\xf8\x82\x3b\x4d\xca\xc4\xd3\xe4\x85\x39\x5b\x6c\xad\x29\x18\x17\xa7\xd9\xaa\x2a\x4e\x95\x35\x83\xfc\xc9\x6a\x9d\xad\x36\x37\xe7\xad\xe6\x35\xdf\xbc\xa2\x3d\xe6\xab\x9c\x11\x1b\x6f\xa0\xe2\x22\xa4\x1c\x82\x77\xfd\xaf\x56\xd3\xdc\x7f\x69\x6d\xfe\xa2\x35\xef\xf9\xf4\xfa\x08\x00\x00\xff\xff\x54\xe3\x53\x62\x6a\x02\x00\x00"

func dataVpcPublicPrivateMainTfBytes() ([]byte, error) {
	return bindataRead(
		_dataVpcPublicPrivateMainTf,
		"data/vpc-public-private/main.tf",
	)
}

func dataVpcPublicPrivateMainTf() (*asset, error) {
	bytes, err := dataVpcPublicPrivateMainTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vpc-public-private/main.tf", size: 618, mode: os.FileMode(420), modTime: time.Unix(1433415322, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataVpcPublicPrivateOutputsTf = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xb1\x6e\xc3\x30\x0c\x44\x77\x7f\xc5\xc1\xe9\xe8\xfa\x0f\xba\x74\xe9\xd8\xa5\x7b\x40\x4b\x74\x2c\x44\x16\x0d\x89\x4a\x10\x04\xf9\xf7\xd2\x36\x82\x4e\x41\xb5\x89\x77\xf7\x78\x90\x0e\xf8\xe2\xc4\x99\x94\x3d\x86\x1b\xbe\x55\xa5\x83\x17\x24\x51\xb0\x0f\x8a\x99\x52\xa5\x18\x6f\x7d\x73\x68\x0e\x9b\x8e\x5a\xb8\x40\xaa\x2e\x55\x0b\xa8\x40\x27\xc6\xcc\x3a\x89\xc7\x28\x19\x9a\x29\x95\x91\x73\x0e\xe9\x04\x4f\x4a\x18\xb3\xcc\xf8\xb1\x09\x99\x3e\x1b\x66\x20\x77\x46\x48\xc6\xda\x80\x3a\x91\xe2\x1a\x62\xc4\xc0\x2b\x7d\xe7\x78\x5e\xa2\xdc\x4a\x87\xb1\x6a\xcd\x6c\xfe\x31\x53\xd1\x5c\xdd\x7a\x35\x8a\x9b\x28\x9d\xb8\x03\xab\xdb\xdb\x7d\xb2\x23\x8b\x43\xc6\xad\x53\x98\x17\xc9\x4a\xc9\x3d\x27\x26\x5d\x28\x56\x6b\xbf\xf2\xf7\xd5\x62\xf8\xe4\x34\x48\xea\xe0\x68\xe3\x96\x49\x6a\xf4\x6b\x17\xa5\x33\x27\x84\x67\xd8\x64\xcc\xe2\xc3\x18\xd8\xf7\x4d\xb3\x3f\x01\xda\xcb\xe2\x8e\xc1\xb7\xb8\x37\xb0\xb3\x6d\xc0\x07\xda\xb7\x3b\x5d\xcb\xd1\xc4\x7e\xa6\x90\xfa\xe0\x1f\x6d\xf3\xf8\x4b\x95\x3a\x24\xd6\xf7\xa5\x0e\x31\xb8\x17\xe1\xdd\xd3\xef\x9e\x97\x84\x1c\x2e\xf6\x7f\xff\x20\x76\xd3\x93\xf1\x1b\x00\x00\xff\xff\x8c\x60\xcc\x82\xf7\x01\x00\x00"

func dataVpcPublicPrivateOutputsTfBytes() ([]byte, error) {
	return bindataRead(
		_dataVpcPublicPrivateOutputsTf,
		"data/vpc-public-private/outputs.tf",
	)
}

func dataVpcPublicPrivateOutputsTf() (*asset, error) {
	bytes, err := dataVpcPublicPrivateOutputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vpc-public-private/outputs.tf", size: 503, mode: os.FileMode(420), modTime: time.Unix(1439235998, 0)}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
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
	"data/vpc-public-private/main.tf": dataVpcPublicPrivateMainTf,
	"data/vpc-public-private/outputs.tf": dataVpcPublicPrivateOutputsTf,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"vpc-public-private": &bintree{nil, map[string]*bintree{
			"main.tf": &bintree{dataVpcPublicPrivateMainTf, map[string]*bintree{
			}},
			"outputs.tf": &bintree{dataVpcPublicPrivateOutputsTf, map[string]*bintree{
			}},
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
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

