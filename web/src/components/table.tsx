import React, { Component } from 'react';
import {
  createStyles,
  makeStyles,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Theme
} from '@material-ui/core';
import { getKey } from '../global-key';
import { ReactNodeLike } from 'prop-types';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%'
    },
    paper: {
      marginTop: theme.spacing(3),
      width: '100%',
      overflowX: 'auto',
      marginBottom: theme.spacing(2)
    },
    table: {
      minWidth: 650
    }
  })
);

function StyledContainer({ children }: { children: ReactNodeLike }) {
  const classes = useStyles();
  return <div className={classes.root}>{children}</div>;
}

function StyledTable({ children }: { children: ReactNodeLike }) {
  const classes = useStyles();
  return (
    <Table className={classes.table} size="small">
      {children}
    </Table>
  );
}

function StyledPaper({ children }: { children: ReactNodeLike }) {
  const classes = useStyles();
  return <Paper className={classes.paper}>{children}</Paper>;
}

export class TableComponent extends Component<ITableComponent> {
  public render() {
    const { rows } = this.props;
    let columns: string[] = [];

    if (rows[0]) {
      Object.keys(rows[0] as object).forEach((key: string) => {
        if (!(rows[0][key] instanceof Object)) {
          columns.push(key);
        }
      });
    }

    return (
      <StyledContainer>
        <StyledPaper>
          <StyledTable>
            <TableHead>
              <TableRow key={getKey()}>
                {columns.map((column: string) => (
                  <TableCell key={getKey()}>{column}</TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {rows.map((row: Whatever) => (
                <TableRow key={getKey()}>
                  {columns.map((column: string) => (
                    <TableCell key={getKey()}>{!(row[column] instanceof Object) && row[column]}</TableCell>
                  ))}
                </TableRow>
              ))}
            </TableBody>
          </StyledTable>
        </StyledPaper>
      </StyledContainer>
    );
  }
}

export interface ITableComponent {
  rows: Whatever[];
}

interface Whatever {
  [key: string]: any;
}
