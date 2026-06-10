# Network Monitor

## Overview

Network Monitor is a Go-based tool for capturing and analysing network traffic, designed for efficiency and compatibility with various platforms, including PCs or Raspberry Pi. It utilises InfluxDB for storing network metrics, providing insights into network performance and health.

## Features

- **Comprehensive Packet Capture**: Logs every packet, detailing size, source, destination, and protocol.
- **Performance Metrics**: Measures latency, jitter, and packet loss to assess network health.
- **Data Aggregation and Visualisation**: Stores metrics in InfluxDB, supporting complex analyses and visualisations with tools like Grafana or directly through the InfluxDB dashboards themselves.

![network-monitor.png](images/packets-by-protocol.png)

## Deployment

Network Monitor can be run on any computer or a Raspberry Pi, requiring minimal setup.

### Prerequisites

- A computer or Raspberry Pi with network access.

### Configuration

## Building Network Monitor

- Install MSYS2
- Set up MSYS2 in system/user path
- Follow https://docs.fyne.io/started/quick/



## Conclusion

Network Monitor, with its detailed packet capturing and integration with InfluxDB, offers a powerful and flexible solution for monitoring network performance and health. Its ability to capture all packets, along with their size and protocol, opens up possibilities for in-depth network analysis and visualisation. Whether you're looking to troubleshoot network issues, optimise performance, or simply gain a better understanding of your network traffic, Network Monitor provides the tools and data you need to achieve your goals.
