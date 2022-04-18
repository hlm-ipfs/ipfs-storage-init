package storage

import (
	"errors"
	"fmt"
	"strings"
)

func InitStorage(opt *Option) (bool, error) {
	//判断挂载点是否已经被挂载
	lsblk, err := ListDisk()
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("lsblk error: %+v ", err.Error())
	}
	mounted := false
	for i := range lsblk.Blockdevices {
		if HasMounted(lsblk.Blockdevices[i], opt.mountPoint) {
			mounted = true
			break
		}
	}
	if mounted {
		return true, fmt.Errorf(" %v point has mounted\n", opt.mountPoint)
	}
	//查找出没挂再的设备
	var defaultDisk *Disk
	initDisk := make([]*Disk, 0)
	for _, device := range lsblk.Blockdevices {
		disk := Disk{
			Name: fmt.Sprintf("/dev/%+v", device.Name),
			Init: true,
		}
		if device.Fstype == "" && device.UUID == "" && device.Children == nil {
			disk.Init = false
		}
		if strings.Compare(opt.diskName, device.Name) == 0 {
			defaultDisk = &disk
		}
		initDisk = append(initDisk, &disk)
	}
	if defaultDisk == nil {
		defaultDisk, err = opt.random(initDisk)
		if err != nil {
			return false, fmt.Errorf("random disk : %+v ", err.Error())
		}
	}
	if !defaultDisk.Init {
		err = FormatDisk(defaultDisk.Name, opt.fstype, "")
		if err != nil {
			return false, fmt.Errorf("format disk %+v err: %+v", defaultDisk.Name, err.Error())
		}
	}
	//mount
	err = Mount(defaultDisk.Name, opt.mountPoint)
	if err != nil {
		return false, fmt.Errorf("mount err: %+v", err.Error())
	}
	return true, errors.New("init ok")
}
