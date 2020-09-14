# Nginx Webserver Multiple Pages

 - This directory has 3 files that deploy a load balancer service using multiple web pages
 - Uses Nginx with multiple web **"location"**
 - Files
   1. **websrv.yaml**: Basic webserver 
   2. **cm-content.yaml**: The config map for index.html for "/", "/service1/", "/service2/"
   3. **cm-config.yaml**: The config map for /etc/nginx/nginx.conf file
