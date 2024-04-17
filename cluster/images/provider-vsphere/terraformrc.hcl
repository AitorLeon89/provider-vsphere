provider_installation {
  filesystem_mirror {
    path    = "/terraform.local/local"
    include = ["*/*"]
  }
  direct {
    exclude = ["*/*"]
  }
}
