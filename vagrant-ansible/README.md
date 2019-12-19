# Calculator made in GO running at Vagrant VM and provisioning with Ansible

A simple go calculator HTTP Service running at Vagrant.

By Hashicorp:

"Vagrant is a tool for building and managing virtual machine environments in a single workflow. With an easy-to-use workflow and focus on automation, Vagrant lowers development environment setup time, increases production parity, and makes the "works on my machine" excuse a relic of the past."

Ansible:

"Ansible is an open-source software provisioning, configuration management, and application-deployment tool."

## Summary

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Check go service status](#check_go_service_status)
- [Author](#author)

## Introduction

Access the folder ~/jts.devops.2019.2/fabio-marsiaj/Tema9/ and run the Vagrant machine:

```
  $vagrant up
```

It will be provisioned at http://99.99.99.9/5000/

Endpoints:

You can make simple calculations such as: sum, sub, mul or div by accessing the following endpoints:

http://99.99.99.9/5000//calc/{operation}/{number1}/{number2}

At {operation} choose: sum, sub, mul or div.

If you want to check out the calculator history just access:
http://99.99.99.9/5000/calc/history

## Requirements

- [Git](https://git-scm.com/)
- [Go](https://golang.org/)
- [Virtual Box]()
- [Vagrant]()

## Installation

- Installing Git:

```
$apt-get install git
```

- Installing Virtual Box:

```
$apt-get install virtualbox
```

- Installing Vagrant:

```
$apt-get install vagrant
```

- Clone the project from github:

```
$git clone https://github.com/fabioqmarsiaj/jts.devops.2019.2/tree/tema9
```

## Check Go Service Status

After you run Vagrant:

```
  $vagrant up
```

You can access it:

```
  $vagrant ssh
```

Then run ls and check all files, run the check_status.sh:

```
  $./check_status.sh
```

If you want to stop the Vagrant:

```
  $vagrant halt
```

If you want to delete the Vagrant:

```
  $vagrant destroy
```

## Author

**Fabio Quinto Marsiaj** - [GitHub](https://github.com/fabioqmarsiaj)

<a href="https://github.com/fabioqmarsiaj">
    <img
    alt="Imagem do Autor Fabio Marsiaj" src="https://avatars0.githubusercontent.com/u/34289167?s=460&v=4" width="100">
</a>
