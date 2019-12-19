# Calculator made in GO and provisioned Terraform

A simple go calculator HTTP Service provided with Terraform by using Amazon Web Services such as: EC2 machine (Elastic Compute Cloud),
ELB (Elastic Load Balancer) and Autoscaling Group.

Terraform by Hashicorp:

"Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. Terraform can manage existing and popular service providers as well as custom in-house solutions."

Ansible:

"Ansible is an open-source software provisioning, configuration management, and application-deployment tool."

Jenkins:

"The leading open source automation server, Jenkins provides hundreds of plugins to support building, deploying and automating any project."

The main goal is to automate the process to create an image with our application using packer and ansible to provisioning.
Then we use terraform to bake and launch our infrastructure provided by AWS: our machine, the load balance and autoscaling.

## Summary

- [Requirements](#requirements)
- [Installation](#installation)
- [AWS Account and Enviroment Variables](#aws-account-and-enviroment-variables)
- [Running the Application](#running-the-application)
- [Author](#author)

## Requirements

- [Git](https://git-scm.com/)
- [Packer]()
- [Ansible]()
- [Terraform]()

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

- Installing terraform

```
wget https://releases.hashicorp.com/terraform/0.11.13/terraform_0.11.13_linux_amd64.zip
sudo unzip ./terraform_0.11.13_linux_amd64.zip -d /usr/local/bin/
```

- Clone the project from github:

```
$git clone https://github.com/fabioqmarsiaj/jovens-talentos/tree/tema13
```

## AWS Account and Enviroment Variables

Create an AWS account so you can create your machine, generate AWS key and secret and your key pairs to generate your machine.

First you must have 4 enviroment variables on your linux: (at .bashrc)

export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export TF_VAR_AWS_ACCESS_KEY_ID=""
export TF_VAR_AWS_SECRET_ACCESS_KEY=""

PS: Both access key and secret are equals.

After you create your AWS Account access yourprofile > my security credentials > access keys

Then create your own so you can set the variables above.

Also, it's important to create your own key pars so the terraform on our application can create a new machine.

So at the aws site logged in on your account:

services > EC2 > Key Pairs (At Network & Security label)

Create your key pairs and then it will download your kp_t2_micro.pem file.

PS: Move the file to ~/.ssh folder and give permission:

```
sudo chmod 400 ~/.ssh/kp_t2_micro.pem
```

## Running the Application

Now your good to go and run the packer file to generate our image with the go application.

- Access ~/3-devops/fabio-marsiaj/Tema13/roles/tasks/

```
packer build build.json
```

- Now you can initialize terraform at ~/3-devops/fabio-marsiaj/Tema13/roles/tasks/

```
terraform init
```

And then:

```
terraform apply
```

Answer 'yes' to keep going.

- Finally you can access your machine to run the Go service!

```
ssh -i ~/.ssh/kp_t2_micro.pem ec2-user@18.222.178.9
```

PS: Note that you must check the machine you want to access IP (In my case it's 18.222.178.9):

At the AWS website:

services > EC2 > Instances (under INSTANCES label)

Then click one of the instances and check the "IPv4 Public IP".
Also you can check all three machines created and load balanced.

Endpoints:

You can make simple calculations such as: sum, sub, mul or div by accessing the following endpoints:

http://yourIP:5000/calc/{operation}/{number1}/{number2}

At {operation} choose: sum, sub, mul or div.

If you want to check out the calculator history just access:
http://yourIP:5000/calc/history

## Author

**Fabio Quinto Marsiaj** - [GitHub](https://github.com/fabioqmarsiaj)

<a href="https://github.com/fabioqmarsiaj">
    <img
    alt="Imagem do Autor Fabio Marsiaj" src="https://avatars0.githubusercontent.com/u/34289167?s=460&v=4" width="100">
</a>
