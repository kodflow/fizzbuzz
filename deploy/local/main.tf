## Configuration
terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2" 
    }
    random = {
      source  = "hashicorp/random"
      version = "3.6.3"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

## Docker images
resource "docker_image" "fizzbuzz_image" {
  name = "kodmain/fizzbuzz:latest"
  keep_locally = false
}

resource "docker_image" "prometheus_image" {
  name         = "prom/prometheus:latest"
  keep_locally = false
}

resource "docker_image" "grafana_image" {
  name         = "grafana/grafana:latest"
  keep_locally = false
}

## Network
resource "docker_network" "app_network" {
  name = "fizzbuzz"
}

## Docker containers
resource "docker_container" "fizzbuzz_container" {
  image = docker_image.fizzbuzz_image.image_id
  name  = "fizzbuzz"
  rm    = true
  networks_advanced {
    name = docker_network.app_network.name
  }

  ports {
    internal = "80"
    external = "80"
  }

  ports {
    internal = "443"
    external = "443"
  }
}

resource "docker_container" "prometheus_container" {
  image = docker_image.prometheus_image.image_id
  name  = "prometheus"
  rm    = true
  networks_advanced {
    name = docker_network.app_network.name
  }

  ports {
    internal = "9090"
    external = "9090"
  }

  volumes {
    container_path = "/etc/prometheus/prometheus.yml"
    host_path      = "${path.cwd}/prometheus/prometheus.yml"
    read_only      = true
  }
}

# Generated per stack, never committed; read it with
# `terraform output -raw grafana_admin_password`. Dashboards stay readable
# without it through the anonymous Viewer role.
resource "random_password" "grafana_admin" {
  length  = 24
  special = false
}

output "grafana_admin_password" {
  value     = random_password.grafana_admin.result
  sensitive = true
}

resource "docker_container" "grafana_container" {
  image = docker_image.grafana_image.image_id
  name  = "grafana"
  rm    = true

  env = [
    "GF_AUTH_ANONYMOUS_ENABLED=true",
    "GF_AUTH_ANONYMOUS_ORG_NAME=Main Org.",
    "GF_AUTH_ANONYMOUS_ORG_ROLE=Viewer",
    "GF_SECURITY_ADMIN_USER=admin",
    "GF_SECURITY_ADMIN_PASSWORD=${random_password.grafana_admin.result}",
    "GF_LOG_MODE=console",
    "GF_PATHS_PROVISIONING=/etc/grafana/provisioning",
    "GF_DASHBOARDS_DEFAULT_HOME_DASHBOARD_PATH=/var/lib/grafana/dashboards/fizzbuzz.json"
  ]

  networks_advanced {
    name = docker_network.app_network.name
  }

  # Loopback only: the admin login must not be reachable from the network.
  ports {
    internal = "3000"
    external = "3000"
    ip       = "127.0.0.1"
  }

  volumes {
    container_path = "/etc/grafana/provisioning"
    host_path      = "${path.cwd}/grafana"
  }

  volumes {
    container_path = "/var/lib/grafana/dashboards"
    host_path      = "${path.cwd}/grafana/imports"
  }

}
