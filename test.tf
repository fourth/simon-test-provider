provider "sjj" {
  user = "simon"
}

resource "sjj_test" "foobar" {
  name = "foobar.txt"
  content = "hello world!!!!!!!!!!\n"
}

resource "sjj_test" "baz" {
  count = 10
  name = "baz${count.index}.txt"
  content = "This is file ${count.index}. ${sjj_test.foobar.name} is ${sjj_test.foobar.size} bytes big\n"
}

output "foobar-size" {
  value = "${sjj_test.foobar.size}"
}
