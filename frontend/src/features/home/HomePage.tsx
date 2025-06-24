import React from 'react'
import { Container, Typography, Box } from '@mui/material'

const HomePage: React.FC = () => {
  return (
    <Container maxWidth="lg">
      <Box textAlign="center" py={8}>
        <Typography variant="h2" component="h1" gutterBottom>
          Welcome to Movie API
        </Typography>
        <Typography variant="h6" color="text.secondary">
          Discover movies and TV shows powered by TMDb
        </Typography>
      </Box>
    </Container>
  )
}

export default HomePage