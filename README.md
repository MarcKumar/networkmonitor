# Networkmonitor - Internet Failure Detection and Powercycle

![Networkmonitor - Internet Failure Detection and Powercycle](link_to_image)

Welcome to the Networkmonitor GitHub project for Internet Failure Detection and Powercycle automation! This Go program allows you to monitor your router's internet connection by checking a specified URL and automatically perform a power cycle on a Tasmota-compatible plug to resolve network issues.

## Table of Contents
1. [Description](#description)
2. [Requirements](#requirements)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [Usage](#usage)
6. [Notes](#notes)
7. [Contributing](#contributing)
8. [License](#license)

## Description

Networkmonitor is a lightweight Go program designed to detect internet failures by periodically checking a user-defined URL. When an internet failure is detected, it automatically triggers a power cycle on a Tasmota-compatible plug connected to the network. This automation eliminates the need for manual intervention and helps to quickly resolve common network issues.

## Requirements

To use Networkmonitor, you need the following:

1. A Tasmota-compatible plug that can be controlled over the local network.
2. A computer or server with Docker installed, where the program will be executed.
3. A working internet connection for the computer or server.

## Installation

1. Clone the GitHub repository to your computer or server:

   ```
   git clone https://github.com/your_username/Networkmonitor.git
   ```

2. Change to the project directory:

   ```
   cd Networkmonitor
   ```

3. Build the Docker container:

   ```
   docker build -t networkmonitor .
   ```

## Configuration

Networkmonitor can be configured using environment variables. Before running the program, you need to set the following environment variables:

- `CHECK_URL`: The URL you want to use for internet connection checks.
- `CHECK_INTERVAL`: The time interval (in seconds) for checking the internet connection.
- `TASMOTA_DEVICE_IP`: The IP address of your Tasmota-compatible plug.

## Usage

Run the program in a Docker container with the following command, specifying the environment variables:

```
docker run --name networkmonitor_container -e CHECK_URL=your_url -e CHECK_INTERVAL=your_interval -e TASMOTA_DEVICE_IP=your_device_ip networkmonitor
```

The program will now monitor the internet connection by checking the specified URL and perform a power cycle on the Tasmota-compatible plug when an internet failure is detected.

## Notes

- Please ensure that you carefully configure the program and the Tasmota plug IP address to avoid any malfunctions.
- We take no responsibility for any network issues or damages caused by the use of this program.

## Contributing

We welcome contributions to the Networkmonitor project! If you have any suggestions or want to report issues, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](link_to_license). For more information, please see the `LICENSE` file.
