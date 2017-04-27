import React, {Component} from 'react';

import AppBar from 'material-ui/AppBar';
import Drawer from 'material-ui/Drawer';
import {List, ListItem, makeSelectable} from 'material-ui/List';
import MenuItem from 'material-ui/MenuItem';
import RaisedButton from 'material-ui/RaisedButton';

const SelectableList = makeSelectable(List);

export default class Navigation extends React.Component {

	constructor(props) {
		super(props);
		this.state = {open: false};
	}

	handleToggle = () => this.setState({open: !this.state.open});

	render() {
		return (
			<div>
				<AppBar
					title="Kube CI"
					onLeftIconButtonTouchTap={this.handleToggle}
				/>
                <Drawer
                    docked={false}
                    width={200}
                    open={this.state.open}
                    onRequestChange={(open) => this.setState({open})}
                >
                    <ListItem
                        primaryText="Activity"
                        primaryTogglesNestedList={true}
                        nestedItems={[
                            <ListItem
                                key="0"
                                primaryText="Currently Active"
                                value="/activity/current"
                                href="#/activity/current"
                            />,
                            <ListItem
                                key="1"
                                primaryText="Recently Run"
                                value="/activity/recent"
                                href="#/activity/recent"
                            />,
                            <ListItem
                                key="2"
                                primaryText="Failed Runs"
                                value="/activity/failed"
                                href="#/activity/failed"
                            />,
                        ]}
                    />
                    <ListItem
                        primaryText="Pipelines"
                        primaryTogglesNestedList={true}
                        nestedItems={[
                            <ListItem
                                key="0"
                                primaryText="View All"
                                value="/pipelines"
                                href="#/pipelines"
                            />,
                            <ListItem
                                key="1"
                                primaryText="Active Instances"
                                value="/pipelines/active"
                                href="#/pipelines/active"
                            />,
                        ]}
                    />
                    <ListItem
                        primaryText="Administration"
                        primaryTogglesNestedList={true}
                        nestedItems={[
                            <ListItem
                                key="0"
                                primaryText="Webhook Targets"
                                value="/admin/webhooks"
                                href="#/admin/webhooks"
                            />,
                            <ListItem
                                key="1"
                                primaryText="Kubernetes Clusters"
                                value="/admin/kubernetes"
                                href="#/admin/kubernetes"
                            />,
                        ]}
                    />
                </Drawer>
			</div>
		);
	}
}
