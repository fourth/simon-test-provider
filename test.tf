provider "sjj" {
  user = "simon"
}

resource "sjj_test" "someresource" {
  name = "foobar.txt"
  content = "hello world 2"
}

resource "sjj_test" "newresource" {
  name = "bar.txt"
  content = "${sjj_test.someresource.name}"
}
