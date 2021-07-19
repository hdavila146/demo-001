package aws

import (
	
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
    "github.com/gruntwork-io/terratest/modules/terraform"
    
)
   // An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
    func TestTerraformAwsS3(t *testing.T) {
	t.Parallel()


	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../demo-001",
		
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

    // Run `terraform output` to get the value of an output variable
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	awsRegion := terraform.Output(t, terraformOptions, "aws_region")
	key1 := terraform.Output(t, terraformOptions, "file_Name1")
	key2 := terraform.Output(t, terraformOptions, "file_Name2")

	// Verify that our Bucket Exists
	aws.AssertS3BucketExistsE(t, awsRegion, bucketID) 

	//Verify files inside the bucket
	aws.GetS3ObjectContentsE(t, awsRegion, bucketID, key1)
	aws.GetS3ObjectContentsE(t, awsRegion, bucketID, key2)
	
	}