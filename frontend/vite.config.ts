import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: true, // Listen on all available network interfaces
    port: 3000,
    strictPort: true, // Fail if port is already in use
  }
})
