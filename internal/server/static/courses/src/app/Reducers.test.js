import * as actions from './Actions.js'
import * as reducers from './Reducers.js'
import {INITIAL_COURSE_FORM_STATE} from "./Reducers";

describe('Test global reducers', () => {
    it('globalReducer should handle initial state', () => {
        expect(reducers.globalStateReducer(undefined, {})).toEqual({
            loading: false,
            message: null
        })
    });
    it('globalReducer should handle setLoading', () => {
        expect(reducers.globalStateReducer(reducers.INITIAL_GLOBAL_STATE, actions.setLoading(true)))
            .toEqual({
                loading: true,
                message: null
            })
    });
    it('globalReducer should handle setError', () => {
        expect(reducers.globalStateReducer(reducers.INITIAL_GLOBAL_STATE, actions.setMessage('REQUEST_FAILED')))
            .toEqual({
                loading: false,
                message: 'REQUEST_FAILED'
            })
    });
});

describe('Test course form reducers', () => {
    it('coursesForm should handle initial state', () => {
        expect(reducers.courseFormStateReducer(undefined, {})).toEqual({
            id: null,
            name: '',
            description: '',
            instructor: '',
            centPrice: 0
        })
    });
    it('coursesForm should handle SET_COURSE_NAME action', () => {
        expect(reducers.courseFormStateReducer(reducers.INITIAL_COURSE_FORM_STATE, actions.setCourseName('Go'))).toEqual({
            id: null,
            name: 'Go',
            description: '',
            instructor: '',
            centPrice: 0
        })
    });
    it('coursesForm should handle SET_COURSE_DESCRIPTION action', () => {
        expect(reducers.courseFormStateReducer(reducers.INITIAL_COURSE_FORM_STATE, actions.setCourseDescription('Go course'))).toEqual({
            id: null,
            name: '',
            description: 'Go course',
            instructor: '',
            centPrice: 0
        })
    });
    it('coursesForm should handle SET_COURSE_INSTRUCTOR action', () => {
        expect(reducers.courseFormStateReducer(reducers.INITIAL_COURSE_FORM_STATE, actions.setCourseInstructor('John'))).toEqual({
            id: null,
            name: '',
            description: '',
            instructor: 'John',
            centPrice: 0
        })
    });
    it('coursesForm should handle SET_COURSE_CENT_PRICE action', () => {
        expect(reducers.courseFormStateReducer(reducers.INITIAL_COURSE_FORM_STATE, actions.setCourseCentPrice(1000))).toEqual({
            id: null,
            name: '',
            description: '',
            instructor: '',
            centPrice: 1000
        })
    });
    it('coursesForm should handle SET_COURSE_FORM action', () => {
        expect(reducers.courseFormStateReducer(reducers.INITIAL_COURSE_FORM_STATE, actions.setCourseForm({
            id: 1,
            name: 'Go',
            description: 'Go course',
            instructor: 'John',
            centPrice: 1000
        }))).toEqual({
            id: 1,
            name: 'Go',
            description: 'Go course',
            instructor: 'John',
            centPrice: 1000
        })
    });
    it('coursesForm should handle RESET_COURSE_FORM action', () => {
        expect(reducers.courseFormStateReducer({
            id: 1,
            Name: 'Go'
        }, actions.resetCourseForm())).toEqual(INITIAL_COURSE_FORM_STATE)
    });
});


describe('Test courses list reducers', () => {
    it('coursesList should handle initial state', () => {
        expect(reducers.coursesListStateReducer(undefined, {})).toEqual([])
    });
    it('coursesList should handle SET_COURSES_LIST action', () => {
        expect(reducers.coursesListStateReducer(reducers.INITIAL_COURSES_LIST_STATE, actions.setCoursesList([{
            id: 1,
            Name: 'Go'
        }]))).toEqual(
            [{id: 1, Name: 'Go'}]
        );
    });
    it('coursesList should handle ADD_TO_COURSES_LIST action', () => {
        expect(reducers.coursesListStateReducer(reducers.INITIAL_COURSES_LIST_STATE, actions.addToCoursesList({
            id: 1,
            Name: 'Go'
        }))).toEqual(
            [{id: 1, Name: 'Go'}]
        );
    });
    it('coursesList should handle ADD_TO_COURSES_LIST action and append course', () => {
        expect(reducers.coursesListStateReducer([{id: 1, Name: 'Go'}], actions.addToCoursesList({
            id: 2,
            Name: 'Go II'
        }))).toEqual([
            {id: 1, Name: 'Go'},
            {id: 2, Name: 'Go II'}
        ])
        ;
    });
    it('coursesList should handle ADD_TO_COURSES_LIST action and update course', () => {
        expect(reducers.coursesListStateReducer([{id: 1, Name: 'Go'}], actions.addToCoursesList({
            id: 1,
            Name: 'Go II'
        }))).toEqual([
            {id: 1, Name: 'Go II'},
        ])
        ;
    });
});
