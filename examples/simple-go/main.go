package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	// kthw "github.com/unstoppablemango/pulumi-kubernetes-the-hard-way/sdk"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// rootCa, err := kthw.NewRootCa("simple", &kthw.NewRootCaArgs{
		// 	Algorithm:          kthw.AlgorithmRSA,
		// 	ValidityPeriodOurs: pulumi.Int(256),
		// })
		// if err != nil {
		// 	return err
		// }

		return nil
	})
}
