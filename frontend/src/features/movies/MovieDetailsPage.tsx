import React from 'react'
import { Container, Typography } from '@mui/material'

const MovieDetailsPage: React.FC = () => {
  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Movie Details
      </Typography>
      <Typography variant="body1">
        Movie details page coming soon...
      </Typography>
    </Container>
  )
}

export default MovieDetailsPage