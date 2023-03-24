# Setup a VM in GCP
These steps will quickly spin up a VM instance in GCP's Compute Engine.
- Setup SSH Keys
    - Run `ssh-keygen -t ed25519 -C my_ssh_key`
    - Copy ssh public key to GCP
        - Navigate to `Compute Engine -> Settings -> Metadata -> SSH Keys`
            - add the new ed25519.pub to SSH keys
            - This will automatically add this key to all VM instances in this GCP project
- Setup firewall rules for prometheus/grafana
    - Navigate to `VPC Network -> Firewall -> Create Firewall Rule`
        - target tags: http-server
        - Source IPv4 Ranges: `your_IP`, or `0.0.0.0/0` for everyone
        - Specific ports: Under TCP add `9090,3000,8545`, for prometheus,grafana, and geth respectively 
- Setup VM Instance
    - Navigate to `Compute Engine -> VM Instances -> Create Instance`
        - Change boot disk image to `Ubuntu 22.04`
        - Change boot disk size to `300` GB
        - Firewalls -> Enable checkbox for `HTTP`
        - Click Create
