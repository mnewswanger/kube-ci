import React from 'react';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';
import RaisedButton from 'material-ui/RaisedButton';

import PipelineForm from '../forms/Pipeline';

/**
 * A modal dialog can only be closed by selecting one of the actions.
 */
export default class AddPipeline extends React.Component {
  state = {
    open: false,
  };

  handleOpen = () => {
    this.setState({open: true});
  };

  handleClose = () => {
    this.setState({open: false});
  };

  render() {
    const actions = [
      <FlatButton
        label="Cancel"
        primary={true}
        onTouchTap={this.handleClose}
      />,
      <FlatButton
        label="Submit"
        primary={true}
        disabled={true}
        onTouchTap={this.handleClose}
      />,
    ];

    return (
      <div>
        <RaisedButton label="Add Pipeline" onTouchTap={this.handleOpen} />
        <Dialog
          title="Create New Pipeline"
          actions={actions}
          modal={true}
          open={this.state.open}
        >
            <PipelineForm />
        </Dialog>
      </div>
    );
  }
}
