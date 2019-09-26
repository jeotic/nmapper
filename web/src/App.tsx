import React, {Component} from 'react';
import './App.css';
import {Grid, Paper, TextField} from "@material-ui/core";
import { RunsComponent } from "./components/runs";

export class App extends Component {
    state: { addr: string|null } = { addr: "81.107.115.203" };
    public render () {
        const { addr } = this.state;

        return (
            <div className="App">
                <header className="App-header">
                    NMap Viewer
                </header>
                <TextField onChange={e => this.setState({addr: e.target.value})}></TextField>

                {addr && <RunsComponent addr={addr} />}
            </div>
        );
    }
}