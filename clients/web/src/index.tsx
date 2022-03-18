import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import StartApp from "./auth";

const appRenderer = () => ReactDOM.render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
    document.getElementById('root')
);

StartApp(appRenderer)
