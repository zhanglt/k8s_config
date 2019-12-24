kubectl describe secret/$(kubectl get secret -n kube-system |grep ^dashboard-admin|awk '{print $1}') -n kube-system   | grep '^token' |awk '{print $2}'


