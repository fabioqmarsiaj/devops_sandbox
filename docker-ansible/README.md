# Calculator made in GO running with Docker and provisioning with Ansible

A simple go calculator HTTP Service containerized with Docker.

By Docker:

"A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another."

Ansible:

"Ansible is an open-source software provisioning, configuration management, and application-deployment tool."

## Summary

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Check go service status](#check_go_service_status)
- [Author](#author)

## Introduction

Access the folder ~/jovens-talentos/3-devops/fabio-marsiaj/Tema10/roles/go/tasks and build the packer file:

```
  $packer build build.json
```

Check the container ID by running:

```
  $docker images
```

Then you need to run the docker image created.

```
  $docker run -d -p 8080:5000 <container-id>
```

It will be provisioned at http://localhost:8080/

Endpoints:

You can make simple calculations such as: sum, sub, mul or div by accessing the following endpoints:

http://localhost:8080/calc/{operation}/{number1}/{number2}

At {operation} choose: sum, sub, mul or div.

If you want to check out the calculator history just access:
http://localhost:8080/calc/history

## Requirements

- [Git](https://git-scm.com/)
- [Docker]()

## Installation

- Installing Git:

```
$apt-get install git
```

- Installing Docker

```
$apt-get install docker
```

- Clone the project from github:

```
$git clone https://github.com/fabioqmarsiaj/jovens-talentos/tree/tema10
```

## Author

**Fabio Quinto Marsiaj** - [GitHub](https://github.com/fabioqmarsiaj)

<a href="https://github.com/fabioqmarsiaj">
    <img
    alt="Imagem do Autor Fabio Marsiaj" src="https://avatars0.githubusercontent.com/u/34289167?s=460&v=4" width="100">
</a>
