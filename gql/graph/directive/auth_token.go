package directive

import (
	"context"
	"fmt"

	"orchid-starter/constants"
	"orchid-starter/internal/bootstrap"
	"orchid-starter/internal/common"

	"github.com/99designs/gqlgen/graphql"
	mbizUtil "github.com/mataharibiz/sange/v2/utils"
	"github.com/mataharibiz/ward/logging"
	"github.com/vektah/gqlparser/v2/gqlerror"
	gqlError "github.com/vektah/gqlparser/v2/gqlerror"
)

type Directive struct {
	DI *bootstrap.DirectInjection
}

func NewDirective(di *bootstrap.DirectInjection) *Directive {
	return &Directive{
		DI: di,
	}
}

type CheckTokenInput struct {
	Injection *bootstrap.DirectInjection
	Ctx       context.Context

	TokenIdentifier string
	TokenState      string
	AuthToken       string
	AppOrigin       string

	SessionUserID uint64
	SessionCompID uint64
}

type CheckThirdPartySessionTokenInput struct {
	Injection *bootstrap.DirectInjection
	Ctx       context.Context
}

func (d *Directive) AuthToken(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {

	var (
		appToken        = common.GetAppTokenFromContext(ctx)
		appOrigin       = common.GetAppOriginFromContext(ctx)
		tokenState      = common.GetTokenStateFromContext(ctx)
		tokenIdentifier = common.GetTokenIdentifierFromContext(ctx)
		userID          = common.GetUserIDFromContext(ctx)
		companyID       = common.GetCompanyIDFromContext(ctx)
	)

	// handle third party authentication
	if isThirdParty := common.GetThirdPartyFromContext(ctx); isThirdParty == "1" {
		if partner := common.GetPartnerFromContext(ctx); partner != "" {
			thirdParty := CheckThirdPartySession(d.DI, ctx)
			if exist := thirdParty.CheckCacheToken(); !exist {
				// authenticate token to auth service
				client := d.DI.GetClient().InternalClient
				if err := client.AuthenticateThirdPartyToken(ctx); err != nil {
					return nil, &gqlError.Error{
						Message: err.Error(),
						Extensions: map[string]interface{}{
							"code": "INVALID_TOKEN",
						},
					}
				}
			}
		}
	}

	if appOrigin == "" || appToken == "" || tokenIdentifier == "" {
		return nil, &gqlError.Error{
			Message: "app origin Or app token Or token identifier header is empty",
			Extensions: map[string]interface{}{
				"code": "INVALID_TOKEN",
			},
		}
	}

	// Bypass non-user token include (system,admin,worker)
	if constants.IsNonUser(appOrigin) {
		return next(ctx)
	}

	// Validate user token
	if userID == 0 || (companyID == 0 && tokenState == "PRE-LOGIN") {
		return nil, &gqlError.Error{
			Message: "user id Or company id is invalid",
			Extensions: map[string]interface{}{
				"code": "INVALID_TOKEN",
			},
		}
	}

	// valid session token
	checkTokenInput := CheckTokenInput{
		Ctx:             ctx,
		Injection:       d.DI,
		TokenIdentifier: tokenIdentifier,
		TokenState:      tokenState,
		AuthToken:       appToken,
		AppOrigin:       appOrigin,
		SessionUserID:   userID,
		SessionCompID:   companyID,
	}
	if errToken := CheckSessionToken(&checkTokenInput); errToken != nil {
		return nil, errToken
	}

	return next(ctx)
}

func CheckTokenIdentifier(redisType string, tokenIdentifier string) *gqlError.Error {
	rdsUtil, err := mbizUtil.NewRedisUtil(redisType)
	if err != nil {
		return &gqlError.Error{
			Message: "cache error. Error: " + err.Error(),
			Extensions: map[string]interface{}{
				"code": "CACHE_ERROR",
			},
		}
	}

	key := "token_identifier_" + tokenIdentifier
	token, errGet := rdsUtil.Get(key)
	if errGet != nil || token == "" {
		return &gqlerror.Error{
			Message: "invalid token identifier",
			Extensions: map[string]interface{}{
				"code": "INVALID_TOKEN",
			},
		}
	}

	return nil
}

func (ck *CheckTokenInput) GetRedisType() string {
	switch ck.TokenState {
	case "PRE-LOGIN":
		return constants.RedisTypePreLoginToken
	case "POST-LOGIN":
		return constants.RedisTypePostLoginToken
	default:
		return ""
	}
}

func (ck *CheckTokenInput) ValidateTokenState() *gqlError.Error {
	if ck.GetRedisType() == "" {
		return &gqlError.Error{
			Message: "invalid token state",
			Extensions: map[string]interface{}{
				"code": "INVALID_TOKEN",
			},
		}
	}
	return nil
}

func CheckSessionToken(input *CheckTokenInput) *gqlError.Error {
	// validate token state
	if err := input.ValidateTokenState(); err != nil {
		return err
	}

	// check token identifier in cache
	if err := CheckTokenIdentifier(input.GetRedisType(), input.TokenIdentifier); err != nil {
		return err
	}

	if exist := input.CheckCacheToken(); !exist {
		// authenticate token to auth service
		client := input.Injection.GetClient().InternalClient
		if err := client.AuthenticateToken(input.Ctx); err != nil {
			return &gqlError.Error{
				Message: err.Error(),
				Extensions: map[string]interface{}{
					"code": "INVALID_TOKEN",
				},
			}
		}
	}

	return nil
}

func (ck *CheckTokenInput) CheckCacheToken() bool {
	rdsUtil, err := mbizUtil.NewRedisUtil(ck.GetRedisType())
	if err != nil {
		ck.Injection.Log.Warn("failed connect to redis", "error", err)
		return false
	}

	redisKey := fmt.Sprintf("COMPANY:%v:USER:%v:TOKENJWT:%v", ck.SessionCompID, ck.SessionUserID, ck.AuthToken)
	token, errGet := rdsUtil.Get(redisKey)
	if errGet != nil || token == "" {
		logging.NewLogger().Warn("token redis value is empty")
		return false
	}

	return true
}

func CheckThirdPartySession(di *bootstrap.DirectInjection, ctx context.Context) *CheckThirdPartySessionTokenInput {
	return &CheckThirdPartySessionTokenInput{
		Injection: di,
		Ctx:       ctx,
	}
}

func (ckt *CheckThirdPartySessionTokenInput) CheckCacheToken() bool {
	var (
		appToken = common.GetAppTokenFromContext(ckt.Ctx)
		clientID = common.GetClientIDFromContext(ckt.Ctx)
	)

	rdsUtil, err := mbizUtil.NewRedisUtil(constants.RedisTypeThirdPartyToken)
	if err != nil {
		ckt.Injection.Log.Warn("failed connect to redis", "error", err)
		return false
	}

	redisKey := fmt.Sprintf("CLIENT:%v:TOKENJWT:%v", clientID, appToken)
	token, errGet := rdsUtil.Get(redisKey)
	if errGet != nil || token == "" {
		logging.NewLogger().Warn("token redis value is empty")
		return false
	}

	return true
}
