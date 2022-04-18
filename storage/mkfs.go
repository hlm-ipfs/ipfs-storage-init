package storage

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	defaultBlockSize = 512
)

// FormatDisk formats and labels a disk
func FormatDisk(path, fstype, label string) error {
	family := fsFamily(fstype)
	mkfsCmd := mkfsCommand(fstype)

	cmd := []string{}

	// Add options for the sector size if it's not the default size
	logSec := sectorSize(path)
	if logSec > defaultBlockSize {
		optSector, err := familyFlag("sectorsize", family)
		if err != nil {
			return err
		} else {
			cmd = append(cmd, optSector)
			cmd = append(cmd, string(logSec))
		}
	}

	// Always set the force option
	optForce, err := familyFlag("force", family)
	if err != nil {
		return err
	} else {
		cmd = append(cmd, optForce)
	}

	if len(label) > 0 {
		// Set the label on the disk
		optLabel, err := familyFlag("label", family)
		if err != nil {
			return err
		} else {
			cmd = append(cmd, optLabel)
			cmd = append(cmd, label)
		}
	}

	// Add the path to the command
	cmd = append(cmd, path)
	// Run the mkfs.<fstype> command
	out, err := exec.Command(mkfsCmd, cmd...).CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func mkfsCommand(fstype string) string {
	mkfsCommands := map[string]string{
		"btrfs":    "mkfs.btrfs",
		"ext2":     "mkfs.ext2",
		"ext3":     "mkfs.ext3",
		"ext4":     "mkfs.ext4",
		"fat":      "mkfs.vfat",
		"fat12":    "mkfs.vfat",
		"fat16":    "mkfs.vfat",
		"fat32":    "mkfs.vfat",
		"vfat":     "mkfs.vfat",
		"jfs":      "jfs_mkfs",
		"ntfs":     "mkntfs",
		"reiserfs": "mkfs.reiserfs",
		"swap":     "mkswap",
		"xfs":      "mkfs.xfs",
	}

	if val, ok := mkfsCommands[fstype]; ok {
		return val
	}
	return "mkfs.ext4"
}

func fsFamily(fstype string) string {
	family := map[string]string{
		"ext2":  "ext",
		"ext3":  "ext",
		"ext4":  "ext",
		"fat12": "fat",
		"fat16": "fat",
		"fat32": "fat",
		"vfat":  "fat",
	}

	if val, ok := family[fstype]; ok {
		return val
	}
	return "ext"
}

func familyFlag(flag, family string) (string, error) {
	switch flag {
	case "force":
		switch family {
		case "ext":
			return "-F", nil
		case "fat":
			return "-I", nil
		case "swap":
			return "--force", nil
		default:
			return "", fmt.Errorf("`force` for family `%s` is not implemented", family)
		}

	case "sectorsize":
		switch family {
		case "ext":
			return "-b", nil
		case "fat":
			return "-S", nil
		default:
			return "", fmt.Errorf("`sectorsize` for family `%s` is not implemented", family)
		}

	case "label":
		switch family {
		case "ext":
			return "-L", nil
		case "fat":
			return "-n", nil
		case "swap":
			return "--label", nil
		default:
			return "", fmt.Errorf("`label` for family `%s` is not implemented", family)
		}

	default:
		return "", fmt.Errorf("flag `%s` is not implemented", flag)
	}
}

func sectorSize(path string) int {
	out, err := exec.Command(
		"blkid", "-i", "-o", "value", "-s", "LOGICAL_SECTOR_SIZE", path).Output()
	if err != nil {
		fmt.Printf("Error fetching sector size for `%s`: %v", path, err)
		return defaultBlockSize
	}

	logSec, err := stringToInt(string(out))
	if err == nil {
		return logSec
	}

	fmt.Printf("  Error fetching sector size for `%s`: %s", path, string(out))
	return defaultBlockSize
}
func stringToInt(s string) (int, error) {
	// Remove any control characters e.g. LF
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return 0, err
	}
	cleaned := reg.ReplaceAllString(s, "")

	return strconv.Atoi(cleaned)
}
