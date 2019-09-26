import React, {Component} from 'react';
import {
    createStyles,
    Grid,
    makeStyles,
    Paper,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableRow, Theme
} from "@material-ui/core";
import {Ajax, IAjaxProps} from "./ajax";
import {Run} from "../interfaces/run";
import {getKey} from "../global-key";
import {ReactNodeLike} from "prop-types";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            width: '100%',
            marginTop: theme.spacing(3),
            overflowX: 'auto',
        },
        table: {
            minWidth: 650,
        },
    }),
);

function StyledTable({ children }: { children: ReactNodeLike }) {
    const classes = useStyles();
    return <Table className={classes.root}>{children}</Table>
}

function StyledPaper({ children }: { children: ReactNodeLike }) {
    const classes = useStyles();
    return <Paper className={classes.root}>{children}</Paper>
}

export class TableComponent extends Component<ITableComponent> {
    public render() {
        const { rows } = this.props;
        let columns: string[] = [];

        if (rows[0]) {
            columns = Object.keys(rows[0] as object);
        }

        return (
            <StyledPaper>
                <StyledTable>
                    <TableHead>
                        <TableRow key={getKey()}>
                        {columns.map((column: string) => (
                            <TableCell>
                                {column}
                            </TableCell>
                        ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <TableRow key={getKey()}>
                        {rows.map((row: Whatever) => {
                            return columns.map((column: string) => (
                                    <TableCell>
                                        {!(row[column] instanceof Object) && row[column]}
                                        {(row[column] instanceof Object) &&
                                        <div>
                                            {column}
                                            <TableComponent rows={[row[column]]} />
                                        </div>
                                        }
                                    </TableCell>
                            ))
                        })}
                        </TableRow>
                    </TableBody>
                </StyledTable>
            </StyledPaper>
        );
    }
}

export interface ITableComponent {
    rows: Whatever[];
}

interface Whatever {
    [key: string]: any;
}