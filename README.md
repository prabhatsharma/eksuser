# eksuser

eksuser is a convenience utility that you can use to manage Amazon EKS users.

It allows you to add, update and delete existing IAM users to EKS.

## Prerequisites

1. An Amazon EKS cluster is installed and running
2. aws-cli is configured
3. kubectl and aws-iam-authenticator are configured
4. Existing kubernetes groups that have access

You can create a group you can create a Role/ClusterRole and then create a binding:

dev-role1.yaml - A Role that gives rights to everything in namespace app1

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: super-developer
  namespace: app1
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]

---

kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: super-developer
  namespace: app1
subjects:
- kind: Group
  name: super-developer
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: super-developer
  apiGroup: rbac.authorization.k8s.io
```

```shell
$ kubectl apply -f dev-role1.yaml
```
admin-cluster-role1.yaml - A ClusterRole that gives supr privileges on cluster

```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: super-admin
rules:
- apiGroups: [ "*" ]
  resources: ["*"]
  verbs: ["*"]
- nonResourceURLs: ["*"]
  verbs: ["*"]

---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: super-admin
subjects:
- kind: Group
  name: super-admin
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: super-admin
  apiGroup: rbac.authorization.k8s.io
```

```shell
$ kubectl apply -f admin-cluster-role1.yaml
```

Now to add an existing IAM user to EKS:

```shell
$ eksuser add --user=prabhat --group=super-admin
$ eksuser add --user=prabhat --group=super-admin,super-developer
```

To update an existing IAM user to EKS:

```shell
$ eksuser update --user=prabhat --group=super-developer
```

To delete an existing IAM user to EKS:

```shell
$ eksuser delete --user=prabhat
```
Remember that it does not delete the IAM user from AWS IAM, just the IAM user entry from EKS.


## Generate kubeconfig file

On user's machine who has been added to EKS, they can configure .kube/config file using the following command:

```shell
$ aws eks update-kubeconfig --name cluster_name
```

# Installation

Download binaries from [releases page](https://github.com/prabhatsharma/eksuser/releases/) and place the binary in PATH
