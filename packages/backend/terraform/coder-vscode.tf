terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.1"
    }
  }
}

provider "docker" {
  host = "unix:///var/run/docker.sock"
}

# Pulls the image
resource "docker_image" "coder-vscode" {
  name = "linuxserver/code-server:latest"
}

# Create a container
resource "docker_container" "coder-vscode" {
  image = docker_image.coder-vscode.image_id
  name  = "coder-vscode"
  
  ports {
    internal = "8443"
    external = "8443"
  }
}