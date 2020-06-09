PROJECT := dictav-net
APP := test-auth
GATEWAY := cors-gateway
IMAGE := gcr.io/$(PROJECT)/$(APP)
GATEWAY_IMAGE := gcr.io/$(PROJECT)/$(GATEWAY)
CMD := gcloud --project=$(PROJECT)
BACKEND := test-auth-iljqu5cu6q-an.a.run.app

TEST_URL := https://cors-gateway-iljqu5cu6q-an.a.run.app

deploy:
	$(CMD) builds submit --tag $(IMAGE) .
	$(CMD) run deploy $(APP) \
	  --platform=managed \
	  --image $(IMAGE) \
	  --region=asia-northeast1 \
	  --no-allow-unauthenticated

.PHONY: gateway
gateway:
	$(CMD) builds submit --tag $(GATEWAY_IMAGE) gateway
	$(CMD) run deploy $(GATEWAY) \
	  --platform managed \
	  --image="$(GATEWAY_IMAGE)" \
	  --region=asia-northeast1 \
	  --allow-unauthenticated \
	  --set-env-vars=BACKEND=$(BACKEND)

curl:
	curl -vH "Authorization: Bearer $$(gcloud auth print-identity-token)" $(TEST_URL)

curl_backend:
	curl -vH "Authorization: Bearer $$(gcloud auth print-identity-token)" https://$(BACKEND)
