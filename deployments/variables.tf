variable "service_name" {
  default     = "sentron-sourcemaps"
  description = "Name of service"
}

variable "environment" {
  default     = "dev"
  description = "Service environment"
}

variable "zip_location" {
  default     = ".build/sentron-sourcemaps.zip"
  description = "Location of the lambda zip"
}

variable "binary_path" {
  default     = "bin/handler"
  description = "Binary path for the lambda to execute"
}

variable "api_domain_name" {
  default     = "sourcemaps.sentron-int.com"
  description = "Domain name for API"
}

variable "hosted_zone_name" {
  default     = "sentron-int.com"
  description = "Hosted zone name"
}