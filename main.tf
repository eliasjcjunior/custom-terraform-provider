terraform {
  required_providers {
    jsonserver = {
      versions = ["0.1"]
      source = "local.com/eliasjcjunior/jsonserver"
    }
  }
}

resource "jsonserver_user" "new" {
  name = "My User"
  email = "myuser@mail.com"
  phone = "+1-202-555-0134"
  username = "myuser"
  website = "myuser.com"
}

data "jsonserver_user" "main" {
  depends_on = [jsonserver_user.new]
  where {
    field = "email"
    value = "myuser@mail.com"
  }
}

output "user_name" {
  value = data.jsonserver_user.main.name
}

output "user_email" {
  value = data.jsonserver_user.main.email
}

output "user_website" {
  value = data.jsonserver_user.main.website
}