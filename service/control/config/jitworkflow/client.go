package jitworkflow

import (
	"github.com/common-fate/ciem/config"
	"github.com/common-fate/ciem/gen/commonfate/control/config/v1alpha1/configv1alpha1connect"
)

func NewFromConfig(cfg *config.Context) configv1alpha1connect.JITWorkflowServiceClient {
	return configv1alpha1connect.NewJITWorkflowServiceClient(cfg.HTTPClient, cfg.APIURL)
}
