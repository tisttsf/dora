#
# Makefile
# @author Lempiy Anton <lempiyada@gmail.com>
#

GOVERSION = 1.11

# Parser def
PARSER_NAME = dora_parser
PARSER_VER = v0.0.1

# Bot def
BOT_NAME = dora_bot
BOT_VER = v0.0.1

# API def
API_NAME = dora_api
API_VER = v0.0.1

MINIKUBE_STOPPED = $(shell minikube status | grep -o Stopped)


ifeq ($(SERVICE),parser)
	name=$(PARSER_NAME)
endif
ifeq ($(SERVICE),bot)
	name=$(BOT_NAME)
endif
ifeq ($(SERVICE),api)
	name=$(API_NAME)
endif

service:
ifndef name
	@echo 'Please provide SERVICE=name (posible variants: parser, bot, api)'
	@exit 1
endif
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif

kube:
ifeq ($(MINIKUBE_STOPPED), Stopped)
	@echo minikube is down. Running minikube ...
	@minikube start
endif

generate:
	@protoc -I . ./pb/bot.proto --go_out=plugins=grpc:.
	@protoc -I . ./pb/prs.proto --go_out=plugins=grpc:.

.PHONY: target

target:
ifndef name
	@echo 'Please provide SERVICE=name (possible variants: parser, bot, api)'
	@exit 1
endif
ifndef VERSION
	@echo 'Please provide VERSION=version (in ex. VERSION=v0.4.2)'
	@exit 1
endif
	@docker rmi -f lempiy/$(name):$(VERSION) || true
	services/$(SERVICE)/build.sh $(VERSION)
	@exit 0

.PHONY: dev

dev: service
	@eval $(shell minikube docker-env); \
	docker rmi -f dev-dora-$(name):latest || true; \
	services/$(SERVICE)/dev.sh
	@kubectl delete pod $(shell kubectl get pods | grep -o "^$(SERVICE)-[a-z0-9]*-[a-z0-9]*-[a-z0-9]*") || true
	@exit 0

refresh: service
	@kubectl delete pod $(shell kubectl get pods | grep -o "^$(SERVICE)-[a-z0-9]*-[a-z0-9]*-[a-z0-9]*")
	@exit 0

logs:
	@kubectl logs deployment/$(SERVICE)-deployment

k8s-show-url: kube
	@minikube service dora-api --url

k8s-clear-dev: kube
	@kubectl delete -f k8s/dora-dev.yaml

k8s-create-dev: kube
	@kubectl create -f k8s/dora-dev.yaml
	
k8s-show-pods:
	@kubectl get pods

k8s-dev-down:
	@minikube stop
