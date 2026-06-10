terraform {
  required_version = ">= 1.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.25"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  default = "us-west-2"
}

variable "environment" {
  default = "dev"
}

variable "cluster_name" {
  default = "escort-crm"
}

output "cluster_endpoint" {
  value = "https://placeholder"
}

output "database_endpoint" {
  value = "postgres://placeholder:5432"
}
