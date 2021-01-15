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

func dataSourceHyperflexVmBackupInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexVmBackupInfoRead,
		Schema: map[string]*schema.Schema{
			"additional_properties": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: SuppressDiffAdditionProps,
			},
			"backup_status": {
				Description: "Description of the backup status of this VmBackupInfo.\n* `InitializingProtection` - Protection has started, but not completed.\n* `Protected` - Protection has completed successfully.\n* `ExceedsInterval` - Protection has not completed successfully in over two times the backup interval.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"cluster_entity_reference": {
				Description: "Entity reference to the cluster this VmBackupInfo is associated with.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"confignum": {
							Description: "Configuration number for this reference.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "Uuid of the entity for this reference.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"idtype": {
							Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the entity for this entity reference.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"cluster_id_protection_info_map": {
				Type:     schema.TypeList,
				Optional: true,
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
						"cluster_id": {
							Description: "The Cluster Id we are using to map to ProtectionInfo.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"protection_info": {
							Description: "ProtectionInfo that is being stored for this Virtual Machine.",
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"vm_current_protection_info": {
										Description: "Current snapshot protection information.",
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
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"replication_status": {
													Description: "Replication status information for this particular snapshot.",
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
															"bytes_replicated": {
																Description: "Number of bytes currently replicated.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"class_id": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"end_time": {
																Description: "Replication end time for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"error": {
																Description: "Error information in case of failure.",
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
																		"message": {
																			Description: "The error message string for this error stack.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"message_id": {
																			Description: "The error message ID for this error stack.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pack_entity_reference": {
																Description: "EntityReference for the Replication Pack.",
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
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"confignum": {
																			Description: "Configuration number for this reference.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"id": {
																			Description: "Uuid of the entity for this reference.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"idtype": {
																			Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"name": {
																			Description: "Name of the entity for this entity reference.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"type": {
																			Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"pct_complete": {
																Description: "Completion percentage for the replication job.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"rpo_status": {
																Description: "Status for timeliness of replication job in relation to the schedule interval.",
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"actual": {
																			Description: "Actual end time for the snapshot.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"additional_properties": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			DiffSuppressFunc: SuppressDiffAdditionProps,
																		},
																		"class_id": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"expected": {
																			Description: "Expected end time for the snapshot.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"rpo_exceeded": {
																			Description: "A flag to determine if snapshot delivery is delayed.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"start_time": {
																Description: "Replication start time for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"status": {
																Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"site": {
													Description: "Cluster site for this snapshot.\n* `PRIMARY` - The cluster site for this backup is primary.\n* `SECONDARY` - The cluster site for this backup is secondary.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"snapshot_status": {
													Description: "Status for this Snapshot Point.",
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"description": {
																Description: "Description of this Snapshot Point.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"error": {
																Description: "Error information in case of failure.",
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
																		"message": {
																			Description: "The error message string for this error stack.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"message_id": {
																			Description: "The error message ID for this error stack.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pct_complete": {
																Description: "Completion percentage for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"status": {
																Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"timestamp": {
																Description: "Timestamp at which the Snapshot is taken.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"used_space": {
																Description: "Space Used by this Snapshot Point.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"vm_snapshot_entity_reference": {
													Description: "EntityReference of this VmSnapshot.",
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"confignum": {
																Description: "Configuration number for this reference.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"id": {
																Description: "Uuid of the entity for this reference.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"idtype": {
																Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"name": {
																Description: "Name of the entity for this entity reference.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"type": {
																Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"vm_last_successful_protection_info": {
										Description: "Last successful snapshot protection information.",
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
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"replication_status": {
													Description: "Replication status information for this particular snapshot.",
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
															"bytes_replicated": {
																Description: "Number of bytes currently replicated.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"class_id": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"end_time": {
																Description: "Replication end time for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"error": {
																Description: "Error information in case of failure.",
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
																		"message": {
																			Description: "The error message string for this error stack.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"message_id": {
																			Description: "The error message ID for this error stack.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pack_entity_reference": {
																Description: "EntityReference for the Replication Pack.",
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
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"confignum": {
																			Description: "Configuration number for this reference.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"id": {
																			Description: "Uuid of the entity for this reference.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"idtype": {
																			Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"name": {
																			Description: "Name of the entity for this entity reference.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"type": {
																			Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"pct_complete": {
																Description: "Completion percentage for the replication job.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"rpo_status": {
																Description: "Status for timeliness of replication job in relation to the schedule interval.",
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"actual": {
																			Description: "Actual end time for the snapshot.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"additional_properties": {
																			Type:             schema.TypeString,
																			Optional:         true,
																			DiffSuppressFunc: SuppressDiffAdditionProps,
																		},
																		"class_id": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																		},
																		"expected": {
																			Description: "Expected end time for the snapshot.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"rpo_exceeded": {
																			Description: "A flag to determine if snapshot delivery is delayed.",
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"start_time": {
																Description: "Replication start time for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"status": {
																Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"site": {
													Description: "Cluster site for this snapshot.\n* `PRIMARY` - The cluster site for this backup is primary.\n* `SECONDARY` - The cluster site for this backup is secondary.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"snapshot_status": {
													Description: "Status for this Snapshot Point.",
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"description": {
																Description: "Description of this Snapshot Point.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"error": {
																Description: "Error information in case of failure.",
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
																		"message": {
																			Description: "The error message string for this error stack.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																		"message_id": {
																			Description: "The error message ID for this error stack.",
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Computed:    true,
																		},
																		"object_type": {
																			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																			Type:        schema.TypeString,
																			Optional:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"pct_complete": {
																Description: "Completion percentage for this snapshot.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"status": {
																Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"timestamp": {
																Description: "Timestamp at which the Snapshot is taken.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"used_space": {
																Description: "Space Used by this Snapshot Point.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
												"vm_snapshot_entity_reference": {
													Description: "EntityReference of this VmSnapshot.",
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
																Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
															},
															"confignum": {
																Description: "Configuration number for this reference.",
																Type:        schema.TypeInt,
																Optional:    true,
																Computed:    true,
															},
															"id": {
																Description: "Uuid of the entity for this reference.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"idtype": {
																Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"name": {
																Description: "Name of the entity for this entity reference.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"object_type": {
																Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"type": {
																Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
									"vm_space_usage": {
										Description: "Protection space usage for this snapshot.",
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
													Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
												},
												"object_type": {
													Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"space_usage": {
													Description: "Space usage of the VM from StDataServiceManager.",
													Type:        schema.TypeInt,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Computed: true,
			},
			"error": {
				Description: "List of errors associated with this VmBackupInfo.",
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
						"message": {
							Description: "The error message string for this error stack.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"message_id": {
							Description: "The error message ID for this error stack.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"protection_status": {
				Description: "Description of the protection status of this VmBackupInfo.\n* `PREPARE_FAILOVER_STARTED` - The protection status is prepare failover started.\n* `PREPARE_FAILOVER_FAILED` - The protection status is prepare failover failed.\n* `PREPARE_FAILOVER_COMPLETED` - The protection status is prepaire failover completed.\n* `FAILOVER_STARTED` - The protection status is failover started.\n* `FAILOVER_FAILED` - The protection status is failover failed.\n* `FAILOVER_COMPLETED` - The protection status is failover completed.\n* `PREPARE_REVERSEPROTECT_STARTED` - The protection status is prepare reverse protect started.\n* `PREPARE_REVERSEPROTECT_FAILED` - The protection status is prepare reverse protect failed.\n* `PREPARE_REVERSEPROTECT_COMPLETED` - The protection status is prepair reverse protect completed.\n* `REVERSEPROTECT_STARTED` - The protection status is reverse protect started.\n* `REVERSEPROTECT_FAILED` - The protection status is reverse protect failed.\n* `ACTIVE` - The protection status is active.\n* `CREATION_IN_PROGRESS` - The protection status is failover in progress.\n* `CREATION_FAILED` - The protection status is creation failed.\n* `LOCAL_RESTORE_STARTED` - The protection status is local restore started.\n* `LOCAL_RESTORE_FAILED` - The protection status is local restore failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Optional: true,
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
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"schedule": {
							Description: "Replication schedule information.",
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
									"backup_interval": {
										Description: "Time interval between two copies in minutes.",
										Type:        schema.TypeInt,
										Optional:    true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"target_cluster_entity_reference": {
							Description: "EntityReference of target cluster.",
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
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"confignum": {
										Description: "Configuration number for this reference.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Description: "Uuid of the entity for this reference.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"idtype": {
										Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the entity for this entity reference.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"type": {
										Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
				Computed: true,
			},
			"src_backup_cluster": {
				Description: "A reference to a hyperflexBackupCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"src_cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"tgt_cluster": {
				Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
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
			"vm_entity_reference": {
				Description: "Reference to the virtual machine this VmBackupInfo is associated with.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"confignum": {
							Description: "Configuration number for this reference.",
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Description: "Uuid of the entity for this reference.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"idtype": {
							Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Description: "Name of the entity for this entity reference.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"vm_info": {
				Description: "More detailed information about this VmBackupInfo including runtime info.",
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
							Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"run_time_info": {
							Description: "Virtual machine runtime details.",
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
									"bios_uuid": {
										Description: "BiosUuid of the Protected Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"class_id": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.\nThe enum values provides the list of concrete types that can be instantiated from this abstract type.",
										Type:        schema.TypeString,
										Optional:    true,
									},
									"connection_state": {
										Description: "Connection state of the VM.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"cpu_usage": {
										Description: "CPU Usage of Virtual Machine.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"folder": {
										Description: "Folder which VM belongs to.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"guest_family": {
										Description: "Guest operating system family, if known.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"guest_full_name": {
										Description: "Guest operating system full name, if known.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"guest_id": {
										Description: "Guest operating system identifier (short name), if known.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"guest_state": {
										Description: "VMware guest reset, powercycle, or boot action states.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"host_name": {
										Description: "Hostname of Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"instance_uuid": {
										Description: "InstanceUuid of the Protected Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"memory_mb": {
										Description: "CPU Memory in MB of Virtual Machine.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"memory_usage": {
										Description: "Memory usage of Virtual Machine.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"moid": {
										Description: "Virtual Machine unique MOID.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Description: "Name of the Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"networks": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString}},
									"num_cpu": {
										Description: "Number of CPUs for the VM.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"object_type": {
										Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"power_state": {
										Description: "Power state of the Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"provisioned_size": {
										Description: "Provisioned Size of Virtual Machine.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"resource_pool": {
										Description: "Resource pool which VM belongs to.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"used_size": {
										Description: "Used Size of Virtual Machine.",
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
									},
									"nr_version": {
										Description: "Version of the Virtual Machine.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
									"vmx_path": {
										Description: "Vmx Path in VC datastore format.",
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"status_code": {
							Description: "Virtual Machine Status Code.\n* `VM_ACCESSIBLE` - This virtual machine is accessible.\n* `VM_INACCESSIBLE` - This virtual machine is not accessible.\n* `VM_NOT_SUPPORTED` - This virtual machine is not supported.\n* `NONE` - This virtual machine does not have a status code.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"uuid": {
							Description: "Virtual machine unique UUID.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHyperflexVmBackupInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexVmBackupInfo{}
	if v, ok := d.GetOk("backup_status"); ok {
		x := (v.(string))
		o.SetBackupStatus(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("protection_status"); ok {
		x := (v.(string))
		o.SetProtectionStatus(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexVmBackupInfo object failed with error : %s", err.Error())
	}
	resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmBackupInfoList(conn.ctx).Filter(getRequestParams(data)).Execute()
	if responseErr.Error() != "" {
		return diag.Errorf("error occurred while fetching HyperflexVmBackupInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}

	x, err := resMo.MarshalJSON()
	if err != nil {
		return diag.Errorf("error occurred while marshalling response for HyperflexVmBackupInfo list: %s", err.Error())
	}
	var s = &models.HyperflexVmBackupInfoList{}
	err = json.Unmarshal(x, s)
	if err != nil {
		return diag.Errorf("error occurred while unmarshalling response to HyperflexVmBackupInfo list: %s", err.Error())
	}
	result := s.GetResults()
	length := len(result)
	if length == 0 {
		return diag.Errorf("your query for HyperflexVmBackupInfo data source did not return results. Please change your search criteria and try again")
	}
	if length > 1 {
		return diag.Errorf("your query for HyperflexVmBackupInfo data source returned more than one result. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s = &models.HyperflexVmBackupInfo{}
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = json.Unmarshal(oo, s); err != nil {
				return diag.Errorf("error occurred while unmarshalling result at index %+v: %s", i, err.Error())
			}
			if err := d.Set("additional_properties", flattenAdditionalProperties(s.AdditionalProperties)); err != nil {
				return diag.Errorf("error occurred while setting property AdditionalProperties: %s", err.Error())
			}
			if err := d.Set("backup_status", (s.GetBackupStatus())); err != nil {
				return diag.Errorf("error occurred while setting property BackupStatus: %s", err.Error())
			}
			if err := d.Set("class_id", (s.GetClassId())); err != nil {
				return diag.Errorf("error occurred while setting property ClassId: %s", err.Error())
			}

			if err := d.Set("cluster_entity_reference", flattenMapHyperflexEntityReference(s.GetClusterEntityReference(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ClusterEntityReference: %s", err.Error())
			}

			if err := d.Set("cluster_id_protection_info_map", flattenListHyperflexMapClusterIdToProtectionInfo(s.GetClusterIdProtectionInfoMap(), d)); err != nil {
				return diag.Errorf("error occurred while setting property ClusterIdProtectionInfoMap: %s", err.Error())
			}

			if err := d.Set("error", flattenMapHyperflexErrorStack(s.GetError(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Error: %s", err.Error())
			}
			if err := d.Set("moid", (s.GetMoid())); err != nil {
				return diag.Errorf("error occurred while setting property Moid: %s", err.Error())
			}
			if err := d.Set("object_type", (s.GetObjectType())); err != nil {
				return diag.Errorf("error occurred while setting property ObjectType: %s", err.Error())
			}
			if err := d.Set("protection_status", (s.GetProtectionStatus())); err != nil {
				return diag.Errorf("error occurred while setting property ProtectionStatus: %s", err.Error())
			}

			if err := d.Set("schedule", flattenListHyperflexReplicationClusterReferenceToSchedule(s.GetSchedule(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Schedule: %s", err.Error())
			}

			if err := d.Set("src_backup_cluster", flattenMapHyperflexBackupClusterRelationship(s.GetSrcBackupCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SrcBackupCluster: %s", err.Error())
			}

			if err := d.Set("src_cluster", flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property SrcCluster: %s", err.Error())
			}

			if err := d.Set("tags", flattenListMoTag(s.GetTags(), d)); err != nil {
				return diag.Errorf("error occurred while setting property Tags: %s", err.Error())
			}

			if err := d.Set("tgt_cluster", flattenMapHyperflexClusterRelationship(s.GetTgtCluster(), d)); err != nil {
				return diag.Errorf("error occurred while setting property TgtCluster: %s", err.Error())
			}

			if err := d.Set("vm_entity_reference", flattenMapHyperflexEntityReference(s.GetVmEntityReference(), d)); err != nil {
				return diag.Errorf("error occurred while setting property VmEntityReference: %s", err.Error())
			}

			if err := d.Set("vm_info", flattenMapHyperflexVirtualMachine(s.GetVmInfo(), d)); err != nil {
				return diag.Errorf("error occurred while setting property VmInfo: %s", err.Error())
			}
			d.SetId(s.GetMoid())
		}
	}
	return de
}
