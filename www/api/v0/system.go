// SPDX-License-Identifier: AGPL-3.0-only
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package v0

import (
	"context"
	"runtime"

	"github.com/swopstar/gokit/ver"
	"golang.org/x/sys/unix"
)

func (h *Handlers) GetInstanceVersion(_ context.Context, _ GetInstanceVersionRequestObject) (GetInstanceVersionResponseObject, error) {
	v := ver.Get()
	s := v.String()
	source := "https://github.com/swopstar/swoptape"
	return GetInstanceVersion200JSONResponse{
		Version: &s,
		Branch:  &v.Branch,
		Commit:  &v.Commit,
		Source:  &source,
	}, nil
}

func (h *Handlers) GetInstanceStatus(_ context.Context, _ GetInstanceStatusRequestObject) (GetInstanceStatusResponseObject, error) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	heapUsed := int(mem.HeapAlloc)
	heapTotal := int(mem.Sys)

	var fs unix.Statfs_t
	_ = unix.Statfs(".", &fs)
	blockSize := uint64(fs.Bsize)
	diskTotal := int(fs.Blocks * blockSize)
	diskFree := int(fs.Bfree * blockSize)
	diskUsed := diskTotal - diskFree

	return GetInstanceStatus200JSONResponse{
		Memory: &struct {
			TotalBytes *int `json:"totalBytes,omitempty"`
			UsedBytes  *int `json:"usedBytes,omitempty"`
		}{
			TotalBytes: &heapTotal,
			UsedBytes:  &heapUsed,
		},
		Storage: &struct {
			TotalBytes *int `json:"totalBytes,omitempty"`
			UsedBytes  *int `json:"usedBytes,omitempty"`
		}{
			TotalBytes: &diskTotal,
			UsedBytes:  &diskUsed,
		},
	}, nil
}
