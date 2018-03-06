import ReactDOM from 'react-dom';
import React from 'react';
import Welcome from './welcome.js';
import ControlledCarousel from './carousel.js';

class ControlledCarousel2 extends React.Component {
        render() {
                return (
                                <div id="ccl">
                                  <ControlledCarousel />
                                </div>
                       )
        }
}
ReactDOM.render(
  <ControlledCarousel2 />,
  document.getElementById('root')
);
