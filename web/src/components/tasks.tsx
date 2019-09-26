import React, { Component } from 'react';
import { Grid } from '@material-ui/core';
import { Ajax } from './ajax';
import { TableComponent } from './table';

export class TasksComponent extends Component<ITasksComponentProp> {
  public render() {
    const { RunId } = this.props;

    return (
      <Ajax url={`/runs/${RunId}/tasks`}>
        {(tasks: object[]) => (
          <Grid item xs={12}>
            Tasks
            <TableComponent rows={tasks} />
          </Grid>
        )}
      </Ajax>
    );
  }
}

export interface ITasksComponentProp {
  RunId: number;
}
