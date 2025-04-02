swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

# Do this after cd-ing into client
swagger-client:
	swagger generate client -f ../swagger.yaml -A product-api