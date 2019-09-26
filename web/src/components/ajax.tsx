import React, { Component, ReactNode } from 'react';
import axios, { AxiosError, AxiosResponse } from 'axios';
import { Paper } from '@material-ui/core';

export class Ajax extends Component<IAjaxProps> {
  state: { loading: boolean; data: unknown; error: AxiosError | null } = {
    loading: true,
    data: undefined,
    error: null
  };

  public constructor(props: IAjaxProps) {
    super(props);
  }

  componentDidMount(): void {
    const loc = window.location;
    axios
      .get(`${loc.protocol}//${loc.hostname}:8000${this.props.url}`)
      .then(response => {
        this.setState({ loading: false, data: response.data, error: null });
      })
      .catch(error => {
        this.setState({ error, loading: false });
      });
  }

  public render() {
    const { children } = this.props;
    const { error, loading, data } = this.state;

    if (error !== null) {
      if (error.response && error.response.status === 404) {
        return <Paper>NOT FOUND</Paper>;
      }

      return <div>ERROR</div>;
    }

    if (loading) {
      return <div>LOADING</div>;
    }

    return <div>{children(data)}</div>;
  }
}

export interface IAjaxProps {
  children: (data: any) => ReactNode;
  url: string;
}
