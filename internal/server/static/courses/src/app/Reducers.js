import {
    ADD_TO_COURSES_LIST,
    RESET_COURSE_FORM,
    SET_COURSE_CENT_PRICE,
    SET_COURSE_DESCRIPTION, SET_COURSE_FORM,
    SET_COURSE_INSTRUCTOR,
    SET_COURSE_NAME, SET_COURSES_LIST,
    SET_LOADING, SET_MESSAGE
} from "./Actions";
import {combineReducers} from "redux";

export const INITIAL_GLOBAL_STATE = {
    loading: false,
    message: null,
};

export const globalStateReducer = (state = INITIAL_GLOBAL_STATE, action) => {
    switch (action.type) {
        case SET_LOADING:
            return {...state, loading: action.loading};
        case SET_MESSAGE:
            return {...state, message: action.message};
        default:
            return state;
    }
};

export const INITIAL_COURSE_FORM_STATE = {
    id: null,
    name: '',
    description: '',
    instructor: '',
    centPrice: 0
};

export const courseFormStateReducer = (state = INITIAL_COURSE_FORM_STATE, action) => {
    switch (action.type) {
        case SET_COURSE_NAME:
            return {...state, name: action.name};
        case SET_COURSE_DESCRIPTION:
            return {...state, description: action.description};
        case SET_COURSE_INSTRUCTOR:
            return {...state, instructor: action.instructor};
        case SET_COURSE_CENT_PRICE:
            return {...state, centPrice: action.centPrice};
        case SET_COURSE_FORM:
            return action.course;
        case RESET_COURSE_FORM:
            return INITIAL_COURSE_FORM_STATE;
        default:
            return state;
    }
};

export const INITIAL_COURSES_LIST_STATE = [];

export const coursesListStateReducer = (state = INITIAL_COURSES_LIST_STATE, action) => {
    switch (action.type) {
        case SET_COURSES_LIST:
            return action.courses;
        case ADD_TO_COURSES_LIST:
            const index = state.findIndex((course) => course.id === action.add.id);
            if (index !== -1) {
                const copy = [...state];
                copy[index] = action.add;
                return copy;
            }
            return [...state, action.add];
        default:
            return state;
    }
};

export const rootReducer = combineReducers({
    global: globalStateReducer,
    courses: coursesListStateReducer,
    form: courseFormStateReducer,
});