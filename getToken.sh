kubectl describe secret/$(kubectl get secret -nkube-system |grep admin|awk '{print $1}') -nkube-system   | grep '^token' |awk '{print $2}'


