package mongo_aks_oidc

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type AksCallback struct {
	transport policy.Transporter
	scope     string
}

func NewAksCallback(options ...func(callback *AksCallback)) *AksCallback {
	aksCallback := &AksCallback{}
	for _, option := range options {
		option(aksCallback)
	}
	return aksCallback
}

func WithTransport(transport policy.Transporter) func(*AksCallback) {
	return func(aksCallback *AksCallback) {
		aksCallback.transport = transport
	}
}

func WithScope(scope string) func(*AksCallback) {
	return func(aksCallback *AksCallback) {
		aksCallback.scope = scope
	}
}

func (a *AksCallback) GetAksCallback() (func(ctx context.Context, _ *options.OIDCArgs) (*options.OIDCCredential, error), error) {
	credentials, err := azidentity.NewDefaultAzureCredential(&azidentity.DefaultAzureCredentialOptions{
		ClientOptions: policy.ClientOptions{
			Transport: a.transport,
		},
	})
	if err != nil {
		zap.L().Error("failed to get default azure credential", zap.Error(err))
		return nil, err
	}

	config := policy.TokenRequestOptions{
		EnableCAE: false,
		Scopes:    []string{"openid", a.scope},
	}

	return func(ctx context.Context, _ *options.OIDCArgs) (*options.OIDCCredential, error) {
		token, r := credentials.GetToken(ctx, config)
		if r != nil {
			zap.L().Error("failed to get token", zap.Error(r))
			return nil, r
		}

		return &options.OIDCCredential{
			AccessToken: token.Token,
			ExpiresAt:   &token.ExpiresOn,
		}, nil
	}, nil
}
