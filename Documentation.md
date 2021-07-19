AWS terraform example demo-001
-Run process

This directory contains an example of a  S3 repository with two text files inside.  Each of those files has been timestamped from the execution in Terraform.
To run the demo-001, clone the repository and run terraform apply within the demo-001’s own directory.

For example:
$ git clone https://github.com/hdavila/demo-001
$ cd demo-001/
$ terraform apply
...
-Test process
With the help of Terratest was created a automation test, which checks the configuration and  creation of the bucket and the tests  files inside. 
To test the demo-001, clone the repository and run go test -v 30m terraform_S3_test.go within the demo-001’s own directory.
For example:
$ cd demo-001/
$ go test -v 30m terraform_S3_test.go 
