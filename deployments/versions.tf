terraform {
  required_version = ">= 0.14.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }

  backend "s3" {
    bucket  = "terraform-state-239591708370"
    key     = "sentron-sourcemaps/terraform.tfstate"
    region  = "eu-west-1"
    profile = "sentron"
  }
}
