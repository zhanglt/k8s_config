kubectl describe secret/$(kubectl get secret -n kube-system |grep ^admin-user-token|awk '{print $1}') -n kube-system   | grep '^token' |awk '{print $2}'


