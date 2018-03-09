import React from 'react';
import { Button } from 'react-bootstrap';

class LoginBtn extends React.Component {
  render () {
    return (
      <Button bsStyle="info">{ this.props.text }</Button>
    )
  }
}

export default LoginBtn
