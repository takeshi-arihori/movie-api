import React from 'react'
import { Container, Typography } from '@mui/material'

const SearchPage: React.FC = () => {
  return (
    <Container maxWidth="lg">
      <Typography variant="h4" component="h1" gutterBottom>
        Search
      </Typography>
      <Typography variant="body1">
        Search functionality coming soon...
      </Typography>
    </Container>
  )
}

export default SearchPage