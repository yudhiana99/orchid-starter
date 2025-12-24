package gql

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"orchid-starter/internal/common"

	"github.com/go-resty/resty/v2"
	"github.com/mataharibiz/sange/v2/gql"
)

type ClientGQLQuery struct {
	Client         *resty.Client
	Query          string
	GQLClientURL   string
	VariablesQuery map[string]any
}

func NewGQLQuery(baseClient *resty.Client, token, appOrigin, GQLClientURL string) *ClientGQLQuery {
	restyClient := baseClient.
		SetAuthScheme("Bearer").
		SetAuthToken(token).
		SetHeader("X-App-Origin", appOrigin).
		SetHeader("Accept", "application/json")

	return &ClientGQLQuery{
		Client:       restyClient,
		GQLClientURL: GQLClientURL,
	}
}

func (q *ClientGQLQuery) SetHeader(key, value string) *ClientGQLQuery {
	client := q.Client
	client.SetHeader(key, value)
	q.Client = client
	return q
}

func (q *ClientGQLQuery) SetBaseQuery(query string) *ClientGQLQuery {
	q.Query = query
	return q
}

func (q *ClientGQLQuery) SetQueryOptions(opt *gql.QueryOption) *ClientGQLQuery {
	if opt != nil {
		q.VariablesQuery = map[string]any{
			"query": opt,
		}
	}
	return q
}

func (q *ClientGQLQuery) SetRespField(selectedFields string) {
	if selectedFields != "" {
		q.Query = strings.Replace(q.Query, "%%_SET_RESPONDS_FIELD%%", selectedFields, 1)
		q.Query = common.CleanString(q.Query)
	}
}

// TODO : use this for input mutation
func (q *ClientGQLQuery) SetVariablesInput() {
	// q.Variables =
}

func (q *ClientGQLQuery) DoRequest(ctx context.Context, debug bool, result any) (err error) {

	payload := gql.GraphQLRequest{
		Query:     q.Query,
		Variables: q.VariablesQuery,
	}

	resp, errValue := q.Client.SetDebug(debug).
		R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&result).
		Post(q.GQLClientURL)

	if errValue != nil {
		err = fmt.Errorf("failed to do request: %w", errValue)
		return
	}

	if resp.IsError() {
		err = errors.New(resp.String())
		return
	}

	return
}
