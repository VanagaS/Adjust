
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

## Metric Categories

### Default metrics
- CPU utilization
- idle time
- speed
- processor time
- Memory utilization
- Packet loss
- Interface traffic / in-out bytes

### Performance
- Transactions per second

  The number of handshakes with or without re-using SSL Session ID. Monitoring and following this number helps us in either planning ahead for future requirements, or identifying an underlying issue.
- Data size influence on performance

  The inflow or outflow of data size (or the HTTP data) embedded in the HTTPS connection, effects the overall performance with the increase in the size. As the data size grows, the encryption and decryption adds additional delay. Patterns in these can be analysed if there are any DOS attacks or how or what application URLs handle such data and if that is accepted.
  
- Key generation performance

  The total time, taken from syn/ack, client hello to the generation of premaster shared key on either end is expensive both in terms of CPU usage and time taken. Helps analyze how frequently or lately the premaster key needs to be generated and whats the performance benefits. Helps analyze how several factors like Cipher Suite etc are effecting the overall performance.
  
- SSL Session ID caching performance

  SSL Session ID caching cuts down the handshake process by limiting the premaster key generation to specific timing or per number of requests. 
  
- Asymmetric key size performance

  Effects of Asymmentric key size selected during the handhsake. The more the key size, the more the delay (but is more secure). Helps analyze and find the frequency of the key size used and their performance
  
- Outbound connection performance

  How outbound connection (to app server) is performing. For security purposes, this as well might have been configured to encrypt (probably with lower key size Cipher) data. Sometimes the application might be sluggish, or throwing errors which inturn effect the overall time. This metric helps in identifying outbound side broblems.

- Certificate chain performance

  If a certificate chain size is greater than the initial congestion window, the server needs to wait until the client confirms the receipt of the first batch of bytes that are sent before resuming the connection. If round-trips are increased before a session is established,  then slower connection is experienced by the end user.

### Security
- DOS/Rate limiting observations

  We can quickly analyze if there are any connections that fall in the Rate limit criteria and accordingly block the respective IPs or dig further of what could be the issue.

- Ciphers active / Safe Ciphers observations

  To be PCI-DSS complaint and for general security purpose, it is advised to use strong ciphers and block any ciphers that are compromised. The metrics targeting the Ciphers helps us in identifing the clients that are using old, weak or compromised Ciphers. 

- Rejected cipher observations

  These metrics allows us to identify clients who are offering un-supported Ciphers and SSL-Offloader is dropping their connections.

- Downstream/outbound connection encryption performance

  In case downstream connection as well needs SSL/TLS to connect to app, these metrics helps in identifying any bottlenecks in initiating handshake and further encryption/decryption process.

### Generic
- Regional statistics and their effects on performance

  Metrics identifying the client locations. Analysing these metrics for slow connections from particular region and their frequency which can block several CPU threads.

- Connection drop statistics

  Metrics on number of connections being dropped clubbed with what is the CPU, i/o and Memory usage, network bandwidth usage during the time of drops helps in identify issues before hand.

- Static data / File caching

  Static data and/or Files cached locally instead of creating connections to the App server saves lots of time. The peformance of serving static files from local setup, comparisions of performance fetching from local vs App server gives better overview.

- Scalable redundancy

  Scalable redundancy monitor helps in analyzing the transactions per second and analyzing performance improvements.

- Percentage of clients using SSL session ID

  Out of the total connections, how many end users have been using SSL session ID. This helps in promoting to end users to switch to using a SSL session ID. Also helps analyze how frequently server is generating Session ID.

- Ciphers offered by client vs Ciphers picked by Server

  The metrics on what Ciphers are being offered by clients when initiating Client hello and which of them is the Server selecting. We'll also know whether server is selecting a lower cpu intensive cipher out of the cipher list or the highest rated one.
  
- Cipher suite selection - CPU intensive vs Security

### Good-to-know
- Regional statistics
- Client information (browser version, OS version, etc)
- Recording and analyzing WAF or IDS blocked pages or requests
- Time taken per Cipher suite

## Nice to do
- Exportable metrics, so that they can be uploaded/integrated with prometheus, datadog etc.,
- Access logs in customizable format to be imported to ELK stack, etc
- Failover setup
- Caching solutions and CDN network
- Rate limiting


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



