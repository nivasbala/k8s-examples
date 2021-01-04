# Nginx Webserver Multiple Pages

- This directory has 3 files that deploy a load balancer service using multiple web pages
- Uses Nginx with multiple web **"location"**
- Files
  1. **websrv.yaml**: Basic webserver 
  2. **cm-content.yaml**: The config map for index.html for "/", "/service1/", "/service2/"
  3. **cm-config.yaml**: The config map for /etc/nginx/nginx.conf file
- To create the deployment
  ```
  # Create Name Space "websrv"
  $ kubectl create ns websrv

  # Apply the yaml configs
  $ kubeclt apply -f cm-context.yaml -f cm-config.yaml -f websrv.yaml
  ```
- To Look at the output of the LB web page from a browser type
  - <lb-public-ip>/  - This take you to home
  - <lb-public-ip>/service1 - Takes you to service 1 page
  - <lb-public-ip>/service2 - Takes you to service 2 page
- To Delete the deployment
  ```
  # Delete the yaml configs
  $ kubeclt delete -f cm-context.yaml -f cm-config.yaml -f websrv.yaml

  # Delete Name Space "websrv"
  $ kubectl delete ns websrv
  ```
