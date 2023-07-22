# Networkmonitor - Internet Failure Detection and Powercycle

[Networkmonitor - Internet Failure Detection and Powercycle](https://hub.docker.com/r/marckumar/networkmonitor)

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

Networkmonitor is a lightweight Go program designed to detect internet failures by periodically checking a user-defined URL. When an internet failure is detected, it can automatically trigger a power cycle on a Tasmota-compatible plug connected to the network. This automation eliminates the need for manual intervention and helps to quickly resolve common network issues.

## Requirements

To use Networkmonitor, you need the following:

1. Go installed on your computer or server.
2. A Tasmota-compatible plug that can be controlled over the local network.

## Installation

You have two options for installation:

**1. Build from Source:**

- Clone the GitHub repository to your computer or server:

  ```
  git clone https://github.com/MarcKumar/networkmonitor.git
  ```

- Change to the project directory:

  ```
  cd networkmonitor
  ```

- Build the program:

  ```
  go build
  ```

**2. Using Docker:**

- Pull the Networkmonitor Docker image from DockerHub. The image is small, and when unpacked, it is approximately 7MB in size:

  ```
  docker pull marckumar/networkmonitor
  ```

- The container requires less than 3MB of RAM to run efficiently, making it lightweight and suitable for resource-constrained environments.

## Configuration

Networkmonitor can be configured using command-line arguments. Before running the program, you need to set the following parameters:

- `check_url`: The URL you want to use for internet connection checks.
- `check_interval`: The time interval for checking the internet connection. Examples: "5s" (5 seconds), "1m" (1 minute), "2h30m" (2 hours and 30 minutes).
- `tasmota_device_ip`: The IP address of your Tasmota-compatible plug.

## Usage

You can run the program with command-line arguments as follows:

```
./Networkmonitor -check_url your_url -check_interval your_interval -tasmota_ip your_device_ip
```

If you are using Docker, you can pass the parameters using environment variables:

```
docker run --name networkmonitor_container -e check_url=your_url -e check_interval=your_interval -e tasmota_ip=your_device_ip marckumar/networkmonitor
```

The program will now monitor the internet connection by checking the specified URL and perform a power cycle on the Tasmota-compatible plug when an internet failure is detected.

## Notes

- Please ensure that you carefully configure the program and the Tasmota plug IP address to avoid any malfunctions.
- We take no responsibility for any network issues or damages caused by the use of this program.

## Contributing

We welcome contributions to the Networkmonitor project! If you have any suggestions or want to report issues, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](https://github.com/MarcKumar/networkmonitor/blob/master/MIT_LICENSE.txt). For more information, please see the `LICENSE` file.

---

Thank you for your interest in our Networkmonitor - Internet Failure Detection and Powercycle project. If you have any questions or need assistance, feel free to contact us. Happy monitoring and using!
