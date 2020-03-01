* CMD
```bash
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | sls admin-user | ForEach-Object { $_ -Split '\s+' } | Select -First 1)
kubectl proxy
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
```
* Referer
https://github.com/kubernetes/dashboard/blob/master/docs/user/access-control/creating-sample-user.md
https://github.com/kubernetes/dashboard