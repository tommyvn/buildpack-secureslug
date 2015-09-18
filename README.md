# buildpack-secureslug

## This isn't secure unless you're running the build in your own environment

This is a tech POC. You need to build your slugs in a trusted environment and only then deploy them to the untrusted environment.

## QuickStart

First set the relevant buildpack for your app followed by this one:
```
$ heroku buildpacks:set https://github.com/heroku/heroku-buildpack-python.git
$ heroku buildpacks:set https://github.com/tommyvn/buildpack-secureslug.git
```

Configure your AWS KMS details
```
heroku config:set AWS_ACCESS_KEY_ID="KEY" AWS_SECRET_ACCESS_KEY="SECRET" AWS_KMS_KEY_ID="KMS CUSTOMER KEY UUID" AWS_REGION="us-east-1"
```

Set the key provider to aws\_kms
```
heroku config:set KEY_PROVIDER=aws_kms
```

From this point on your app will be encrypted at build time and decrypted at deploy time using your KMS key and all your log lines will be prepended with a sha256 hash of the combination of your kms key and the log line.

You can see key access during deploys in CloudTrail and revoke keys to block deploys but because the data keys are used your customer key is only transiently exposed to the PaaS, giving you an audit trial for decryption actions and the ability to revoke your key and prevent further access to your slug.
