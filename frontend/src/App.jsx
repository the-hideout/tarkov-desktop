import logo from './assets/images/logo-universal.png';
import React, { useState, useEffect } from "react";
import './App.css';
import { QueueScanner, RestartQueueScanner, StopQueueScanner, GetQueueLogs, DeviceID } from "../wailsjs/go/main/App";
import SyntaxHighLighter from 'react-syntax-highlighter';
import { dracula } from 'react-syntax-highlighter/dist/cjs/styles/hljs';
import { QRCodeSVG } from 'qrcode.react';

function App() {

    function startQueueScanner() {
        QueueScanner();
    }
    function stopQueueScanner() {
        StopQueueScanner();
    }
    function restartQueueScanner() {
        RestartQueueScanner();
    }

    const deviceId = DeviceID();

    const [queueLogs, getQueueLogs] = useState("");
    const [isShown, setIsShown] = useState(false);

    const handleClick = _ => {
        setIsShown(current => !current);
    };

    useEffect(() => {
        setInterval(async () => {
            getQueueLogs(await GetQueueLogs());
        }, 1000);
    }, []);

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="input" className="input-box">
                <button className="btn" onClick={startQueueScanner}>Start</button>
                <button className="btn" onClick={stopQueueScanner}>Stop</button>
                <button className="btn" onClick={restartQueueScanner}>Restart</button>
                <button className="btn" onClick={handleClick}>QR</button>
            </div>

            {isShown && (
                <QRCodeSVG value={deviceId} />
            )}

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
