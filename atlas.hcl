data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./atlas/loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url

  # uncomment bellow if you use local .env 
  # url = local.envfile["DATABASE_URL"]
  # dev = local.envfile["DATABASE_DEV_URL"]

  # also delete url & dev bellow
  url = var.database_url
  dev = var.database_dev_url

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

# delete or comment this variable if you use .env file
variable "database_url" {
  type    = string
  default = getenv("DATABASE_URL")
}
# delete or comment this variable if you use .env file
variable "database_dev_url" {
  type    = string
  default = getenv("DATABASE_DEV_URL")
}

# https://atlasgo.io/faq/dotenv-files#write-an-hcl-expression-to-load-the-file-into-atlas
# uncomment bellow if you use local .env file
# variable "envfile" {
#    type    = string
#    default = ".env"
#}

# uncomment bellow if you use local .env file
#locals {
#    envfile = {
#        for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
#        if !startswith(line, "#") && length(split("=", line)) > 1
#    }
#}