package system

import (
	"context"

	"practice/pkg/util/tls"
	v1 "practice/pkg/vsock/system/v1"
)

type SystemService struct {
	caManager tls.ClientCAManager
}

func (s SystemService) CABundle(ctx context.Context, _ *v1.EmptyRequest) (*v1.Bundle, error) {
	raw, err := s.caManager.GetCurrentRaw()
	if err != nil {
		return nil, err
	}
	return &v1.Bundle{Raw: raw}, nil
}

func NewSystemService(mgr tls.ClientCAManager) *SystemService {
	return &SystemService{caManager: mgr}
}
