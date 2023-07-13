# Gin Gonic on Kubernetes

## Development

### Docker

1. Install Docker
2. Spin up the containers
   - `docker compose $REPO_PATH`
3. Reach the API on `http://localhost:8080`

### Kubernetes

#### Minikube

1. Install Minikube
2. Start the cluster
   - `minikube start`
3. Install Ingress controller
   - `minikube addons enable ingress`
4. Install ArgoCD controller
   - `kubectl create namespace argocd`
   - `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`
5. Install Sealed Secrets controller
   - `kubectl apply -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.22.0/controller.yaml`
6. Encrypt Secrets
   - `kubeseal --scope cluster-wide -f $REPO_PATH/infra/charts/database/templates/secret.yaml.bak -o yaml > $REPO_PATH/infra/charts/database/templates/sealed-secret.yaml`
   - `kubeseal --scope cluster-wide -f $REPO_PATH/infra/charts/backend/templates/secret.yaml.bak -o yaml > $REPO_PATH/infra/charts/backend/templates/sealed-secret.yaml`
7. Create resources
   - With Helm
     - `kubectl create namespace namaste-rodox`
     - `helm install $HELM_PACKAGE_NAME $REPO_PATH/infra`
   - With ArgoCD
     - `k apply -f $REPO_PATH/application.yaml`
8. Reach the API on `http://api.mayflower.de`

#### Haufen

1. Setup Kubernetes configuration
   - Copy the following code to `~/.kube/config`
     ```yaml
     apiVersion: v1
     clusters:
     - cluster:
         certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURkRENDQWx5Z0F3SUJBZ0lVSGhCcERZTzFjYmFLclJLRjFDcEp5c3MwcXM0d0RRWUpLb1pJaHZjTkFRRUwKQlFBd1VqRVhNQlVHQTFVRUJ4TU9ZWFYwYnkxblpXNWxjbUYwWldReERqQU1CZ05WQkFvVEJVNXBlRTlUTVNjdwpKUVlEVlFRTEV4NXpaWEoyYVdObGN5NXJkV0psY201bGRHVnpMbkJyYVM1allWTndaV013SGhjTk1qSXdPREF4Ck1UWTBOekF3V2hjTk1qY3dOek14TVRZME56QXdXakJTTVJjd0ZRWURWUVFIRXc1aGRYUnZMV2RsYm1WeVlYUmwKWkRFT01Bd0dBMVVFQ2hNRlRtbDRUMU14SnpBbEJnTlZCQXNUSG5ObGNuWnBZMlZ6TG10MVltVnlibVYwWlhNdQpjR3RwTG1OaFUzQmxZekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNLzlnQ3FsCkRkRURhRmFzUWJBdmNUa2ZaY294cFZuZ2RzNkZWNHRGa0ZPOStPYjd1NFpob245eWNlUmxzZmlwakszNUUzZXkKZXYzOFc1YnBRbkdoYjIraUdBUkdyclQyUWpCR2FsN2NBYTdRU2pwaG9LZTliY1VxbXQzZEdvc1hmNkphU2ZyNQp3NmVrV0FRcG4rWEtTMmMrODF1ODI2aUNrODhVSE5ZYjA5cWJFRitySXpIT0R2eXhZaDMwRVJvTityZ3hGTnN2CjJza2ZxNTVyKzRIaHYyYWY4OVF2VDBUajVaT05HNlFZZUxTMVY2MHFMTE1vZDcrUEd2WEMrdk1hL0s3cmZHYkMKRHBkcEJFenV0c3JZdTJpdXM5VnYxWjQ3ZzhNR3huZGRHZ3VoZkFub1dxRkk0OG1SL3puOWROMDFpbFJZWUJtVApSNjhmRTJ3SmNiU2kzMWtDQXdFQUFhTkNNRUF3RGdZRFZSMFBBUUgvQkFRREFnRUdNQThHQTFVZEV3RUIvd1FGCk1BTUJBZjh3SFFZRFZSME9CQllFRkxuUVRZRmhpY3JEVERuejVPbmdVb2hqYXhIdU1BMEdDU3FHU0liM0RRRUIKQ3dVQUE0SUJBUUJMUGhGYTJXWm0wM0lnYkR2OXJ6SFFTOUIwdmVhQ1BsM3RFRTR6T04zaVBrNDRQSm16eFliLwp0amEyNmNqS0pjeUV2d0ZkQzhWdGlWQlYyOGlRNS9rK3RGMWpSdFU1Y096VVRTazhKcDU4OW9MQm9ybmh1emVpCkhKRTBSTFpkWkh4M2dpR1ZHMnhIMWNUWXY0YjNjMERrWHhoN1JSeUczSGFIL2hOU1hXcDh4SkR0YnBibHlLYmQKbmlCbFI2MklMbXlGakZhTEFXR1hkdjA5MWxIcjhBY2lPWkd6dkNDT2QzdHlqaHFaMG5TL2tRZW9mYS9LcjBTeApmMnFscVdacVNYSmp3b3J5cUdVM0c3bkIwUjFmMldWQ21pd09qNDFqSFRZRmNpMFl4Yi9HMERVeVRZemlRU1FkClFvRC9MREpwWTgzRXJOTk1YZWp3eGpDcHo0MkhvbVc0Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
         server: https://master.mayflower.cloud:6443
       name: mf
     contexts:
     - context:
         cluster: mf
         user: oidc
       name: mf
     current-context: mf
     kind: Config
     preferences: {}
     users:
     - name: oidc
       user:
         exec:
           apiVersion: client.authentication.k8s.io/v1beta1
           args:
           - oidc-login
           - get-token
           - --oidc-issuer-url=https://auth.mayflower.de/application/o/haufen/
           - --oidc-client-id=1qIzXCQ33PgqpCM91ToR7EcDhsUFFCr7mixiwtCK
           - --oidc-extra-scope="email profile openid"
           command: kubectl
           env: null
           interactiveMode: IfAvailable
           provideClusterInfo: false
     ```
2. Start the cluster
   - `minikube start`
3. Install Ingress controller
   - `minikube addons enable ingress`
4. Install ArgoCD controller
   - `kubectl create namespace argocd`
   - `kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml`
5. Install Sealed Secrets controller
   - `kubectl apply -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.22.0/controller.yaml`
6. Encrypt Secrets
   - `kubeseal --scope cluster-wide -f $REPO_PATH/infra/charts/database/templates/secret.yaml.bak -o yaml > $REPO_PATH/infra/charts/database/templates/sealed-secret.yaml`
   - `kubeseal --scope cluster-wide -f $REPO_PATH/infra/charts/backend/templates/secret.yaml.bak -o yaml > $REPO_PATH/infra/charts/backend/templates/sealed-secret.yaml`
7. Create resources
   - With Helm
     - `kubectl create namespace namaste-rodox`
     - `helm install $HELM_PACKAGE_NAME $REPO_PATH/infra`
   - With ArgoCD
     - `k apply -f $REPO_PATH/application.yaml`
8. Reach the API on `http://api.mayflower.de`

## Production

2. Reach the API on `https://api.rodox.mayflower.cloud`
