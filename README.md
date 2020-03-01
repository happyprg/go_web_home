# go_web_home
* How to run in Minikube
```bash
minikube stop && minikube start --cpus 16 --memory 15192
kubectl run go-web-home --image=hongsgo/gowebhome:latest --port=1323 --restart=Always
minikube service go-web-home --url
```
* Kill pid by port
```bash
netstat -ano | findstr :1323
taskkill /PID 17420 /F
```