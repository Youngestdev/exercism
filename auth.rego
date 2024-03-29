package httpapi.authz
import input
default allow = false



allow {
  input.request_path == "v1/collections/obs"
  input.company == "geobeyond"
  input.request_method == "POST"
}

allow {
  input.request_path == "v1/collections/geo"
  input.company == "osgeo"
  input.request_method == "DELETE"
}

