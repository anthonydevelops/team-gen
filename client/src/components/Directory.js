import React from 'react'
import { Nav, NavItem, NavLink } from 'reactstrap';

export default class Directory extends React.Component {
    render() {
        return (
            <div>
                <Nav vertical>
                    <NavItem>
                        <NavLink href="#">Dashboard</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink href="#">Agenda</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink href="#">Teams</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink href="#">Discussion</NavLink>
                    </NavItem>
                </Nav>
            </div>
        )
    }
}