import axios from "axios";
import * as messages from './Messages.js'

export const Client = axios.create({timeout: 10000});

export const SET_LOADING = 'SET_LOADING';
export const SET_MESSAGE = 'SET_MESSAGE';
export const setLoading = (bool) => {
    return {
        type: SET_LOADING,
        loading: bool
    }
};

export const setMessage = (message) => {
    return {
        type: SET_MESSAGE,
        message: message
    }
};

export const SET_COURSE_NAME = 'SET_COURSE_NAME';
export const SET_COURSE_DESCRIPTION = 'SET_COURSE_DESCRIPTION';
export const SET_COURSE_INSTRUCTOR = 'SET_COURSE_INSTRUCTOR';
export const SET_COURSE_CENT_PRICE = 'SET_COURSE_CENT_PRICE';
export const SET_COURSE_FORM = 'SET_COURSE_FORM';
export const RESET_COURSE_FORM = 'RESET_COURSE_FORM'

export const setCourseName = (name) => {
    return {
        type: SET_COURSE_NAME,
        name: name
    }
};

export const setCourseDescription = (description) => {
    return {
        type: SET_COURSE_DESCRIPTION,
        description: description
    }
};

export const setCourseInstructor = (instructor) => {
    return {
        type: SET_COURSE_INSTRUCTOR,
        instructor: instructor
    }
};

export const setCourseCentPrice = (centPrice) => {
    return {
        type: SET_COURSE_CENT_PRICE,
        centPrice: centPrice
    }
};

export const setCourseForm = (course) => {
    return {
        type: SET_COURSE_FORM,
        course: course
    }
};

export const resetCourseForm = () => {
    return {
        type: RESET_COURSE_FORM
    }
};

export const SET_COURSES_LIST = 'SET_COURSES_LIST';
export const ADD_TO_COURSES_LIST = 'ADD_TO_COURSES_LIST';

export const setCoursesList = (courses) => {
    return {
        type: SET_COURSES_LIST,
        courses: courses
    }
};

export const addToCoursesList = (course) => {
    return {
        type: ADD_TO_COURSES_LIST,
        add: course
    }
};

export const getCoursesList = () => (dispatch) => {
    dispatch(setLoading(true));

    return Client.get('api/courses').then((res) => {
        dispatch(setLoading(false));

        if (res.status !== 200){
            dispatch(setMessage(messages.ERR_GET_COURSES_LIST_FAILED));
            dispatch(setCoursesList([]));
            return res;
        }

        if (res.data){
            dispatch(setCoursesList(res.data));
        }

        return res;
    }).catch((error) => {

        dispatch(setLoading(false));
        dispatch(setMessage(messages.ERR_GET_COURSES_LIST_FAILED));
        dispatch(setCoursesList([]));

        return error
    });
};


export const postCourse = course => (dispatch) => {
    dispatch(setLoading(true));

    return Client.post('api/courses', course).then((res) => {
        dispatch(setLoading(false));

        if (res.status !== 201){
            dispatch(setMessage(messages.ERR_POST_COURSE_FAILED));
            return res;
        }

        dispatch(setMessage(messages.POST_COURSE_SUCCESSFUL));
        dispatch(addToCoursesList({...course, id: Number(res.data)}));
        dispatch(resetCourseForm());
        return res;
    }).catch((error) =>{
        dispatch(setLoading(false));
        dispatch(setMessage(messages.ERR_POST_COURSE_FAILED));
        return error;
    })
};

