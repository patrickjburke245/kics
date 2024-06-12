package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	doc := input.document[i]
	resource := doc.resource.tencentcloud_instance[name]

	vpcName := split(resource.vpc_id, ".")[1]
	vpc := input.document[_].resource.tencentcloud_instance[vpcName]

	contains(lower(vpc.vpc_id), "default")

	result := {
		"documentId": doc.id,
		"resourceType": "tencentcloud_instance",
		"resourceName": tf_lib.get_resource_name(resource, name),
		"searchKey": sprintf("tencentcloud_instance[%s].vpc_id", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("tencentcloud_instance[%s].vpc_id should not be associated with a default VPC", [name]),
		"keyActualValue": sprintf("tencentcloud_instance[%s].vpc_id is associated with a default VPC", [name]),
		"searchLine": common_lib.build_search_line(["resource", "tencentcloud_instance", name, "vpc_id"], []),
	}
}

CxPolicy[result] {
	doc := input.document[i]
	resource := doc.resource.tencentcloud_instance[name]

	subnetName := split(resource.subnet_id, ".")[1]
	subnet := input.document[_].resource.tencentcloud_instance[sbName]

	contains(lower(subnet.subnet_id), "default")

	result := {
		"documentId": doc.id,
        "resourceType": "tencentcloud_instance",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("tencentcloud_instance[%s].subnet_id", [name]),
        "issueType": "IncorrectValue",
        "keyExpectedValue": sprintf("tencentcloud_instance[%s].subnet_id should not be associated with a default Subnet", [name]),
        "keyActualValue": sprintf("tencentcloud_instance[%s].subnet_id is associated with a default Subnet", [name]),
        "searchLine": common_lib.build_search_line(["resource", "tencentcloud_instance", name, "subnet_id"], []),
	}
}
