package provider

import (
	log "github.com/sirupsen/logrus"
	"github.com/szuecs/kube-static-egress-controller/provider/aws"
	"github.com/szuecs/kube-static-egress-controller/provider/noop"
)

type Provider interface {
	Execute([]string) error
	String() string
}

func NewProvider(name string, natCidrBlocks, availabilityZones []string) Provider {
	switch name {
	case aws.ProviderName:
		return aws.NewAwsProvider(natCidrBlocks, availabilityZones)
	case noop.ProviderName:
		return noop.NewNoopProvider()
	default:
		log.Fatalf("Unkown provider: %s", name)
	}
	return nil
}
