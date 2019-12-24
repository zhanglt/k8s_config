#!/bin/bash 
# Modify Author: zhanglt
 # Modify Date: 2019-06-23
 # Version:
 #***************************************************************#
 # Initialize the machine. This needs to be executed on every machine.
 
 # Add host domain name.
 cat >> /etc/hosts << EOF
 172.24.8.74 k8snode01
 172.24.8.75 k8snode02
 172.24.8.76 k8snode03
 EOF
 
 # Add docker user
 useradd -m docker
 
 # Disable the SELinux.
 sed -i 's/^SELINUX=.*/SELINUX=disabled/' /etc/selinux/config
 
 # Turn off and disable the firewalld.
 systemctl stop firewalld
 systemctl disable firewalld
 
 # Modify related kernel parameters & Disable the swap.
 cat > /etc/sysctl.d/k8s.conf << EOF
 net.ipv4.ip_forward = 1
 net.bridge.bridge-nf-call-ip6tables = 1
 net.bridge.bridge-nf-call-iptables = 1
 net.ipv4.tcp_tw_recycle = 0
 vm.swappiness = 0
 vm.overcommit_memory = 1
 vm.panic_on_oom = 0
 net.ipv6.conf.all.disable_ipv6 = 1
 EOF
 sysctl -p /etc/sysctl.d/k8s.conf >&/dev/null
 swapoff -a
 sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
 modprobe br_netfilter
 
 # Add ipvs modules
 cat > /etc/sysconfig/modules/ipvs.modules <<EOF
 #!/bin/bash
 modprobe -- ip_vs
 modprobe -- ip_vs_rr
 modprobe -- ip_vs_wrr
 modprobe -- ip_vs_sh
 modprobe -- nf_conntrack_ipv4
 EOF
 chmod 755 /etc/sysconfig/modules/ipvs.modules
 bash /etc/sysconfig/modules/ipvs.modules
 
 # Install rpm
 yum install -y conntrack git ntpdate ntp ipvsadm ipset jq 
#iptables curl sysstat libseccomp wget gcc gcc-c++ make openssl-devel
 
 # Install Docker Compose
 sudo curl -L "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
 sudo chmod +x /usr/local/bin/docker-compose
 
 # Update kernel
 rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org
 rpm -Uvh https://www.elrepo.org/elrepo-release-7.0-4.el7.elrepo.noarch.rpm
 yum --disablerepo="*" --enablerepo="elrepo-kernel" install -y kernel-ml-5.4.1-1.el7.elrepo
 sed -i 's/^GRUB_DEFAULT=.*/GRUB_DEFAULT=0/' /etc/default/grub
 grub2-mkconfig -o /boot/grub2/grub.cfg
 yum update -y
 
 # Reboot the machine.
 #reboot
