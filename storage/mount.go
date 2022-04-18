package storage

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Mount mounts the device at a path
func Mount(device, target string) error {
	_ = os.MkdirAll(target, os.ModePerm)
	// Unmount the device, just in case
	_ = Unmount(device)

	out, err := exec.Command("mount", device, target).CombinedOutput()
	if len(out) > 0 {
		fmt.Println(string(out))
		return nil
	}
	return err
}

// Unmount unmounts the device
func Unmount(device string) error {
	out, err := exec.Command("umount", device).Output()
	if len(out) > 0 {
		fmt.Println(string(out))
		return nil
	}
	return err
}


func HasMounted(dev BlockDevice,mountPoint string) bool  {
	if dev.Mountpoint!=nil {
		if point,ok:=dev.Mountpoint.(string);ok {
			if strings.Compare(point,mountPoint)==0{
				return true
			}
		}
	}
	if dev.Children!=nil && len(dev.Children)>0 {
		for _,child:=range dev.Children{
			if HasMounted(child,mountPoint) {
				return true
			}
		}
	}
	return false
}