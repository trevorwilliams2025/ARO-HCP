package pipeline

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/ARO-Tools/pipelines/types"
	"github.com/go-logr/logr"

	"github.com/Azure/ARO-Tools/pipelines/graph"
)

func runIstioUpgradeStep(id graph.Identifier, step *types.IstioUpgradeStep, ctx context.Context, options *StepRunOptions, executionTarget ExecutionTarget) error {
	logger := logr.FromContextOrDiscard(ctx)

	logger.Info("Istio upgrade step running (Stage 1 — no-op)",
		"cluster", step.AKSCluster,
		"dryRun", step.DryRun,
	)

	kubeconfigFile, err := KubeConfig(ctx, executionTarget.GetSubscriptionID(), executionTarget.GetResourceGroup(), step.AKSCluster)
	if err != nil {
		return fmt.Errorf("failed to get kubeconfig: %w", err)
	}
	defer func() {
		if err := os.Remove(kubeconfigFile); err != nil {
			logger.Error(err, "failed to remove kubeconfig file")
		}
	}()

	versions, err := options.Configuration.GetByPath("svc.istio.versions")
	if err != nil {
		return fmt.Errorf("failed to get svc.istio.versions from config: %w", err)
	}

	logger.Info("Istio upgrade step completed (Stage 1 — no action taken)",
		"cluster", step.AKSCluster,
		"versions", fmt.Sprintf("%v", versions),
	)

	return nil
}
