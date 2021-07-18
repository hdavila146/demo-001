
resource "null_resource" "write_1" {
  provisioner "local-exec" {
    command = "sh timestamp.sh >> test1.txt"
    interpreter = ["/bin/bash","-c"]
   
  }
}
resource "null_resource" "write_2" {
  provisioner "local-exec" {
    command = "sh timestamp.sh >> test2.txt"
    interpreter = ["/bin/bash","-c"]
  }
}
