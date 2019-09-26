import React, {Component} from 'react';
import {Grid, Paper, Table, TableBody, TableCell, TableHead, TableRow} from "@material-ui/core";
import {Ajax, IAjaxProps} from "./ajax";
import {Run} from "../interfaces/run";
import {TableComponent} from "./table";

export class HostsComponent extends Component<IHostsComponentProp> {
    public render() {
        const { RunId } = this.props;

        return (
            <Ajax url={`/runs/${RunId}/hosts`}>
                {
                    (hosts: any[]) => (
                        <Grid xs={12}>
                            Hosts
                            {hosts.map(host => (
                                <Grid xs={12}>
                                    <TableComponent rows={[host]} />
                                    <Ajax url={`/runs/${RunId}/hosts/${host.Id}/ports`}>
                                        {
                                            (ports: object[]) => (
                                                <div>
                                                    Ports
                                                    <TableComponent rows={ports} />
                                                </div>
                                            )
                                        }
                                    </Ajax>
                                    <Ajax url={`/runs/${RunId}/hosts/${host.Id}/names`}>
                                        {
                                            (hostNames: object[]) => (
                                                <div>
                                                    Names:
                                                    <TableComponent rows={hostNames} />
                                                </div>
                                            )
                                        }
                                    </Ajax>
                                </Grid>
                            ))}
                        </Grid>
                    )
                }
            </Ajax>
        );
    }
}

export interface IHostsComponentProp {
    RunId: number;
}