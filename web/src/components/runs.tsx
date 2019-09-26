import React, { Component } from 'react';
import { Grid, makeStyles, Paper, Table, TableBody, TableCell, TableHead, TableRow } from '@material-ui/core';
import { Ajax, IAjaxProps } from './ajax';
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
            {runs.map(run => (
              <Grid xs={12}>
                <Grid xs={12}>
                  Run
                  <TableComponent rows={[run]} />
                </Grid>
                <TasksComponent RunId={run.Id} />
                <HostsComponent RunId={run.Id} />
              </Grid>
            ))}
          </Grid>
        )}
      </Ajax>
    );
  }
}

export interface IRunsComponentProps {
  addr: string;
}
