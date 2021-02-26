package intersight

import (
	"context"
	"encoding/json"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceApplianceRemoteFileImport() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplianceRemoteFileImportRead,
		Schema: map[string]*schema.Schema{
			"account": {
				Description: "A reference to a iamAccount resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"class_id": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the remote type referred by this relationship.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filename": {
				Description: "The name of the file to be imported.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"hostname": {
				Description: "The hostname of the machine where the file is located.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"is_password_set": {
				Description: "Indicates whether the value of the 'password' property has been set.",
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Description: "Password for remote requiest.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"path": {
				Description: "The port that should be used for the remote request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"port": {
				Description: "The port that should be used for the remote request.",
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"protocol": {
				Description: "Protocol for the remote request.\n* `scp` - Secure Copy Protocol (SCP) to access the file server.\n* `sftp` - SSH File Transfer Protocol (SFTP) to access file server.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
			},
			"username": {
				Description: "The username for the remote request.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func dataSourceApplianceRemoteFileImportRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.ApplianceRemoteFileImport{}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("filename"); ok {
		x := (v.(string))
		o.SetFilename(x)
	}
	if v, ok := d.GetOk("hostname"); ok {
		x := (v.(string))
		o.SetHostname(x)
	}
	if v, ok := d.GetOk("is_password_set"); ok {
		x := (v.(bool))
		o.SetIsPasswordSet(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("password"); ok {
		x := (v.(string))
		o.SetPassword(x)
	}
	if v, ok := d.GetOk("path"); ok {
		x := (v.(string))
		o.SetPath(x)
	}
	if v, ok := d.GetOk("port"); ok {
		x := int64(v.(int))
		o.SetPort(x)
	}
	if v, ok := d.GetOk("protocol"); ok {
		x := (v.(string))
		o.SetProtocol(x)
	}
	if v, ok := d.GetOk("username"); ok {
		x := (v.(string))
		o.SetUsername(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of ApplianceRemoteFileImport object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.ApplianceApi.GetApplianceRemoteFileImportList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching ApplianceRemoteFileImport: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for ApplianceRemoteFileImport list: %s", err.Error())
	}
	var s = &models.ApplianceRemoteFileImportList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to ApplianceRemoteFileImport list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for ApplianceRemoteFileImport data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for ApplianceRemoteFileImport data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.ApplianceRemoteFileImport{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}

			if err := d.Set("account", flattenMapIamAccountRelationship(s.GetAccount(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Account: %s", err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}
			if err := d.Set("filename", (s.GetFilename())); err != nil {
				return diag.Errorf("error occurred while setting property Filename: %s", err.Error())
			}
			if err := d.Set("hostname", (s.GetHostname())); err != nil {
				return diag.Errorf("error occurred while setting property Hostname: %s", err.Error())
			}
			if err := d.Set("is_password_set", (s.GetIsPasswordSet())); err != nil {
				return diag.Errorf("error occurred while setting property IsPasswordSet: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("path", (s.GetPath())); err != nil {
				return diag.Errorf("error occurred while setting property Path: %s", err.Error())
			}
			if err := d.Set("port", (s.GetPort())); err != nil {
				return diag.Errorf("error occurred while setting property Port: %s", err.Error())
			}
			if err := d.Set("protocol", (s.GetProtocol())); err != nil {
				return diag.Errorf("error occurred while setting property Protocol: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}
			if err := d.Set("username", (s.GetUsername())); err != nil {
				return diag.Errorf("error occurred while setting property Username: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
