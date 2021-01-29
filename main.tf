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