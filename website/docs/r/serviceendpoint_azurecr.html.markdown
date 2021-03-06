---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_azurecr"
description: |-
  Manages a Azure Container Registry service endpoint within Azure DevOps organization.
---

# azuredevops_serviceendpoint_azurecr

Manages a Azure Container Registry service endpoint within Azure DevOps.

## Example Usage

```hcl
resource "azuredevops_project" "project" {
  name       = "Sample Project"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

# azure container registry service connection
resource "azuredevops_serviceendpoint_azurecr" "azurecr" {
  project_id             = azuredevops_project.project.id
  service_endpoint_name  = "Sample AzureCR"
  resource_group            = "sample-rg"
  azurecr_spn_tenantid      = "72f987tg-95f1-87af-91bh-2d8jd091db47"
  azurecr_name              = "sampleAcr"
  azurecr_subscription_id   = "f7ooi795-c577-6210-9886-a5e898uue3gc"
  azurecr_subscription_name = "sample"
}
```

## Argument Reference

The following arguments are supported:

- `project_id` - (Required) The project ID or project name.
- `service_endpoint_name` - (Required) The name you will use to refer to this service connection in task inputs.
- `resource_group` - (Required) The resource group to which the container registry belongs.
- `azurecr_spn_tenantid` - (Required) The tenant id of the service principal.
- `azurecr_name` - (Required) The Azure container registry name.
- `azurecr_subscription_id` - (Required) The subscription id of the Azure targets.
- `azurecr_subscription_name` - (Required) The subscription name of the Azure targets.
- `description` - (Optional) The name you will use to refer to this service connection in task inputs.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service endpoint.
- `project_id` - The project ID or project name.
- `service_endpoint_name` - The Service Endpoint name.

## Relevant Links

- [Azure DevOps Service REST API 5.1 - Service Endpoints](https://docs.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints?view=azure-devops-rest-5.1)
- [Azure Container Registry REST API](https://docs.microsoft.com/en-us/rest/api/containerregistry/)
