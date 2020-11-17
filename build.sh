sh ./clean.sh

docker-compose  build


docker tag cnm-retail_main_app narenrana02/cnm-retail_main_app
docker tag cnm-retail_cnm_payment narenrana02/cnm-retail_cnm_payment
#Reuse same image for all microservices
docker tag cnm-retail_cnm_server narenrana02/cnm-retail_cnm_server


docker push narenrana02/cnm-retail_main_app
docker push narenrana02/cnm-retail_cnm_payment
docker push narenrana02/cnm-retail_cnm_server

#kubectl apply -f  ./deployment-cnm-server.yaml
kubectl apply -f  ./deployment-cnm-auth.yaml
kubectl apply -f  ./deployment-cnm-carts.yaml
kubectl apply -f  ./deployment-cnm-coupons.yaml
kubectl apply -f  ./deployment-cnm-offers.yaml
kubectl apply -f  ./deployment-cnm-orders.yaml
kubectl apply -f  ./deployment-cnm-payments.yaml
kubectl apply -f  ./deployment-cnm-products.yaml
kubectl apply -f  ./deployment-cnm-users.yaml
#Main Ui App
kubectl apply -f  ./deployment-main-app.yaml
#Main payment App
kubectl apply -f  ./deployment-payment-app.yaml
#Router
kubectl apply -f  ./deployment-ingress.yaml


kubectl describe namespace cnm-auth
kubectl describe service cnm-auth
kubectl describe deployment cnm-auth

kubectl describe namespace cnm-carts
kubectl describe service cnm-cats
kubectl describe deployment cnm-carts

kubectl describe namespace cnm-coupons
kubectl describe service cnm-coupons
kubectl describe deployment cnm-coupons


kubectl describe namespace cnm-offers
kubectl describe service cnm-offers
kubectl describe deployment cnm-offers

kubectl describe namespace cnm-orders
kubectl describe service cnm-orders
kubectl describe deployment cnm-orders

kubectl describe namespace cnm-payments
kubectl describe service cnm-payments
kubectl describe deployment cnm-payments

kubectl describe namespace cnm-products
kubectl describe service cnm-products
kubectl describe deployment cnm-products

kubectl describe namespace cnm-users
kubectl describe service cnm-users
kubectl describe deployment cnm-users


kubectl describe namespace main-app
kubectl describe service main-app
kubectl describe deployment main-app

kubectl describe namespace payment-app
kubectl describe service payment-app
kubectl describe deployment payment-app

kubectl describe ingress cnm-ingress


