package jsonserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Filter struct {
	Field string
	Value string
}

func buildFilter(set *schema.Set) *Filter {
	var filter *Filter
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		filter = &Filter{
			Field: m["field"].(string),
			Value: m["value"].(string),
		}
	}
	return filter
}

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"where": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
						},

						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"website": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics
	where := d.Get("where").(*schema.Set)
	filter := buildFilter(where)

	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/users?%[2]s=%[3]s", baseURL, filter.Field, filter.Value), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()

	var user []*User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", user[0].Name)
	d.Set("email", user[0].Email)
	d.Set("phone", user[0].Phone)
	d.Set("username", user[0].Username)
	d.Set("website", user[0].Website)
	d.SetId(strconv.Itoa(user[0].ID))

	return diags
}
