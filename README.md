# ipfs-storage-init

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