package filecrud

import (
	//"fmt"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFileCrud() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFileCrudCreate,
		ReadContext:   resourceFileCrudRead,
		UpdateContext: resourceFileCrudUpdate,
		DeleteContext: resourceFileCrudDelete,
		Schema: map[string]*schema.Schema{
			"folder": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Optional: false,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceFileCrudCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	folder := d.Get("folder").(string)
	content := d.Get("content").(string)

	fc := FileCrud{folder}
	id := fc.Create(content)

	d.SetId(id)

	return diags
}

func resourceFileCrudRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	folder := d.Get("folder").(string)
	fc := FileCrud{folder}
	if fc.Exist(id) {
		content := fc.Read(id)
		d.Set("content", content)
	} else {
		d.Set("content", "")
	}
	return diags
}

func resourceFileCrudUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	folder := d.Get("folder").(string)
	content := d.Get("content").(string)
	fc := FileCrud{folder}
	fc.Update(id, content)
	return diags
}

func resourceFileCrudDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	id := d.Id()
	folder := d.Get("folder").(string)
	fc := FileCrud{folder}
	fc.Delete(id)
	return diags
}
