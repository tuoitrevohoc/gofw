import { ThemeProvider } from './theme/theme-provider'  
import { Box, Typography } from '@mui/material'

function App() {

  return (
    <ThemeProvider>
      <Box>
        <Typography variant="h1">Hello World</Typography>
      </Box>
    </ThemeProvider>
  )
}

export default App
