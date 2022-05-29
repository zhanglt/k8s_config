[root@master gitlab]# helm upgrade --install gitlab ./gitlab --timeout 600s  --set q1global.hosts.domain=gitlab.code.local  --set certmanager-issuer.email=kitsdk@163.com
Release "gitlab" does not exist. Installing it now.
NAME: gitlab
LAST DEPLOYED: Tue Apr  5 20:00:48 2022
NAMESPACE: postgres
STATUS: deployed
REVISION: 1
NOTES:
NOTICE: The minimum required version of PostgreSQL is now 12. See https://gitlab.com/gitlab-org/charts/gitlab/-/blob/master/doc/installation/upgrade.md for more details.

WARNING: Automatic TLS certificate generation with cert-manager is disabled and no TLS certificates were provided. Self-signed certificates were generated.

You may retrieve the CA root for these certificates from the `gitlab-wildcard-tls-ca` secret, via the following command. It can then be imported to a web browser or system store.

    kubectl get secret gitlab-wildcard-tls-ca -ojsonpath='{.data.cfssl_ca}' | base64 --decode > gitlab.gitlab.code.local.ca.pem

If you do not wish to use self-signed certificates, please set the following properties:
  - global.ingress.tls.secretName
  OR
  - global.ingress.tls.enabled (set to `true`)
  - gitlab.webservice.ingress.tls.secretName
  - registry.ingress.tls.secretName
  - minio.ingress.tls.secretName
WARNING: Automatic TLS certificate generation with cert-manager is disabled and no TLS certificates were provided. Self-signed certificates were generated that do not work with gitlab-runner. Please either disable gitlab-runner by setting `gitlab-runner.install=false` or provide valid certificates.

NOTICE: You've installed GitLab Runner without the ability to use 'docker in docker'.
The GitLab Runner chart (gitlab/gitlab-runner) is deployed without the `privileged` flag by default for security purposes. This can be changed by setting `gitlab-runner.runners.privileged` to `true`. Before doing so, please read the GitLab Runner chart's documentation on why we
chose not to enable this by default. See https://docs.gitlab.com/runner/install/kubernetes.html#running-docker-in-docker-containers-with-gitlab-runners
Help us improve the installation experience, let us know how we did with a 1 minute survey:
https://gitlab.fra1.qualtrics.com/jfe/form/SV_6kVqZANThUQ1bZb?installation=helm&release=14-9

NOTICE: The in-chart NGINX Ingress Controller has the following requirements:
    - Kubernetes version must be 1.19 or newer.
    - Ingress objects must be in group/version `networking.k8s.io/v1`.
[root@master gitlab]#
[root@master gitlab]#

