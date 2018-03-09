import React from 'react';
import Menu from './navbar.js';
import ControlledJumbotron from './jumbotron.js'
import ControlledThumbnail from './thumbnail.js'
import { Row, Grid, Col } from 'react-bootstrap';

class App extends React.Component {
  render() {
    return (
      <div className="app">
        <Grid>
          <Row>
            <Col>
              <div>
                <Menu />
              </div>
            </Col>
          </Row>
          <Row>
            <Col>
              <div>
                <ControlledJumbotron />
              </div>
            </Col>
          </Row>
          <Row>
            <Col md={4}>
              <ControlledThumbnail />
            </Col>
            <Col md={4}>
              <ControlledThumbnail />
            </Col>
            <Col md={4}>
              <ControlledThumbnail />
            </Col>
          </Row>
        </Grid>
      </div>
    );
  }
}

export default App;
