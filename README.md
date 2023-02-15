# kubernetesStudy

just study kubernetes knowledge

# Cluster type to identify the deployment type
    - name: CLUSTER_TYPE
        value: "k8s,bgp"
    - name: IP_AUTODETECTION_METHOD
        value: "interface=eth0"

# set hostname

    $ sudo hostnamectl set-hostname "k8s-master"       // Run on master node
    $ sudo hostnamectl set-hostname "k8s-worker1"      // Run on 1st worker node
    $ sudo hostnamectl set-hostname "k8s-worker2"      // Run on 2nd worker node
    $ sudo hostnamectl set-hostname "k8s-worker3"      // Run on 3rd worker node
```
172.22.16.12 k8s-master
172.22.16.14 k8s-worker1
172.22.16.10 k8s-worker2
172.22.16.5 k8s-worker3
```