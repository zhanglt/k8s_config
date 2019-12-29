NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace kube-monitor grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   grafana.kube-monitor.svc.cluster.local

   Get the Grafana URL to visit by running these commands in the same shell:

     export POD_NAME=$(kubectl get pods --namespace kube-monitor -l "app=grafana,release=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace kube-monitor port-forward $POD_NAME 3000

3. Login with the password from step 1 and the username: admin

