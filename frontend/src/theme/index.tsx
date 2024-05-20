import { createTheme } from "@mui/material/styles";

const grey = {
  0: '#FFFFFF',
  100: '#F9FAFB',
  200: '#EBECEE',
  300: '#D7D8DC',
  400: '#C4C5CB',
  500: '#9C9FA8',
  600: '#616574',
  700: '#393E51',
  800: '#282B39',
  900: '#171920',
}

const primary = {
  lighter: "#C7D0FA",
  light: "#8EA0F5",
  main: "#5671F0",
  dark: "#364EBE",
  darker: "#223177",
  contrastText: grey[0]
}

const secondary = {
  lighter: "#D5D5EC",
  light: "#B6B6DD",
  main: "#9796CF",
  dark: "#6A6991",
  darker: "#3C3C53",
  contrastText: grey[0]
}

const info = {
  lighter: "#C9EEF2",
  light: "#A1E2E9",
  main: "#78D5DF",
  dark: "#54959C",
  darker: "#305559",
  contrastText: grey[700]
}

const warning = {
  lighter: "#FFEAC5",
  light: "#FFDB9A",
  main: "#FFCB6F",
  dark: "#B38E4E",
  darker: "#66512C",
  contrastText: grey[700]
}

const success = {
  lighter: "#DBF2E7",
  light: "#C1E9D6",
  main: "#A6DFC4",
  dark: "#749C89",
  darker: "#42594E",
  contrastText: grey[700]
}

const error = {
  lighter: "#FCC8C8",
  light: "#FA9E9E",
  main: "#F87575",
  dark: "#AE5252",
  darker: "#632F2F",
  contrastText: grey[0]
}

export const theme = createTheme({
  palette: {
    primary,
    secondary,
    info,
    warning,
    success,
    error,
    grey
  },
  typography: {
    allVariants:{
      fontFamily: [
        '"Poppins"',
        'sans-serif'
      ].join(',')
    }
  },
  components: {
    MuiButtonBase: {
      defaultProps: {
        style: {
          textTransform: "none"
        }
      }
    }
  }
});

