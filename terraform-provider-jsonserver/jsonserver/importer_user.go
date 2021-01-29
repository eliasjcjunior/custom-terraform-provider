package jsonserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserImportState(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	id := d.Id()

	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/users/%[2]s", baseURL, id), nil)
	if err != nil {
		return nil, err
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var user *User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("phone", user.Phone)
	d.Set("username", user.Username)
	d.Set("website", user.Website)
	d.SetId(strconv.Itoa(user.ID))

	return []*schema.ResourceData{d}, nil
}
