
import configureMockStore from 'redux-mock-store';
import thunk from 'redux-thunk';
import MockAdapter from 'axios-mock-adapter'

import * as actions from './Actions.js'
import * as messages from './Messages.js'

const mockStore = configureMockStore([thunk]);

describe('Test course global actions', () => {
    it('setLoading should create a SET_LOADING action', () => {
        expect(actions.setLoading(true)).toEqual({
            type: actions.SET_LOADING,
            loading: true,
        })
    });
    it('setMessage should create a SET_MESSAGE action', () => {
        expect(actions.setMessage('REQUEST_FAILED')).toEqual({
            type: actions.SET_MESSAGE,
            message: 'REQUEST_FAILED',
        })
    });
});

describe('Test courses form actions', () => {
    it('setCourseName should create a SET_COURSE_NAME action', () => {
        expect(actions.setCourseName('Go')).toEqual({
            type: actions.SET_COURSE_NAME,
            name: 'Go',
        })
    });
    it('setCourseDescription should create a SET_COURSE_DESCRIPTION action', () => {
        expect(actions.setCourseDescription('Go course')).toEqual({
            type: actions.SET_COURSE_DESCRIPTION,
            description: 'Go course',
        })
    });
    it('setCourseInstructor should create a SET_COURSE_INSTRUCTOR action', () => {
        expect(actions.setCourseInstructor('John')).toEqual({
            type: actions.SET_COURSE_INSTRUCTOR,
            instructor: 'John',
        })
    });
    it('setCourseCentPrice should create a SET_COURSE_CENT_PRICE action', () => {
        expect(actions.setCourseCentPrice(1000)).toEqual({
            type: actions.SET_COURSE_CENT_PRICE,
            centPrice: 1000,
        })
    });
    it('setCourseForm should create a SET_COURSE_FORM action', () => {
        expect(actions.setCourseForm({
            id: 1,
            name: 'Go',
            description: 'Go course',
            instructor: 'John',
            centPrice: 1000
        })).toEqual({
            type: actions.SET_COURSE_FORM,
            course:{
                id: 1,
                name: 'Go',
                description: 'Go course',
                instructor: 'John',
                centPrice: 1000
            }
        })
    });
    it('resetCourseForm should create a RESET_COURSE_FORM action', () => {
        expect(actions.resetCourseForm()).toEqual({
            type: actions.RESET_COURSE_FORM
        })
    });
});

describe('Test courses list actions', () => {
    it('setCoursesList should create a SET_COURSES_LIST action', () => {
        expect(actions.setCoursesList([{id: 1, name: 'Go'}])).toEqual({
            type: actions.SET_COURSES_LIST,
            courses: [{id: 1, name: 'Go'}],
        })
    });
    it('addToCoursesList should create a ADD_TO_COURSES_LIST action', () => {
        expect(actions.addToCoursesList({id: 1, name: 'Go'})).toEqual({
            type: actions.ADD_TO_COURSES_LIST,
            add: {id: 1, name: 'Go'},
        })
    });
});


describe('Test async courses actions', () => {
    it('getCoursesList should succeed and update courses', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
            { type: actions.SET_COURSES_LIST, courses: [{id: 1, name: 'Go'}]}
        ];

        const api = new MockAdapter(actions.Client);
        api.onGet('/api/courses').reply(200, [{id: 1, name: 'Go'}]);

        const store = mockStore({});
        return store.dispatch(actions.getCoursesList()).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('getCoursesList should succeed not update on no data', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
        ];

        const api = new MockAdapter(actions.Client);
        api.onGet('/api/courses').reply(200);

        const store = mockStore({});
        return store.dispatch(actions.getCoursesList()).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('getCoursesList should fail on 201 and dispatch empty list', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
            { type: actions.SET_MESSAGE, message: messages.ERR_GET_COURSES_LIST_FAILED},
            { type: actions.SET_COURSES_LIST, courses: []}
        ];

        const api = new MockAdapter(actions.Client);
        api.onGet('/api/courses').reply(201, null);

        const store = mockStore({});
        return store.dispatch(actions.getCoursesList()).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('getCoursesList should fail on 500 and dispatch empty list', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false},
            { type: actions.SET_MESSAGE, message: messages.ERR_GET_COURSES_LIST_FAILED},
            { type: actions.SET_COURSES_LIST, courses: []}
        ];

        const api = new MockAdapter(actions.Client);
        api.onGet('/api/courses').reply(400);

        const store = mockStore({});
        return store.dispatch(actions.getCoursesList()).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('postCourse should succeed and add course to courses', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
            { type: actions.SET_MESSAGE, message: messages.POST_COURSE_SUCCESSFUL},
            { type: actions.ADD_TO_COURSES_LIST, add: {id: 1, name: 'Go'}},
            { type: actions.RESET_COURSE_FORM}
        ];

        const api = new MockAdapter(actions.Client);
        api.onPost('/api/courses').reply(201, '1');

        const store = mockStore({});
        return store.dispatch(actions.postCourse({id: 1, name: 'Go'})).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('postCourse should fail on 200 and show error message', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
            { type: actions.SET_MESSAGE, message: messages.ERR_POST_COURSE_FAILED},
        ];

        const api = new MockAdapter(actions.Client);
        api.onPost('/api/courses').reply(200);

        const store = mockStore({});
        return store.dispatch(actions.postCourse({id: 1, name: 'Go'})).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
    it('postCourse should fail on 500 and show error message', () => {
        const expActions = [
            { type: actions.SET_LOADING, loading: true},
            { type: actions.SET_LOADING, loading: false },
            { type: actions.SET_MESSAGE, message: messages.ERR_POST_COURSE_FAILED},
        ];

        const api = new MockAdapter(actions.Client);
        api.onPost('/api/courses').reply(500);

        const store = mockStore({});
        return store.dispatch(actions.postCourse({id: 1, name: 'Go'})).then(() =>{
            expect(store.getActions()).toEqual(expActions);
        });
    });
});
