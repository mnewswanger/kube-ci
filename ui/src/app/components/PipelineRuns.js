import React, {Component} from 'react';

import {green500} from 'material-ui/styles/colors';
import ActionHome from 'material-ui/svg-icons/action/home';

import RaisedButton from 'material-ui/RaisedButton';
import {
  Table,
  TableBody,
  TableFooter,
  TableHeader,
  TableHeaderColumn,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';
import TextField from 'material-ui/TextField';

import AddPipeline from './modals/AddPipeline';

const styles = {
  propContainer: {
    width: 200,
    overflow: 'hidden',
    margin: '20px auto 0',
  },
  propToggleHeader: {
    margin: '20px auto 10px',
  },
};

const tableData = [
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"},
    {id: "01234567789abcdef", name: "Name", status: "Running", start_time: "2017-04-01 00:00:00", end_time: "2017-04-01 00:00:00"}
];

/**
 * A more complex example, allowing the table height to be set, and key boolean properties to be toggled.
 */
export default class PipelineRuns extends Component {
  handleChange = (event) => {
    this.setState({height: event.target.value});
  };

  render() {
    return (
        <Table
          height="300px"
          fixedHeader={true}
          fixedFooter={true}
          selectable={false}
          showCheckboxes={false}
        >
          <TableHeader
            displaySelectAll={false}
            adjustForCheckbox={false}
            enableSelectAll={false}
          >
            <TableRow>
              <TableHeaderColumn colSpan="6" style={{textAlign: 'center'}}>
                <h1>Pipeline Runs</h1>
              </TableHeaderColumn>
            </TableRow>
            <TableRow>
                <TableRowColumn></TableRowColumn>
              <TableHeaderColumn tooltip="Pipeline Instance ID">ID</TableHeaderColumn>
              <TableHeaderColumn tooltip="Pipeline Name">Pipeline</TableHeaderColumn>
              <TableHeaderColumn tooltip="Status of the Pipeline Instance">Status</TableHeaderColumn>
              <TableHeaderColumn tooltip="Start time of the Pipeline Instance">Start Time</TableHeaderColumn>
              <TableHeaderColumn tooltip="Completion time of the Pipeline Instance">Completion Time</TableHeaderColumn>
            </TableRow>
          </TableHeader>
          <TableBody
            displayRowCheckbox={false}
            showRowHover={true}
            stripedRows={false}
          >
            {tableData.map( (row, index) => (
              <TableRow key={index}>
              <TableHeaderColumn><ActionHome /></TableHeaderColumn>
                <TableRowColumn>{row.id}</TableRowColumn>
                <TableRowColumn>{row.name}</TableRowColumn>
                <TableRowColumn>{row.status}</TableRowColumn>
                <TableRowColumn>{row.start_time}</TableRowColumn>
                <TableRowColumn>{row.end_time}</TableRowColumn>
              </TableRow>
              ))}
          </TableBody>
          <TableFooter>
            <TableRow>
              <TableRowColumn colSpan="6" style={{textAlign: 'center'}}>
                <AddPipeline />
              </TableRowColumn>
            </TableRow>
          </TableFooter>
        </Table>
    );
  }
}

