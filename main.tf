terraform {
  required_providers {
    jsonserver = {
      versions = ["0.1"]
      source = "local.com/eliasjcjunior/jsonserver"
    }
  }
}

resource "jsonserver_user" "new" {
  name = ""
  email = ""
  phone = ""
  username = ""
  website = ""
}