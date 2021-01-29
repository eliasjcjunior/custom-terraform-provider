package jsonserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const baseURL = "http://localhost:3000"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	name := d.Get("name").(string)
	email := d.Get("email").(string)
	phone := d.Get("phone").(string)
	website := d.Get("website").(string)
	username := d.Get("username").(string)

	user := &User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Website:  website,
		Username: username,
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", baseURL), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	var response User
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return err
	}

	d.Set("name", response.Name)
	d.Set("email", response.Email)
	d.Set("phone", response.Phone)
	d.Set("username", response.Username)
	d.Set("website", response.Website)

	d.SetId(strconv.Itoa(response.ID))

	return nil
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%[1]s/users/%[2]s", baseURL, d.Id()), nil)
	if err != nil {
		return err
	}

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	var user *User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("phone", user.Phone)
	d.Set("username", user.Username)
	d.Set("website", user.Website)

	d.SetId(d.Id())

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	name := d.Get("name").(string)
	email := d.Get("email").(string)
	phone := d.Get("phone").(string)
	website := d.Get("website").(string)
	username := d.Get("username").(string)

	user := &User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Website:  website,
		Username: username,
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%[1]s/users/%[2]s", baseURL, d.Id()), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	var response User
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return err
	}

	d.Set("name", response.Name)
	d.Set("email", response.Email)
	d.Set("phone", response.Phone)
	d.Set("username", response.Username)
	d.Set("website", response.Website)

	d.SetId(d.Id())

	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%[1]s/users/%[2]s", baseURL, d.Id()), nil)
	if err != nil {
		return err
	}

	r, err := client.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	var user *User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"phone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"website": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
