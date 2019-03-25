import React from 'react';
import {Button, Form} from "react-bootstrap";
import PropTypes from 'prop-types'
import {connect} from "react-redux";
import {postCourse, setCourseCentPrice, setCourseDescription, setCourseInstructor, setCourseName} from "../Actions";

const CoursesForm = ({dispatch, form, loading}) => {
    return (
        <Form onSubmit={(e) => {
            e.preventDefault();
            dispatch(postCourse(form))
        }}>
            <Form.Group controlId="formCoursesName">
                <Form.Label>Name:</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Name"
                    value={form.name}
                    onChange={(e) => dispatch(setCourseName(e.target.value))}
                    required
                />
            </Form.Group>

            <Form.Group controlId="formCoursesDescription">
                <Form.Label>Description:</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Description"
                    value={form.description}
                    onChange={(e) => dispatch(setCourseDescription(e.target.value))}
                    required
                />
            </Form.Group>

            <Form.Group controlId="formCoursesInstructor">
                <Form.Label>Instructor:</Form.Label>
                <Form.Control
                    type="text"
                    placeholder="Instructor"
                    value={form.instructor}
                    onChange={(e) => dispatch(setCourseInstructor(e.target.value))}
                    required
                />
            </Form.Group>

            <Form.Group controlId="formCoursesPrice">
                <Form.Label>Price:</Form.Label>
                <Form.Control
                    type="number"
                    step="0.01"
                    placeholder="Price"
                    value={toPrice(form.centPrice)}
                    onChange={(e) => dispatch(setCourseCentPrice(toCentPrice(e.target.value)))}
                />
            </Form.Group>

            <Button disabled={loading} variant="primary" type="submit">
                Submit
            </Button>
        </Form>
    )
};

CoursesForm.propTypes = {
    dispatch: PropTypes.func.isRequired
};

const toPrice = (centPrice) => {
    return (centPrice / 100)
};

const toCentPrice = (price) => {
    return price * 100;
};

const mapStateToProps = state => {
    const {global, form} = state;
    const {loading} = global;
    return {
        loading,
        form
    }
};


export default connect(mapStateToProps)(CoursesForm);
