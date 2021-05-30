terraform {
	required_providers {
		filecrud = {
			source = "emirates.com/golang/filecrud"
			version = "0.1"
		}
	}
}

provider "filecrud" { }

resource "filecrud" "foo" {
	folder = "/tmp"
	content = "some content 3\n"
}
