# Nginx Ingress Controller TCP port Setup Info 

## Commands 
 
- command to get cat nginx controller config inside nginx ingress controller deployment
  ```
  kubectl exec -it -n ingress-nginx deploy/nginx-ingress-controller -- cat /etc/nginx/nginx.conf
  ```
- command to edit ingress controller svc
  ```
  kubectl edit svc nginx-ingress-ontroller -n ingress-nginx
  ```
- command to edit the ingress controller deployment 
