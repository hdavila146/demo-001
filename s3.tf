
resource "aws_s3_bucket" "bucket1" {
  bucket = var.bucket_name
  acl    = "private"
depends_on = [
    null_resource.write_1, null_resource.write_2
  ]
  tags = {
    Name        = "My bucket"
  }
}
  resource "aws_s3_bucket_object" "file_upload_test1" {
  bucket = var.bucket_name
  key    = var.file_name1
  source = "test1.txt"
  depends_on = [
    aws_s3_bucket.bucket1
  ]
}
resource "aws_s3_bucket_object" "file_upload_test2" {
  bucket = var.bucket_name
  key    = var.file_name2
  source = "test2.txt"
depends_on = [
    aws_s3_bucket.bucket1
  ]
}
output "bucket_id" {
  value = aws_s3_bucket.bucket1.id
  
}
output "aws_region" {
  value = var.aws_region
  
}
output "file_Name1" {
  value = var.file_name1
}
output "file_Name2" {
  value = var.file_name2
  
}
