provider "aws" {
  region            = "${var.region}"
  access_key        = "${var.AWS_ACCESS_KEY_ID}"
  secret_key        = "${var.AWS_SECRET_ACCESS_KEY}"
  aws_region_config = "${var.aws_region_config}"

}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["devops-example"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["self"]
}

resource "aws_launch_configuration" "as_conf" {
  name          = "web_config"
  image_id      = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"
}

resource "aws_security_group" "allow_tls" {
  name        = "allow_tls"
  description = "Allow TLS inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["191.32.54.163/32"]
  }

  ingress {
    from_port   = 5000
    to_port     = 5000
    protocol    = "tcp"
    cidr_blocks = ["191.32.54.163/32"]
  }
}

resource "aws_instance" "web" {
  ami             = "${data.aws_ami.ubuntu.id}"
  instance_type   = "t2.micro"
  security_groups = ["${aws_security_group.allow_tls.name}"]
}

resource "aws_elb" "go-elb" {
  name               = "aws-elb-terraform"
  availability_zones = ["${var.aws_region_config}"]

  listener {
    instance_port     = 5000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }

  instances = ["${aws_instance.web.id}"]
}

resource "aws_autoscaling_group" "asw-asg" {
  name                 = "aws-awg-terraform"
  availability_zones   = ["${var.aws_region_config}"]
  max_size             = 1
  min_size             = 1
  launch_configuration = "${aws_launch_configuration.as_conf.name}"
  load_balancers       = ["${aws_elb.go-elb.id}"]
  tag {
    key                 = "Name"
    value               = "devops-example"
    propagate_at_launch = true
  }
}
