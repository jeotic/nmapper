import React, { Component } from 'react';
import { Grid, Paper } from '@material-ui/core';
import { Ajax } from './ajax';
import { Run } from '../interfaces/run';
import { TasksComponent } from './tasks';
import { HostsComponent } from './hosts';
import { getKey } from '../global-key';
import { TableComponent } from './table';

export class RunsComponent extends Component<IRunsComponentProps> {
  public render() {
    const { addr } = this.props;

    return (
      <Ajax url={`/runs?addr=${addr}`}>
        {(runs: Run[]) => (
          <Grid container>
            {runs.map((run: any) => (
              <Grid container spacing={1} direction="column" alignItems="center" justify="center" key={getKey()}>
                <Paper key={getKey()}>
                  Run {run.Id}
                  <TableComponent rows={[run]} />
                  <Grid item xs={12} key={getKey()}>
                    {run.ScanInfo && run.ScanInfo.id && <TableComponent rows={[run.ScanInfo]} />}
                    {run.Verbose && run.Verbose.id && <TableComponent rows={[run.Verbose]} />}
                    {run.Debugging && run.Debugging.id && <TableComponent rows={[run.Debugging]} />}

                    <TasksComponent RunId={run.Id} />
                    <HostsComponent RunId={run.Id} />
                  </Grid>
                </Paper>
              </Grid>
            ))}
          </Grid>
        )}
      </Ajax>
    );
  }
}

/**
 *
 <TasksComponent RunId={run.Id} />
 <HostsComponent RunId={run.Id} />
 */

export interface IRunsComponentProps {
  addr: string;
}
