# Geth Prsym Ansible Configuration Management

## About
This repo has two folders:
- **ansible** - This is `Part 1` of the challenge. It contains an ansible playbook that will provision a VM with:
    1. An Eth node using Geth and Prysm
    1. Monitoring with prometheus and grafana
- **go** - This is `Part 2` of the challenge. It contain a Go script that will continuously poll a target eth node's `admin_peers` RPC method and display the node's connected peers.

## Prerequisites
1. Ansible installed locally
1. A target VM with `Ubuntu 22.04` installed and SSH access. [gcp_guide.md](gcp_guide.md) contains instructions for setting up a quick VM in GCP. 

## Part 1 - Eth Node with geth/prysm
### Setup
1. Update `my_eth_node` in inventory file `ansible/inventory/hosts.yml`:
    - Set `ansible_host` to your VM's IP address
    - Set `ansible_user` to the SSH user that has access to the VM
1. Update `monitoring_ip` to your current IP address, in file `ansible/roles/eth/vars/main.yml`. This is the IP that is given firewall access in UFW for grafana, prometheus and geth.

### Execution
- Run `ansible-playbook -i inventory playbook.yml` from the `ansible` directory

### Results
After executing the ansible playbook:
- `Grafana` is accessible at port `3000`, with default user/pass: `admin/admin`
    - There is an example dashboard called `Geth` which will show system and geth metrics
- `Prmetheus` is accessible at port `9090` for direct access to metrics
- `Geth` RPC API is accessible at port `8545`
- `Service logs` can be checked with:
    - `sudo journalctl -f -u [geth|prysm]`

## Part 2 - Go program for polling geth peers
### Execution
1. From the `go` directory, run:
```
`./pollpeers -addr=[RPC_IP] -port=[RPC_PORT] -pollrate=[POLL_RATE]`
```
- where:
    - `RPC_IP` = IP of Eth Node, `default=localhost`
    - `RPC_PORT` = Port of Eth Node, `default=8545`
    - `POLL_RATE` = How often to poll (in seconds), `default=3`
