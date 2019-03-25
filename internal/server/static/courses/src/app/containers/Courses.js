import React from 'react';
import {Container, Navbar} from "react-bootstrap";
import PropTypes from 'prop-types'
import CoursesForm from "./CoursesForm";
import CoursesTable from "../components/CoursesTable";
import Progress from "../components/Progress";
import Messages from "../components/Messages";
import {getCoursesList, setCourseForm} from "../Actions";
import {connect} from "react-redux";

class Courses extends React.Component {
    static propTypes = {
        dispatch: PropTypes.func.isRequired
    };

    componentDidMount() {
        const {dispatch} = this.props;
        dispatch(getCoursesList())
    }

    render() {
        const {loading, message, courses} = this.props;
        return (
            <div>
                <Navbar bg="primary" variant="dark">
                    <Navbar.Brand>Courses App</Navbar.Brand>
                </Navbar>
                <div className="mb-3"/>

                <Container>
                    <Progress loading={loading}/>
                    <div className="mb-3"/>

                    <Messages message={message}/>
                    <div className="mb-3"/>

                    <CoursesTable loading={loading} courses={courses} onEdit={this.handleEdit}/>
                    <div className="mb-3"/>

                    <CoursesForm/>
                </Container>
            </div>
        )
    }

    handleEdit = course => {
        this.props.dispatch(setCourseForm(course))
    }
}

const mapStateToProps = state => {
    const {global, courses} = state;
    const{loading, message} = global;
    return{
        loading,
        message,
        courses,
    }
};

export default connect(mapStateToProps)(Courses);
