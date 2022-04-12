> cd $HOME/learn/Kubernetes/learn-kubernetes-goapp/app
> go build . 
> ./goapp
> http://localhost:8080/
> docker build -t goapp .
> docker images
> docker run --rm -it -p 8080:8080 goapp
> http://localhost:8080/
> cd ../database
> docker build -t goapp-database .
> docker run --rm -it -p 5000:5432 --name goapp-database -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres goapp-database

> 
export POSTGRES_HOST=0.0.0.0
export POSTGRES_PORT=5000
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=password
export POSTGRES_DB=goapp

> cd ../app
> ./goapp
> http://localhost:8080/
> http://localhost:8080/worker?id=123
> docker exec -it goapp-database psql -d goapp -U postgres
> select * from workers;


> minikube start
> kubectl get namespaces
> eval $(minikube docker-env)
> cd $HOME/learn/Kubernetes/learn-kubernetes-goapp/kubernetes
> kubectl create namespace goapp-namespace
> kubectl apply -f configmap.yaml --namespace=goapp-namespace
> kubectl apply -f goapp-secret.yaml --namespace=goapp-namespace
> kubectl apply -f goapp-database.yaml --namespace=goapp-namespace
> kubectl apply -f goapp.yaml --namespace=goapp-namespace
> kubectl get all -n goapp-namespace
> kubectl logs -f --namespace=goapp-namespace pod/goapp-database-deployment-756dc74c4c-zxcbb
> kubectl logs -f --namespace=goapp-namespace pod/goapp-deployment-6ff48bcc7f-f9j4l
> kubectl cluster-info
> http://192.168.49.2:30001/
> http://192.168.49.2:30001/worker?id=77
> kubectl exec -it --namespace goapp-namespace pod/goapp-database-deployment-756dc74c4c-66tqj -- psql -d goapp -U admin
> select * from workers;
>> * update replicas in goapp.yaml *
> kubectl apply -f goapp.yaml --namespace=goapp-namespace
> kubectl get all -n goapp-namespace

> export KUBE_EDITOR='code --wait'
> kubectl edit deployment -n goapp-namespace goapp-deployment


> kubectl get namespaces
> kubectl delete namespace goapp-namespace
> kubectl create namespace argocd
> eval $(minikube docker-env)
> kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
> kubectl get svc -n argocd
> kubectl port-forward -n argocd svc/argocd-server 8081:443
> https://localhost:8081/
> login: 
    user: admin 
    password: kubectl get secret argocd-initial-admin-secret -n argocd -o yaml
> echo b2dZRjBDTE0yUG1XMUJMQ== | base64 --decode
> eval $(minikube docker-env)
> cd ../argocd
> kubectl apply -f argocd.yaml
> kubectl get all --namespace goapp-namespace
> http://192.168.49.2:30001/
> http://192.168.49.2:30001/worker?id=22
> export KUBE_EDITOR='code --wait'
> kubectl edit deployment -n goapp-namespace goapp-deployment
> kubectl get all -n goapp-namespace
>> * update git repo in goapp.yaml *
> kubectl exec -it  --namespace goapp-namespace pod/goapp-database-deployment-756dc74c4c-66tqj -- psql -d goapp -U admin
> select * from workers;