package provider

import (
	"fmt"
	"testing"

	"github.com/pulumi/providertest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestExamplesUpgrades(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		runExampleParallel(t, "simple-ts")
	})
}

func runExampleParallel(t *testing.T, example string, opts ...providertest.Option) {
	t.Parallel()
	test(fmt.Sprintf("../examples/%s", example), opts...).Run(t)
}

func test(dir string, opts ...providertest.Option) *providertest.ProviderTest {
	opts = append(opts,
		providertest.WithProviderName("kubernetes-the-hard-way"),
		providertest.WithE2eOptions(func(option *integration.ProgramTestOptions) {
			// This is killing us
			// https://github.com/pulumi/providertest/blob/f71d02aa2656e4b3b7c964656e14b4c06b96673b/upgrade.go#L514
			option.Dependencies = []string{"@unmango/kubernetes-the-hard-way"}
		}),
		providertest.WithSkippedUpgradeTestMode(
			providertest.UpgradeTestMode_Quick,
			"Quick mode is only supported for providers written in Go at the moment"),
		providertest.WithBaselineVersion("0.0.4"),
		providertest.WithExtraBaselineDependencies(map[string]string{
			"command": "0.9.1",
			"random":  "4.16.0",
			"tls":     "5.0.1",
		}))

	return providertest.NewProviderTest(dir, opts...)
}
