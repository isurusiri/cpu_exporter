# CPU Exporter

A [Prometheus](https://prometheus.io/) custom metrics exporter written to expose CPU statistics of linux machines (similar to [Node exporter](https://github.com/prometheus/node_exporter), but only for CPU stats).   


This exporter exposes following metrics,   
- Idle, `cpu_exporter_cpu_idle_time`, Shows the CPU idle time    
- Total, `cpu_exporter_cpu_total_usage`, Shows the total CPU availability   
- Utilization, `cpu_exporter_cpu_utilization`, Shows the total CPU utilization   


```
         Scrape       File read
   ......                      ......
   :    :  -->   ----   -->    :    :
   :    :        |  |          :    :
   :    :  <--   ----   <--    :    :
   ......      exporter        ......
 Prometheus                  /proc/stat
```

