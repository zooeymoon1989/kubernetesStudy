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
10.0.0.14 k8s-master
10.0.0.8 k8s-worker1
10.0.0.15 k8s-worker2
10.0.0.6 k8s-worker3
```