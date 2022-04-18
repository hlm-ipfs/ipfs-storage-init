# ipfs-storage-init

## go.mod
```
require hlm-ipfs/ipfs-storage-init v0.0.0

replace hlm-ipfs/ipfs-storage-init => github.com/hlm-ipfs/ipfs-storage-init v0.0.0-20220418070932-80177a771bb8
```
## example
```go

package main

import (
	"fmt"
	"hlm-ipfs/ipfs-storage-init/storage"
)

func main() {
	opt := storage.NewOption(
		storage.WithMountPoint("/mnt/sdb"),
		storage.WithDefaultDisk("sdc"),
	)
	ok, err := storage.InitStorage(opt)
	if !ok {
		fmt.Errorf("init storage err: %+v\n", err.Error())
		return
	} else {
		if err != nil {
			fmt.Errorf("init %+v \n", err.Error())
		}
	}
}
```
