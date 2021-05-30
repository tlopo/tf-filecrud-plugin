package filecrud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"filecrud": resourceFileCrud(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
