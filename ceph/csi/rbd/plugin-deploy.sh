#!/bin/bash
# 执行本脚本就可安装完成
deployment_base="${1}"

if [[ -z $deployment_base ]]; then
	deployment_base="../deploy/rbd/kubernetes/v1.14+"
#echo "---------------------------------------"
fi

cd "$deployment_base" || exit 1

objects=(csi-provisioner-rbac csi-nodeplugin-rbac  csi-rbdplugin-provisioner csi-rbdplugin)

for obj in "${objects[@]}"; do
	kubectl create -f "./$obj.yaml"
#echo "========="
done
