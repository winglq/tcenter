import React from 'react';
import { Navbar, Nav, NavItem } from 'react-bootstrap';
import LoginBtn from './loginbtn.js'

class Menu extends React.Component {

  render() {
    return (
      <Navbar collapseOnSelect>
        <Navbar.Header>
          <Navbar.Brand className="nav-brand">
            <a href="#brand">小伙伴</a>
          </Navbar.Brand>
          <Navbar.Toggle />
        </Navbar.Header>
        <Navbar.Collapse>
          <Nav pullRight>
            <NavItem className="nav-item" eventKey={1} href="#">
              所有课程
            </NavItem>
            <NavItem eventKey={2} href="#">
              <LoginBtn text={"  登录  |  注册  "} />
            </NavItem>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    );
  }
}

export default Menu;
