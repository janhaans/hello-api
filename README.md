# Hello API
 
## Release Milestones
 
### V0 (1 day)
- [ ] Onboarding Documentation
- [ ] Simple API response (hello world!)
- [ ] Unit tests
- [ ] Running somewhere other than the dev machine
 
### V1 (7 days)
- [ ] Create translation endpoint
- [ ] Store translations in short-term storage
- [ ] Call existing service for translation
- [ ] Move towards long-term storage

## Dependencies
 
- Go version 1.21
 
## Setup
 
## Release Milestones
..

# Chapter 4 - Deploy hello-api to GCP Cloud Functions

See `deploy-function` job in Github workflow. This job automatically deploys the app on GCP Cloud Functions

Prerequisites:
- Create Service Account `hello-api` that will be used by GitHub to deploy the app. This Service Account should have Roles:
```
App Engine Admin
App Engine Deployer
Cloud Build Editor
Cloud Functions Admin
Cloud Functions Developer
Storage Admin
```
- Create keys for Service Account `hello-api` (credentials JSON file) and save the contents of the credentials JSON file in a GitHub secret with name GCP_CREDENTIALS
- [Enable Access to API's](https://console.cloud.google.com/flows/enableapi?apiid=cloudfunctions,cloudbuild.googleapis.com,artifactregistry.googleapis.com,run.googleapis.com,logging.googleapis.com&redirect=https://cloud.google.com/functions/docs/create-deploy-gcloud&_ga=2.126917573.955066507.1704387230-1676995943.1703927418). If you enabled the API's recently, wait a few minutes for the action to propagate to our systems. The following API's are enabled:
```
Cloud Functions API
Cloud Build API
Artifact Registry API
Cloud Run Admin API
Cloud Logging API
```
- Grant the role `roles/iam.serviceAccountUser` to Service Account`hello-api` on Service Account `AppEngine Default Service Account`. Open Cloud Shell from GCP Console:
```
gcloud iam service-accounts add-iam-policy-binding hello-api-409709@appspot.gserviceaccount.com --member serviceAccount:hello-api@hello-api-409709.iam.gserviceaccount.com --role roles/iam.serviceAccountUser

Updated IAM policy for serviceAccount [hello-api-409709@appspot.gserviceaccount.com].
bindings:
- members:
  - serviceAccount:hello-api@hello-api-409709.iam.gserviceaccount.com
  role: roles/iam.serviceAccountUser
etag: BwYOJmkvrQM=
version: 1
``` 
- Make the Cloud Function public by granting "allUsers" the role "Cloud Function Invoker"

See `deploy-paas` job in Github workflow. This job automatically deploys the app on GCP App Engine

- Enable `App Engine Admin API` by visiting https://console.developers.google.com/apis/api/appengine.googleapis.com/overview?project=875030482222
