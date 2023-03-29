terraform {
	required_providers {
		filecrud = {
			source = "tlopo/golang/filecrud"
			version = "0.1"
		}
	}
}

provider "filecrud" { }

resource "filecrud" "foo" {
	folder = "/tmp"
	content = "some content 3\n"
}
