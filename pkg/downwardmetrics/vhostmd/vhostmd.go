package vhostmd

import "practice/pkg/downwardmetrics/vhostmd/api"

type MetricsIO interface {
	Create() error
	Read() (*api.Metrics, error)
	Write(metrics *api.Metrics) error
}

func NewMetricsIODisk(filePath string) *vhostmd {
	return &vhostmd{filePath: filePath}
}
