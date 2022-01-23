import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#ff8538",
    },
    secondary: {
      main: "#ebebeb",
    },
    mode: "dark",
  },
  components: {
    MuiAppBar: {
      defaultProps: {
        enableColorOnDark: true,
      },
    },
  },
});

export default theme;