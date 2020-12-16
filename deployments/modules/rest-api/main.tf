data "aws_acm_certificate" "issued" {
  domain   = var.hosted_zone_name
  statuses = ["ISSUED"]
}

data "aws_route53_zone" "selected" {
  name         = var.hosted_zone_name
  private_zone = false
}

resource "aws_api_gateway_rest_api" "main" {
  name = var.service_name
}

resource "aws_api_gateway_domain_name" "main" {
  domain_name              = var.domain_name
  regional_certificate_arn = data.aws_acm_certificate.issued.arn

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_route53_record" "main" {
  name    = aws_api_gateway_domain_name.main.domain_name
  type    = "A"
  zone_id = data.aws_route53_zone.selected.id

  alias {
    evaluate_target_health = true
    name                   = aws_api_gateway_domain_name.main.regional_domain_name
    zone_id                = aws_api_gateway_domain_name.main.regional_zone_id
  }
}

output "id" {
  value = aws_api_gateway_rest_api.main.id
}

output "execution_arn" {
  value = aws_api_gateway_rest_api.main.execution_arn
}

output "root_resource_id" {
  value = aws_api_gateway_rest_api.main.root_resource_id
}