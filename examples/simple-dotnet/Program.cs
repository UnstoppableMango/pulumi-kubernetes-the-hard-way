using System.Collections.Generic;
using Pulumi;
using UnMango.KubernetesTheHardWay;

return await Deployment.RunAsync(() =>
{
    var ca = new RootCa("simple", new RootCaArgs {
        Algorithm = Algorithm.RSA,
        ValidityPeriodHours = 256,
    });

    // Export outputs here
    return new Dictionary<string, object?>
    {
        ["outputKey"] = "outputValue"
    };
});
