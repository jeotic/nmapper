import React, { Component } from 'react';
import { Grid, Paper, Table, TableBody, TableCell, TableHead, TableRow } from '@material-ui/core';
import { Ajax, IAjaxProps } from './ajax';
import { Run } from '../interfaces/run';
import { TableComponent } from './table';

export class TasksComponent extends Component<ITasksComponentProp> {
  public render() {
    const { RunId } = this.props;

    return (
      <Ajax url={`/runs/${RunId}/tasks`}>
        {(tasks: object[]) => (
          <Grid item xs={12}>
            Tasks
            <TableComponent rows={tasks}></TableComponent>
          </Grid>
        )}
      </Ajax>
    );
  }
}

export interface ITasksComponentProp {
  RunId: number;
}
