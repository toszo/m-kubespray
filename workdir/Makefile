include $(M_RESOURCES)/consts.mk
include $(M_RESOURCES)/defaults.mk

export

#medatada method is printing static metadata information about module
metadata: guard-M_RESOURCES
	#k8sprom | metadata | should print component metadata
	@echo "$$M_METADATA_CONTENT"

#init method is used to initialize module configuration and check if state is providing strong (and weak) dependencies
#TODO should also validate state if strong requirements are met
init: guard-M_RESOURCES guard-M_SHARED setup ensure-state-file init_vars init_inventory initialize-state-file 

#plan method would get config file and environment state file and compare them and calculate what would be done o apply stage

#plan: guard-M_RESOURCES guard-M_SHARED setup validate-config validate-state template-tfvars module-plan terraform-plan

#apply method runs module provider logic using config file

apply: guard-M_RESOURCES guard-M_SHARED copy_vars copy_inventory print_inventory_location
#apply: guard-M_RESOURCES guard-M_SHARED setup module-plan terraform-apply update-state-after-apply terraform-output

#audit method should call logic to check if remote components are in "known" state
#TODO implement validation if remote resources are as expected, possibly with terraform plan

#audit:
	#kubespray | audit | should output current state of remote components

#destroy: template-tfvars terraform-destroy update-state-after-destroy

#plan-destroy: template-tfvars terraform-plan-destroy

#output: terraform-output

setup: $(M_SHARED)/$(M_MODULE_SHORT)
	#kubespray | setup | ensure required directories

ensure-state-file: $(M_SHARED)/$(M_STATE_FILE_NAME)
	#kubespray | ensure-state-file | checks if state file exists

init_vars:
	#kubespray | copy_vars | will copy vars folder from resources
	@cp -rf $(M_RESOURCES)/group_vars $(M_SHARED)/$(M_MODULE_SHORT)/

init_inventory:
	#kubespray | copy_vars | will copy vars folder from resources
	@cp -rf $(M_RESOURCES)/hosts.yaml $(M_SHARED)/$(M_MODULE_SHORT)/

copy_vars:
	#kubespray | copy_vars | will copy vars folder from resources
	@cp -rfp $(M_SHARED)/$(M_MODULE_SHORT)/group_vars $(M_WORKDIR)/kubespray/inventory/$(M_NAME)/

copy_inventory:
	#kubespray | copy_vars | will copy vars folder from resources
	@cp -rfp $(M_SHARED)/$(M_MODULE_SHORT)/hosts.yaml $(M_WORKDIR)/kubespray/inventory/$(M_NAME)/

print_inventory_location:
	#kubespray | print_inventory | will print inventory folder
	@ls $(M_WORKDIR)/kubespray/inventory/$(M_NAME)

initialize-state-file:
	#kubespray | initialize-state-file | will initialize state file
	@echo "$$M_STATE_INITIAL" > $(M_SHARED)/$(M_MODULE_SHORT)/epikubespray.tmp
	@yq m -i -x $(M_SHARED)/$(M_STATE_FILE_NAME) $(M_SHARED)/$(M_MODULE_SHORT)/epikubespray.tmp
	@rm $(M_SHARED)/$(M_MODULE_SHORT)/epikubespray.tmp

guard-%:
	@if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

$(M_SHARED)/$(M_STATE_FILE_NAME):
	@mkdir -p $(dir $@)
	@touch $@

$(M_SHARED)/$(M_MODULE_SHORT):
	@mkdir -p $@