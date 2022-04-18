package storage

import (
	"encoding/json"
	"os/exec"
)

type LSBLK struct {
	Blockdevices []BlockDevice
}

type BlockDevice struct {
	// lsblk from util-linux 2.38
	Alignment    string        `json:"alignment"`
	DiscAln      string        `json:"disc-aln"`
	Dax          string        `json:"dax"`
	DiscGran     string        `json:"disc-gran"`
	DiscMax      string        `json:"disc-max"`
	DiscZero     string        `json:"disc-zero"`
	Fsavail      interface{}   `json:"fsavail"`
	Fsroots      []interface{} `json:"fsroots"`
	Fssize       interface{}   `json:"fssize"`
	Fstype       string        `json:"fstype"`
	Fsused       interface{}   `json:"fsused"`
	Fsuse        interface{}   `json:"fsuse%"`
	Fsver        string        `json:"fsver"`
	Group        string        `json:"group"`
	Hctl         string        `json:"hctl"`
	Hotplug      string        `json:"hotplug"`
	Kname        string        `json:"kname"`
	Label        string        `json:"label"`
	LogSec       string        `json:"log-sec"`
	MajMin       string        `json:"maj:min"`
	MinIo        string        `json:"min-io"`
	Mode         string        `json:"mode"`
	Model        string        `json:"model"`
	Name         string        `json:"name"`
	OptIo        string        `json:"opt-io"`
	Owner        string        `json:"owner"`
	Partflags    interface{}   `json:"partflags"`
	Partlabel    interface{}   `json:"partlabel"`
	Parttype     interface{}   `json:"parttype"`
	Parttypename interface{}   `json:"parttypename"`
	Partuuid     interface{}   `json:"partuuid"`
	Path         string        `json:"path"`
	PhySec       string        `json:"phy-sec"`
	Pkname       interface{}   `json:"pkname"`
	Pttype       string        `json:"pttype"`
	Ptuuid       string        `json:"ptuuid"`
	Ra           string        `json:"ra"`
	Rand         string        `json:"rand"`
	Rev          string        `json:"rev"`
	Rm           string        `json:"rm"`
	Ro           string        `json:"ro"`
	Rota         string        `json:"rota"`
	RqSize       string        `json:"rq-size"`
	Sched        string        `json:"sched"`
	Serial       string        `json:"serial"`
	Size         string        `json:"size"`
	Start        interface{}   `json:"start,omitempty"`
	State        string        `json:"state"`
	Subsystems   string        `json:"subsystems"`
	Mountpoint   interface{}   `json:"mountpoint"`
	Mountpoints  []interface{} `json:"mountpoints"`
	Tran         string        `json:"tran"`
	Type         string        `json:"type"`
	UUID         string        `json:"uuid"`
	Vendor       string        `json:"vendor"`
	Wsame        string        `json:"wsame"`
	Wwn          interface{}   `json:"wwn"`
	Zoned        string        `json:"zoned"`
	ZoneSz       string        `json:"zone-sz,omitempty"`
	ZoneWgran    string        `json:"zone-wgran,omitempty"`
	ZoneApp      string        `json:"zone-app,omitempty"`
	ZoneNr       string        `json:"zone-nr,omitempty"`
	ZoneOmax     string        `json:"zone-omax,omitempty"`
	ZoneAmax     string        `json:"zone-amax,omitempty"`
	Children     []BlockDevice `json:"children,omitempty"`
}

func ListDisk() (*LSBLK, error) {
	lsblk, err := exec.Command(
		"lsblk",
		"-J",
		"-O",
	).Output()
	if err != nil {
		return nil, err
	}
	lsblkInfo := LSBLK{}
	err = json.Unmarshal(lsblk, &lsblkInfo)
	if err != nil {
		return nil, err
	}
	return &lsblkInfo, nil
}
