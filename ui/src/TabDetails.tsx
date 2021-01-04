import {Box, Button, Tab, Tabs, Typography} from "@material-ui/core";
import React, {FC} from "react";
import {Call} from "./Call";

interface TabPanelProps {
    children?: React.ReactNode;
    index: any;
    value: any;
}

function TabPanel(props: TabPanelProps) {
    const {children, value, index, ...other} = props;

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
        id: `${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
}

interface TabDetailsProps {
    currentCall: Call
}

export const TabDetails:FC<TabDetailsProps> = ({currentCall}) => {
    const [value, setValue] = React.useState(0);
    const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
        setValue(newValue);
    };

    const handleReplyClicked = () => {
        alert("not implemented")
    }
    return (
        <div>
            <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                <Tab label="Request" {...a11yProps(0)} />
                <Tab label="Response" {...a11yProps(1)} />
            </Tabs>
            <TabPanel value={value} index={0}>
                {JSON.stringify(currentCall.request)}
                <br/>
                <Button variant="contained" onClick={handleReplyClicked}>Reply</Button>
            </TabPanel>
            <TabPanel value={value} index={1}>
                {JSON.stringify(currentCall.response)}
            </TabPanel>
        </div>
    )
}
