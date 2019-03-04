package foo

//#cgo LDFLAGS: -lpthread
//#cgo darwin CFLAGS: -framework CoreFoundation -framework Security
import "C"
