Pre-requisities
===============

1. Sign-up for Google Cloud Platform (https://cloud.google.com/compute/).

2. Ask permission for project `quality-farm-1247`.

3. Install Google Cloud SDK (which gives `gcloud` command): https://cloud.google.com/sdk/#Quick_Start

4. Initialize to sign-in with `gcloud init`.

5. Set default zone with `gcloud config set compute/zone europe-west1-d`.

6. Go to https://console.cloud.google.com/kubernetes/list?project=quality-farm-1247 and create Container cluster, eg. `firstname-cluster`.

6. Fetch configuration for `kubectl` tool: `gcloud container clusters get-credentials quality-farm`.

7. Verify that kubectl is working: `kubectl cluster-info`.

Build and Deploy
================

1. Login with Docker in Google Container Registry: `docker login -e your.email@gmail.com -u _token -p "$(gcloud auth print-access-token)" https://gcr.io`.

2. Build Docker image with proper tag: `docker build -t gcr.io/quality-farm-1247/quality-farm:v1 .`.

3. Push Docker image: `docker push gcr.io/quality-farm-1247/quality-farm:v1`.

4. Create new pod: `kubectl run quality-farm --image=gcr.io/quality-farm-1247/quality-farm:v1 --port=8080` and verify that it works:

```sh
$ kubectl get pods
NAME                 READY     STATUS    RESTARTS   AGE
quality-farm-kd45x   0/1       Pending   0          27s
```

5. Replicate to more pods: `kubectl scale rc quality-farm --replicas=4`.

6. Create a load balancer for our service: `kubectl expose rc quality-farm --type="LoadBalancer"` and verify that is works:

```sh
$ kubectl get services quality-farm 
NAME           CLUSTER_IP     EXTERNAL_IP   PORT(S)    SELECTOR           AGE
quality-farm   10.3.250.114                 8080/TCP   run=quality-farm   49s

$ kubectl get services quality-farm 
NAME           CLUSTER_IP     EXTERNAL_IP      PORT(S)    SELECTOR           AGE
quality-farm   10.3.250.114   104.155.80.178   8080/TCP   run=quality-farm   1m
```

Update and Deploy
=================

1. Do some changes and rebuild Docker image: `docker build -t gcr.io/quality-farm-1247/quality-farm:v2 .`.

2. Push new Docker image: `docker push gcr.io/quality-farm-1247/quality-farm:v2`.

3. Roll-out the update: `kubectl rolling-update quality-farm --image=gcr.io/quality-farm-1247/quality-farm:v2 --update-period=2s`.


Kubernetes Dashboard
====================

1. Get password from `kubectl config view | grep "password"`.

2. Get URL from `kubectl cluster-info | grep UI`.

3. Login with user `admin`.

Tear Down
=========

1. Delete all services: `kubectl delete services quality-farm`.

2. Delete replication controller running our pods: `kubectl delete rc quality-farm`.

3. Delete the cluster: `gcloud container clusters delete quality-farm`.

4. Delete Docker registry storage: `gsutil rm -r gs://artifacts.quality-farm-1247.appspot.com/`
