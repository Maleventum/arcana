APP := arcana

verify:
	golint
	go test ./...

docker: verify
	$(eval branch = $(shell git branch --show-current) )
	@echo current branch: $(branch)
	docker build . -t $(APP):$(branch)
	#docker image tag $(APP) $(APP):$(branch)

openAPI:
	swagger generate spec -o ./swagger.json


#swaggerAPI:
	#docker run -p 80:8080 -e BASE_URL=/arcana -e SWAGGER_JSON=./swagger.json swaggerapi/swagger-ui
	#docker run -p 80:8080 -e BASE_URL=/ -e SWAGGER_JSON=swagger.json swaggerapi/swagger-ui
