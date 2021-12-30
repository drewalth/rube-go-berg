# rube-go-berg

A small [Rube Goldberg machine](https://en.wikipedia.org/wiki/Rube_Goldberg_machine
) to get acquainted with Go and experiment with GitHub Webhooks and play with a neglected Raspberry Pi. The idea is to show some sort of indicator, in addition to the small red 'x' in the Github UI, when an app's build fails.

I'm using a raspberry pi and an awful-sounding buzzer.

![Image](./public/pi.png)

## Requirements

- Raspberry Pi
- Some sort of Buzzer or LED module for Pi
- GPIO breadboard, circuit, or some other means of connecting module
- Golang
- Text Editor
- Some sort of CI pipeline/workflow. I'm using Travis CI.

## Getting Started

- Clone the project
- Install dependencies

```bash
go get ./...
```

- Update variables.

```go
// L46-47
username := "drewalth"             // change to whatever GitHub username
buildErrorTitle := "Build Errored" // change to what your CI titles failed builds
```

- Build circuit with module of choice, and update GPIO Pin if needed.

```go
// L31
pin := rpio.Pin(17)
```

- Upload project files to Pi
- Install Go on Pi

```bash
sudo apt-get install golang
```

- Install dependencies

```bash
go get ./...
```

- Run

```go
go run main.go
```

- Setup [GitHub Webhook](https://docs.github.com/en/developers/webhooks-and-events/webhooks/about-webhooks)
  - Select `check_run`
  - Point to `<my_ip_address>:5000/check-run`
  - Note. Chances are you will need to do some port forwarding on your router to reach the Pi.

## Notes

- You could do all the cloning/editing directly on the Pi if you want to. Just need to install Git and do editing with vim or nano.
