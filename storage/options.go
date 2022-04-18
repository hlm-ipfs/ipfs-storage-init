package storage

import "errors"

type Disk struct {
	Name string
	Init bool
}

type RandomHandler func([]*Disk) (*Disk, error)
type Option struct {
	mountPoint string //挂载点
	diskName   string //盘符
	nvmeEnable bool   //是否允许nvme盘
	random     RandomHandler
	fstype     string
}

func NewOption(options ...func(*Option)) *Option {
	svr := &Option{
		mountPoint: "/mnt/sdb",
		diskName:   "sdb",
		nvmeEnable: false,
		fstype:     "ext4",
		random: func(disks []*Disk) (*Disk, error) {
			if len(disks) > 0 {
				return disks[0], nil
			}
			return nil, errors.New("disk not found")
		},
	}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithMountPoint(mountPoint string) func(*Option) {
	return func(s *Option) {
		s.mountPoint = mountPoint
	}
}

func WithDefaultDisk(diskName string) func(*Option) {
	return func(s *Option) {
		s.diskName = diskName
	}
}

func WithNvmeEnable(nvmeEnable bool) func(*Option) {
	return func(s *Option) {
		s.nvmeEnable = nvmeEnable
	}
}

func WithRandom(random RandomHandler) func(*Option) {
	return func(s *Option) {
		s.random = random
	}
}

func WithFsType(fstype string) func(*Option) {
	return func(s *Option) {
		s.fstype = fstype
	}
}