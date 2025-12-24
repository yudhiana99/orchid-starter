package ports

import (
	"context"

	modelPorts "orchid-starter/internal/ports/model"
)

type PortsUsecaseAdaptor interface {
	GenerateCompanyData(ctx context.Context, companyID uint64) (result modelPorts.CompanyData, err error)
}
