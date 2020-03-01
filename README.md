# go_web_home
* How to run in Minikube
```bash
kubectl run go-web-home --image=hongsgo/go_web_home:latest --port=1323 --restart=Always


minikube service go-web-home --url
```