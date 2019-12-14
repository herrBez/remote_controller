# Remote Controller

Remote controller is a light remote controller for your laptop. It allows
you to increase the volume of your laptop and to pause a video from your
smart phone.

## Original use case

As a lazy netflix watcher I want to be able to skip an intro, increase or decrease the volume, while lying down on my bed.

## Getting Started

Make sure that your laptop and your smart-phone are in the same
network

> :warning: run this utility only in a trusted network, where
you known all the connected devices, i.e., your home network.
Currently, the traffic is neither encrypted nor the
server requires any form of authentication.

### Windows

On your laptop (:computer:):

1. Clone the repository
```sh
git clone https://github.com/herrBez/remote_controller.git
```
2. `cd remote_controller`
3. Download the nircmd utility from this [link](https://www.nirsoft.net/utils/index.html):
either add it to your environment or simply copy the executable in the same same directory
4. Compile the program `go build server.go`
5. Run the `server.exe` executable

On your smartphone (:iphone:):

1. Open your browser on the page `http://<IP-Address-of-your-laptop>:8080`
2. Use your smart-phone to control remotely your pc

## How does it work

The web server responds to the different endpoints by calling
a command in the shell of your laptop. In Windows systems this
utility is [nircmd](https://www.nirsoft.net/utils/index.html). Thus, in windows system it is a wrapper for nircmd.

## License

[MIT License](LICENSE.md)