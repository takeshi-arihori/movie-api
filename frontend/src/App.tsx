import React from 'react'
import { Routes, Route } from 'react-router-dom'
import { Container, Box, Typography, CircularProgress } from '@mui/material'
import { useSelector, useDispatch } from 'react-redux'

// Components
import Header from '@/components/layout/Header'
import Footer from '@/components/layout/Footer'
import ErrorBoundary from '@/components/common/ErrorBoundary'

// Features
import HomePage from '@/features/home/HomePage'
import SearchPage from '@/features/search/SearchPage'
import MovieDetailsPage from '@/features/movies/MovieDetailsPage'
import TVDetailsPage from '@/features/tv/TVDetailsPage'
import NotFoundPage from '@/features/common/NotFoundPage'

// Types
import type { RootState } from '@/app/store'

// App Component
function App() {
  const dispatch = useDispatch()
  const { isLoading, error } = useSelector((state: RootState) => state.app)

  // Test API connection on mount
  React.useEffect(() => {
    const testApiConnection = async () => {
      try {
        const response = await fetch('/api/health')
        if (response.ok) {
          const data = await response.json()
          console.info('✅ Backend API connected:', data)
        }
      } catch (error) {
        console.warn('⚠️ Backend API connection failed:', error)
      }
    }

    testApiConnection()
  }, [])

  // Global loading state
  if (isLoading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="100vh"
        flexDirection="column"
        gap={2}
      >
        <CircularProgress size={60} />
        <Typography variant="h6" color="text.secondary">
          Loading Movie API...
        </Typography>
      </Box>
    )
  }

  // Global error state
  if (error) {
    return (
      <Container maxWidth="md">
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
          flexDirection="column"
          gap={2}
        >
          <Typography variant="h4" color="error" gutterBottom>
            Something went wrong
          </Typography>
          <Typography variant="body1" color="text.secondary" textAlign="center">
            {error}
          </Typography>
        </Box>
      </Container>
    )
  }

  return (
    <ErrorBoundary>
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          minHeight: '100vh',
        }}
      >
        {/* Header */}
        <Header />

        {/* Main Content */}
        <Box
          component="main"
          sx={{
            flex: 1,
            py: 3,
          }}
        >
          <Routes>
            {/* Home */}
            <Route path="/" element={<HomePage />} />
            
            {/* Search */}
            <Route path="/search" element={<SearchPage />} />
            
            {/* Movie Details */}
            <Route path="/movie/:id" element={<MovieDetailsPage />} />
            
            {/* TV Show Details */}
            <Route path="/tv/:id" element={<TVDetailsPage />} />
            
            {/* 404 Not Found */}
            <Route path="*" element={<NotFoundPage />} />
          </Routes>
        </Box>

        {/* Footer */}
        <Footer />
      </Box>
    </ErrorBoundary>
  )
}

export default App