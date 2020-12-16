data "aws_api_gateway_rest_api" "main" {
  name = var.service_name
}

resource "aws_cloudwatch_log_group" "lambda_logs" {
  name = "/aws/lambda/${var.function_name}"
}

resource "aws_lambda_function" "main" {
  filename         = var.zip_location
  function_name    = var.function_name
  role             = var.iam_role
  handler          = var.binary_path
  source_code_hash = filebase64sha256(var.zip_location)
  runtime          = "go1.x"
  publish          = true
}

resource "aws_api_gateway_resource" "main" {
  rest_api_id = data.aws_api_gateway_rest_api.main.id
  parent_id   = data.aws_api_gateway_rest_api.main.root_resource_id
  path_part   = var.path
}

resource "aws_api_gateway_method" "main" {
  rest_api_id   = data.aws_api_gateway_rest_api.main.id
  resource_id   = aws_api_gateway_resource.main.id
  http_method   = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "main" {
  rest_api_id             = data.aws_api_gateway_rest_api.main.id
  resource_id             = aws_api_gateway_resource.main.id
  http_method             = aws_api_gateway_method.main.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  timeout_milliseconds    = 20000
  uri                     = aws_lambda_function.main.invoke_arn
}

resource "aws_lambda_permission" "allow_agw" {
  function_name = aws_lambda_function.main.arn
  action        = "lambda:InvokeFunction"
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${data.aws_api_gateway_rest_api.main.execution_arn}/*/*"
}