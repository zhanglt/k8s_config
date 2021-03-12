kubeadm init --kubernetes-version=v1.20.4 \
	--pod-network-cidr=10.244.0.0/16 \
	--ignore-preflight-errors=swap --v=6 \
	--apiserver-advertise-address=10.0.0.81


#kubeadm join 10.0.0.61:6443 --token xmdiyt.ot3gehcorr234fkl \
#	    --discovery-token-ca-cert-hash sha256:38a5a8d1b364dbd9aa17d4b9076d3c94074e01b5d504bd3304fb4cb814f10db4
#kubeadm join 10.0.0.61:6443 --token gxt7ni.uq379ys4vne7f8gl \
#	    --discovery-token-ca-cert-hash sha256:c1e20c0792d205347a3d2bac011bcb129bb6abf195003ee577296c0988187137 

