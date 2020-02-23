
# SSL Offloading metrics
## Server Specification
- 4 times Intel(R) Xeon(R) CPU E7-4830 v4 @ 2.00GHz  
- 64GB of RAM
- 2 TB HDD disk space  
- 2 x 10 Gbit/s nics  

### SSL/TLS - ASIC 
- SSL/TLS specific ASIC not present/not mentioned

### IPv6 support
- IPv6 support not mentioned

### NICs
- 2 x 10 Gbit/s NIC
- One connected to the external network/cloud to the end users and other connected to the internal network to the actual application
### Total CPU(s)
Intel(R) Xeon(R) CPU E7-4830 v4 @ 2.00GHz  
- 2 threads / core
- 14 cores / socket
- 4 sockets

Total 112 On-line CPUs

### Given load
- 25000 requests per second
- 


## Categorizing Metrics

### Default metrics
- CPU utilization
- idle time
- speed
- processor time
- Memory utilization
- Packet loss
- Availability 
- Interface traffic / in-out bytes
- 
### Performance
- Transactions per second
- Data size influence on performance
- Key generation performance
- SSL Session ID caching performance
- Asymmetric key size performance
- Outbound connection performance

### Security
- DOS/Rate limiting observations
- Ciphers active / Safe Ciphers observations
- Rejected cipher observations
- Downstream/outbound connection encryption performance

### Generic
- Regional statistics and their effects on performance
- Connection drop statistics
- Static data / File caching
- Scalable redundancy
- Percentage of clients using SSL session ID
- Ciphers offered by client vs Ciphers picked by Server
- Cipher suite selection - CPU intensive vs Security

### Good-to-know
- Regional statistics
- Client information (browser version, OS version, etc)
- WAF or IDS blocks (recorded)
- Time taken per Cipher suite

# 
- Exportable metrics, so that they can be uploaded/integrated with prometheus, datadog etc.,
- Access logs in customizable format to be imported to ELK stack, etc
- Failover setup
- Caching solutions and CDN network
- Rate limiting

### Certificate chain
It's not uncommon to have chained certificates. If the certificate chain size is greater than the initial congestion window, the server needs to wait until the client confirms the receipt of the first batch of bytes that are sent before resuming the connection. If round-trips are increased before a session is established,  then slower connection is experienced by the end user.

## Challenges of Monitoring
Biggest challenge with monitoring is to set it up in such a way that we can identify any threat beforehand. In a complete setup, there can be any resource that is acting weird or overloaded and signs are clearly shown. A network device can discard packets depending on their memory, hence by affecting the performance. Discards increase application latency. Excessive discards needs to be monitored.

### Network
- Packet drops: -insufficient bandwidth allocation
- Network errors: DNS issues, TCP timeouts, server did not respond
- WAN availability and reliability
- Round trip time
### Alerts
- Alert fatigue
- Actionable Alerts
- Proper thresholds

Unless proper thresholds are applied to the metrics, there will be too many alerts to monitor which would result in an alert fatigue. To avoid this, we may get rid of Non-actionable alerts and prepare thresholds that lead to some actionable alerts. 

### Disk
The logs needs to be configured to be printed to stdout. Afterwhich monitoring tools needs to be configured to read from the stdout stream. If that is not the case, the HDD with size 2 TB would become look miniscule. Also, the type of disk provided is an HDD where an SSD could've been used. SSD's provide better read/write statistics.



