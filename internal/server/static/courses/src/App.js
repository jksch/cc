import React from 'react';
import './App.css';
import Courses from "./app/containers/Courses";
import {applyMiddleware, createStore} from "redux";
import {rootReducer} from "./app/Reducers";
import {Provider} from "react-redux";
import thunk from "redux-thunk";

const store = createStore(
    rootReducer,
    applyMiddleware(thunk)
);

const App = () => {
    return (
        <Provider store={store}>
            <Courses/>
        </Provider>
    )
};

export default App;
