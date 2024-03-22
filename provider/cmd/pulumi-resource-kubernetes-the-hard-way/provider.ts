import * as pulumi from "@pulumi/pulumi";
import * as provider from "@pulumi/pulumi/provider";

export class Provider implements provider.Provider {
    constructor(readonly version: string, readonly schema: string) { }

    async construct(name: string, type: string, inputs: pulumi.Inputs,
        options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {

        // TODO: Add support for additional component resources here.
        switch (type) {
            // case "kubernetes-the-hard-way:index:StaticPage":
            //     return await constructStaticPage(name, inputs, options);
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    }
}

// async function constructStaticPage(name: string, inputs: pulumi.Inputs,
//     options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {

//     // Create the component resource.
//     const staticPage = new StaticPage(name, inputs as StaticPageArgs, options);

//     // Return the component resource's URN and outputs as its state.
//     return {
//         urn: staticPage.urn,
//         state: {
//             bucket: staticPage.bucket,
//             websiteUrl: staticPage.websiteUrl,
//         },
//     };
// }
