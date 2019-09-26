import React, { Component } from 'react';
import { Grid } from '@material-ui/core';
import { Ajax } from './ajax';
import { TableComponent } from './table';
import { getKey } from '../global-key';

export class HostsComponent extends Component<IHostsComponentProp> {
  public render() {
    const { RunId } = this.props;

    return (
      <Ajax url={`/runs/${RunId}/hosts`}>
        {(hosts: any[]) => (
          <Grid container key={getKey()}>
            Hosts
            {hosts.map(host => (
              <Grid item xs={12} key={getKey()}>
                <TableComponent rows={[host]} />
                <Ajax url={`/runs/${RunId}/hosts/${host.Id}/ports`}>
                  {(ports: object[]) => (
                    <div>
                      Ports
                      <TableComponent rows={ports} />
                    </div>
                  )}
                </Ajax>
                <Ajax url={`/runs/${RunId}/hosts/${host.Id}/names`}>
                  {(hostNames: object[]) => (
                    <div>
                      Names:
                      <TableComponent rows={hostNames} />
                    </div>
                  )}
                </Ajax>
              </Grid>
            ))}
          </Grid>
        )}
      </Ajax>
    );
  }
}

export interface IHostsComponentProp {
  RunId: number;
}
