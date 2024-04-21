import * as pulumiSchema from "./vendor/pulumi-schema";
import * as fs from "fs";
import * as ts from "typescript";
import path = require("path");

/*
 TODO:
 - Type docs
 - Property docs
 */

const headerWarning = `This file was automatically generated by pulumi-provider-scripts.
DO NOT MODIFY IT BY HAND. Instead, modify the source Pulumi Schema file,
and run "pulumi-provider-scripts gen-provider-types" to regenerate this file.`;
interface ExternalRef {
  import: string;
  name: string;
}

type Direction = "Input" | "Output";

const projectDir = path.join(__dirname, "..", "cmd", "pulumi-resource-kubernetes-the-hard-way");

const externalRefs = (() => {
  const packageJson = JSON.parse(fs.readFileSync(path.join(projectDir, "package.json"), "utf-8"));

  const externalRefs: Record<string, ExternalRef> = {};
  const addRef = (name: string) => {
    const importName = "@pulumi/" + name;
    const version = packageJson.dependencies[importName];
    externalRefs[`/${name}/v${version}/schema.json`] = {
      import: importName,
      name,
    };
  };
  addRef("command");
  addRef("kubernetes");
  addRef("random");
  addRef("tls");
  return externalRefs;
})();

const resolveRef = (ref: unknown, direction: Direction): ts.TypeNode => {
  if (typeof ref !== "string") {
    return ts.factory.createKeywordTypeNode(ts.SyntaxKind.UnknownKeyword);
  }
  if (ref === "pulumi.json#/Any") {
    return ts.factory.createKeywordTypeNode(ts.SyntaxKind.AnyKeyword);
  }
  const typesPrefix = "#/types/";
  const resourcesPrefix = "#/resources/";
  if (ref.startsWith(typesPrefix)) {
    const typeName = ref.substring(typesPrefix.length).split(":");
    return ts.factory.createTypeReferenceNode(typeName[2] + direction + "s");
  }
  if (ref.startsWith(resourcesPrefix)) {
    const typeName = ref.substring(resourcesPrefix.length);
    const typeParts = typeName.split(":");
    const resourceName = typeParts[2];
    const path = decodeURIComponent(typeParts[1]).split("/");
    path.pop(); // Last section is same as the resource name
    return ts.factory.createTypeReferenceNode([...path, resourceName].join("."));
  }
  const externalName = ref.split("#")[0];
  const externalRef = externalRefs[externalName];
  if (externalRef !== undefined) {
    const relativeRef = ref.substring(externalName.length);
    if (relativeRef.startsWith(typesPrefix)) {
      const typeName = relativeRef.substring(typesPrefix.length);
      const typeParts = typeName.split(":");
      const resourceName = typeParts[2];
      const path = decodeURIComponent(typeParts[1]).split("/");
      return ts.factory.createTypeReferenceNode(
        [
          externalRef.name,
          "types",
          direction.toLowerCase(),
          ...path.filter(x => !['index', resourceName].includes(x)),
          resourceName,
        ].join("."),
      );
    }
    if (relativeRef.startsWith(resourcesPrefix)) {
      const typeName = relativeRef.substring(resourcesPrefix.length);
      const typeParts = typeName.split(":");
      const resourceName = typeParts[2];
      const path = decodeURIComponent(typeParts[1]).split("/");
      return ts.factory.createTypeReferenceNode(
        [
          externalRef.name,
          ...path.filter(x => ![
            'index',
            resourceName.toLowerCase(),
          ].includes(x.toLowerCase())),
          resourceName,
        ].join("."),
      );
    }
  }

  console.warn("Unresolvable ref:", ref);
  return ts.factory.createKeywordTypeNode(ts.SyntaxKind.UnknownKeyword);
};

const getPlainType = (
  typeDefinition: pulumiSchema.TypeReference,
  direction: Direction,
): ts.TypeNode => {
  switch (typeDefinition.type) {
    case "string":
      return ts.factory.createKeywordTypeNode(ts.SyntaxKind.StringKeyword);
    case "integer":
    case "number":
      return ts.factory.createKeywordTypeNode(ts.SyntaxKind.NumberKeyword);
    case "boolean":
      return ts.factory.createKeywordTypeNode(ts.SyntaxKind.BooleanKeyword);
    case "array":
      return ts.factory.createArrayTypeNode(
        getType(typeDefinition.items as any, direction, direction === "Output"),
      );
    case "object":
      return ts.factory.createTypeReferenceNode("Record", [
        ts.factory.createKeywordTypeNode(ts.SyntaxKind.StringKeyword),
        getType(typeDefinition.additionalProperties as any, direction, direction === "Output"),
      ]);
  }
  return resolveRef(typeDefinition.$ref, direction);
};

const getType = (
  typeDefinition: pulumiSchema.TypeReference,
  direction: Direction,
  plain?: boolean,
): ts.TypeNode => {
  const plainType = getPlainType(typeDefinition, direction);
  if (plain || typeDefinition.plain) {
    return plainType;
  }

  let type = plainType;
  if (typeDefinition.oneOf) {
    const types = typeDefinition.oneOf.map(t => {
      return getPlainType(t, direction);
    });

    type = ts.factory.createUnionTypeNode(types);
  }

  return ts.factory.createTypeReferenceNode("pulumi." + direction, [type]);
};

function genTypeProperties(
  properties: Record<string, pulumiSchema.TypeReference> | undefined,
  required: string[] | undefined,
  direction: Direction,
): ts.TypeElement[] {
  if (properties === undefined) {
    return [];
  }
  const requiredLookup = new Set(required);
  return Object.entries(properties).map(([propKey, typeDefinition]): ts.TypeElement => {
    const type = getType(typeDefinition, direction);

    return ts.factory.createPropertySignature(
      [ts.factory.createModifier(ts.SyntaxKind.ReadonlyKeyword)],
      propKey,
      requiredLookup.has(propKey) ? undefined : ts.factory.createToken(ts.SyntaxKind.QuestionToken),
      type,
    );
  });
}

function genClassProperties(
  properties?: Record<string, pulumiSchema.TypeReference>,
  required?: string[],
): ts.PropertyDeclaration[] {
  if (properties === undefined) {
    return [];
  }
  const requiredLookup = new Set(required);
  return Object.entries(properties).map(([propKey, typeDefinition]) => {
    const type = ts.factory.createUnionTypeNode([
      getType(typeDefinition, "Output", true),
      getType(typeDefinition, "Output", false),
    ]);
    return ts.factory.createPropertyDeclaration(
      undefined,
      [ts.factory.createModifier(ts.SyntaxKind.PublicKeyword)],
      propKey,
      requiredLookup.has(propKey)
        ? ts.factory.createToken(ts.SyntaxKind.ExclamationToken)
        : ts.factory.createToken(ts.SyntaxKind.QuestionToken),
      type,
      undefined,
    );
  });
}

const genResourceArgs = (
  typeName: string,
  resource: pulumiSchema.ObjectTypeDefinition,
): ts.InterfaceDeclaration => {
  const inputProperties = genTypeProperties(
    resource.inputProperties as any,
    resource.requiredInputs as any,
    "Input",
  );
  const inputs = ts.factory.createInterfaceDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    typeName + "Args",
    undefined,
    undefined,
    inputProperties,
  );

  return inputs;
};

const genResourceAbstractType = (
  typeToken: string,
  typeName: string,
  resource: pulumiSchema.ObjectTypeDefinition,
) => {
  const heritage = ts.factory.createHeritageClause(ts.SyntaxKind.ExtendsKeyword, [
    ts.factory.createExpressionWithTypeArguments(
      ts.factory.createPropertyAccessChain(
        ts.factory.createIdentifier("pulumi"),
        undefined,
        "ComponentResource",
      ),
      [ts.factory.createTypeReferenceNode(ts.factory.createIdentifier("TData"), undefined)],
    ),
  ]);
  const constructor = ts.factory.createConstructorDeclaration(
    undefined,
    undefined,
    [
      ts.factory.createParameterDeclaration(
        undefined,
        undefined,
        undefined,
        "name",
        undefined,
        ts.factory.createKeywordTypeNode(ts.SyntaxKind.StringKeyword),
      ),
      ts.factory.createParameterDeclaration(
        undefined,
        undefined,
        undefined,
        "args",
        undefined,
        ts.factory.createTypeReferenceNode("pulumi.Inputs"),
      ),
      ts.factory.createParameterDeclaration(
        undefined,
        undefined,
        undefined,
        "opts",
        undefined,
        ts.factory.createTypeReferenceNode("pulumi.ComponentResourceOptions"),
        ts.factory.createObjectLiteralExpression(undefined),
      ),
    ],
    ts.factory.createBlock(
      [
        ts.factory.createExpressionStatement(
          ts.factory.createCallExpression(ts.factory.createSuper(), undefined, [
            ts.factory.createStringLiteral(typeToken),
            ts.factory.createIdentifier("name"),
            ts.factory.createConditionalExpression(
              ts.factory.createPropertyAccessExpression(
                ts.factory.createIdentifier("opts"),
                ts.factory.createIdentifier("urn"),
              ),
              ts.factory.createToken(ts.SyntaxKind.QuestionToken),
              ts.factory.createObjectLiteralExpression(
                Object.keys((resource?.properties ?? {}) as any).map((prop) =>
                  ts.factory.createPropertyAssignment(
                    ts.factory.createIdentifier(prop),
                    ts.factory.createIdentifier("undefined"),
                  ),
                ),
                false,
              ),
              ts.factory.createToken(ts.SyntaxKind.ColonToken),
              ts.factory.createObjectLiteralExpression(
                [
                  ts.factory.createShorthandPropertyAssignment(
                    ts.factory.createIdentifier("name"),
                    undefined,
                  ),
                  ts.factory.createShorthandPropertyAssignment(
                    ts.factory.createIdentifier("args"),
                    undefined,
                  ),
                  ts.factory.createShorthandPropertyAssignment(
                    ts.factory.createIdentifier("opts"),
                    undefined,
                  ),
                ],
                false,
              ),
            ),
            ts.factory.createIdentifier("opts"),
          ]),
        ),
      ],
      true,
    ),
  );
  const resourceType = ts.factory.createClassDeclaration(
    undefined,
    [
      ts.factory.createModifier(ts.SyntaxKind.ExportKeyword),
      ts.factory.createModifier(ts.SyntaxKind.AbstractKeyword),
    ],
    typeName,
    [
      ts.factory.createTypeParameterDeclaration(
        ts.factory.createIdentifier("TData"),
        undefined,
        ts.factory.createKeywordTypeNode(ts.SyntaxKind.AnyKeyword),
      ),
    ],
    [heritage],
    [...genClassProperties(resource.properties as any, resource.required as any), constructor],
  );

  return resourceType;
};

const genTypeInterfaces = (
  typeName: string,
  resource: pulumiSchema.TypeDefinition,
): (ts.InterfaceDeclaration | ts.TypeAliasDeclaration)[] => {
  if (resource.enum) {
    const unionTypeMembers = (resource.enum as any).map((x: any) =>
      ts.factory.createLiteralTypeNode(ts.factory.createStringLiteral(x.value)),
    );

    const genEnumType = (suffix: string) =>
      ts.factory.createTypeAliasDeclaration(
        undefined,
        [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
        ts.factory.createIdentifier(typeName + suffix),
        undefined,
        ts.factory.createUnionTypeNode(unionTypeMembers),
      );

    return [genEnumType("Inputs"), genEnumType("Outputs")];
  }

  const inputProperties = genTypeProperties(
    resource.properties as any,
    resource.required as any,
    "Input",
  );
  const inputs = ts.factory.createInterfaceDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    typeName + "Inputs",
    undefined,
    undefined,
    inputProperties,
  );

  const outputProperties = genTypeProperties(
    resource.properties as any,
    resource.required as any,
    "Output",
  );
  const outputs = ts.factory.createInterfaceDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    typeName + "Outputs",
    undefined,
    undefined,
    outputProperties,
  );

  return [inputs, outputs];
};

function genFunctionInputs(
  functionName: string,
  objectDef: pulumiSchema.ObjectTypeDetails1 | undefined,
) {
  const inputProperties = genTypeProperties(objectDef?.properties, objectDef?.required, "Input");
  return ts.factory.createInterfaceDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    functionName + "Inputs",
    undefined,
    undefined,
    inputProperties,
  );
}

function genFunctionOutputs(
  functionName: string,
  objectDef: pulumiSchema.ObjectTypeDetails2 | undefined,
) {
  const outputProperties = genTypeProperties(objectDef?.properties, objectDef?.required, "Output");
  return ts.factory.createInterfaceDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    functionName + "Outputs",
    undefined,
    undefined,
    outputProperties,
  );
}

function getTypeName(typeToken: string) {
  const tokenParts = typeToken.split(":");
  const typeName = tokenParts[2];
  return typeName;
}

function getFunctionName(typeToken: string) {
  return getTypeName(typeToken).replace("/", "_");
}

function genResources(resources: pulumiSchema.PulumiPackageMetaschema["resources"]) {
  return Object.entries(resources ?? {}).flatMap(([typeToken, resource]) => {
    const typeName = getTypeName(typeToken);
    return [
      genResourceAbstractType(typeToken, typeName, resource),
      genResourceArgs(typeName, resource),
    ];
  });
}

function genTypes(resources: pulumiSchema.PulumiPackageMetaschema["types"]) {
  return Object.entries(resources ?? {}).flatMap(([key, value]) =>
    genTypeInterfaces(getTypeName(key), value),
  );
}

function genFunctions(functions: pulumiSchema.PulumiPackageMetaschema["functions"]) {
  return Object.entries(functions ?? {}).flatMap(([token, value]) => {
    const functionName = getFunctionName(token);
    return [
      genFunctionInputs(functionName, value.inputs),
      genFunctionOutputs(functionName, value.outputs),
    ];
  });
}

function genFunctionCallsType(functions: pulumiSchema.PulumiPackageMetaschema["functions"]) {
  return ts.factory.createTypeAliasDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    "Functions",
    undefined,
    ts.factory.createTypeLiteralNode(
      Object.keys(functions ?? {}).map((token) => {
        const functionName = getFunctionName(token);
        return ts.factory.createPropertySignature(
          undefined,
          ts.factory.createStringLiteral(token),
          undefined,
          ts.factory.createFunctionTypeNode(
            undefined,
            [
              ts.factory.createParameterDeclaration(
                undefined,
                undefined,
                undefined,
                "inputs",
                undefined,
                ts.factory.createTypeReferenceNode(functionName + "Inputs"),
              ),
            ],
            ts.factory.createTypeReferenceNode("Promise", [
              ts.factory.createTypeReferenceNode(functionName + "Outputs"),
            ]),
          ),
        );
      }),
    ),
  );
}

function genResourceConstructors(resources: pulumiSchema.PulumiPackageMetaschema["resources"]) {
  return ts.factory.createTypeAliasDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    ts.factory.createIdentifier("ResourceConstructor"),
    undefined,
    ts.factory.createTypeLiteralNode(
      Object.entries(resources ?? {}).map(([typeToken, resource]) => {
        const typeName = getTypeName(typeToken);
        return ts.factory.createPropertySignature(
          [ts.factory.createModifier(ts.SyntaxKind.ReadonlyKeyword)],
          ts.factory.createStringLiteral(typeToken),
          undefined,
          ts.factory.createTypeReferenceNode(ts.factory.createIdentifier("ConstructComponent"), [
            ts.factory.createTypeReferenceNode(ts.factory.createIdentifier(typeName), undefined),
          ]),
        );
      }),
    ),
  );
}

function resourceConstructorType() {
  const tType = ts.factory.createIdentifier("T");
  const pulumiId = ts.factory.createIdentifier("pulumi");
  const componentResourceRef = ts.factory.createTypeReferenceNode(
    ts.factory.createQualifiedName(pulumiId, ts.factory.createIdentifier("ComponentResource")),
    undefined,
  );
  const param = (name: string, type: ts.TypeNode) =>
    ts.factory.createParameterDeclaration(
      undefined,
      undefined,
      undefined,
      ts.factory.createIdentifier(name),
      undefined,
      type,
    );
  return ts.factory.createTypeAliasDeclaration(
    undefined,
    [ts.factory.createModifier(ts.SyntaxKind.ExportKeyword)],
    ts.factory.createIdentifier("ConstructComponent"),
    [ts.factory.createTypeParameterDeclaration(tType, componentResourceRef, componentResourceRef)],
    ts.factory.createFunctionTypeNode(
      undefined,
      [
        param("name", ts.factory.createKeywordTypeNode(ts.SyntaxKind.StringKeyword)),
        param("inputs", ts.factory.createKeywordTypeNode(ts.SyntaxKind.AnyKeyword)),
        param(
          "options",
          ts.factory.createTypeReferenceNode(
            ts.factory.createQualifiedName(
              pulumiId,
              ts.factory.createIdentifier("ComponentResourceOptions"),
            ),
            undefined,
          ),
        ),
      ],
      ts.factory.createTypeReferenceNode(tType, undefined),
    ),
  );
}

export function generateProviderTypes(args: { schema: string; out: string }) {
  const schemaPath = path.resolve(args.schema);
  const schemaText = fs.readFileSync(schemaPath, { encoding: "utf-8" });
  const schema: pulumiSchema.PulumiPackageMetaschema = JSON.parse(schemaText);

  const nodes = ts.factory.createNodeArray([
    ts.factory.createJSDocComment(headerWarning),
    ts.factory.createImportDeclaration(
      undefined,
      undefined,
      ts.factory.createImportClause(false, ts.factory.createIdentifier("* as pulumi"), undefined),
      ts.factory.createStringLiteral("@pulumi/pulumi"),
    ),
    resourceConstructorType(),
    genResourceConstructors(schema.resources),
    genFunctionCallsType(schema.functions),
    ...Object.values(externalRefs).map((externalRef) =>
      ts.factory.createImportDeclaration(
        undefined,
        undefined,
        ts.factory.createImportClause(
          false,
          ts.factory.createIdentifier("* as " + externalRef.name),
          undefined,
        ),
        ts.factory.createStringLiteral(externalRef.import),
      ),
    ),
    ...genResources(schema.resources),
    ...genTypes(schema.types),
    ...genFunctions(schema.functions),
  ]);
  const sourceFile = ts.createSourceFile(
    "provider-types.d.ts",
    "",
    ts.ScriptTarget.ES2019,
    undefined,
    ts.ScriptKind.TS,
  );

  const printer = ts.createPrinter({ newLine: ts.NewLineKind.LineFeed });
  const result = printer.printList(ts.ListFormat.MultiLine, nodes, sourceFile);

  const outPath = path.resolve(args.out);
  fs.writeFileSync(outPath, "/* tslint:disable */\n" + result);
}

generateProviderTypes({
  schema: path.join(projectDir, "schema.json"),
  out: path.join(projectDir, "schema-types.ts"),
});
