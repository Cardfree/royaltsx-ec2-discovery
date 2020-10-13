royaltsx-ec2-discovery
---

Build:
```
go build .

# main executable will be the build output
```

```
# Cardfree aka nonprod using your aws default profile
AWS_PROFILE=default AWS_REGION=us-west-2 ./main -rdp_credential_name CARDFREE  -ssh_credential_name CARDFREE

# Prdcf aka production using your aws prdcf profile
AWS_PROFILE=prdcf AWS_REGION=us-west-2 ./main -rdp_credential_name PRDCF  -ssh_credential_name PRDCF
```