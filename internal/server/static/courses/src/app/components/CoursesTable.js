import {Button, Card, Table} from "react-bootstrap";
import React from "react";
import PropTypes from 'prop-types'

const CoursesTable = ({loading, courses, onEdit}) => {
        if (!courses || courses.length < 1) {
            return (
                <Card className="text-center">
                    <Card.Body>
                        <Card.Title>Empty list</Card.Title>
                    </Card.Body>
                </Card>
            )
        }
        const content = courses.map((course, index) =>
            <tr key={course.id}>
                <td>{index + 1}</td>
                <td>{course.name}</td>
                <td>{course.description}</td>
                <td>{course.instructor}</td>
                <td>{toPrice(course.centPrice)}</td>
                <td><Button onClick={() => onEdit(course)} disabled={loading}>Edit</Button></td>
            </tr>
        );
        return (
            <Table striped bordered hover>
                <thead>
                <tr>
                    <th>#</th>
                    <th>Name:</th>
                    <th>Description:</th>
                    <th>Instructor:</th>
                    <th>Price:</th>
                    <th>Options:</th>
                </tr>
                </thead>
                <tbody>
                {content}
                </tbody>
            </Table>
        )
    }
;

const toPrice = (centPrice) => {
    const price = centPrice / 100;
    return price.toFixed(2) + ' €'
};

CoursesTable.propTypes = {
    loading: PropTypes.bool.isRequired,
    courses: PropTypes.array.isRequired,
    onEdit: PropTypes.func.isRequired,
};

export default CoursesTable;