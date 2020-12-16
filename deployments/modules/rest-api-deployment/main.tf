resource "aws_api_gateway_deployment" "sourcemaps" {
  rest_api_id = var.rest_api_id
  stage_name  = var.environment

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_base_path_mapping" "sourcemaps" {
  api_id      = var.rest_api_id
  stage_name  = var.environment
  domain_name = var.domain_name

  depends_on = [
    aws_api_gateway_deployment.sourcemaps
  ]
}