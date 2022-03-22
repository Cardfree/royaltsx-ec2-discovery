royaltsx-ec2-discovery
---

Get RoyalTSX: [Downloads](https://royalapps.com/ts/mac/download)


Build:
```
go build .

# main executable will be the build output
```

Run locally:
```
# Cardfree aka nonprod using your aws default profile
AWS_PROFILE=default AWS_REGION=us-west-2 ./royaltsx-ec2-discovery -rdp_credential_name CARDFREE  -ssh_credential_name CARDFREE

# Prdcf aka production using your aws prdcf profile
AWS_PROFILE=prdcf AWS_REGION=us-west-2 ./royaltsx-ec2-discovery -rdp_credential_name PRDCF  -ssh_credential_name PRDCF
```

RoyalTSX Setup:
- Add Credential to RoyalTSX
    - `nonprod` create `CARDFREE` credentials
    - `prdcf` create `PRDCF` credentials
- Add Document to RoyalTSX
    - Call document `Connections`
    - Add `Dynamic Folder` to `Connections` document called `nonprod`.
    - Add `Dynamic Folder` to `Connections` document called `prdcf`.
    - Open `Properties` of the newly created `Dynamic Folder` and select `Dynamic Folder Script` section.
        - For `nonprod` `Dynamic Folder Script` add:
          ```
          # Cardfree aka nonprod using your aws default profile
          # Note: Must point to fully qualified path to royaltsx-ec2-discovery binary. Replace `<username>` with your username if using the same location.
          AWS_PROFILE=default AWS_REGION=us-west-2 /Users/<username>/Documents/Github/Cardfree/royaltsx-ec2-discovery/royaltsx-ec2-discovery -rdp_credential_name CARDFREE  -ssh_credential_name CARDFREE
          ```
        - For `prdcf` `Dynamic Folder Script` add:
          ```
          # Prdcf aka production using your aws prdcf profile
          # Note: Must point to fully qualified path to royaltsx-ec2-discovery binary. Replace `<username>` with your username if using the same location.
          AWS_PROFILE=prdcf AWS_REGION=us-west-2 /Users/<username>/Documents/Github/Cardfree/royaltsx-ec2-discovery/royaltsx-ec2-discovery -rdp_credential_name PRDCF  -ssh_credential_name PRDCF
          ```

References: https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html