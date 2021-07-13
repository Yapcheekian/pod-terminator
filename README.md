## Environment Variables

The tool is mainly configured through environment variables. These are:
- namespace (required): provide working namespace
- label (required): filter pod by label
- duration (required): delete a pod if the creation timestamp exceed the allowed duration

## Required Kubernetes Service Account

You will need to setup a service account with permissions to list/delete pods. Ideally, you should give this service account the minimal amount of permissions needed to do its job. An example of this minimal permissions setup can be found in example/serviceAccount.yaml. You can use this apply this configuration directly as follows:

```
kubectl apply -f example/serviceAccount.yaml
```

## Deploy the cronjob

You'll need to Deploy the cron job using the example yaml file in example/cronjob.yaml:

```
kubectl apply -f example/cronjob.yaml
```
