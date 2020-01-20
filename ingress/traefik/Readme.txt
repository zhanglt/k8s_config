# Setup domain, user and password first.
$ export USER=admin
$ export DOMAIN=traefik.ingress.com
$ htpasswd -c auth $USER
New password:
Re-type new password:
Adding password for user user
$ PASSWORD=$(cat auth| awk -F: '{print $2}')

# Deploy with helm.
helm install stable/traefik --name traefik-ingress  --namespace kube-ingress  --set rbac.enabled=true,acme.enabled=true,dashboard.enabled=true,acme.staging=false,acme.email=admin@$DOMAIN,dashboard.domain=ui.$DOMAIN,ssl.enabled=true,acme.challengeType=http-01,acme.persistence.storageClass=fast-rbd,dashboard.auth.basic.$USER=$PASSWORD

注意：storageClass 要提前部署
Get Traefik's load balancer IP/hostname:

     NOTE: It may take a few minutes for this to become available.

     You can watch the status by running:

         $ kubectl get svc traefik-ingress --namespace kube-ingress -w

     Once 'EXTERNAL-IP' is no longer '<pending>':

         $ kubectl describe svc traefik-ingress --namespace kube-ingress | grep Ingress | awk '{print $3}'
