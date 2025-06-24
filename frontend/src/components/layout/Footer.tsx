import React from 'react'
import { Box, Typography } from '@mui/material'

const Footer: React.FC = () => {
  return (
    <Box component="footer" sx={{ py: 2, px: 2, mt: 'auto', backgroundColor: 'grey.100' }}>
      <Typography variant="body2" color="text.secondary" align="center">
        © 2025 Movie API. Powered by TMDb.
      </Typography>
    </Box>
  )
}

export default Footer