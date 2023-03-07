module github.com/pulumi/pulumi-github/provider

go 1.13

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

require (
	github.com/hashicorp/go-getter v1.7.0 // indirect
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191212124732-c6ae6269b9d7 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.9.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.1.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.1-0.20200424001829-090f390d7b1a
	github.com/terraform-providers/terraform-provider-bitbucket v1.2.0
)
