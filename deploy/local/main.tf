terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2" 
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

resource "docker_image" "fizzbuzz_image" {
  name = "kodmain/fizzbuzz:latest"
  keep_locally = false
}

resource "docker_container" "fizzbuzz_container" {
  image = docker_image.fizzbuzz_image.image_id
  name  = "fizzbuzz"
  rm    = true

  ports {
    internal = "80"
    external = "80"
  }

  ports {
    internal = "443"
    external = "443"
  }
}
