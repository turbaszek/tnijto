![CI](https://github.com/turbaszek/tnijto/workflows/CI/badge.svg?branch=master)

# TnijTo

Easy to deploy link shortener.


<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Development](#development)
- [Deployment](#deployment)
- [Contributing](#contributing)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Development

To run the app just do:
```shell
PROJECT_ID="your-project-id" make build
```

## Deployment

First authenticate using:
```shell
gcloud auth login
```

To deploy TnijTo on Google Cloud you have to enable [Cloud Build](https://cloud.google.com/cloud-build)
and [Cloud Run](https://cloud.google.com/run) services:
```shell
make setup
```
then go to https://console.cloud.google.com/firestore and enable native Firestore mode.

Deploy the image using Cloud Build and deploy using Cloud Run:
```shell
PROJECT_ID="your-project-id" REGION="europe-west1" make deploy
```

<!--
If you wish to limit access to authenticated user run
```shell
gcloud run services add-iam-policy-binding $SERVICE \
  --member="allAuthenticatedUsers" \
  --role="roles/run.invoker" \
  --platform managed \
  --region $REGION
```

To limit access to people from single domain run
```shell
export DOMAIN="google.com"
gcloud run services add-iam-policy-binding $SERVICE \
  --member="domain:${DOMAIN} \
  --role="roles/run.invoker" \
  --platform managed \
  --region $REGION
```
-->

## Contributing

We welcome all contributions! Please submit an issue or PR no matter if it's a bug or a typo.

This project is using [pre-commits](https://pre-commit.com) to ensure the
quality of the code. To install pre-commits just do:
```bash
pip install pre-commit
# or
brew install pre-commit
```
Then from project directory run `pre-commit install`.
