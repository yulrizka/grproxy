import React, {FC} from 'react';
import {List, ListItem, ListItemProps, ListItemText} from "@material-ui/core";
import {Call} from "./Call";

function ListItemLink(props: ListItemProps<'a', { button?: true }>) {
    return <ListItem button divider component="a" {...props}/>;
}

interface ListCallProps {
    calls: Array<Call>
    currentCall:Call
    setCurrentCall: React.Dispatch<React.SetStateAction<Call>>
}

export const ListCall:FC<ListCallProps> = ({calls, currentCall, setCurrentCall}) => {
    const handleClick  = (c:Call) => {
            setCurrentCall(c)
    }

    return (
        <List component="nav" aria-label="secondary mailbox folders">
            {calls.map((call) => {
                return (
                    <ListItemLink key={call.id} onClick={() => handleClick(call)} selected={currentCall.id === call.id}>
                        <ListItemText primary={`Foo.bar.GetUser ${call.id}`}/>
                    </ListItemLink>
                )
            })}
        </List>
    )
}
