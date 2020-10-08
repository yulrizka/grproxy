import React from 'react';
import './App.css';
import {
    AppBar,
    Box,
    createMuiTheme,
    createStyles, Divider,
    Grid,
    IconButton,
    List,
    ListItem,
    ListItemProps,
    ListItemText,
    MuiThemeProvider,
    Tab,
    Tabs,
    Theme,
    Toolbar,
    Typography
} from "@material-ui/core";
import MenuIcon from '@material-ui/icons/Menu';
import {makeStyles} from "@material-ui/core/styles";

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

interface TabPanelProps {
    children?: React.ReactNode;
    index: any;
    value: any;
}

function TabPanel(props: TabPanelProps) {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`simple-tabpanel-${index}`}
            aria-labelledby={`simple-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box p={3}>
                    <Typography>{children}</Typography>
                </Box>
            )}
        </div>
    );
}

function a11yProps(index: any) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

function ListItemLink(props: ListItemProps<'a', { button?: true }>) {
    return <ListItem button divider component="a" {...props} />;
}


function App() {
    const classes = useStyles();
    const darkTheme = createMuiTheme({
        palette: {
            type: 'light',
        },
    });
    const [value, setValue] = React.useState(0);
    const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
        setValue(newValue);
    };

    return (
        <MuiThemeProvider theme={darkTheme}>
            <AppBar position="static">
                <Toolbar>
                    <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
                        <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" className={classes.title}>
                        grproxy
                    </Typography>
                </Toolbar>
            </AppBar>
            <Grid container spacing={0} alignItems={"stretch"} style={{}}>
                <Grid item xs={3} style={{overflow: "auto", height:"100vh"}}>
                    <List component="nav" aria-label="secondary mailbox folders">
                        <ListItem button>
                            <ListItemText primary="Trash" />
                        </ListItem>
                        <Divider/>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <Divider/>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                    </List>
                </Grid>
                <Grid item xs={9}>
                    <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                        <Tab label="Request" {...a11yProps(0)} />
                        <Tab label="Response" {...a11yProps(1)} />
                    </Tabs>
                    <TabPanel value={value} index={0}>
                        Request
                    </TabPanel>
                    <TabPanel value={value} index={1}>
                        <pre style={{overflow: "auto", height:"100vh"}}>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        <ListItemLink href="#simple-list">
                            <ListItemText primary="Spam" />
                        </ListItemLink>
                        </pre>
                    </TabPanel>
                </Grid>
            </Grid>
        </MuiThemeProvider>
  );
}

export default App;
