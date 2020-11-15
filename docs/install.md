# install

## install NGSI Go binary

Install NGSI Go binary in `/usr/local/bin`.

```
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.1.0/ngsi-v0.1.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.1.0-linux-amd64.tar.gz -C /usr/local/bin
```

## install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

## Other binaries

-    ngsi-v0.1.0-linux-arm.tar.gz
-    ngsi-v0.1.0-linux-arm64.tar.gz
-    ngsi-v0.1.0-darwin-amd64.tar.gz