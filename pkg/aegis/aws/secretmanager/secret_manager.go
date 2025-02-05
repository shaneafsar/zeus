package aegis_aws_secretmanager

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/rs/zerolog/log"
	aws_aegis_auth "github.com/zeus-fyi/zeus/pkg/aegis/aws/auth"
)

type SecretsManagerAuthAWS struct {
	*secretsmanager.Client
}

type SecretInfo struct {
	Region string `json:"region,omitempty"`
	Name   string `json:"name"`
	Key    string `json:"key,omitempty"`
}

func InitSecretsManager(ctx context.Context, auth aws_aegis_auth.AuthAWS) (SecretsManagerAuthAWS, error) {
	creds := credentials.NewStaticCredentialsProvider(auth.AccessKey, auth.SecretKey, "")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(creds), config.WithRegion(auth.Region))
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SecretsManagerAuthAWS: error loading config")
		return SecretsManagerAuthAWS{}, err
	}
	cfg.Region = auth.Region
	log.Ctx(ctx).Info().Interface("region", auth.Region).Msg("InitSecretsManager")
	secretsManagerClient := secretsmanager.NewFromConfig(cfg)
	return SecretsManagerAuthAWS{secretsManagerClient}, err
}
