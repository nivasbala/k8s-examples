# K8S Nginx Ingress Controller to allow TCP Ports
---

## 1. Content

- Has 3 Directories
  1. **docker-img**: Base Docker Image to create basic response (golang based)
     - Displays the host name of the pod 
  2. **http-port-dep**: HTTP port based service deployment (port 80)
     - Has the Yaml for Deployment, service and ingress
  3. **tcp-port-dep**: TCP port based service deployment (port 9000)
     - Has the Yaml for Deployment, service and ingress

## 2. Create and push Docker Images

- Login to Docker hub
  ```
  $ docker login
  ``` 
- Build the docker image and push to docker hub
- The Repository has to be a public repo
  ```
  $ docker build -t nivasbala/go-app-http .
  $ docker build -t nivasbala/go-app-tcp .

  $ docker push nivasbala/go-app-http
  $ docker push nivasbala/go-app-tcp
  ```

## 3. Edit Nginx Ingress Service to add TCP Port

- Use the following commands to edit the ingress controller service to add the **TCP port**
  ```
  $ kubectl -n ingress-nginx edit svc ingress-nginx
  ``` 
- Edit the Ingress Service as shown below
  ```
  spec:
    clusterIP: 10.xx.yy.zz
    externalTrafficPolicy: Cluster
    ports:
    - name: http
      nodePort: 30240
      port: 80
      protocol: TCP
      targetPort: http
    - name: https
      nodePort: 32065
      port: 443
      protocol: TCP
      targetPort: https
    - name: proxied-tcp-9000   <---- These lines have to be added for the port being used
      nodePort: 31896
      port: 9000
      protocol: TCP
      targetPort: 9000
  ```
- Save the above configuration

## 4. Create the Nginx TCP Port Configuration

- From the tcp-port directory do the following
  - Use file tcp-config-map to create **`"config map"`** for tcp port under **ingress namespace**
    - This is the port that will be used to access the service
    ```
    $ kubectl -n ingress-nginx apply -f tcp-config-map.yaml
    ```
  - Verify this by displaying ingress service - Should show the TCP port (9000)
    ```
    $ kubectl -n ingress-nginx get svc
    NAME            TYPE           CLUSTER-IP     EXTERNAL-IP      PORT(S)                                     AGE
    ingress-nginx   LoadBalancer   a.b.c.d        w.x.y.z   80:30240/TCP,443:32065/TCP,9000:31896/TCP   36d

    ```

## 5. Create the TCP based deployment

- Create TLS certificate to use in the ingress controller
  ```
  $ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=nginxsvc/O=nginxsvc"
  
  $ kubectl create secret tls tls-secret --key tls.key --cert tls.crt
  ```
- Apply the K8S yaml file under tcp-port directory
  ```
  $ kubectl apply -f go-app-tcp-ing.yaml -f go-app-tcp.yaml
  ```

## 6. Create the HTTP based deployment

- Create TLS certificate to use in the ingress controller
  ```
  $ openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=nginxsvc/O=nginxsvc"
  
  $ kubectl create secret tls tls-secret --key tls.key --cert tls.crt
  ```
- Apply the K8S yaml file under http-port directory
  ```
  $ kubectl apply -f go-app-http-ing.yaml -f go-app-http.yaml
  ``` 

## 7. Test output

- Get the Ingress Service IP
  ``` 
  $ kubectl -n ingress-nginx get svc 
  NAME            TYPE           CLUSTER-IP     EXTERNAL-IP      PORT(S)                                     AGE
  ingress-nginx   LoadBalancer   a.b.c.d        x.y.z.w   80:30240/TCP,443:32065/TCP,9000:31896/TCP   36d
  ```
- For accessing the TCP based service use <external-ip>:9000
  - This is the TCP port and service which is being accessed on a non **HTTP|HTTPS** port
- For accessing the HTTP based service use <external-ip>/test
  - This is the path provided in the ingress
- To check inside the ingress controller deployment **`"exec"`** into the pod of the nginx deployment and cat **`"nginx.conf"`**
  ```
  kubectl exec -it -n ingress-nginx deploy/nginx-ingress-controller -- cat /etc/nginx/nginx.conf
  ```
  - This needs to be done after the K8S ingress manifest object has been created
  - Will show the TCP port based configuration in the conf file
