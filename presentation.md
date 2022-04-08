> cd $HOME/learn/Kubernetes/learn-kubernetes-goapp/app
> go build . 
> ./goapp
> http://localhost:8080/
> docker build -t goapp .
> docker images
> docker run --rm -it -p 8080:8080 goapp
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

> kubectl get namespaces
> kubectl delete namespace goapp-namespace
> eval $(minikube docker-env)
> kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
> kubectl get svc -n argocd
> kubectl port-forward -n argocd svc/argocd-server 8081:443
> login: 
    user: admin 
    password: kubectl get secret argocd-initial-admin-secret -n argocd -o yaml
> echo b2dZRjBDTE0yUG1XMUJMQ== | base64 --decode
> kubectl apply -f argocdgoapp.yaml
> kubectl edit deployment -n goapp goapp-deployment
> kubectl port-forward -n goapp service/goapp-service 30001:8080