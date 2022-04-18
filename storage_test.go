package ipfs_storage_init

import (
	"fmt"
	"hlm-ipfs/ipfs-storage-init/storage"
	"testing"
)

func TestStorageInit(t *testing.T) {
	opt := storage.NewOption(
		storage.WithMountPoint("/mnt/sdb"),
		storage.WithDefaultDisk("sdc"),
	)
	ok, err := storage.InitStorage(opt)
	if !ok {
		fmt.Errorf("init storage err: %+v", err.Error())
		return
	} else {
		if err != nil {
			fmt.Errorf("init %+v ", err.Error())
		}
	}
}
