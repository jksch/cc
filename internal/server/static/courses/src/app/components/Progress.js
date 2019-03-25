import React from "react";
import {ProgressBar} from "react-bootstrap";
import PropTypes from 'prop-types'

const Progress = ({loading}) => {
    if (!loading) {
        return null;
    }
    return (
        <ProgressBar animated now={100}/>
    )
};

Progress.propTypes = {
    loading: PropTypes.bool.isRequired
};

export default Progress;