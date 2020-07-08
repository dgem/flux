// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 7265,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xdd\x6f\x1b\x37\x12\x7f\xf7\x5f\x31\x50\x0e\x48\x0c\x48\x2b\xbb\x6e\x8b\xc3\xf6\x54\x5c\x9a\x0f\x37\x97\x26\x35\xec\xe4\x0e\x7d\xaa\x29\xee\x48\x4b\x88\x4b\xee\x71\xb8\x52\x05\xa1\xff\xfb\x61\xc8\xfd\xe0\xca\xb2\x53\xe4\xed\xf2\x10\xdb\xbb\xc3\xe1\xcc\x6f\xbe\x7e\xe4\xce\x66\xb3\x33\x51\xab\x7f\xa3\x23\x65\x4d\x0e\xa2\xae\x69\xbe\xbd\x3c\xdb\x28\x53\xe4\xf0\x1a\x6b\x6d\xf7\x15\x1a\x7f\x56\xa1\x17\x85\xf0\x22\x3f\x03\x30\xa2\xc2\x1c\x56\xba\xf9\xe3\x70\x00\xb5\x82\xec\xa3\xa8\x90\x6a\x21\x11\xfe\xfc\xb3\x7d\x1f\xfe\xcc\xe1\x70\x18\xbf\x3d\x1c\x00\x4d\xc1\x62\x54\xa3\x64\x65\x0e\x6b\xad\xa4\xa0\x1c\x2e\xcf\x00\x08\x35\x4a\x6f\x1d\xbf\x01\xa8\x84\x97\xe5\x2f\x62\x89\x9a\xe2\x83\x74\x6f\x96\xf6\x4e\x78\x5c\xef\xe3\x4b\xbf\xaf\x31\x87\x5b\x94\x0e\x85\xc7\x33\x00\x8f\x55\xad\x85\xc7\x56\x59\xe2\x01\xff\x13\xc6\x58\x2f\xbc\xb2\xa6\x57\x0e\x50\x3b\x5b\xa1\x2f\xb1\xa1\x4c\xd9\x79\x6d\x9d\xcf\x61\x72\x75\x71\x75\x39\x81\x67\xe0\x51\xeb\x44\x02\xbc\x05\x92\x4e\xd4\x08\xf3\x0a\xbd\x53\x92\xd8\xb9\xda\x2a\xe3\x9f\x13\xf0\xe2\xac\x55\xac\x47\x3e\x1c\x79\x01\xd0\x61\x11\x5e\xd9\x02\xef\x46\x28\xf0\xbf\x25\x7a\x91\x6d\x9a\x25\x3a\x83\x1e\x83\x71\x96\x72\xd0\xca\xb4\x2a\x18\x3a\xb7\x55\x12\x5f\x4a\x69\x1b\xe3\x3f\x8e\x77\x00\xd8\x5a\xdd\x54\xd8\xdb\x30\x6b\x6d\x58\x2b\x3f\xdb\xe0\xbe\xdf\x88\x18\x3e\x3f\x6c\xdc\x3d\x19\xf4\xcd\x78\x49\x11\x32\x23\x91\x2a\x70\x25\x1a\xed\x3f\xd8\x02\x73\xb8\xf8\xf6\xe2\x02\x9e\xc1\xae\x44\x03\x15\x5b\x83\x05\x38\x14\xc5\xcc\x1a\xbd\x9f\xc2\x0e\x61\x67\xcd\x73\x0f\x4b\x04\xb1\xd4\xc8\x40\xca\xb2\xb2\xc5\x59\xab\xf0\x19\x7c\x2a\x15\x81\x22\x10\xe0\xab\x7a\x45\xd0\x10\x16\xb0\xb2\x0e\xd6\x68\xd0\x09\xaf\xcc\x1a\xee\xee\x7e\x86\x0d\xee\x29\x83\x77\x06\xde\xff\x9d\xe0\xc7\x05\x5c\x66\x97\x17\xd3\x5e\x4b\xb7\x77\x74\x81\x40\x38\x4c\xed\x20\xcb\xa6\x18\xc4\x02\x04\x10\xd6\x82\xb3\xa9\x05\x0a\x76\xd8\xab\x91\xc2\xc0\xce\x29\xcf\x86\x66\xa7\xf1\x5b\xa3\xe9\xc1\xc0\xaa\xf6\xfb\xd7\xca\xa5\x20\x56\x58\xa8\xa6\xca\xe1\x03\x56\xd6\xed\x53\x3f\x11\x56\x56\x6b\xbb\x63\x8f\xda\xad\x15\x05\x57\x1b\xe2\x67\x02\x64\x43\xde\x56\x8a\x11\xd8\x18\xbb\x33\xbf\x97\x96\x3c\xf5\x2a\x56\x4a\xe3\x14\x76\xa5\x92\x25\xec\x6d\x03\x3b\xa5\x75\x74\xca\x5b\x28\x2c\x17\x28\x3f\xe6\x45\xfc\x8b\x03\xbb\x33\x6c\x76\xaf\xc0\x61\x6d\xc1\x09\x5f\xa2\x03\x5f\x0a\xd3\x6e\xbc\x56\xbe\x6c\x96\x60\xf9\x21\x82\x56\x1b\xcc\xe0\x37\xdb\x3c\xd7\x1a\x84\x26\xdb\x6d\x31\x06\x1b\x94\x07\x65\xbc\x0d\x6b\xa4\x35\x5e\x28\x83\x6e\x0a\x4b\xd4\x76\x97\xc1\x1d\x0e\xa8\x96\xde\xd7\x94\xcf\xe7\x85\x95\x94\x71\x62\xc9\x82\xd3\x1a\xcd\x9c\x6b\x96\xfc\x7c\xdd\xa8\x02\x69\xde\x10\xce\x6a\xa7\xb6\xc2\x63\x48\x3d\x76\xa4\x57\xd2\x85\x81\xa8\x9c\x49\x6b\x56\x6a\xdd\xbf\x02\x88\x0f\x3e\x88\x3a\x4f\x1e\xa6\xc5\x37\x4b\x96\x7d\x6d\x48\x42\x55\xce\xa3\x92\x21\xf3\xbe\x18\x8e\x9d\xa2\x92\x9f\x94\x62\x8b\x20\xa0\x50\xab\x15\x3a\x6e\xb4\x9d\x86\xb6\xa0\x86\x66\x1a\xd0\x8f\xea\x52\xfc\xb9\x21\x6d\x55\x81\x1d\xe2\x2b\xb5\xae\x44\x3d\x18\xa2\x7c\x09\xc2\x00\x1a\xef\xf6\xc1\x87\xfb\x28\x74\x3f\x05\x61\x0a\x68\x8c\xb4\x15\x77\xf8\xb0\x3e\x7a\xfb\x21\x44\x52\x98\xa2\xd7\x82\x66\x1b\x34\x28\xa4\x36\x94\x0f\x22\xc0\x30\x7c\x45\x04\x92\x65\x5f\x8c\x40\x68\x02\xde\x82\xaa\xb8\xb7\xc2\xf5\xcd\x75\xa8\x7f\x78\xc1\x6e\x91\x5a\x1b\x65\x86\xcd\xd9\xb9\x2d\x3a\xb5\x52\x32\x34\x79\xa8\x1b\x57\x5b\x42\x3a\xff\x0b\x40\xf6\x5a\x62\xe7\x88\x28\x32\x40\xbc\xdf\x5f\x00\x0e\x84\x5b\x0f\x15\xfa\x08\x62\xeb\x7a\xcd\xad\x83\x12\x68\xc6\xdd\xf7\xd9\x23\xfd\xf7\xe1\xba\x13\xfd\xb7\x83\xb3\x2f\xc2\x07\xad\x3f\x19\x0e\x2d\xea\x0e\x43\x8b\x34\x16\x26\x79\x2c\xc2\x09\xa8\x4a\xac\x31\x66\x3f\x2f\xc8\xe0\xad\x32\x45\xf0\xb9\xe2\x8e\xe2\x50\x0e\x59\x1b\xbb\x89\x46\x41\xc8\x7d\x23\x2c\xe5\x20\x30\xb7\x00\xe1\xfb\x92\x2f\x9b\x65\x56\x58\xb9\x41\x97\x49\x5b\xcd\xdd\x3c\x96\x7f\xf8\x31\xf7\xa2\x87\xae\x8b\x23\x73\x04\xe6\x0f\xbc\xab\x17\x6b\x60\x4b\xb3\x5e\x26\x6c\x93\x43\xab\x50\xd9\x54\x5b\x7e\x99\x7d\x73\x91\x5d\x8c\x65\x6f\x1a\xad\x6f\xac\x56\x72\x9f\xc3\xbb\xd5\x47\xeb\x6f\x1c\x52\xea\x85\x43\xb2\x8d\x93\x48\x69\x0b\x77\xf8\xdf\x06\xc9\x8f\x9e\x01\xc8\xba\xc9\xe1\xbb\x8b\x6a\xf4\xb0\x0a\x5d\x3e\x87\xef\xbf\xfd\xa0\x06\x6a\x61\x5d\xba\x78\x36\x44\xe6\x26\xd0\x8c\xab\x8b\x2b\x1e\x9a\xca\xac\xac\xab\x42\xca\x0a\xdd\x4b\x6b\xb5\x45\x83\x44\x37\xce\x2e\x31\xb5\x80\x21\xbd\x1e\x0f\xec\xb8\x55\x54\x38\x7e\x2c\x7c\x99\xc3\x5c\xd4\x2a\x22\xbd\xfd\x7e\xae\x0a\x34\x5e\xf9\x7d\x56\x37\xcb\x44\x56\x19\xe5\x95\xd0\xaf\x51\x8b\xfd\x1d\xd7\x67\x41\x39\x7c\x97\x08\x78\x55\xa1\x6d\xfc\x89\x77\x3c\x5f\xd5\xff\x87\xa9\x49\xd1\x8e\x02\x73\x9a\x19\x41\x9c\x70\x37\xd1\x32\xf4\x32\x58\x56\xcc\x89\x4a\xe6\x86\x36\xb2\x55\xd0\xb6\xed\x37\x6b\x0e\x19\x28\x13\x73\xee\x39\xc5\x35\x44\xe5\x7c\xd4\x26\x3b\xcc\x7e\x35\x7a\x9f\x83\x77\x0d\xb2\x36\xa6\x3f\xa1\x43\x2d\xdb\xc6\xce\x25\x55\xa3\x5b\x59\x27\x91\x95\x46\xbe\xc3\x74\xe7\x31\xc3\x53\x4a\x32\xb6\x7d\x2b\x5c\x6b\x7b\x14\xfb\x3a\xf3\x93\x1a\x7d\x67\xa4\x6e\x42\xe7\x64\xd6\x16\x07\x5c\xd7\x55\x23\x2d\xf8\x02\x8b\xe9\x78\xcc\x0f\xbc\xf4\x88\x61\xf4\xdd\x15\x0a\x94\x5a\x38\x66\x6b\x4b\xbb\x4d\x1a\xc0\x13\x34\x20\xb6\xc7\xd4\x79\x67\xad\x9f\x67\x44\xe5\xa3\x0e\x08\x33\xda\x75\x32\x8c\xa8\x49\xdc\x79\xda\x89\x24\x1a\xd0\x6c\x95\xb3\x26\x0c\x84\x38\x6b\x27\xef\x3f\xff\xf4\xe6\xd5\xaf\x1f\xdf\xbe\xbb\x9e\xc4\x11\x30\x65\x3c\xec\x16\x9d\x1b\xcf\xeb\x44\x4d\x18\x71\xcb\x7d\x9c\xa6\x5e\x9f\xf2\xf1\xc1\xa0\x7d\xe8\xe3\x90\x9c\x2c\xfc\xa8\xa3\x3c\xf3\xf8\xb0\xd2\xed\xc6\x2d\x3a\xa1\x22\xad\x75\x21\x26\x89\x8a\x63\x42\x93\x06\x3d\xb0\x99\x8e\x75\x0b\x03\x42\x7b\x74\x86\x59\xf5\x03\x8b\x57\xce\x56\x9c\x16\x1d\x63\x99\x82\x20\x4e\xb7\x76\xaa\x32\x0c\xda\xca\x0d\x3d\x0c\x36\x9a\x6d\x7e\x02\x97\x01\xee\x11\x2e\x5b\xa1\x1b\x7c\x80\xc9\x97\x92\xf8\x38\x07\xba\x99\xfb\x44\x06\xf0\xc8\x1f\x8f\xfa\x27\x86\xfd\x23\x79\xc9\x52\x91\xdd\x8c\xe4\xc6\xfd\xe1\x4b\x95\xb7\x13\x4c\x4a\x2c\x50\x53\xd7\x7a\x0f\x3f\x7f\xfa\x74\x03\x4b\x41\x4a\x82\x68\x7c\x09\xd2\x61\xe8\xa4\x42\xc7\xa9\x3e\x1c\x05\x58\xe1\x56\x89\xe0\xf8\xfd\xf5\xbb\x4f\xbf\xbf\xfc\xfc\xe9\xe7\xcf\x77\x6f\x6e\xef\x83\xbb\xfd\xa3\xf7\x6f\x7e\xbb\x1f\x25\xfc\x56\x38\xc5\x07\x39\xea\x08\x72\xa2\x30\xd2\x97\xa3\xf8\xbd\x75\xb6\x1a\xc7\x30\x8a\xdd\xe2\x2a\x1f\x79\x3e\xe2\x8a\xdc\xd8\xd8\x85\x01\x00\xc6\x3c\x1f\xe1\x11\x21\x88\xc7\x53\x2c\x78\x12\x4b\x21\x4b\x2c\x38\xb5\xd2\xdc\xee\x69\x35\x23\xc5\xda\xa7\x89\x16\xeb\x5a\xde\x9c\x2c\x68\x8f\xd7\x61\xe1\x34\x6c\xc2\xc7\xc2\x16\x63\x5f\x22\xa5\xb9\x30\xb0\x57\xbf\xb3\x6c\x65\xc3\x38\x85\x8a\x0b\x97\x08\x21\x11\xa1\xb4\xbb\x70\xf4\xb5\xc6\xa0\x0c\x21\x53\x7e\x9c\x3b\xb3\x59\xef\x40\x38\xf7\xf0\xe6\x8b\xfe\x51\xd6\x92\xbe\x8c\xb6\x32\x93\xba\x21\x8f\x2e\xe3\x06\xae\x53\x48\x3e\x53\xec\x35\x03\x14\xaf\xa2\xe8\xbb\x9b\x91\x53\xdc\x76\x08\x7d\x38\x5a\x8f\x33\x7b\xb0\xa1\x93\xe7\xec\xf2\x8e\x25\xc3\x61\x37\x19\x41\xa9\xc5\xad\xf4\xe2\x6c\xc4\x32\x15\x41\xd5\x50\x38\xfc\x07\xf4\x14\x16\xb1\x9c\x96\x61\xb0\x05\x8e\x17\xce\xfc\x2f\xba\x83\xf4\x79\x6a\x4b\xd7\x5c\x62\x19\x72\x02\x27\x47\xff\x91\x21\x3c\x0c\xe2\x80\x9b\x15\xca\x2d\x1e\x8c\xbd\xd4\xac\xdb\x84\x61\x0e\xc1\xfb\x7c\xfb\x4b\xbc\x9b\x10\x66\x1d\xdf\x5d\x2b\x1f\xce\xcb\xa4\xbc\x75\xfb\xbe\x5d\xbf\x65\x66\x9c\xa8\x7b\xaa\xe6\x38\x6d\x12\xdf\xdb\x92\x39\x59\x4e\x69\x2d\x74\xdc\xf9\x6f\x2f\xd2\xca\x3c\xcf\x87\xbf\xdf\xbf\xf9\xed\xfc\x9f\xf1\xd4\x1e\x68\x75\x43\xe8\xe6\x83\xb1\x59\x5a\xe8\x8c\x0f\x97\x53\xe3\xf4\xe2\x70\x80\xec\x5a\x79\x76\x36\x5c\xdf\x8d\x25\x96\x4e\x18\x59\x76\x42\x3f\x85\xbf\xe2\x45\x9e\x5a\x85\x47\xdc\xbf\xe8\xd4\x4a\xe6\x70\xbc\xee\x2e\x64\x0a\xfd\xcb\x2a\x93\x2c\x98\x4c\x27\xed\x7d\xa0\x26\x4c\x97\x3f\xdd\xd4\x1c\x72\xe2\xc9\x78\xea\xaa\x84\x51\x2b\xe6\xe4\x5c\x43\xa4\x0a\x74\x31\x1c\x47\x27\x9b\x70\x1d\x61\x09\xa1\x31\x05\xba\xa3\x18\x3b\xd4\xc2\xab\x2d\x06\xca\x49\x5d\x06\xae\x47\x71\x3e\xaa\xc9\xde\x39\x6a\x96\x85\x72\x97\xd3\xf8\xf3\x9b\xfe\x72\x73\x00\x27\x5c\x5e\x9e\x02\x27\xdc\x08\x76\xa8\x76\x52\x27\x14\x7c\x26\x74\xa7\xd6\x73\x70\xfb\xc8\xb1\x0c\x9c\x5e\xff\xa6\x12\xea\xa4\x01\xc8\x2f\x3a\x0d\x9d\xd4\x70\x3d\x7b\x32\x1c\xc8\xad\x64\x67\x19\x50\x34\xe1\xe6\x8e\x71\xe2\x89\xad\xfc\xd1\x01\x3c\xc5\xaa\x9d\x7d\xed\x64\x5b\x3c\x31\xea\xba\x15\xad\x2e\x5e\xb5\xf8\xc7\x06\xf7\xa0\x8a\x1f\x7b\xb1\x27\xe8\x4c\x62\x15\xab\x10\xbe\x71\x38\xba\x05\x38\xb1\x57\x78\xbd\x9f\xf5\xf2\x34\x6a\x57\x5d\xb7\x06\xe5\xa1\x14\x14\x46\xb1\x35\x7a\x0f\x42\x4a\xa4\xd8\xd1\x4b\x8c\x77\x68\x2f\xba\x3b\x9b\xfb\x95\xd0\x84\xf7\xe7\x67\x87\xc3\xac\x0b\xc4\x6d\x3b\xc3\x4f\xc5\xa2\x53\x1a\xe4\x1f\xd6\xc3\x69\xb1\x13\x71\x22\xef\x1a\xe9\xa3\xbd\xbb\x70\x9c\x67\x8a\xd7\x78\xa0\xbd\x91\xb0\xb4\x76\xb3\x41\xac\x39\xeb\x7b\x53\x27\x6b\xe5\x27\x53\xa8\x50\x30\xe0\xdc\xd0\x40\x84\x33\x76\x5b\x08\x4d\x4d\xde\xa1\xa8\xfa\x8a\x38\x3f\x32\x8c\x55\xcf\xc8\x0b\x8f\x0b\x6e\x30\x8f\xe6\x8d\xc1\x3f\x7c\x97\x3c\xc9\xc4\x13\x06\x26\xdd\x1e\x93\x6e\x1e\x25\x4a\x5e\x60\xb6\xce\xa6\xf0\x1f\x64\x66\xf9\x4a\xdb\xa6\x38\xcf\xc2\x05\x91\xb7\x1b\x3e\x9f\x10\xd4\xc2\x79\x25\x1b\x2d\x5c\x17\x8c\x56\xcb\xf1\x28\x6d\x77\x5d\xec\x88\xfb\xa8\x64\x5d\xd9\x8e\xf5\x66\x3b\xeb\x36\xd4\x1f\x36\x8f\x96\x85\x8d\x16\x62\x29\x2f\xbf\xb9\x7a\xf8\x7f\xea\xf0\x9b\x98\x7d\x5d\x57\xea\xef\xaa\xad\x79\x22\x35\x3e\xb4\xd2\xd7\x83\xf0\x51\x86\x74\xfa\x66\x83\xbe\x45\xe0\x81\x8f\x67\xcb\xa9\x25\x61\xe3\x47\x52\xe7\x0e\xdd\xf6\xc4\x57\x0c\x3e\x10\x0c\x0c\x88\x6b\xf5\x87\x74\x14\x8b\x0d\x8f\xb1\x98\x65\x84\x3e\xf9\x34\xf2\x3c\xf9\xba\x92\x7c\x26\xe1\xe0\x84\xab\xbb\x40\xca\xb3\x91\x97\x5a\x91\x47\x33\x6b\x4d\x58\xe4\x57\x17\x57\x97\x3d\x48\xb7\xb8\x56\xe4\xdd\xfe\xb5\x22\x86\xf8\x4e\x0a\x13\xd2\xf5\x08\x29\xd7\x8a\xcd\x8a\x28\x37\xa3\x56\x30\x75\xbb\xed\x8d\x2f\x8b\x42\xc5\x4b\x16\x1e\xde\x2f\x99\xbc\x8f\x60\x1c\xde\x0f\xfc\xed\x70\x00\x17\xa8\xc0\x17\x56\xcf\xc2\x77\xaf\x51\x3f\x1d\x7e\xeb\x36\xf8\xb5\x6e\xd5\xbf\xfe\x78\xd7\x11\x2f\x9a\xb6\x07\xa2\xc6\xb5\x34\x0c\x4c\x61\x3d\x81\x0d\xc2\x50\x89\x7d\xb8\x9c\xd2\xdb\xe1\x8a\xd2\x90\xb6\x76\xd3\xd4\xa0\x88\x1a\x24\xb0\x06\xc8\x56\x08\xef\xfb\xaf\x45\xac\xbd\xa9\x69\xb8\x81\x2c\x0c\x75\xf7\x5f\x93\x8f\xd6\xe0\x24\x7d\xf3\x2a\x18\x90\xde\x41\xc6\xcd\x69\x7c\x2d\xd9\x1d\x6c\x82\x7d\xa3\x37\xfd\x99\x6b\x72\x39\x39\xfb\x5f\x00\x00\x00\xff\xff\x1e\x21\x4d\x4f\x61\x1c\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 967,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\x4d\x6f\x9b\x40\x10\xbd\xf3\x2b\x9e\xe4\x6b\x21\x25\x52\x2e\xdc\xa2\xa6\xad\x22\xb5\x91\xa5\x28\xbd\x8f\x97\x81\xac\xb2\x5f\xdd\x9d\x75\x4d\x51\xfe\x7b\x05\x76\x6c\x68\x32\x27\xe0\xbd\x79\xf3\x76\xf6\x51\x96\x65\xb1\x81\x65\xab\x48\x3d\x73\x8b\x96\x83\xf1\x83\x65\x27\xc8\x89\x5b\xec\x06\x7c\x33\xf9\x00\xf1\x98\x19\xc5\x06\xca\x3b\x21\xed\x38\x42\x5b\xea\x19\x96\x85\x5a\x12\xaa\x0a\x0a\xfa\x17\xc7\xa4\xbd\x6b\x40\x21\xa4\xab\x7d\x5d\xbc\x68\xd7\x36\xb8\x3b\xcb\x16\x6f\xf4\xa6\x00\x1c\x59\x6e\x2e\xd3\xc7\x11\xba\x43\xf5\x40\x96\x53\x20\xc5\x78\x7d\x3d\x91\xe6\xd7\x06\xe3\xb8\x46\xc7\x11\xec\xda\x89\x96\x02\xab\x49\x31\x72\x30\x5a\x51\x6a\x50\x17\x40\x62\xc3\x4a\x7c\x9c\x10\xc0\x92\xa8\xe7\x1f\xb4\x63\x93\x8e\x1f\xde\x19\x28\x00\x61\x1b\x0c\x09\x9f\x5a\x16\x66\xa7\x32\xab\xee\x8f\xfa\x81\x37\x2b\x33\xee\x5b\x7e\x5c\x99\x98\x6a\xc7\x42\xd5\x4b\xde\x71\x74\x2c\x9c\x2a\xed\xaf\x7c\x6a\x60\xb4\xcb\x87\x13\xe9\xbc\xe4\xf3\xb0\xf2\xc3\x61\x53\xcd\xd7\xb0\x00\x9a\xba\xba\xa9\xae\x3f\xaf\xf1\x6d\x36\x66\xeb\x8d\x56\x43\x83\xfb\xee\xc1\xcb\x36\x72\x9a\xee\xe3\x8d\x45\xb1\x5f\x1c\xac\x44\x69\x71\x53\x5f\x03\xd8\xe0\x27\x1d\xb4\xcd\x76\x9a\xe0\xe3\x30\x65\x21\x27\xfe\x04\xed\x60\xb9\xa7\xdd\x20\x9c\x96\x8d\xf7\xb8\xb1\x58\x35\x26\xfd\x97\xd1\xf9\x08\xef\x18\x5a\xd8\x2e\xe9\x01\x75\x7d\x5d\xd7\xd8\xe0\x8e\x3b\xca\x46\x10\x7c\xbc\xf8\xda\x4c\x9c\xfd\xfe\xf8\xf8\xe4\x94\xb7\x73\x3a\xc5\xa3\x67\x81\xf1\x7d\x82\xef\xc0\xa4\x9e\x11\xf9\x77\xe6\x24\x20\xd7\x22\x72\x0a\xde\x25\xae\xce\x42\x93\xea\xea\x84\xc7\x7d\x2a\xa3\xd9\xc9\xe5\x00\x8b\xdd\x6f\x7d\x94\xe6\xe8\xee\x14\xcd\xdb\xb6\x7d\x64\x95\xa3\x96\xe1\x8b\x77\xc2\x07\x99\x23\x7a\xac\xb4\x46\x9a\x85\x64\xcc\xee\x36\x3d\x25\x8e\x27\xb9\xff\xa1\xef\xd1\xe7\xf0\x1e\x23\x63\xfc\x9f\x6d\xd4\x7b\x6d\xb8\xe7\xaf\x49\x91\x21\x99\x7f\xaf\x8e\x4c\xe2\x4b\xfc\xff\x05\x00\x00\xff\xff\xec\xf7\xcf\x04\xc7\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
