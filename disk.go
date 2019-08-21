package proc

import (
	"syscall"
)

type Disk struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

func (self *Disk) DiskTotal() uint64 {
	return self.Total
}

func (self *Disk) DiskFree() uint64 {
	return self.Free
}

func (self *Disk) DiskUsed() uint64 {
	return self.Used
}

func ReadDisk(path string) (*Disk, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	disk := Disk{}
	disk.Total = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.Total - disk.Free
	return &disk, nil
}
