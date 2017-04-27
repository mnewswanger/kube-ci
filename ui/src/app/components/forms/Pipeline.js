import React from 'react';
import TextField from 'material-ui/TextField';

const PipelineForm = () => (
  <div>
    <TextField
      defaultValue=""
      hintText="Descriptive name for the pipeline"
      floatingLabelText="Pipeline Name"
    /><br />

    <h3>Steps</h3>

    <h3>Labels</h3>
  </div>
);

export default PipelineForm;
