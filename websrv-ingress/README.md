# Ingress Based service with multiple webpages using multiple services
---

 - Uses both **Traefik** Ingress and **Nginx** Ingress
 - Multi-Path Routing
   - Creates Ingress and directs traffic to "/", "/service1/" and "/service2/"
 - Files
   - WebServer (installed under a separate namespace (websrv)
     1. websrv-main.yaml: Has Nginx deploy,configmap (config/content), service for main page
     2. websrv-s1.yaml: Same for Service 1
     3. websrv-s2.yaml: Same for Service 2
   - Traefix Ingress
     1. traefik-ing-rbac.yaml: Create NS and RBAC (rolebind default role for NS to cluster-admin)
     2. ingress-traefik.yaml: Create path based route for main, s1 and s2 using traefik ingress controller
   - Nginx Ingress
     1. nginx-ing-install.yaml: This is from Oracle's website (installed under ingress-nginx)
        - Creates, RBAC, Nginx Deployment, service etc.,
     2. ingress-nginx.yaml: Creates path based routes for main, s1 and s2 using nginx ingress controller
   - Cookie Persistence
     1. hello-world-ing-cookie.yaml
        - Added cookie field to ingress definition for nginx
        - The cookie need to be specified in curl **(curl -b route=12345 <nginx ing svc IP>)**
          - When a backend is hit, that session affinity will be maintained everytime
          - Without the cookie (without -b) replicas will be hit in round-robin
     2. hello-world.yaml
        - Simple deployment (with 3 replicas) which displays the container id
        - [Scott Baldwin Hello World](https://github.com/scottsbaldwin/docker-hello-world)

## 1. Traefik Based Ingress

### 1.1 RBAC Setup

 - Traefik is installed in a separate namespace ***"traefik-ing"***
 - The default service account **"default"** for this namespace
   - This had to be modified to allow clusterRole **"cluster-admin"**
    - See the file traefik-ing.yaml
 > Need to figure out how to install helm with a particular Service Account (mapped to a Role|ClusterRole)

### 1.2 Traefik Install using HELM

 - Helm Chart was used to install traefik ingress. The following commands were used
   ```
       helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
       helm install my-ing stable/traefik -n traefik-ing
       helm search repo stable
       helm list -n <namespace>
       helm delete -n <namespace> <chart-name>  <- From helm list
   ```

### 1.3 Command to run

 - Step 1: Create NS and Rolebind default to cluster-admin
   ```
      kubectl apply -f traefik-ing-rbac.yaml
   ```
 - Step 2: Install Traefik Ingress in NS traefik-ing
   ```
      helm search repo stable | grep traefik
      helm list -n traefik-ing

      # Verify Install and Get Public LB IP
      kubectl get deploy,pod,svc,cm,ingress -o wide -n traefik-ing
   ```
 - Step 3: Create Deployments and Ingress *(in NS websrv)*
   ```
      kubectl -n traefik-ing apply -f ingress-traefik.yaml -f websrv-main.yaml -f websrv-s1.yaml -f websrv-s2.yaml
   ```
 - Step 4: Verify Deployment
   ```
      kubectl get deploy,pod,svc,cm,ingress -o wide -n websrv
      curl <Ingres LB IP>/
      curl <Ingres LB IP>/s1/
      curl <Ingres LB IP>/s2/
   ```

## 2. Nginx Ingress Install

 - Step 1: Apply nginx-ing-install.yaml
   - Installs Nginx Ingress Controller in NS ingress-nginx
   - Added Load balancer shape parameter (other params can be added)
   - Creates, RBAC, Deployment, Svc etc.,
 - Step 2: Apply ingress-nginx.yaml
   - Creates Ingress Path
 - Step 3: Apply websrv files
   - Webserver

## 3. Cookie Persistence

  - Step 1: Install nginx-ing-install.yaml (same as above)
  - Step 2: Install Ingress (hello-world-ing-cookie.yaml)
  - step 3: Install Deployment (hello-world.yaml)
  - Step 4: Curl to access cookie based access (cookie name is  <u>**route**</u>)
    - Cookie set using -b (cookie will be set in the client browser)
      - When hit will cookie, it will always go to the same backend (session affinity)
      ```
          curl -b route=12345 <nginx-ing-svc-ip>
      ```
      - When hit without cookie, it will go to any of the replicas in round-robin

