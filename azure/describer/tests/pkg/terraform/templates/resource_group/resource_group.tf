resource "azurerm_resource_group" "describer-test-rg" {
  count = var.resourceCount
  name     = format("%s-%d",var.resource_group_name,count.index)
  location = var.location              
}