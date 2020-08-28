SERVICE ?= tnijto

PROJECT_ID ?= "your-project-id"
REGION ?= "europe-west1"
SERVICE ?= ?"tnijto"
TAG ?= "gcr.io/${PROJECT_ID}/tnijto"

all: setup deploy

deploy: publish run

# Setup
setup:
	gcloud services enable cloudbuild.googleapis.com
	gcloud services enable run.googleapis.com

# Build on Cloud Build
publish:
	gcloud builds submit --tag ${TAG}

# Deploy on Cloud Run
run:
	gcloud run deploy ${SERVICE} \
	  --image ${TAG} \
	  --region ${REGION} \
	  --set-env-vars "PROJECT_ID=${PROJECT_ID}" \
	  --platform managed \
	  --allow-unauthenticated

# Build locally
build:
	go build -v ./pkg/tnijto.go

# Build local docker image
docker:
	docker build -t tnijto .
