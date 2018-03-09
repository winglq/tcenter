import React from 'react';
import { Thumbnail, Button } from 'react-bootstrap';

class ControlledThumbnail extends React.Component {
  render() {
    return (
      <Thumbnail src="/thumbnaildiv.png" alt="242x200">
        <h3>Thumbnail label</h3>
        <p>Description</p>
        <p>
          <Button bsStyle="primary">Button</Button>&nbsp;
          <Button bsStyle="default">Button</Button>
        </p>
      </Thumbnail>
      );
  }
}

export default ControlledThumbnail
