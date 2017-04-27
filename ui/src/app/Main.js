import React, {Component} from 'react';

// Theming
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
// import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
// import getMuiTheme from 'material-ui/styles/getMuiTheme';
//import {deepOrange500} from 'material-ui/styles/colors';

// Material-UI Elements
import Paper from 'material-ui/Paper';

// Application Elements
import Navigation from './components/Navigation';
import PipelineRuns from './components/PipelineRuns';

// const muiTheme = getMuiTheme(darkBaseTheme);

class Main extends Component {
    constructor(props, context) {
    super(props, context);

    this.state = {
        open: false,
    };
}

handleRequestClose = () => {
    this.setState({
        open: false,
    });
}

handleTouchTap = () => {
    this.setState({
        open: true,
    });
}

render() {
    return (
        <MuiThemeProvider>
            <div>
                <Navigation />
                <Paper style={{padding: 20, margin: 20,}} zDepth={2}>
                    <PipelineRuns />
                </Paper>
            </div>
        </MuiThemeProvider>
    );
    }
}

export default Main;
