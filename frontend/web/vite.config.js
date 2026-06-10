import { defineConfig } from 'vite'

export default defineConfig({
  plugins: [],
  root: '.',
  publicDir: 'public',
  build: {
    outDir: 'dist',
    emptyOutDir: true,
  },
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/ws': {
        target: 'ws://localhost:8089',
        ws: true,
      },
    },
  },
  test: {
    globals: true,
    environment: 'happy-dom',
  },
})
