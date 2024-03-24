package provider

import (
	"fmt"
	"testing"

	"github.com/pulumi/providertest"
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
		providertest.WithSkippedUpgradeTestMode(
			providertest.UpgradeTestMode_Quick,
			"Quick mode is only supported for providers written in Go at the moment"),
		providertest.WithExtraBaselineDependencies(map[string]string{
			"command": "0.9.1",
			"tls":     "5.0.1",
		}))

	return providertest.NewProviderTest(dir, opts...)
}
