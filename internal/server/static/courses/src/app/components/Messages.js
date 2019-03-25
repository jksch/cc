import React from "react";
import {Alert} from "react-bootstrap";

const Messages = ({message}) => {
    if (!message){
        return null;
    }
    return(
        <Alert variant={'primary'}>
            {message}
        </Alert>
    )
};

export default Messages;