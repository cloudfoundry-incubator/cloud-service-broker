variable db_name { type = string }
variable server { type = string }
variable server_credentials { type = map }
variable labels { type = map }
variable pricing_tier { type = string }
variable cores { type = number }
variable storage_gb { type = number }

data "azurerm_sql_server" "azure_sql_db_server" {
  name                         = var.server_credentials[var.server].server_name
  resource_group_name          = var.server_credentials[var.server].server_resource_group
}

resource "azurerm_sql_database" "azure_sql_db" {
  name                = var.db_name
  resource_group_name = data.azurerm_sql_server.azure_sql_db_server.resource_group_name
  location            = data.azurerm_sql_server.azure_sql_db_server.location
  server_name         = data.azurerm_sql_server.azure_sql_db_server.name
  requested_service_objective_name = format("%s_Gen5_%d", var.pricing_tier, var.cores)
  max_size_bytes      = var.storage_gb * 1024 * 1024 * 1024
  tags                = var.labels
}

locals {
  serverFQDN = data.azurerm_sql_server.azure_sql_db_server.fqdn
}

output "sqldbName" {value = azurerm_sql_database.azure_sql_db.name}
output "sqlServerName" {value = data.azurerm_sql_server.azure_sql_db_server.name}
output "sqlServerFullyQualifiedDomainName" {value = local.serverFQDN}
output "hostname" {value = local.serverFQDN}
output "port" {value = 1433}
output "name" {value = azurerm_sql_database.azure_sql_db.name}
output "username" {value = var.server_credentials[var.server].admin_username}
output "password" {value = var.server_credentials[var.server].admin_password}
