import React from 'react'
import { Collapse, Navbar, NavbarToggler, NavbarBrand, Nav, NavItem, NavLink } from 'reactstrap';

export default class Header extends React.Component {
    constructor(props) {
        super(props);

        this.toggleNavbar = this.toggleNavbar.bind(this);
        this.state = {
            collapsed: true
        };
    }

    toggleNavbar() {
        this.setState({
            collapsed: !this.state.collapsed
        });
    }

    render() {
        return (
            <div className="nav-wrapper">
                <div>
                    <Navbar
                        className="fixed-top"
                        style={{ backgroundColor: "rgba(0, 0, 0, 0.5)" }}
                        dark
                        expand="md"
                        role="navigation"
                    >
                        <NavbarBrand href="/">TeamGen</NavbarBrand>
                        <NavbarToggler onClick={this.toggleNavbar} className="mr-2" />
                        <Collapse isOpen={!this.state.collapsed} navbar>
                            <Nav className="ml-auto" navbar>
                                <NavItem>
                                    <NavLink href="#about">Create Team</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="#software">Join Team</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="#projects">Settings</NavLink>
                                </NavItem>
                            </Nav>
                        </Collapse>
                    </Navbar>
                </div>
            </div>
        )
    }

}