kubeadm  init --apiserver-advertise-address=10.0.0.51  --pod-network-cidr=10.244.0.0/16  --kubernetes-version=v1.15.6

1、kubeadm init --config kubeadm-config.yaml
2、mkdir -p $HOME/.kube
   cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   chown $(id -u):$(id -g) $HOME/.kube/config
3 、 kubectl apply -f https://docs.projectcalico.org/v3.3/getting-started/kubernetes/installation/hosted/rbac-kdd.yaml
     kubectl apply -f https://docs.projectcalico.org/v3.3/getting-started/kubernetes/installation/hosted/kubernetes-datastore/calico-networking/1.7/calico.yaml




 kubeadm join 172.18.0.2:6443 --token 7serau.jblqirx4136w2sl6 --discovery-token-ca-cert-hash sha256:18ec60489a6dcd65486ab3a90767527de7f8107e32edb5a12b43bce362693995   --experimental-control-plane


>1.在master上执行命令：kubeadm token create
>2.获取ca证书`sha256`编码hash值:openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null |   openssl dgst -sha256 -hex | sed 's/^.* //'




kubectl create -f kubernetes-dashboard.yaml
kubectl create serviceaccount dashboard-admin -n kube-system
kubectl create clusterrolebinding cluster-dashboard-admin --clusterrole=cluster-admin --serviceaccount=kube-system:dashboard-admin


grep 'client-certificate-data' /etc/kubernetes/admin.conf | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.crt
grep 'client-key-data' /etc/kubernetes/admin.conf | head -n 1 | awk '{print $2}' | base64 -d >> kubecfg.key
openssl pkcs12 -export -clcerts -inkey kubecfg.key -in kubecfg.crt -out kubecfg.p12 -name "kubernetes-client"



kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'

