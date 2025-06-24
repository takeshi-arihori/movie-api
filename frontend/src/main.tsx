import React from 'react'
import ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'
import { BrowserRouter } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
// Development imports
import { ThemeProvider } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'

// App imports
import App from './App'
import { store } from './app/store'
import { theme } from './app/theme'

// Styles
import './styles/index.css'

// Query client configuration
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 5 * 60 * 1000, // 5 minutes
      gcTime: 10 * 60 * 1000, // 10 minutes (formerly cacheTime)
      retry: (failureCount, error) => {
        // Don't retry on 4xx errors
        if (error instanceof Error && 'status' in error && typeof error.status === 'number') {
          if (error.status >= 400 && error.status < 500) {
            return false
          }
        }
        // Retry up to 3 times for other errors
        return failureCount < 3
      },
      refetchOnWindowFocus: false,
      refetchOnReconnect: 'always',
    },
    mutations: {
      retry: 1,
    },
  },
})

// Error boundary for development
if (import.meta.env.DEV) {
  // Enable React DevTools
  ;(window as any).__REACT_DEVTOOLS_GLOBAL_HOOK__?.setSettings?.({
    hideConsoleLogsInStrictMode: false,
  })
}

// Performance monitoring
const startTime = performance.now()

// Main App Component
function MainApp() {
  React.useEffect(() => {
    // Log performance metrics
    const loadTime = performance.now() - startTime
    console.info(`🚀 Movie API loaded in ${loadTime.toFixed(2)}ms`)
    
    // Log environment info
    console.info('🌍 Environment:', {
      mode: import.meta.env.MODE,
      dev: import.meta.env.DEV,
      prod: import.meta.env.PROD,
      apiUrl: import.meta.env.VITE_API_BASE_URL,
    })
  }, [])

  return (
    <React.StrictMode>
      <Provider store={store}>
        <QueryClientProvider client={queryClient}>
          <BrowserRouter>
            <ThemeProvider theme={theme}>
              <CssBaseline />
              <App />
            </ThemeProvider>
          </BrowserRouter>
          
          {/* Development tools */}
          {import.meta.env.DEV && (
            <div style={{ 
              position: 'fixed', 
              bottom: '10px', 
              right: '10px',
              fontSize: '12px',
              color: '#666',
              background: 'rgba(0,0,0,0.1)',
              padding: '4px 8px',
              borderRadius: '4px'
            }}>
              Dev Mode
            </div>
          )}
        </QueryClientProvider>
      </Provider>
    </React.StrictMode>
  )
}

// Mount the app
const rootElement = document.getElementById('root')
if (!rootElement) {
  throw new Error('Root element not found')
}

const root = ReactDOM.createRoot(rootElement)
root.render(<MainApp />)

// Hot module replacement
if (import.meta.hot) {
  import.meta.hot.accept()
}