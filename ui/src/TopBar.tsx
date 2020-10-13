import {AppBar, createStyles, IconButton, Theme, Toolbar, Typography} from "@material-ui/core";
import MenuIcon from "@material-ui/icons/Menu";
import React from "react";
import {makeStyles} from "@material-ui/core/styles";

export const TopBar = () => {
    const useStyles = makeStyles((theme: Theme) =>
        createStyles({
            root: {
                flexGrow: 1,
                height: '100vh',
            },
            menuButton: {
                marginRight: theme.spacing(2),
            },
            title: {
                flexGrow: 1,
            },
            paper: {
                padding: theme.spacing(0),
                textAlign: 'center',
                color: theme.palette.text.secondary,
            },
        }),
    );
    const classes = useStyles();

    return(
        <AppBar position="static">
        <Toolbar>
            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
        <MenuIcon />
        </IconButton>
        <Typography variant="h6" className={classes.title}>
        grproxy - Inspect
        </Typography>
        </Toolbar>
        </AppBar>
)
}
