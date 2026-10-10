Pangolin (because I like Pangolins and also it's in Go so, you know, panGOlin) is a simple tool for personal use which queries the Proxmox REST API for existing VMs, gets their state and leverages the installed qemu guest agent to pull their hostnames and IP addresses. 
I use this tool occasionally when I need to access some of my running services by IP and not by internal URL via my reverse proxy. This tool makes it relatively fast to get the IPs of the VMs without
needing to log into Proxmox and access the console to find the IP
<img width="1920" height="1040" alt="pangolin" src="https://github.com/user-attachments/assets/22ac5c91-9aad-4ca7-8b23-4829f3d12348" />
