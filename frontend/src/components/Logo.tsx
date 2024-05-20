import { Typography, useTheme } from "@mui/material";

export default function Logo() {
  const theme = useTheme();
  const info = theme.palette.info.main

  return (
    <Typography
      variant="h5"
      noWrap
      component="a"
      href="/"
      fontWeight={700}
      sx={{
        color: info,
        mr: 2,
        display: "flex",
        flexGrow: 0,
        letterSpacing: '.1rem',
        textDecoration: 'none',
      }}
    >
      The Weekly Byte
    </Typography>
  );
}