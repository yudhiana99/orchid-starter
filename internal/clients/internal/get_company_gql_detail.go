package internalClient

import (
	"context"
	"fmt"
	"os"

	"orchid-starter/internal/common"
	"orchid-starter/pkg/gql"

	modelCommon "orchid-starter/internal/common/model"

	"github.com/mataharibiz/sange/v2"
	sangeGQL "github.com/mataharibiz/sange/v2/gql"
)

func (ic *InternalClient) GetCompanyGQLDetail(ctx context.Context, companyID uint64, selected string) (result modelCommon.CompanyDetailGQLResponse, err error) {
	token := common.GetAppTokenFromContext(ctx)
	appOrigin := common.GetAppOriginFromContext(ctx)

	if common.GetAppTokenFromContext(ctx) == "" {
		token = ic.TokenSystem
		appOrigin = ic.AppOriginSystem
	}

	GraphQL := gql.NewGQLQuery(GetRestyClient(), token, appOrigin, os.Getenv("RADIANCE_GQL"))
	GraphQL.SetBaseQuery(gql.CompanyDetail).SetQueryOptions(&sangeGQL.QueryOption{
		ExtraParams: []sangeGQL.ExtraParams{
			{
				Key:   "compID",
				Value: companyID,
			},
		},
	}).SetRespField(selected)

	err = GraphQL.DoRequest(ctx, ic.Debug, &result)
	if err != nil {
		return result, sange.NewError(sange.CallServicesFailed, err, "failed to get company detail (GraphQL)", "internalClient", "GetCompanyGQLDetail")
	}

	if len(result.Errors) > 0 {
		return result, sange.NewError(sange.CallServicesFailed, fmt.Errorf("failed to get company detail (GraphQL): %s", result.Errors[0].Message), "failed to get company detail (GraphQL)", "internalClient", "GetCompanyGQLDetail")
	}

	return result, nil
}
