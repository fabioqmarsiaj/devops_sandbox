# Calculator made in GO automated with Jenkins

A simple go calculator HTTP Service containerized with Docker using Jenkins.

By Docker:

"A container is a standard unit of software that packages up code and all its dependencies so the application runs quickly and reliably from one computing environment to another."

Ansible:

"Ansible is an open-source software provisioning, configuration management, and application-deployment tool."

Jenkins:

"The leading open source automation server, Jenkins provides hundreds of plugins to support building, deploying and automating any project."

## Summary

- [Requirements](#requirements)
- [Installation](#installation)
- [Running Jenkins](#running-jenkins)
- [Author](#author)

## Requirements

- [Git](https://git-scm.com/)
- [Jenkins](https://jenkins.io/)
- [Packer]()
- [Ansible]()
- [Docker]()

## Installation

- Installing Git:

```
$apt-get install git
```

- Installing Packer:

```
$sudo apt-get install packer
```

- Installing Ansible:

```
$sudo apt-get install ansible
```

- Installing Docker

```
$apt-get install docker
```

- Add Jenkins to Docker Group:

```
$sudo usermod -a -G docker jenkins
```

- Give access to .docker/ folder on your sistem:

```
sudo chmod -R 777 .docker/*
```

- Clone the project from github:

```
$git clone https://github.com/fabioqmarsiaj/jovens-talentos/tree/tema12
```

## Running Jenkins

When you install Jenkins make sure to install all recommended plugins, like Git.

After installing Jenkins, it should be available at http://localhost:8080

You should be able to create your own jobs accessing a specific git respository, to do so, follow the steps at:

https://subscription.packtpub.com/book/application_development/9781783553471/2/ch02lvl1sec23/creating-a-new-build-job-in-jenkins-with-git

The calculator will be provisioned at http://localhost:8082/

Endpoints:

You can make simple calculations such as: sum, sub, mul or div by accessing the following endpoints:

http://localhost:8082/calc/{operation}/{number1}/{number2}

At {operation} choose: sum, sub, mul or div.

If you want to check out the calculator history just access:
http://localhost:8082/calc/history

## Author

**Fabio Quinto Marsiaj** - [GitHub](https://github.com/fabioqmarsiaj)

<a href="https://github.com/fabioqmarsiaj">
    <img
    alt="Imagem do Autor Fabio Marsiaj" src="https://avatars0.githubusercontent.com/u/34289167?s=460&v=4" width="100">
</a>
