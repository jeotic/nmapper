import React, {Component} from 'react';
import './App.css';
import {Button, Grid, TextField} from "@material-ui/core";
import { RunsComponent } from "./components/runs";

export class App extends Component {
    state: { addr: string|null, inputAddr: string } = { addr: null, inputAddr: "81.107.115.203" };

    public render () {
        const { addr, inputAddr } = this.state;

        return (
            <div className="App">
                NOTE: File Upload goes to API. Simply go back to get here. This should be moved to an Ajax call

                <header className="App-header">
                    NMap Viewer

                    <form
                        encType="multipart/form-data"
                        action="http://localhost:8000/upload"
                        method="post"
                    >
                        <input type="file" name="xml"/>
                        <Button variant="contained" color="primary" type="submit">Upload</Button>
                    </form>
                </header>
                <Grid container
                      spacing={1}
                      direction="column"
                      alignItems="center"
                      justify="center" >
                    <Grid item>
                        <Grid container>
                            <Grid item xs={5}>
                                Enter IP Address:
                            </Grid>
                            <Grid item xs={4}>
                                <TextField value={inputAddr} onChange={e => this.setState({inputAddr: e.target.value})}></TextField>
                            </Grid>
                            <Grid item xs={3}>
                                <Button variant="contained" color="primary" onClick={() => this.setState({ addr: inputAddr })}>
                                    Enter
                                </Button>
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>

                {addr && <RunsComponent addr={addr} />}
            </div>
        );
    }
}

// 81.107.115.203
