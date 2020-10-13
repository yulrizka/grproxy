import React, {FC, useEffect, useState} from 'react';
import './App.css';
import {createMuiTheme, Grid, MuiThemeProvider} from "@material-ui/core";
import {ListCall} from './ListCall';
import {TopBar} from "./TopBar";
import {TabDetails} from "./TabDetails";
import {Call} from "./Call";


export const App: FC = () => {
    const darkTheme = createMuiTheme({
        palette: {
            type: 'light',
        },
    });

    let [calls, setCalls] = useState<Array<Call>>([{id:1,request:{request_id:1}, response:{response_id:1}}])
    let [currentCall, setCurrentCall] = useState<Call>(calls[0])

    useEffect(() => {
        // component mount
        let i = 1;
        setInterval(()=>{
            if (i < 50) {
                i++
                setCalls(old => [...old, {id:i, request:{request_id:i}, response:{response_id:i}}])
            }
        },4000)
    }, [])

    return (
        <MuiThemeProvider theme={darkTheme}>
            <TopBar/>
            <Grid container spacing={0} alignItems={"stretch"} style={{}}>
                <Grid item xs={3} style={{overflow: "auto", height: "100vh"}}>
                    <ListCall calls={calls} currentCall={currentCall} setCurrentCall={setCurrentCall}/>
                </Grid>
                <Grid item xs={9}>
                    <TabDetails currentCall={currentCall}/>
                </Grid>
            </Grid>
        </MuiThemeProvider>
    );
}

