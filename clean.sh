#docker-compose  build

echo "Deleting cnm-auth services and deployments"
kubectl delete deployment cnm-auth -n app
kubectl delete service cnm-auth -n app

echo "Deleting cnm-carts services and deployments"
kubectl delete deployment cnm-carts -n app
kubectl delete service cnm-cats -n app

echo "Deleting cnm-coupons services and deployments"
kubectl delete deployment cnm-coupons -n app
kubectl delete service cnm-coupons -n app

echo "Deleting cnm-offers services and deployments"
kubectl delete deployment cnm-offers -n app
kubectl delete service cnm-offers -n app

echo "Deleting cnm-orders services and deployments"
kubectl delete deployment cnm-orders -n app
kubectl delete service cnm-orders -n app

echo "Deleting cnm-payments services and deployments"
kubectl delete deployment cnm-payments -n app
kubectl delete service cnm-payments -n app

echo "Deleting cnm-products services and deployments"
kubectl delete deployment cnm-products -n app
kubectl delete service cnm-products -n app

echo "Deleting cnm-users services and deployments"
kubectl delete deployment cnm-users -n app
kubectl delete service cnm-users -n app

echo "Deleting main-app services and deployments"
kubectl delete deployment main-app -n app
kubectl delete service main-app -n app

echo "Deleting payment-app services and deployments"
kubectl delete deployment payment-app -n app
kubectl delete service payment-app -n app

echo "Deleting cnm-ingress services and deployments"
kubectl delete ingress cnm-ingress -n app

echo "Deleting app services and deployments"
kubectl delete namespace app

