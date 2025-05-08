# Add kubernetes-dashboard repository
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
# Deploy a Helm Release named "kubernetes-dashboard" using the kubernetes-dashboard chart
helm upgrade --install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard --create-namespace --namespace kubernetes-dashboard

# Create service account and sercret

kubectl apply -f .\kube-dashboard\dashboard-adminuser.yaml

# Get token to login to Kubernetes Dashboard

kubectl -n kubernetes-dashboard create token admin-user   