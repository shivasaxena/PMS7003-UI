# PMS7003-UI

A real-time air quality monitoring application with a graphical user interface for the PMS7003 particulate matter sensor. Built with Go and Fyne, this application displays PM1.0, PM2.5, and PM10.0 readings with color-coded air quality indicators.

> **📚 This project serves as a practical how-to-use guide and reference implementation for the [PMS7003 Go driver](https://github.com/shivasaxena/PMS7003).** If you're looking to integrate the PMS7003 sensor into your own Go applications, this project demonstrates how to do it.

![Application Demo](docs/static_files/screenshots/2025-10-20%2014-14-04.gif)

## Features

- **Real-time Monitoring**: Continuously reads and displays air quality data from PMS7003 sensor
- **Visual Feedback**: Color-coded display based on air quality index (AQI) standards:
  - 🟢 Green (0-50): Good
  - 🟡 Yellow (51-100): Moderate
  - 🟠 Orange (101-150): Unhealthy for Sensitive Groups
  - 🔴 Red (151-200): Unhealthy
  - 🟣 Purple (201-300): Very Unhealthy
  - 🔴 Maroon (300+): Hazardous
- **Multiple PM Readings**: Displays PM1.0, PM2.5, and PM10.0 atmospheric concentrations
- **Fullscreen Mode**: Toggle fullscreen using the toolbar button or `Ctrl+F11`
- **Clean UI**: Modern, responsive interface built with Fyne toolkit

## Prerequisites

- Go 1.22.0 or higher
- PMS7003 particulate matter sensor connected to your device
- Graphics libraries (for Linux)

## Installation

### Linux

1. Clone the repository:
```bash
git clone https://github.com/shivasaxena/PMS7003-UI.git
cd PMS7003-UI
```

2. Install required graphics libraries:
```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev
```

3. Build the application:
```bash
go build
```

### Windows

1. Clone the repository:
```powershell
git clone https://github.com/shivasaxena/PMS7003-UI.git
cd PMS7003-UI
```

2. Build the application:
```powershell
go build
```

## Usage

### Basic Usage

Run the application with default settings (device: `/dev/ttyAMA0`):

```bash
./PMS7003-UI
```

### Command Line Options

- **`-device`**: Specify the serial device name (default: `/dev/ttyAMA0`)
  ```bash
  ./PMS7003-UI -device /dev/ttyUSB0
  ```

- **`-fullScreen`**: Launch the application in fullscreen mode (default: `false`)
  ```bash
  ./PMS7003-UI -fullScreen
  ```

### Example

```bash
./PMS7003-UI -device /dev/ttyUSB0 -fullScreen
```

### Controls

- **Fullscreen Toggle**: Click the fullscreen icon in the toolbar or press `Ctrl+F11`

## Dependencies

This project uses the following Go modules:

- [fyne.io/fyne/v2](https://github.com/fyne-io/fyne) - Cross-platform GUI toolkit
- [github.com/shivasaxena/PMS7003](https://github.com/shivasaxena/PMS7003) - PMS7003 sensor driver

## Project Structure

```
PMS7003-UI/
├── main.go              # Main application code
├── go.mod              # Go module dependencies
├── go.sum              # Dependency checksums
├── README.md           # Project documentation
└── docs/
    └── static_files/
        └── screenshots/ # Application screenshots
```

## How It Works

This application demonstrates the practical usage of the [PMS7003 Go driver](https://github.com/shivasaxena/PMS7003):

1. **Initialize the sensor**: Opens a connection to the PMS7003 sensor in active mode using `PMS7003.Open()`
2. **Background reading**: Reads sensor data every second in a separate goroutine using `device.Read()`
3. **UI updates**: Updates the display with the latest PM1.0, PM2.5, and PM10.0 atmospheric readings
4. **Visual feedback**: Applies color coding based on the AQI values for easy interpretation
5. **Data presentation**: Values are displayed in μg/m³ (micrograms per cubic meter)

### Code Example

Here's the core implementation showing how to use the PMS7003 driver:

```go
// Open the sensor device in active mode
device, err := PMS7003.Open("/dev/ttyAMA0", PMS7003.ActiveMode)
if err != nil {
    panic(err)
}
defer device.Close()

// Read sensor values continuously
for range time.Tick(time.Second) {
    sensorValue, err := device.Read()
    if err != nil {
        panic(err)
    }
    
    // Access the atmospheric concentration values
    pm10 := sensorValue.PM10Atmospheric
    pm25 := sensorValue.PM25Atmospheric
    pm100 := sensorValue.PM100Atmospheric
}
```

## Development

### Building from Source

```bash
go mod download
go build -o PMS7003-UI
```

### Running in Development Mode

```bash
go run main.go -device /dev/ttyAMA0
```

## Troubleshooting

### Permission Denied Error

If you encounter permission issues accessing the serial device:

```bash
sudo chmod 666 /dev/ttyAMA0
# or add your user to the dialout group
sudo usermod -a -G dialout $USER
```

### Device Not Found

Ensure your PMS7003 sensor is properly connected and the device path is correct. Common device paths:
- Raspberry Pi: `/dev/ttyAMA0` or `/dev/serial0`
- USB adapter: `/dev/ttyUSB0` or `/dev/ttyUSB1`
- Windows: `COM3`, `COM4`, etc.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source. Please check the repository for license details.

## Author

**Shiva Saxena**
- GitHub: [@shivasaxena](https://github.com/shivasaxena)

## Learning from This Project

This project is designed to help you understand how to:
- Initialize and configure the PMS7003 sensor
- Read data continuously from the sensor
- Handle sensor data in Go applications
- Build responsive UIs for sensor monitoring
- Implement proper error handling and resource cleanup

Whether you're building an IoT project, environmental monitoring system, or just learning about sensor integration in Go, this project provides a complete, working example.

## Related Projects

- **[PMS7003 Driver](https://github.com/shivasaxena/PMS7003)** - The underlying Go driver library that this UI demonstrates

## Acknowledgments

- [Fyne](https://fyne.io/) - For the excellent cross-platform GUI toolkit
