resource "google_project" "my_project" {
  name       = "Cartesian"
  project_id = "cartesian-test"
  org_id     = "cartesian-test"
}

resource "google_app_engine_application" "app" {
  project     = google_project.my_project.project_id
  location_id = "us-central"
}
