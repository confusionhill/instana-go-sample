services-up:
	@echo "SVCNAME=${SVCNAME}"
	docker-compose -f docker-compose.yaml up --build ${SVCNAME}
rest-run:
	go run main.go