provider "aws" {
  region  = "eu-west-1"
  profile = "sentron"
}

locals {
  prefix = "tf-${var.service_name}-${var.environment}"
}

module "lambda_iam_role" {
  source       = "./modules/lambda-iam-role"
  service_name = var.service_name
}

module "rest_api" {
  source           = "./modules/rest-api"
  service_name     = var.service_name
  domain_name      = var.api_domain_name
  hosted_zone_name = var.hosted_zone_name
}

module "generate_source_extract" {
  source = "./modules/rest-api-endpoint"

  path          = "generate-source-extract"
  method        = "POST"
  function_name = "${local.prefix}-generate-source-extract"
  iam_role      = module.lambda_iam_role.role_arn
  binary_path   = var.binary_path
  zip_location  = var.zip_location
  service_name  = var.service_name

  depends_on = [module.rest_api]
}

module "locate_sourcemap" {
  source = "./modules/rest-api-endpoint"

  path          = "locate-sourcemap"
  method        = "POST"
  function_name = "${local.prefix}-locate-sourcemap"
  iam_role      = module.lambda_iam_role.role_arn
  binary_path   = var.binary_path
  zip_location  = var.zip_location
  service_name  = var.service_name

  depends_on = [module.rest_api]
}

module "deployment" {
  source      = "./modules/rest-api-deployment"
  rest_api_id = module.rest_api.id
  environment = var.environment
  domain_name = var.api_domain_name

  depends_on = [
    module.generate_source_extract,
    module.locate_sourcemap
  ]
}
