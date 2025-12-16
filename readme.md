# Mongo AKS OIDC

Using Mongo with OIDC authentication, we are able to connect to the database without the need for a username or password.
This project demonstrates how to use Azure Kubernetes Service (AKS) with OpenID Connect (OIDC) for authentication with MongoDB. By leveraging OIDC, we can securely authenticate with MongoDB without exposing sensitive credentials.

You need to have your mongo database configured for OIDC Workload authentication.
https://www.mongodb.com/docs/drivers/csharp/current/security/authentication/oidc/

Second, you need to configure your AKS cluster to use Azure workload identity.
https://learn.microsoft.com/en-us/azure/aks/workload-identity-overview?tabs=go

Once this is done, you can create a workload identity in azure.
https://learn.microsoft.com/en-us/azure/aks/workload-identity-deploy-cluster?tabs=new-cluster#create-a-managed-identity

and using the details for the workload identity, create a service account in AKS
https://learn.microsoft.com/en-us/azure/aks/workload-identity-deploy-cluster?tabs=new-cluster#create-a-kubernetes-service-account

Your workload (pod, deployment, ect... ) needs to use the service account, dont forget to set the labels:
https://learn.microsoft.com/en-us/azure/aks/workload-identity-deploy-cluster?tabs=new-cluster#deploy-a-verification-pod-and-test-access

In mongo, Ensure that the new service account has the correct access as required. 

See the aksCallback_example.go for an example of how to use the workload identity and go.

## References
- [Azure Kubernetes Service (AKS)](https://docs.microsoft.com/en-us/azure/aks/)
- [OpenID Connect (OIDC)](https://openid.net/connect/)
- [MongoDB OIDC Authentication](https://www.mongodb.com/docs/manual/core/security-oidc/)
- [Azure Identity](https://github.com/Azure/azure-sdk-for-go/tree/main/sdk/azidentity)

