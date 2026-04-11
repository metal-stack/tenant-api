.PHONY: generate
generate: go-mocks
	# we remove files explicitly to ensure files really get generated
	rm -f ../go/permissions/servicepermissions.go
	rm -f ../go/client/client.go
	rm -f ../go/tests/mock_clients.go

	go run ./generate.go

	go run ./genscanvaluer.go -package apiv1 -type Project -output ../go/api/v1/project_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type ProjectMember -output ../go/api/v1/project_member_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type Tenant -output ../go/api/v1/tenant_scanvaluer.go
	go run ./genscanvaluer.go -package apiv1 -type TenantMember -output ../go/api/v1/tenant_member_scanvaluer.go

.PHONY: go-mocks
go-mocks:
	# we remove files explicitly to ensure files really get generated
	rm -rf ../go/tests/mocks

	docker run --rm \
		--user $$(id -u):$$(id -g) \
		--tmpfs /.cache:uid=$$(id -u),gid=$$(id -g) \
		-w /work \
		-v ${PWD}:/work \
		vektra/mockery:v3.7.0
