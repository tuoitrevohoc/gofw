# Go Framework Project

## Development Setup

### Prerequisites
1. Install [Visual Studio Code](https://code.visualstudio.com/)
2. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)
3. Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) in VS Code

### Opening in Dev Container

1. Clone the repository:
   ```bash
   git clone https://github.com/tuoitrevohoc/gofw.git
   cd gofw
   ```

2. Open in VS Code:
   ```bash
   code .
   ```

3. When VS Code opens, you should see a notification in the bottom right corner asking if you want to reopen the project in a Dev Container. Click "Reopen in Container".

   Alternatively, you can:
   - Press `F1` or `Cmd/Ctrl + Shift + P`
   - Type "Dev Containers: Reopen in Container"
   - Press Enter

4. Wait for the container to build. This may take a few minutes the first time.

### What's Included
- Go development environment
- PostgreSQL database
- Redis cache
- Prometheus metrics
- Loki for log aggregation
- Grafana for visualization

### Development URLs
- Backend API: http://localhost:3000
- Grafana Dashboard: http://localhost:3200
- Prometheus: http://localhost:9090

### Troubleshooting
If you encounter any issues:
1. Try rebuilding the container: `F1` > "Dev Containers: Rebuild Container"
2. Check Docker Desktop is running
3. Ensure all ports are available on your machine 