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

  # if you use local .env uncomment bellow
  # url = local.envfile["DATABASE_URL"]
  # dev = local.envfile["DATABASE_DEV_URL"]

  # also delete url & dev bellow
  url = getenv["DATABASE_URL"]
  dev = getenv["DATABASE_DEV_URL"]

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

# https://atlasgo.io/faq/dotenv-files#write-an-hcl-expression-to-load-the-file-into-atlas
# if you use local .env uncomment bellow until end this page 
# variable "envfile" {
#    type    = string
#    default = ".env"
#}

#locals {
#    envfile = {
#        for line in split("\n", file(var.envfile)): split("=", line)[0] => regex("=(.*)", line)[0]
#        if !startswith(line, "#") && length(split("=", line)) > 1
#    }
#}