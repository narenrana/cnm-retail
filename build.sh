sh ./clean.sh

echo "Build docker images cnm-retail_main_app, cnm-retail_cnm_payment, cnm-retail_cnm_server"
docker-compose  build

echo "Tag docker images cnm-retail_main_app, cnm-retail_cnm_payment, cnm-retail_cnm_server"
docker tag cnm-retail_main_app narenrana02/cnm-retail_main_app
docker tag cnm-retail_cnm_payment narenrana02/cnm-retail_cnm_payment
docker tag cnm-retail_cnm_server narenrana02/cnm-retail_cnm_server

echo "Pushing docker images cnm-retail_main_app, cnm-retail_cnm_payment, cnm-retail_cnm_server to docker repository"
echo "Warning - Use local or custom remote repository"
docker push narenrana02/cnm-retail_main_app
docker push narenrana02/cnm-retail_cnm_payment
docker push narenrana02/cnm-retail_cnm_server

echo "Creating services and deployments"
echo "Step-1 - init"
kubectl apply -f  ./deployment-init.yaml
echo "Step-2 - cnm-auth"
kubectl apply -f  ./cnm-auth/deployment.yaml
echo "Step-3 - cnm-carts"
kubectl apply -f  ./cnm-carts/deployment.yaml
echo "Step-4 - cnm-coupons"
kubectl apply -f  ./cnm-coupons/deployment.yaml
echo "Step-5 - cnm-offers"
kubectl apply -f  ./cnm-offers/deployment.yaml
echo "Step-6 - cnm-orders"
kubectl apply -f  ./cnm-orders/deployment.yaml
echo "Step-7 - cnm-payments"
kubectl apply -f  ./cnm-payments/deployment.yaml
echo "Step-8 - cnm-products"
kubectl apply -f  ./cnm-products/deployment.yaml
echo "Step-9 - cnm-users"
kubectl apply -f  ./cnm-users/deployment.yaml
#Main Ui App
echo "Step-10 - cnm-app"
kubectl apply -f  ./cnm-app/deployment.yaml
#Main payment App
echo "Step-11 - payment-app"
kubectl apply -f  ./payment-app/deployment.yaml
#Router
echo "Step-12 - deployment-ingress"
kubectl apply -f  ./deployment-ingress.yaml

echo "Describe services and deployments"
kubectl describe service cnm-auth -n app
kubectl describe deployment cnm-auth -n app


kubectl describe service cnm-cats -n app
kubectl describe deployment cnm-carts -n app


kubectl describe service cnm-coupons -n app
kubectl describe deployment cnm-coupons -n app



kubectl describe service cnm-offers -n app
kubectl describe deployment cnm-offers -n app



kubectl describe service cnm-orders -n app
kubectl describe deployment cnm-orders -n app


kubectl describe service cnm-payments -n app
kubectl describe deployment cnm-payments -n app

kubectl describe service cnm-products -n app
kubectl describe deployment cnm-products -n app

kubectl describe service cnm-users -n app
kubectl describe deployment cnm-users -n app



kubectl describe service main-app -n app
kubectl describe deployment main-app -n app


kubectl describe service payment-app -n app
kubectl describe deployment payment-app -n app

kubectl describe ingress cnm-ingress -n app


