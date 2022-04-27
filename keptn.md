https://agardner.net/keptn-hello-world/


> helm repo add keptn https://charts.keptn.sh && helm repo update
> helm install keptn https://github.com/keptn/keptn/releases/download/0.12.0/keptn-0.12.0.tgz -n keptn --create-namespace --wait
> helm install jmeter-service keptn/jmeter-service -n keptn
> helm install helm-service keptn/helm-service -n keptn
> kubectl get all -n keptn
# get the port from 
## kubectl get service -n keptn api-gateway-nginx

http://192.168.49.2:31478/bridge/dashboard
> login: 
    user: keptn 
    password: kubectl get secrets -n keptn bridge-credentials -o yaml
> echo b2dZRjBDTE0yUG1XMUJMQ== | base64 --decode
> keptn auth --endpoint=192.168.49.2:30932
> keptn auth --endpoint=192.168.49.2:30932 --api-token=$KEPTN_API_TOKEN

GIT_USER=tarekbadrshalaan
GIT_TOKEN=...
GIT_REMOTE_URL=https://github.com/tarekbadrshalaan/sockshop
GIT_REMOTE_URL=https://github.com/tarek-keptn/sockshop

> keptn create project sockshop --shipyard=./shipyard.yaml --git-user=$GIT_USER --git-token=$GIT_TOKEN --git-remote-url=$GIT_REMOTE_URL




> curl -sL https://get.keptn.sh | KEPTN_VERSION=0.12.0 bash
> helm install keptn https://github.com/keptn/keptn/releases/download/0.12.0/keptn-0.12.0.tgz -n keptn --create-namespace --wait --set=control-plane.apiGatewayNginx.type=LoadBalancer

> helm repo add keptn https://charts.keptn.sh && helm repo update
> helm install keptn keptn/keptn -n keptn --create-namespace --wait --set=control-plane.apiGatewayNginx.type=LoadBalancer

> helm upgrade keptn https://github.com/keptn/keptn/releases/download/0.12.0/keptn-0.12.0.tgz -n keptn --set=control-plane.apiGatewayNginx.type=LoadBalancer --wait

> export KEPTN_ENDPOINT=$(kubectl get services -n keptn api-gateway-nginx -o=jsonpath='{.status.loadBalancer.ingress[0].ip}')
> echo "Keptn Available at: http://$KEPTN_ENDPOINT"
> keptn auth --endpoint=$KEPTN_ENDPOINT











