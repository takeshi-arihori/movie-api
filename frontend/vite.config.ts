import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      // Enable React Fast Refresh
      fastRefresh: true,
      // Include .jsx files in Fast Refresh
      include: '**/*.{jsx,tsx}',
    }),
  ],
  
  // Path resolution
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
      '@/components': resolve(__dirname, './src/components'),
      '@/features': resolve(__dirname, './src/features'),
      '@/hooks': resolve(__dirname, './src/hooks'),
      '@/stores': resolve(__dirname, './src/stores'),
      '@/types': resolve(__dirname, './src/types'),
      '@/utils': resolve(__dirname, './src/utils'),
      '@/app': resolve(__dirname, './src/app'),
    },
  },
  
  // Development server configuration
  server: {
    port: 3000,
    host: true, // Listen on all addresses
    strictPort: true, // Exit if port is already in use
    open: false, // Don't open browser automatically
    cors: true,
    proxy: {
      // Proxy API requests to backend
      '/api': {
        target: process.env.VITE_API_BASE_URL || 'http://localhost:8080',
        changeOrigin: true,
        secure: false,
        ws: true, // Enable WebSocket proxying
      },
    },
  },
  
  // Build configuration
  build: {
    target: 'esnext',
    outDir: 'dist',
    assetsDir: 'assets',
    sourcemap: true,
    minify: 'esbuild',
    
    // Rollup options
    rollupOptions: {
      output: {
        manualChunks: {
          // Vendor chunk for large libraries
          vendor: ['react', 'react-dom'],
          mui: ['@mui/material', '@mui/icons-material'],
          router: ['react-router-dom'],
          redux: ['@reduxjs/toolkit', 'react-redux'],
        },
      },
    },
    
    // Performance optimizations
    chunkSizeWarningLimit: 1000,
  },
  
  // CSS configuration
  css: {
    modules: {
      localsConvention: 'camelCase',
    },
    preprocessorOptions: {
      scss: {
        additionalData: `@import \"@/styles/variables.scss\";`,
      },
    },
  },
  
  // Testing configuration
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: ['./src/test/setup.ts'],
    css: true,
  },
  
  // Environment variables
  define: {
    __DEV__: JSON.stringify(process.env.NODE_ENV === 'development'),
    __PROD__: JSON.stringify(process.env.NODE_ENV === 'production'),
  },
  
  // Optimization
  optimizeDeps: {
    include: [
      'react',
      'react-dom',
      'react-router-dom',
      '@mui/material',
      '@reduxjs/toolkit',
      'react-redux',
    ],
  },
})