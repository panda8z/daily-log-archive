```
zcj@PandaMac-2 ~ % curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 4572k  100 4572k    0     0   201k      0  0:00:22  0:00:22 --:--:--  773k
WARN: adding root certificate via internet: https://tiup-mirrors.pingcap.com/root.json
You can revoke this by remove /Users/zcj/.tiup/bin/7b8e153f2e2d0928.root.json
Set mirror to https://tiup-mirrors.pingcap.com success
Detected shell: zsh
Shell profile:  /Users/zcj/.zshrc
/Users/zcj/.zshrc has been modified to to add tiup to PATH
open a new terminal or source /Users/zcj/.zshrc to use it
Installed path: /Users/zcj/.tiup/bin/tiup
===============================================
Have a try:     tiup playground
===============================================


```

```

zcj@PandaMac-2 ~ % source .zshrc
zcj@PandaMac-2 ~ % tiup playground
The component `playground` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz download https://tiup-mirrors.pingcap.com/playground-v1.1.download https://tiup-mirrors.pingcap.com/playground-v1.1.2-darwin-amd64.tar.gz 9.03 MiB / 9.03 MiB 100.00% 2.92 MiB p/s                                                              
Starting component `playground`: /Users/zcj/.tiup/components/playground/v1.1.2/tiup-playground
Use the latest stable version: v4.0.6

    Specify version manually:   tiup playground <version>
    The stable version:         tiup playground v4.0.0
    The nightly version:        tiup playground nightly

Playground Bootstrapping...
The component `prometheus` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/prometheus-v4.0.6-darwin-amd64.tar.gz 39.78 MiB / 39.78 MiB 100.00% 1.69 MiB p/s                                                             
download https://tiup-mirrors.pingcap.com/grafana-v4.0.6-darwin-amd64.tar.gz 43.87 MiB / 43.87 MiB 100.00% 1.59 MiB p/s                                                                
Start pd instance...
The component `pd` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/pd-v4.0.6-darwin-amd64.tar.gz 39.23 MiB / 39.23 MiB 100.00% 2.55 MiB p/s                                                                     
Start tikv instance...
The component `tikv` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/tikv-v4.0.6-darwin-amd64.tar.gz 15.91 MiB / 15.91 MiB 100.00% 2.53 MiB p/s                                                                   
Start tidb instance...
The component `tidb` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/tidb-v4.0.6-darwin-amd64.tar.gz 39.26 MiB / 39.26 MiB 100.00% 2.47 MiB p/s                                                                   
...
Waiting for tikv 127.0.0.1:20160 ready 
Start tiflash instance...
The component `tiflash` is not installed; downloading from repository.
download https://tiup-mirrors.pingcap.com/tiflash-v4.0.6-darwin-amd64.tar.gz 57.55 MiB / 57.55 MiB 100.00% 1.89 MiB p/s                                                                
Waiting for tiflash 127.0.0.1:3930 ready ..
CLUSTER START SUCCESSFULLY, Enjoy it ^-^
To connect TiDB: mysql --host 127.0.0.1 --port 4000 -u root
To view the dashboard: http://127.0.0.1:2379/dashboard
To view the Prometheus: http://127.0.0.1:60887
To view the Grafana: http://127.0.0.1:3000


```

