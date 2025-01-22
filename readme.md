# Server Monitor Dashboard

This project is a server monitoring dashboard that collects and displays various system metrics. It is built using Go for the backend and React with TypeScript for the frontend. The project includes the following key components:  
- **Agent**: Collects system metrics and sends them to the central server through Apache Kafka.
- **Backend**: Collects system metrics such as CPU usage, memory usage, disk usage, network statistics, and system information.
- **Frontend**: Displays the collected metrics in a user-friendly dashboard.

## Features

- **CPU Metrics**: Collects context switches, interrupts, and load averages.
- **Disk Usage**: Displays total, free, used space, and used percentage.
- **Network Info**: Shows network interface statistics including bytes sent/received, packets sent/received, errors, and drops.
- **System Info**: Provides details about the host such as hostname, OS, platform version, kernel version, uptime, and boot time.
- **Process Monitoring**: Lists top processes by CPU and memory usage.
- **Temperature Monitoring**: Reads system temperatures from thermal zones.
## Project Structure
  ### Agent (Go)
- **goServerAgent/CollectMetrics.go**: This file contains the logic for collecting various system metrics using the gopsutil library. It gathers CPU times, context switches, interrupts, load averages, disk usage, network information, and system information.  
- **go.mod**: The Go module file that defines the module path and lists the dependencies required for the project. 
### Agent (Go)
- **backend/**: This directory contains the backend code responsible for receiving metrics from the agent, storing them in the database, and serving them to the frontend
- **go.mod**: The Go module file that defines the module path and lists the dependencies required for the project.
### Frontend (React with TypeScript)
- **server-monitor-dashboard/**: This directory contains the frontend code for the server monitoring dashboard.
- **package.json**: Defines the dependencies and scripts for the frontend project.  
- **src/**: Contains the source code for the React application.  
- **components/**: Contains React components used in the application.
- **App.tsx**: The main application component.
- **index.tsx**: The entry point for the React application.

## Prerequisites

- Go 1.16 or later
- Node.js and npm
- lm-sensors (for temperature monitoring on Linux)

## Setup

1. **Clone the repository**:
    ```sh
    git clone https://github.com/Niladri2003/server-monitor.git
    cd server-monitor-dashboard
    ```

2. **Backend Setup**:
    - Navigate to the `goServerAgent` directory:
        ```sh
        cd goServerAgent
        ```
    - Install dependencies:
        ```sh
        go mod tidy
        ```
    - Run the server:
        ```sh
        go run main.go
        ```

3. **Frontend Setup**:
    - Navigate to the `client` directory:
        ```sh
        cd ../client
        ```
    - Install dependencies:
        ```sh
        npm install
        ```
    - Start the development server:
        ```sh
        npm start
        ```

## Usage

- The backend server will start collecting metrics and expose them via an API.
- The frontend React application will display the collected metrics in a user-friendly dashboard.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License.
