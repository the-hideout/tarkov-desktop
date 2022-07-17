import logo from './assets/images/logo-universal.png';
import React, { useState, useEffect } from "react";
import './App.css';
import { QueueScanner, GetQueueLogs } from "../wailsjs/go/main/App";
import SyntaxHighLighter from 'react-syntax-highlighter';
import { dracula } from 'react-syntax-highlighter/dist/cjs/styles/hljs';

function App() {

    function queueScanner() {
        QueueScanner();
    }

    const [queueLogs, getQueueLogs] = useState("");

    useEffect(() => {
        setInterval(async () => {
            getQueueLogs(await GetQueueLogs());
        }, 1000);
    }, []);

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="input" className="input-box">
                <button className="btn" onClick={queueScanner}>Start</button>
            </div>
            <p>Console Output:</p>
            <div className='console-output'>
                <SyntaxHighLighter language="javascript" style={dracula}>
                    {queueLogs}
                </SyntaxHighLighter>
            </div>
        </div>
    )
}

export default App
