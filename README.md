# Raspi-TV-Control

[![Compilation check](https://github.com/Binozo/Raspi-TV-Control/actions/workflows/compile-check.yaml/badge.svg)](https://github.com/Binozo/Raspi-TV-Control/actions/workflows/compile-check.yaml)

A HTTP Server to control a TV via CEC.
Specifically designed for the [Raspi-TV-Control-App](https://github.com/Binozo/Raspi_TV_Control_App).

## Install
```bash
wget https://raw.githubusercontent.com/Binozo/Raspi-TV-Control/master/install.sh -O install.sh && bash install.sh
```

## HTTP Endpoints
### Note: The HTTP Server runs on port 4828
### `/`
Returns some basic info.
Example:
```json
{
  "LibInfo": "compiled on Linux-5.10.63-v8+ ... , features: P8_USB, DRM, P8_detect, randr, RPi, Exynos, Linux, AOCEC",
  "Note": "Target TV has been detected",
  "OS": "linux",
  "Uptime": "23121.426173",
  "Uptime_Date": "2022-08-19 16:12:12",
  "Version": "0.1"
}
```


