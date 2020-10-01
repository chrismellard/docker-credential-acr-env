# ACR Docker Credential Helper

The ACR docker credential helper is an alternative to the existing file store based ACR helper 
located [here](https://github.com/Azure/acr-docker-credential-helper) which relies on `az` command
line and is not optimised for use in CI environments. Primary use case for this helper is for use
with kaniko and other tools running in CI scenarios wishing to push to Azure Container Registry

