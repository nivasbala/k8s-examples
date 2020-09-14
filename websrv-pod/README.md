# Basic WebSrv Pod

 - This file has basic nginx websrv but with a config file specified by /etc/nginx/nginx.conf
 - Create under **websrv** namespace

## Access the Pod

 - Apply the pod and create the config map
   ```
      kubectl apply -f nginx-conf.yaml -f nginx-pod.yaml
   ```
 - Exec into the pod
   ```
      kubectl -n websrv exec -it <pod name> -- /bin/bash
      curl localhost
   ```
