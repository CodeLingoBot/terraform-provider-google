// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitoringGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitoringGroupCreate,
		Read:   resourceMonitoringGroupRead,
		Update: resourceMonitoringGroupUpdate,
		Delete: resourceMonitoringGroupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMonitoringGroupImport,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parent_name": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMonitoringGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	parentNameProp, err := expandMonitoringGroupParentName(d.Get("parent_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_name"); !isEmptyValue(reflect.ValueOf(parentNameProp)) && (ok || !reflect.DeepEqual(v, parentNameProp)) {
		obj["parentName"] = parentNameProp
	}
	isClusterProp, err := expandMonitoringGroupIsCluster(d.Get("is_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_cluster"); !isEmptyValue(reflect.ValueOf(isClusterProp)) && (ok || !reflect.DeepEqual(v, isClusterProp)) {
		obj["isCluster"] = isClusterProp
	}
	displayNameProp, err := expandMonitoringGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	filterProp, err := expandMonitoringGroupFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/projects/{{project}}/groups")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Group: %#v", obj)
	res, err := sendRequest(config, "POST", url, obj)
	if err != nil {
		return fmt.Errorf("error creating Group: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Group %q: %#v", d.Id(), res)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("create response didn't contain critical fields. Create may not have succeeded.")
	}
	d.Set("name", name.(string))
	d.SetId(name.(string))

	return resourceMonitoringGroupRead(d, meta)
}

func resourceMonitoringGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	res, err := sendRequest(config, "GET", url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("MonitoringGroup %q", d.Id()))
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}

	if err := d.Set("parent_name", flattenMonitoringGroupParentName(res["parentName"], d)); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}
	if err := d.Set("name", flattenMonitoringGroupName(res["name"], d)); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}
	if err := d.Set("is_cluster", flattenMonitoringGroupIsCluster(res["isCluster"], d)); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}
	if err := d.Set("display_name", flattenMonitoringGroupDisplayName(res["displayName"], d)); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}
	if err := d.Set("filter", flattenMonitoringGroupFilter(res["filter"], d)); err != nil {
		return fmt.Errorf("error reading Group: %s", err)
	}

	return nil
}

func resourceMonitoringGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	parentNameProp, err := expandMonitoringGroupParentName(d.Get("parent_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parentNameProp)) {
		obj["parentName"] = parentNameProp
	}
	isClusterProp, err := expandMonitoringGroupIsCluster(d.Get("is_cluster"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_cluster"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, isClusterProp)) {
		obj["isCluster"] = isClusterProp
	}
	displayNameProp, err := expandMonitoringGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	filterProp, err := expandMonitoringGroupFilter(d.Get("filter"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("filter"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Group %q: %#v", d.Id(), obj)
	_, err = sendRequest(config, "PUT", url, obj)

	if err != nil {
		return fmt.Errorf("error updating Group %q: %s", d.Id(), err)
	}

	return resourceMonitoringGroupRead(d, meta)
}

func resourceMonitoringGroupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	lockName, err := replaceVars(d, config, "stackdriver/groups/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "https://monitoring.googleapis.com/v3/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Group %q", d.Id())
	res, err := sendRequest(config, "DELETE", url, obj)
	if err != nil {
		return handleNotFoundError(err, d, "Group")
	}

	log.Printf("[DEBUG] Finished deleting Group %q: %#v", d.Id(), res)
	return nil
}

func resourceMonitoringGroupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	config := meta.(*Config)

	// current import_formats can't import id's with forward slashes in them.
	parseImportId([]string{"(?P<name>.+)"}, d, config)

	return []*schema.ResourceData{d}, nil
}

func flattenMonitoringGroupParentName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringGroupName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringGroupIsCluster(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringGroupDisplayName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenMonitoringGroupFilter(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func expandMonitoringGroupParentName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupIsCluster(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupDisplayName(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupFilter(v interface{}, d *schema.ResourceData, config *Config) (interface{}, error) {
	return v, nil
}
