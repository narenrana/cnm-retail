#docker-compose  build

kubectl delete namespace app
#kubectl delete namespace cnm-server
kubectl delete service cnm-auth
kubectl delete deployment cnm-auth


kubectl delete service cnm-cats
kubectl delete deployment cnm-carts


kubectl delete service cnm-coupons
kubectl delete deployment cnm-coupons



kubectl delete service cnm-offers
kubectl delete deployment cnm-offers


kubectl delete service cnm-orders
kubectl delete deployment cnm-orders


kubectl delete service cnm-payments
kubectl delete deployment cnm-payments


kubectl delete service cnm-products
kubectl delete deployment cnm-products


kubectl delete service cnm-users
kubectl delete deployment cnm-users

#kubectl delete namespace main-app
kubectl delete service main-app
kubectl delete deployment main-app

#kubectl delete namespace payment-app
kubectl delete service payment-app
kubectl delete deployment payment-app

kubectl delete ingress cnm-ingress

